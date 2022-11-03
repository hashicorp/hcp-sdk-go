package auth

import (
	"context"
	"testing"
	"time"

	assertpkg "github.com/stretchr/testify/assert"
	requirepkg "github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestGetToken_ExistingToken(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	ctx := context.Background()
	conf := &oauth2.Config{}

	now := time.Now()
	cache := Cache{
		AccessToken:       "TestAccessToken",
		RefreshToken:      "TestRefreshToken",
		AccessTokenExpiry: now,
		SessionExpiry:     time.Now().Add(time.Hour * 24),
	}

	err := Write(cache)
	require.NoError(err)

	userSession := UserSession{}

	tok, err := userSession.GetToken(ctx, conf)
	require.NoError(err)

	assert.Equal(cache.AccessToken, tok.AccessToken)
	assert.Equal(cache.RefreshToken, tok.RefreshToken)
	assert.Equal(cache.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"), tok.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))
}
