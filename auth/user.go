// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
)

var (
	// ErrorNoLocalCredsFound is returned if no local auth methods were found and the invoker created the config with the option WithoutBrowserLogin
	ErrorNoLocalCredsFound = errors.New("there were no credentials found present on the machine")
)

// UserSession implements the auth package's Session interface
type UserSession struct {
	browser        Browser
	NoBrowserLogin bool
}

// GetToken returns an access token obtained from either an existing session or new browser login.
func (s *UserSession) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {

	// Check for existing token in auth cache, if it exists.
	cache, readErr := Read()
	if readErr != nil {
		log.Printf("Failed to read cache from file: %s", readErr.Error())
	}

	var tok *oauth2.Token
	var err error

	// Check the expiry of the retrieved token.
	// If session expiry or the AccessTokenExpiry has passed, then reauthenticate with browser login and reassign token.
	if readErr != nil || cache.SessionExpiry.Before(time.Now()) || cache.AccessTokenExpiry.Before(time.Now()) {

		// This is a configuration option set by WithoutBrowserLogin to provide control over when or if the browser is opened
		// In the case of no valid current login information is found on the system via ClientID:ClientSecret pairing or a previous valid browser login
		if s.NoBrowserLogin {
			return nil, ErrorNoLocalCredsFound
		}

		// Login with browser.
		log.Print("No credentials found, proceeding with browser login.")

		if s.browser == nil {
			s.browser = &oauthBrowser{}
		}

		tok, err = s.browser.GetTokenFromBrowser(ctx, conf)
		if err != nil {
			return nil, fmt.Errorf("failed to get access token: %w", err)
		}

		// Update cache with newly obtained token.
		newCache := Cache{
			AccessToken:       tok.AccessToken,
			RefreshToken:      tok.RefreshToken,
			AccessTokenExpiry: tok.Expiry,
			SessionExpiry:     time.Now().Add(SessionMaxAge),
		}

		err = Write(newCache)
		if err != nil {
			log.Printf("Failed to write cache to file: %s", err.Error())
		}

	} else { // Otherwise return existing, unexpired token to continue existing session.
		tok = &oauth2.Token{
			AccessToken:  cache.AccessToken,
			RefreshToken: cache.RefreshToken,
			Expiry:       cache.AccessTokenExpiry,
		}
	}

	return tok, nil
}
