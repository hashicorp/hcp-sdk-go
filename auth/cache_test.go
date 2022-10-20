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

func TestWrite_NoDirectoryNoFile(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	now := time.Now()
	tok := oauth2.Token{
		AccessToken:  "TopSecret!",
		RefreshToken: "SoRefreshing:)",
		Expiry:       now,
	}

	assert.NoError(Write(tok))
	assert.DirExists(credentialDirectory)
	assert.FileExists(credentialPath)

	rawJSON, err := os.ReadFile(credentialPath)

	var cacheFromJSON Cache

	err = json.Unmarshal(rawJSON, &cacheFromJSON)
	require.NoError(err)

	expectedCache := Cache{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry,
		MaxAge:       MaxAge,
	}
	assert.Equal(expectedCache.AccessToken, cacheFromJSON.AccessToken)
	assert.Equal(expectedCache.RefreshToken, cacheFromJSON.RefreshToken)
	assert.Equal(expectedCache.Expiry.Format("2006-01-02T15:04:05 -07:00:00"), cacheFromJSON.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))
	assert.Equal(expectedCache.MaxAge.String(), cacheFromJSON.MaxAge.String())
	require.NoError(destroy())
}

func TestWrite_DirectoryExistsNoFile(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	err = os.MkdirAll(testDirectory, directoryPermissions)
	require.NoError(err)

	now := time.Now()
	tok := oauth2.Token{
		AccessToken:  "TopSecret!",
		RefreshToken: "SoRefreshing:)",
		Expiry:       now,
	}

	assert.NoError(Write(tok))
	assert.DirExists(credentialDirectory)
	assert.FileExists(credentialPath)

	rawJSON, err := os.ReadFile(credentialPath)

	var cacheFromJSON Cache

	err = json.Unmarshal(rawJSON, &cacheFromJSON)
	require.NoError(err)

	expectedCache := Cache{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry,
		MaxAge:       MaxAge,
	}
	assert.Equal(expectedCache.AccessToken, cacheFromJSON.AccessToken)
	assert.Equal(expectedCache.RefreshToken, cacheFromJSON.RefreshToken)
	assert.Equal(expectedCache.Expiry.Format("2006-01-02T15:04:05 -07:00:00"), cacheFromJSON.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))
	assert.Equal(expectedCache.MaxAge.String(), cacheFromJSON.MaxAge.String())
	require.NoError(destroy())
}

func TestWrite_DirectoryExistsFileExists(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	err = os.MkdirAll(testDirectory, directoryPermissions)
	require.NoError(err)

	now := time.Now()
	tok := oauth2.Token{
		AccessToken:  "TopSecret!",
		RefreshToken: "SoRefreshing:)",
		Expiry:       now,
	}
	assert.NoError(Write(tok))

	tok2 := oauth2.Token{
		AccessToken:  "AnotherTopSecret!",
		RefreshToken: "StillSoRefreshing:)",
		Expiry:       now,
	}
	assert.NoError(Write(tok2))

	assert.DirExists(credentialDirectory)
	assert.FileExists(credentialPath)

	rawJSON, err := os.ReadFile(credentialPath)

	var cacheFromJSON Cache

	err = json.Unmarshal(rawJSON, &cacheFromJSON)
	require.NoError(err)

	expectedCache := Cache{
		AccessToken:  tok2.AccessToken,
		RefreshToken: tok2.RefreshToken,
		Expiry:       tok2.Expiry,
		MaxAge:       MaxAge,
	}
	assert.Equal(expectedCache.AccessToken, cacheFromJSON.AccessToken)
	assert.Equal(expectedCache.RefreshToken, cacheFromJSON.RefreshToken)
	assert.Equal(expectedCache.Expiry.Format("2006-01-02T15:04:05 -07:00:00"), cacheFromJSON.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))
	assert.Equal(expectedCache.MaxAge.String(), cacheFromJSON.MaxAge.String())
	require.NoError(destroy())
}

