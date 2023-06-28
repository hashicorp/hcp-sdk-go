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
	// ErrorNoLocalCredsFound is returned if no client or user credentials were found and the invoker created the config with the option WithoutBrowserLogin
	ErrorNoLocalCredsFound = errors.New("there were no credentials found present on the machine")
)

// UserSession implements the auth package's Session interface
type UserSession struct {
	browser        Browser
	NoBrowserLogin bool

	// refreshTokenSource is encapsulated in this closure in order to mock it in
	// tests
	refreshTokenSourceFunc func(context.Context, *oauth2.Config, *oauth2.Token) oauth2.TokenSource
}

// GetToken returns an access token obtained from either an existing session or new browser login.
func (s *UserSession) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {

	// Check for existing token in auth cache, if it exists.
	cache, readErr := Read()
	if readErr != nil {
		log.Printf("Failed to read cache from file: %s", readErr.Error())
	}

	// Check the expiry of the retrieved token.
	switch {
	// If the file wasn't found or the session is fully expired then
	// reauthenticate with browser login and rewrite token.
	case readErr != nil, cache.SessionExpiry.Before(time.Now()):

		return s.doBrowserLogin(ctx, conf)

	// The access token is expired, but the refresh token should still be
	// valid. Fetch a new token without having to open the browser and
	// rewrite the cache file.
	case cache.AccessTokenExpiry.Before(time.Now()):

		// This allows us to overide with a mock in tests
		if s.refreshTokenSourceFunc == nil {
			s.refreshTokenSourceFunc = func(ctx context.Context, conf *oauth2.Config, token *oauth2.Token) oauth2.TokenSource {
				return conf.TokenSource(ctx, token)
			}
		}

		// Refresh the token.
		refreshed, err := s.refreshTokenSourceFunc(ctx, conf, &oauth2.Token{
			AccessToken:  cache.AccessToken,
			RefreshToken: cache.RefreshToken,
			Expiry:       cache.AccessTokenExpiry,
		}).Token()
		if err != nil {
			// If we fail to refresh, fall back to browser login.
			log.Printf("Failed to refresh token, falling back to browser login: %s", err.Error())
			return s.doBrowserLogin(ctx, conf)
		}

		// Update cache with newly obtained token. Use the new token details,
		// but keep the same session expiry.
		newCache := Cache{
			AccessToken:       refreshed.AccessToken,
			RefreshToken:      refreshed.RefreshToken,
			AccessTokenExpiry: refreshed.Expiry,
			SessionExpiry:     cache.SessionExpiry,
		}

		err = Write(newCache)
		if err != nil {
			log.Printf("Failed to write cache to file: %s", err.Error())
		}

		return refreshed, nil

	// Otherwise return existing, unexpired token to continue existing session.
	default:
		return &oauth2.Token{
			AccessToken:  cache.AccessToken,
			RefreshToken: cache.RefreshToken,
			Expiry:       cache.AccessTokenExpiry,
		}, nil
	}
}

func (s *UserSession) doBrowserLogin(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	// This is a configuration option set by WithoutBrowserLogin to provide
	// control over when or if the browser is opened. When the flag is true and
	// no valid credentials are found on the system via ClientID:ClientSecret
	// pairing or a previous unexpired browser login, the client returns an
	// error instead of automatically opening a browser
	if s.NoBrowserLogin {
		return nil, ErrorNoLocalCredsFound
	}

	// Login with browser.
	log.Print("No credentials found, proceeding with browser login.")

	if s.browser == nil {
		s.browser = &oauthBrowser{}
	}

	tok, err := s.browser.GetTokenFromBrowser(ctx, conf)
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

	return tok, nil
}
