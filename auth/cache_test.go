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

func TestRead(t *testing.T) {

	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name          string
		caseSetup     func(string, string) error
		expectedError string
	}{
		{
			name:      "No Directory, No File",
			caseSetup: func(dirPath, credPath string) error { return nil },
			//expectedError: "failed to read file from user's credential path: open /Users/jolisabrown/hcptest/credentials.json: no such file or directory",
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
				//define and initialize unexpected struct type
				type redHerring struct {
					AccessToken, anotherField, lastField string
				}
				var randomData = redHerring{AccessToken: "TopSecret!", anotherField: "SoRefreshing:)", lastField: "field"}

				randomDataJSON, err := json.Marshal(randomData)
				require.NoError(err)

				//write unexpected struct to config file
				err = os.MkdirAll(dirPath, directoryPermissions)
				require.NoError(err)
				err = os.WriteFile(credPath, randomDataJSON, directoryPermissions)
				require.NoError(err)

				return nil
			},
			expectedError: "bad format: failed to get cache access token",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			credDir, credPath, err := testSetup()
			require.NoError(err)
			require.NotNil(credDir)

			// Runs specific test case setup.
			err = testCase.caseSetup(credDir, credPath)
			require.NoError(err)

			// Run the test.
			_, err = Read()

			// Make assertions
			require.Error(err)
			type error interface {
				Error() string
			}
			//TODO: consider upgrading testify to use ErrorContains method here
			assert.Contains(err.Error(), testCase.expectedError)
			require.NoError(destroy())

		})

	}

}

func TestRead_ValidFormat(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDir, _, err := testSetup()
	require.NoError(err)
	require.NotNil(credentialDir)

	now := time.Now()
	cache := Cache{
		AccessToken:       "TopSecret!",
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
	require.NoError(destroy())
}

func TestWrite(t *testing.T) {

	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name      string
		caseSetup func(string, string) error
	}{
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Sets up empty directory, no files.
			credentialDir, credentialPath, err := testSetup()
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
			require.NoError(destroy())
		})
	}

}

func TestWrite_SessionExpiryValid(t *testing.T) {

	require := requirepkg.New(t)
	cache := Cache{
		AccessToken:       "coolAccessToken",
		RefreshToken:      "coolRefreshToken",
		AccessTokenExpiry: time.Now(),
		SessionExpiry:     time.Now().Add(SessionMaxAge),
	}

	err := Write(cache)
	require.NoError(err)

}

func TestWrite_SessionExpiryInvalid(t *testing.T) {

	require := requirepkg.New(t)
	cache := Cache{
		AccessToken:       "coolAccessToken",
		RefreshToken:      "coolRefreshToken",
		AccessTokenExpiry: time.Now(),
		SessionExpiry:     time.Now().Add(time.Hour * 30),
	}

	err := Write(cache)
	require.Error(err)
	require.EqualError(err, "session expiry greater than 24 hours")

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
			expectedError: "failed to unmarshal the raw data to json: parsing time \"\\\"\\\"\" as \"\\\"2006-01-02T15:04:05Z07:00\\\"\": cannot parse \"\\\"\" as \"2006\"",
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
			expectedError: "failed to unmarshal the raw data to json: parsing time \"\\\"\\\"\" as \"\\\"2006-01-02T15:04:05Z07:00\\\"\": cannot parse \"\\\"\" as \"2006\"",
		},
		{
			name:          "empty session expiry",
			rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "myrefreshtoken", "access_token_expiry": "2022-11-20T17:10:59.273429-04:00", "session_expiry": ""}`),
			expectedError: "failed to unmarshal the raw data to json: parsing time \"\\\"\\\"\" as \"\\\"2006-01-02T15:04:05Z07:00\\\"\": cannot parse \"\\\"\" as \"2006\"",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			_, err := jsonToCache(testCase.rawJSON)
			require.Error(err)
			require.EqualError(err, testCase.expectedError)
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

//TODO: create testHelpers file to include testsetup and destroy
func TestSetup() (credentialDir, credentialPath string, err error) {
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

func Destroy() error {
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
