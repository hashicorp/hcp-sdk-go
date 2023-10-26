package tokencache

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"time"
)

// sourceType identities the type of token source.
type sourceType = string

// cachingTokenSource acts as a read-through cache for token information received from token sources and oauth configurations.
type cachingTokenSource struct {
	cacheFile        string
	login            bool
	forceLogin       bool
	sourceType       sourceType
	sourceIdentifier string
	oauthTokenSource oauth2.TokenSource
	oauthConfig      oAuth2Config
	fetchTimeout     time.Duration
}

// Token implements the oauth2.TokenSource interface. It will read cached tokens from a file and based on their validity
// return, refresh or replace them.
func (source *cachingTokenSource) Token() (*oauth2.Token, error) {
	// Read the cache information from the file, if it exists
	cachedTokens, err := readCache(source.cacheFile)
	if err != nil {
		log.Println(err)
	}

	// Garbage collect expired tokens
	cachedTokens.removeExpiredTokens()

	// Handle login tokens
	if source.login {
		return source.loginToken(cachedTokens)
	}

	// Handle non-login tokens
	switch source.sourceType {
	case sourceTypeServicePrincipal:
		return source.servicePrincipalToken(cachedTokens)
	case sourceTypeWorkload:
		return source.workloadToken(cachedTokens)
	default:
		return nil, fmt.Errorf("invalid source type: %q", source.sourceType)
	}
}

func (source *cachingTokenSource) evaluate(hitEntry *cacheEntry) (*oauth2.Token, error) {
	var token *oauth2.Token
	if hitEntry != nil {
		token = hitEntry.token()
	}

	// Return the access token if it is still valid
	if token != nil && token.Expiry.After(time.Now()) {
		return token, nil
	}

	if source.oauthConfig != nil {
		// Try to refresh the token if an oauth2.Config was provided
		ctx, cancel := context.WithTimeout(context.Background(), source.fetchTimeout)
		defer cancel()

		token, err := source.oauthConfig.TokenSource(ctx, token).Token()
		if err == nil {
			return token, err
		}

		// Try to fetch the token if an error occurred
		log.Printf("failed to refresh the token: %s\n", err)
	}

	if source.oauthTokenSource != nil {
		// Try to get a new token
		return source.oauthTokenSource.Token()
	}

	return nil, fmt.Errorf("no valid credential source available")
}
