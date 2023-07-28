// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	assertpkg "github.com/stretchr/testify/assert"
	requirepkg "github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestRead_ValidFormat(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDir, _, err := CacheSetup(t)
	require.NoError(err)
	require.NotNil(credentialDir)

	now := time.Now()
	cache := Cache{
		AccessToken:       "some.valid.jwt",
		RefreshToken:      "SoRefreshing:)",
		AccessTokenExpiry: now,
		SessionExpiry:     now,
	}

	assert.NoError(Write(cache))

	cachePointer, err := Read()
	require.NoError(err)

	assert.Equal(cache.AccessToken, cachePointer.AccessToken)
	assert.Equal(cache.RefreshToken, cachePointer.RefreshToken)
	assert.Equal(cache.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"), cachePointer.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"))
	assert.Equal(cache.SessionExpiry.Format("2006-01-02T15:04:05 -07:00:00"), cachePointer.SessionExpiry.Format("2006-01-02T15:04:05 -07:00:00"))

	require.NoError(CacheDestroy(t))
}

func TestRead(t *testing.T) {

	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name          string
		caseSetup     func(string, string) error
		expectedError string
	}{
		{
			name:          "No Directory, No File",
			caseSetup:     func(dirPath, credPath string) error { return nil },
			expectedError: "no such file or directory",
		},
		{
			name: "Directory Exists, File Not Exists",
			caseSetup: func(dirPath, credPath string) error {
				err := os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)
				return nil
			},
			expectedError: "no such file or directory",
		},
		{
			name: "Directory Exists, Empty File Exists",
			caseSetup: func(dirPath, credPath string) error {
				err := os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)

				file, err := os.Create(credPath)
				require.NoError(err)
				require.NotNil(file)
				return nil
			},
			expectedError: "bad format: failed to unmarshal the raw data to json: unexpected end of JSON input",
		},
		{
			name: "Directory Exists, File with Unexpected Content Exists",
			caseSetup: func(dirPath, credPath string) error {
				// Define and initialize unexpected struct type.
				type redHerring struct {
					AccessToken, anotherField, lastField string
				}
				var randomData = redHerring{AccessToken: "TopSecret!", anotherField: "SoRefreshing:)", lastField: "field"}

				randomDataJSON, err := json.Marshal(randomData)
				require.NoError(err)

				// Write unexpected struct to config file.
				err = os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)
				err = os.WriteFile(credPath, randomDataJSON, directoryPermissions)
				require.NoError(err)

				return nil
			},
			expectedError: "bad format: failed to get cache access token",
		},
		{
			name: "Directory Exists, File with Badly Formatted Token",
			caseSetup: func(dirPath, credPath string) error {

				var randomData = Cache{
					AccessToken:       "TopSecret!",
					RefreshToken:      "SoRefreshing:)",
					AccessTokenExpiry: time.Now(),
					SessionExpiry:     time.Now(),
				}

				randomDataJSON, err := json.Marshal(randomData)
				require.NoError(err)

				err = os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)
				err = os.WriteFile(credPath, randomDataJSON, directoryPermissions)
				require.NoError(err)

				return nil
			},
			expectedError: "bad format: access token is not a valid JWT",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			credDir, credPath, err := CacheSetup(t)
			require.NoError(err)
			require.NotNil(credDir)

			// Runs specific test case setup.
			err = testCase.caseSetup(credDir, credPath)
			require.NoError(err)

			// Run the test.
			_, err = Read()

			// Make assertions
			assert.Error(err)
			assert.ErrorContains(err, testCase.expectedError)

			// Clean up.
			require.NoError(CacheDestroy(t))

		})

	}
}

