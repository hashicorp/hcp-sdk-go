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
	// DO we need 0755 permissions for directory, or is 0660 sufficient? Can we set read/write permissions as a constant?

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	err = os.MkdirAll(testDirectory, 0755)
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
	// DO we need 0755 permissions for directory, or is 0660 sufficient? Can we set read/write permissions as a constant?

	credentialDirectory, credentialPath, err := setup()
	require.NoError(err)
	require.NotNil(credentialDirectory)

	err = os.MkdirAll(testDirectory, 0755)
	require.NoError(err)

	//initialize empty file with no contents
	//create and initialize separate file that already has cache in it and that cache is overwritten rather than appended to

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

func setup() (credentialDirectory, credentialPath string, err error) {

	os.Setenv(envVarCacheTestMode, "true")

	// TODO: replace with GetDirectory helper
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

	// TODO: replace with GetDirectory helper
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
