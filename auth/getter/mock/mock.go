package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type MockGetter struct {
}

func (g *MockGetter) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {}
