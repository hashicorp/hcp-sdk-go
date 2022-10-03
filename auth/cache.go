package auth

import "time"

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

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

// 1. Write to json file in a specific location in home directory (OS-dependent)

// Setter for access_token

// Setter for refresh_token

// Setter for expiration time

// 2. Read from json file in a specific location in home directory (OS-dependent)

// Getter for access_token

// Getter for refresh_token

// Getter for expiration time

// Constant for max age - 24 hours

// Constant for file name {HOME}"/hcp/credentials.json"

// 3. helpers

// check if directory exists

// create directory + file

// tokenToJson

// JsonToToken
