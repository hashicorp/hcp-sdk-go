// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"fmt"
	"log"

	"golang.org/x/oauth2"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
)

const sourceTypeWorkload = sourceType("workload")

// NewWorkloadTokenSource will create a token source that caches workload identity provider tokens. The tokens will be
// cached under `workloads` in the cache file.
func NewWorkloadTokenSource(
	cacheFile string,
	providerResourceName string,
	oauthTokenSource oauth2.TokenSource,
	geo string,
) (oauth2.TokenSource, error) {
	if !geography.ValidateGeo(geography.Geo(geo)) {
		return nil, fmt.Errorf("login geography %s invalid. Supported: %v", geo, geography.Geographies)
	}
	return &cachingTokenSource{
		cacheFile:        cacheFile,
		sourceIdentifier: providerResourceName,
		sourceType:       sourceTypeWorkload,
		oauthTokenSource: oauthTokenSource,
		geography:        geo,
	}, nil
}

func (source *cachingTokenSource) workloadToken(cachedTokens *cache) (*oauth2.Token, error) {
	// Check if there is the cache has an entry for the workload
	var hitEntry *cacheEntry
	if entry, ok := cachedTokens.Workloads[source.sourceIdentifier]; ok {
		hitEntry = &entry
	}

	// Check the token for validity, try to refresh it and otherwise get a new token
	token, err := source.getValidToken(hitEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to get new token: %w", err)
	}

	// Cache the new token
	cachedTokens.Workloads[source.sourceIdentifier] = *cacheEntryFromToken(token, source.geography)

	// Write the cache back to the file
	if err = cachedTokens.write(source.cacheFile); err != nil {
		log.Printf("failed to write credentials to cache: %s\n", err)
	}

	return token, nil
}
