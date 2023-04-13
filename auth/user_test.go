// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

	_, _, err := CacheSetup(t)
	require.NoError(err)
	ctx := context.Background()
	conf := &oauth2.Config{}

	now := time.Now()
	cache := Cache{
		AccessToken:       "Test.Access.Token",
		RefreshToken:      "TestRefreshToken",
		AccessTokenExpiry: now.Add(time.Hour * 2),
		SessionExpiry:     time.Now().Add(time.Hour * 24),
	}

	err = Write(cache)
	require.NoError(err)

	userSession := UserSession{}

	tok, err := userSession.GetToken(ctx, conf)
	require.NoError(err)

	assert.Equal(cache.AccessToken, tok.AccessToken)
	assert.Equal(cache.RefreshToken, tok.RefreshToken)
	assert.Equal(cache.AccessTokenExpiry.Format("2006-01-02T15:04:05 -07:00:00"), tok.Expiry.Format("2006-01-02T15:04:05 -07:00:00"))

	err = CacheDestroy(t)
	require.NoError(err)
}

func TestGetToken_BrowserFlow(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name      string
		caseSetup func() error
	}{
		{
			name:      "No Token",
			caseSetup: func() error { return nil },
		},
		{
			name: "Expired Token",
			caseSetup: func() error {
				now := time.Now()
				cache := Cache{
					AccessToken:       "ExpiredAccessToken",
					RefreshToken:      "ExpiredRefreshToken",
					AccessTokenExpiry: now,
					SessionExpiry:     time.Now().Add(time.Hour * -2),
				}

				err := Write(cache)
				require.NoError(err)

				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			_, _, err := CacheSetup(t)
			require.NoError(err)

			// Runs specific test case setup.
			err = testCase.caseSetup()
			require.NoError(err)

			// Run the test.
			userSession := UserSession{}
			userSession.browser = &mockBrowser{}

			ctx := context.Background()
			conf := oauth2.Config{}

			tok, err := userSession.GetToken(ctx, &conf)
			require.NoError(err)
			require.NotNil(tok)

			// Make assertions about token retrieved.
			assert.Equal("SomeNewAccessToken", tok.AccessToken)
			assert.Equal("SomeNewRefreshToken", tok.RefreshToken)
			assert.WithinDuration(time.Now().Add(time.Hour*1), tok.Expiry, 10*time.Second)

			// Cleanup.
			require.NoError(CacheDestroy(t))
		})
	}
}
