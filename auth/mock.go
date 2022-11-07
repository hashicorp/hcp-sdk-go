package auth

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

type MockSession struct{}

// GetToken returns some mock token wiht static data.
func (s *MockSession) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	tok := oauth2.Token{
		AccessToken:  "SomeAccessToken",
		RefreshToken: "SomeRefreshToken",
		Expiry:       time.Now(),
	}
	return &tok, nil
}

type MockBrowser struct {
	mock.Mock
}
