package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
)

const (
	MaxAge = time.Hour * 24

	defaultDirectory     = ".config/hcp"
	testDirectory        = "hcptest"
	fileName             = "credentials.json"
	directoryPermissions = 0755

	envVarCacheTestMode = "HCP_CACHE_TEST_MODE"
)

type Cache struct {
	// AccessToken is the bearer token used to authenticate to HCP.
	AccessToken string `json:"access_token,omitempty"`

	// RefreshToken is used to get a new access token.
	RefreshToken string `json:"refresh_token,omitempty"`

	// Expiry is when the access token will expire.
	Expiry time.Time `json:"expiry,omitempty"`

	// MaxAge is the session limit.
	MaxAge time.Duration `json:"max_age,omitempty"`
}

// Write saves HCP auth data in a common location in the home directory.
func Write(token oauth2.Token) error {
	credentialPath, credentialDirectory, err := getCredentialPaths()
	if err != nil {
		return fmt.Errorf("failed to retrieve credential path and directory: %v", err)
	}

	err = os.MkdirAll(credentialDirectory, directoryPermissions)
	if err != nil {
		return fmt.Errorf("failed to create credential directory: %v", err)
	}

	// Write access token, refresh_token, expiry, max age to credentials file.
	credentials := &Cache{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		MaxAge:       MaxAge,
	}
	credentialsJson, err := json.Marshal(credentials)
	if err != nil {
		return fmt.Errorf("failed to marshal the struct to json: %v", err)
	}

	err = os.WriteFile(credentialPath, credentialsJson, directoryPermissions)
	if err != nil {
		return fmt.Errorf("failed to write credentials to the cache file: %v", err)
	}

	return nil
}

// Read opens the saved HCP auth data and returns the token.
func Read() (*Cache, error) {
	credentialPath, _, err := getCredentialPaths()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve credential path and directory: %v", err)
	}

	rawJSON, err := os.ReadFile(credentialPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from user's credential path: %v", err)
	}

	cacheFromJSON, err := jsonToToken(rawJSON)
	if err != nil {
		return nil, fmt.Errorf("bad format: %v", err)
	}

	return cacheFromJSON, nil
}

// getDirectory returns the complete credential path and directory.
func getCredentialPaths() (credentialPath string, credentialDirectory string, err error) {
	// Get the user's home directory.
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve user's home directory path: %v", err)
	}

	directoryName := defaultDirectory
	// If in test mode, use test directory.
	if testMode, ok := os.LookupEnv(envVarCacheTestMode); ok {
		if testMode == "true" {
			directoryName = testDirectory
		}
	}

	// Determine absolute path to config file and directory.
	credentialDirectory = filepath.Join(userHome, directoryName)
	credentialPath = filepath.Join(userHome, directoryName, fileName)

	return credentialPath, credentialDirectory, nil
}

// jsonToToken converts the raw JSON into the Cache struct.
func jsonToToken(rawData []byte) (*Cache, error) {

	var cacheFromJSON Cache
	err := json.Unmarshal(rawData, &cacheFromJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the raw data to json: %v", err)
	}

	if cacheFromJSON.AccessToken == "" {
		return nil, errors.New("failed to get cache access token")
	}

	if cacheFromJSON.RefreshToken == "" {
		return nil, errors.New("failed to get cache refresh token")
	}

	return &cacheFromJSON, nil
}

// tokenToJson converts the OAuth2 token to raw JSON.
func tokenToJSON(token *oauth2.Token) ([]byte, error) {

	if token.AccessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	if token.RefreshToken == "" {
		return nil, errors.New("refresh token cannot be empty")
	}

	if token.Expiry.IsZero() {
		return nil, errors.New("token expiry cannot be empty")
	}

	credentials := &Cache{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		MaxAge:       MaxAge,
	}
	credentialsJson, err := json.Marshal(credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the struct to json: %v", err)
	}

	return credentialsJson, nil
}
