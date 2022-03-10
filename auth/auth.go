package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mitchellh/colorstring"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	callbackURL string = "http://127.0.0.1:8888"
	authzPath   string = "/authorize"
	tokenPath   string = "/oauth/token"
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

// WithBrowserLogin returns an http client with access and refresh tokens obtained via a user browser login.
func WithBrowserLogin(ctx context.Context, authURL, auth0ClientID string) (client *http.Client, err error) {
	conf := &oauth2.Config{
		ClientID: auth0ClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL + authzPath,
			TokenURL: authURL + tokenPath,
		},
		RedirectURL: callbackURL,
		Scopes:      []string{"openid", "offline_access"},
	}

	// On first-time login, the user must log in through the browser and receive an ID token and refresh token.
	var tok *oauth2.Token
	tok, err = getTokenFromBrowser(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	// The http client is the same one attached to the context,
	// only now it will be able to authenticate with the token obtained from a successful browser login.
	client = conf.Client(ctx, tok)

	return client, nil
}

// callbackEndpoint exposes the confiugration for the callback server.
type callbackEndpoint struct {
	server         *http.Server
	code           string
	shutdownSignal chan error
}

// callbackEndpoint endpoint ServeHTTP confirms if an Authorization code was received from Auth0.
func (h *callbackEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code != "" {
		h.code = code
		fmt.Fprintln(w, "Login is successful. You may close the browser and return to the command line")
		colorstring.Println("[bold][yellow]Success!")
	} else {
		fmt.Fprintln(w, "Login is not successful. You may close the browser and try again")
	}
	h.shutdownSignal <- nil
}

// getTokenFromBrowser opens a browser window for the user to log in and handles the OAuth2 flow to obtain a token.
func getTokenFromBrowser(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {

	// Launch a request to Auth0's authorization endpoint.
	colorstring.Printf("[bold][yellow]The default web browser has been opened at %s. Please continue the login in the web browser.", conf.Endpoint.AuthURL)

	// Prepare the /authorize request with randomly generated state and offline access option
	authzURL := conf.AuthCodeURL(generateRandomString(32), oauth2.AccessTypeOffline)

	if err := open.Start(authzURL); err != nil {
		return nil, fmt.Errorf("failed to open browser at URL %q: %w", authzURL, err)
	}

	// Start callback server
	callbackEndpoint := &callbackEndpoint{}
	callbackEndpoint.shutdownSignal = make(chan error)
	server := &http.Server{
		Addr:           ":8888",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	callbackEndpoint.server = server
	http.Handle("/", callbackEndpoint)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			callbackEndpoint.shutdownSignal <- fmt.Errorf("failed to start callback server: %w", err)
		}
	}()
	err := <-callbackEndpoint.shutdownSignal
	if err != nil {
		return nil, err
	}

	err = callbackEndpoint.server.Shutdown(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to shutdown callback server: %w", err)
	}

	// Exchange the code returned in the callback for a token.
	tok, err := conf.Exchange(ctx, callbackEndpoint.code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	return tok, nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func generateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(b)
}