func TestRead_InvalidFormat(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	err = os.MkdirAll(credentialDirectory, directoryPermissions)
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

	err = os.WriteFile(credentialPath, randomDataJSON, directoryPermissions)
	require.NoError(err)

	_, err = Read()
	require.Error(err)
	assert.EqualError(err, "bad format: failed to get cache access token")
	require.NoError(destroy())
}

func TestRead_ValidFormat(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	credentialDirectory, _, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	now := time.Now()
	tok := oauth2.Token{
		AccessToken:  "TopSecret!",
		RefreshToken: "SoRefreshing:)",
		Expiry:       now,
	}

	assert.NoError(Write(tok))

	cachePointer, err := Read()
	require.NoError(err)

	expectedCache := Cache{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry,
		MaxAge:       MaxAge,
	}

	assert.Equal(expectedCache.AccessToken, cachePointer.AccessToken)
	assert.Equal(expectedCache.RefreshToken, cachePointer.RefreshToken)
	assert.Equal(expectedCache.Expiry.Format("2006-01-02T15:04:05 -07:00:00"), cachePointer.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))
	assert.Equal(expectedCache.MaxAge.String(), cachePointer.MaxAge.String())
	require.NoError(destroy())
}

func TestGetCredentialPaths_ReturnsPaths(t *testing.T) {
	os.Setenv(envVarCacheTestMode, "true")

	assert := assertpkg.New(t)

	credentialPath, credentialDir, err := getCredentialPaths()

	assert.NoError(err)

	userHome, err := os.UserHomeDir()
	expectedDirectory := filepath.Join(userHome, testDirectory)
	expectedPath := filepath.Join(userHome, testDirectory, fileName)

	assert.Equal(expectedDirectory, credentialDir)
	assert.Equal(expectedPath, credentialPath)
}

func TestJsonToToken_BadJSONThrowsError(t *testing.T) {

	testCases := []struct {
		name          string
		rawJSON       []byte
		expectedError string
	}{
		//TODO: finish test cases
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
			rawJSON:       []byte(`{ "access_token": "", "refresh_token": "", "expiry": "", "max_age": "" }`),
			expectedError: "failed to unmarshal the raw data to json: parsing time \"\\\"\\\"\" as \"\\\"2006-01-02T15:04:05Z07:00\\\"\": cannot parse \"\\\"\" as \"2006\"",
		},
		//why is the below case passing? With a correct time, should it fail to unmarshall?
		// {
		// 	name:          "invalid JSON",
		// 	rawJSON:       []byte(`{ "access_token": "myaccesstoken", "refresh_token": "", "expiry": "2006-01-02T15:04:05 -07:00:00", "max_age": "86400000000000" }`),
		// 	expectedError: "failed to unmarshal the raw data to json: parsing time \"\\\"2006-01-02T15:04:05 -07:00:00\\\"\" as \"\\\"2006-01-02T15:04:05Z07:00\\\"\": cannot parse \" -07:00:00\\\"\" as \"Z07:00\"",
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require := requirepkg.New(t)

			_, err := jsonToToken(testCase.rawJSON)
			require.Error(err)
			require.EqualError(err, testCase.expectedError)
		})
	}
}

func TestTokenToJson_BadJSONThrowsError(t *testing.T) {
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

func setup() (credentialDirectory, credentialPath string, err error) {
	os.Setenv(envVarCacheTestMode, "true")

	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve test directory: %v", err)
	}
	credentialDirectory = filepath.Join(userHome, testDirectory)
	credentialPath = filepath.Join(credentialDirectory, fileName)

	err = os.RemoveAll(credentialDirectory)
	if err != nil {
		return "", "", fmt.Errorf("failed to clean up test directory: %v", err)
	}

	return credentialDirectory, credentialPath, nil
}

func destroy() error {
	os.Unsetenv(envVarCacheTestMode)

	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to retrieve test directory: %v", err)
	}
	credentialDirectory := filepath.Join(userHome, testDirectory)

	err = os.RemoveAll(credentialDirectory)
	if err != nil {
		return fmt.Errorf("failed to clean up test directory: %v", err)
	}

	return nil
}
