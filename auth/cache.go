package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

const (
	MaxAge        = time.Hour * 24
	directoryName = "hcp"
	fileName      = "credentials.json"
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

// Write saves HCP auth data in a common location in the home directory
func Write(token oauth2.Token) error {
	// get the user's home directory
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to retrieve user's home directory path: %v", err)
	}

	// check if the hcp/credentials.json exists and if not, create it
	credentialDirectory := filepath.Join(userHome, directoryName)
	credentialPath := filepath.Join(userHome, directoryName, fileName)

	err = os.MkdirAll(credentialDirectory, 0755)
	if err != nil {
		return fmt.Errorf("failed to create credential directory: %v", err)
	}

	// write access token, refresh_token, expiry, max age to credentials file
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

	err = os.WriteFile(credentialPath, credentialsJson, 0660)
	if err != nil {
		return fmt.Errorf("failed to write credentials to the cache file: %v", err)
	}

	return nil
}

// Setter for access_token

// Setter for refresh_token

// Setter for expiration time

// 2. Read from json file in a specific location in home directory (OS-dependent)

// Getter for access_token

// Getter for refresh_token

// Getter for expiration time

// 3. helpers

// check if directory exists

// create directory + file

// tokenToJson

// JsonToToken
