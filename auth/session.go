package auth

import (
	"context"

	"golang.org/x/oauth2"
)

// Session handles managing a user's session.
type Session interface {
	GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error)
}
