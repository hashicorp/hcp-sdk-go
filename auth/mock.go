package auth

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

type MockGetter struct{}

// GetToken returns some mock token wiht static data.
func (g *MockGetter) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	tok := oauth2.Token{
		AccessToken:  "SomeAccessToken",
		RefreshToken: "SomeRefreshToken",
		Expiry:       time.Now(),
	}
	return &tok, nil
}
