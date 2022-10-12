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

	// TODO: assert json matches our input token
	rawJSON, err := os.ReadFile(credentialPath)

	var tokFromJSON oauth2.Token
	json.Unmarshal([]byte(rawJSON), tokFromJSON)
	assert.Equal(tok, tokFromJSON)

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
