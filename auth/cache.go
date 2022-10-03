package auth

import (
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
	MaxAge time.Time `json:"max_age,omitempty"`
}

// Write saves HCP auth data in a common location in the home directory
func Write(token oauth2.Token) error {
	// get the user's home directory
	// check err
	// check if the hcp/credentials.json exists
	// create it if it doesn't

	// write access token, refresh_token, expiry, max age to credentials file

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
