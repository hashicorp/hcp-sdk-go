package auth

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	tokenPath string = "/oauth/token"
)

// WithClientCredentials returns an http client with an access token obtained using the configured client credentials.
func WithClientCredentials(ctx context.Context, clientID, clientSecret, authURL string) (*http.Client, error) {
	// The audience is the API identifier configured in the auth provider
	// and must be provided when requesting an access token for the API.
	aud := "https://api.hashicorp.cloud"

	conf := &clientcredentials.Config{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		TokenURL:       authURL + tokenPath,
		EndpointParams: url.Values{"audience": {aud}},
	}

	// The http client is the same one attached to the context,
	// only now it will be able to authenticate with the token obtained using the client credentials.
	client := conf.Client(ctx)

	return client, nil
}
