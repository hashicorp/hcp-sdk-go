// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"fmt"
	"log"

	"golang.org/x/oauth2"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
)

const sourceTypeServicePrincipal = sourceType("service-principal")

// NewServicePrincipalTokenSource will create a token source that caches service principal tokens. The tokens will be
// cached under `service-principals` in the cache file.
func NewServicePrincipalTokenSource(
	cacheFile string,
	clientID string,
	oauthTokenSource oauth2.TokenSource,
	geo string,
) (oauth2.TokenSource, error) {
	if !geography.ValidateGeo(geography.Geo(geo)) {
		return nil, fmt.Errorf("login geography %s invalid. Supported: %v", geo, geography.Geographies)
	}
	return &cachingTokenSource{
		cacheFile:        cacheFile,
		sourceIdentifier: clientID,
		sourceType:       sourceTypeServicePrincipal,
		oauthTokenSource: oauthTokenSource,
		geography:        geo,
	}, nil
}

func (source *cachingTokenSource) servicePrincipalToken(cachedTokens *cache) (*oauth2.Token, error) {
	// Check if there is the cache has an entry for the service principal
	var hitEntry *cacheEntry
	if entry, ok := cachedTokens.ServicePrincipals[source.sourceIdentifier]; ok {
		hitEntry = &entry
	}

	// Check the token for validity, try to refresh it and otherwise get a new token
	token, err := source.getValidToken(hitEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to get new token: %w", err)
	}

	// Cache the new token
	cachedTokens.ServicePrincipals[source.sourceIdentifier] = *cacheEntryFromToken(token, source.geography)

	// Write the cache back to the file
	if err = cachedTokens.write(source.cacheFile); err != nil {
		log.Printf("failed to write credentials to cache: %s\n", err)
	}

	return token, nil
}
