// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

type MockSession struct{}

// GetToken returns some mock token with static data.
func (s *MockSession) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	tok := oauth2.Token{
		AccessToken:  "Some.Access.Token",
		RefreshToken: "SomeRefreshToken",
		Expiry:       time.Now().Add(time.Hour * 1),
	}
	return &tok, nil
}

// mockBrowser provides a mocked response of the OAuth2 login flow.
type mockBrowser struct{}

func (b *mockBrowser) GetTokenFromBrowser(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	tok := oauth2.Token{
		// Cache read expects the access token to look JWT-like
		AccessToken:  "New.Access.Token",
		RefreshToken: "NewRefreshToken",
		Expiry:       time.Now().Add(time.Hour * 1),
	}
	return &tok, nil
}

type mockRefreshTokenSource struct {
	err bool
}

func (m *mockRefreshTokenSource) Token() (*oauth2.Token, error) {
	if m.err {
		return nil, fmt.Errorf("error")
	}

	return &oauth2.Token{
		// Cache read expects the access token to look JWT-like
		AccessToken:  "Refreshed.Access.Token",
		RefreshToken: "RefreshedRefreshToken",
		Expiry:       time.Now().Add(time.Hour * 1),
	}, nil
}