func TestWrite(t *testing.T) {

	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name      string
		caseSetup func(string, string) error
	}{
		{
			name: "Directory Exists, File with Content Exists",
			caseSetup: func(dirPath, credPath string) error {
				err := os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)

				now := time.Now()
				cache := Cache{
					AccessToken:       "AnotherTopSecret!",
					RefreshToken:      "StillSoRefreshing:)",
					AccessTokenExpiry: now,
					SessionExpiry:     now,
				}
				require.NoError(Write(cache))
				return nil
			},
		},
		{
			name:      "No Directory, No File",
			caseSetup: func(string, string) error { return nil },
		},
		{
			name: "Directory Exists, No File",
			caseSetup: func(dirPath, credPath string) error {
				err := os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)
				return nil
			},
		},
		{
			name: "Directory Exists, Empty File Exists",
			caseSetup: func(dirPath, credPath string) error {
				err := os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)

				file, err := os.Create(credPath)
				require.NoError(err)
				require.NotNil(file)
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Sets up empty directory, no files.
			credentialDir, credentialPath, err := CacheSetup(t)
			require.NoError(err)
			require.NotNil(credentialDir)
			require.NotNil(credentialPath)

			// Runs specific test case setup.
			err = testCase.caseSetup(credentialDir, credentialPath)
			require.NoError(err)

			// Run the test.
			now := time.Now()
			cache := Cache{
				AccessToken:       "TopSecret!",
				RefreshToken:      "SoRefreshing:)",
				AccessTokenExpiry: now,
				SessionExpiry:     now,
			}

			err = Write(cache)

			// Make assertions.
			assert.NoError(err)
			assert.DirExists(credentialDir)
			assert.FileExists(credentialPath)

			rawJSON, err := os.ReadFile(credentialPath)
			require.NoError(err)

			var cacheFromJSON Cache

			err = json.Unmarshal(rawJSON, &cacheFromJSON)
			require.NoError(err)

			assert.Equal(cache.AccessToken, cacheFromJSON.AccessToken)
			assert.Equal(cache.RefreshToken, cacheFromJSON.RefreshToken)
			assert.Equal(cache.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"), cacheFromJSON.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"))
			assert.Equal(cache.SessionExpiry.Format("2006-01-02T15:04:05 -07:00:00"), cacheFromJSON.SessionExpiry.Format("2006-01-02T15:04:05 -07:00:00"))

			// Cleanup.
			require.NoError(CacheDestroy(t))
		})
	}
}

func TestWrite_SessionExpiryValid(t *testing.T) {
	require := requirepkg.New(t)

	_, _, err := CacheSetup(t)
	require.NoError(err)

	cache := Cache{
		AccessToken:       "coolAccessToken",
		RefreshToken:      "coolRefreshToken",
		AccessTokenExpiry: time.Now(),
		SessionExpiry:     time.Now().Add(SessionMaxAge),
	}

	err = Write(cache)
	require.NoError(err)

	require.NoError(CacheDestroy(t))
}

func TestWrite_SessionExpiryInvalid(t *testing.T) {
	require := requirepkg.New(t)

	_, _, err := CacheSetup(t)
	require.NoError(err)

	cache := Cache{
		AccessToken:       "coolAccessToken",
		RefreshToken:      "coolRefreshToken",
		AccessTokenExpiry: time.Now(),
		SessionExpiry:     time.Now().Add(time.Hour * 30),
	}

	err = Write(cache)
	require.Error(err)
	require.EqualError(err, "session expiry greater than 24 hours")

	require.NoError(CacheDestroy(t))
}

func TestGetCredentialPaths(t *testing.T) {
	os.Setenv(envVarCacheTestMode, "true")

	assert := assertpkg.New(t)

	credentialPath, credentialDir, err := getCredentialPaths()
	assert.NoError(err)

	userHome, err := os.UserHomeDir()
	assert.NoError(err)
	expectedDirectory := filepath.Join(userHome, testDirectory)
	expectedPath := filepath.Join(userHome, testDirectory, fileName)

	assert.Equal(expectedDirectory, credentialDir)
	assert.Equal(expectedPath, credentialPath)
}

