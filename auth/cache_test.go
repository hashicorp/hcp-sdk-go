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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			credentialDir, credentialPath, err := testSetup()
			require.NoError(err)
			require.NotNil(credentialDir)

			err = os.MkdirAll(credentialDir, directoryPermissions)
			if err != nil {
				fmt.Printf("not able to make test directory :%v", err)
			}
			require.NoError(err)
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

			// Runs specifict test case setup.
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

func TestRead_InvalidFormat(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDir, credentialPath, err := testSetup()
	require.NoError(err)
	require.NotNil(credentialDir)

	err = os.MkdirAll(credentialDir, directoryPermissions)
	if err != nil {
		fmt.Printf("not able to make test directory :%v", err)
	}
	require.NoError(err)

	type redHerring struct {
		AccessToken  string
		anotherField string
		lastField    string
	}

	randomData := redHerring{
		AccessToken:  "TopSecret!",
		anotherField: "SoRefreshing:)",
		lastField:    "field",
	}

	randomDataJSON, err := json.Marshal(randomData)
	require.NoError(err)

	err = os.WriteFile(credentialPath, randomDataJSON, directoryPermissions)
	require.NoError(err)

	_, err = Read()
	require.Error(err)
	assert.EqualError(err, "bad format: failed to get cache access token")
	require.NoError(destroy())
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
			name:          "emptyexpiry",
			rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "myrefreshtoken", "access_token_expiry": "", "session_expiry": "2022-11-20T17:10:59.273429-04:00"}`),
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

func testSetup() (credentialDir, credentialPath string, err error) {
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

func destroy() error {
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
