// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
	"golang.org/x/oauth2"
)

const sourceTypeLogin = sourceType("login")

// NewLoginTokenSource will create a token source that caches login tokens. Only one login token will be cached at a
// time.
//
// The tokens will be cached under `login` in the cache file.
func NewLoginTokenSource(
	cacheFile string,
	oauthTokenSource oauth2.TokenSource,
	oauthConfig oAuth2Config,
	geo string,
) (oauth2.TokenSource, error) {
	if !geography.ValidateGeo(geography.Geo(geo)) {
		return nil, fmt.Errorf("login geography %s invalid. Supported: %v", geo, geography.Geographies)
	}

	return &cachingTokenSource{
		cacheFile:        cacheFile,
		sourceType:       sourceTypeLogin,
		oauthTokenSource: oauthTokenSource,
		oauthConfig:      oauthConfig,
		geography:        geo,
	}, nil
}

func (source *cachingTokenSource) loginToken(cachedTokens *cache) (*oauth2.Token, error) {
	var hitEntry *cacheEntry

	// First check if credentials exist in the `login` field
	if cachedTokens.Login != nil {
		hitEntry = cachedTokens.Login
	}

	// Check the token for validity, try to refresh it and otherwise get a new token
	token, err := source.getValidToken(hitEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to get new token: %w", err)
	}

	// For backwards compatibility, if geography is not set in the cache entry,
	// we force an update to store the geography in the cached file
	if hitEntry != nil && hitEntry.AccessToken == token.AccessToken &&
		matchGeography(hitEntry.Geography, source.geography) {
		// The cached entry was used,  the cache does not need to be updated
		return token, nil
	}

	// Cache the new token
	cachedTokens.Login = cacheEntryFromToken(token, source.geography)

	// Write the cache back to the file
	if err = cachedTokens.write(source.cacheFile); err != nil {
		log.Printf("failed to write credentials to cache: %s\n", err)
	}

	return token, nil
}