func TestJsonToCache_InvalidFormat(t *testing.T) {

	testCases := []struct {
		name          string
		rawJSON       []byte
		expectedError string
	}{
		{
			name:          "empty JSON",
			rawJSON:       []byte("{}"),
			expectedError: "failed to get cache access token",
		},
		{
			name:          "invalid JSON",
			rawJSON:       []byte("Any random string"),
			expectedError: "failed to unmarshal the raw data to json: invalid character 'A' looking for beginning of value",
		},
		{
			name:          "empty values",
			rawJSON:       []byte(`{ "access_token": "", "refresh_token": "", "access_token_expiry": "", "session_expiry": "" }`),
			expectedError: "failed to unmarshal the raw data to json: parsing time",
		},
		{
			name:          "empty access token",
			rawJSON:       []byte(`{ "access_token": "", "refresh_token": "myrefreshtoken", "access_token_expiry": "2022-10-20T17:10:59.273429-04:00", "session_expiry": "2022-11-20T17:10:59.273429-04:00"}`),
			expectedError: "failed to get cache access token",
		},
		{
			name:          "empty refresh token",
			rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "", "access_token_expiry": "2022-10-20T17:10:59.273429-04:00", "session_expiry": "2022-11-20T17:10:59.273429-04:00" }`),
			expectedError: "failed to get cache refresh token",
		},
		{
			name:          "empty access token expiry",
			rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "myrefreshtoken", "access_token_expiry": "", "session_expiry": "2022-11-20T17:10:59.273429-04:00"}`),
			expectedError: "failed to unmarshal the raw data to json: parsing time",
		},
		{
			name:          "empty session expiry",
			rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "myrefreshtoken", "access_token_expiry": "2022-11-20T17:10:59.273429-04:00", "session_expiry": ""}`),
			expectedError: "failed to unmarshal the raw data to json: parsing time",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			_, err := jsonToCache(testCase.rawJSON)
			require.Error(err)
			require.ErrorContains(err, testCase.expectedError)
		})
	}
}

func TestTokenToJson_InvalidFormat(t *testing.T) {
	testCases := []struct {
		name          string
		token         oauth2.Token
		expectedError string
	}{
		{
			name: "no access token",
			token: oauth2.Token{
				AccessToken:  "",
				RefreshToken: "exampleRefreshToken",
				Expiry:       time.Now(),
			},
			expectedError: "access token cannot be empty",
		},
		{
			name: "no refresh token",
			token: oauth2.Token{
				AccessToken:  "exampleAccessToken",
				RefreshToken: "",
				Expiry:       time.Now(),
			},
			expectedError: "refresh token cannot be empty",
		},
		{
			name: "no token expiry",
			token: oauth2.Token{
				AccessToken:  "exampleAccessToken",
				RefreshToken: "exampleRefreshToken",
				Expiry:       time.Time{},
			},
			expectedError: "token expiry cannot be empty",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			_, err := tokenToJSON(&testCase.token)
			require.EqualError(err, testCase.expectedError)
		})
	}
}

// TODO: create testHelpers file to include testsetup and destroy
func CacheSetup(t *testing.T) (credentialDir, credentialPath string, err error) {
	os.Setenv(envVarCacheTestMode, "true")

	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve test directory: %v", err)
	}
	credentialDir = filepath.Join(userHome, testDirectory)
	credentialPath = filepath.Join(credentialDir, fileName)

	err = os.RemoveAll(credentialDir)
	if err != nil {
		return "", "", fmt.Errorf("failed to clean up test directory: %v", err)
	}

	return credentialDir, credentialPath, nil
}

func CacheDestroy(t *testing.T) error {
	os.Unsetenv(envVarCacheTestMode)

	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to retrieve test directory: %v", err)
	}
	credentialDir := filepath.Join(userHome, testDirectory)

	err = os.RemoveAll(credentialDir)
	if err != nil {
		return fmt.Errorf("failed to clean up test directory: %v", err)
	}

	return nil
}
