package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mitchellh/colorstring"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

// callbackEndpoint exposes the confiugration for the callback server.
type callbackEndpoint struct {
	server         *http.Server
	code           string
	shutdownSignal chan error
}

type BrowserGetter struct {
}

// callbackEndpoint endpoint ServeHTTP confirms if an Authorization code was received from Auth0.
func (h *callbackEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code != "" {
		h.code = code
		fmt.Fprintln(w, "Login is successful. You may close the browser and return to the command line.")
		colorstring.Println("[bold][green]Success!")
	} else {
		fmt.Fprintln(w, "Login is not successful. You may close the browser and try again.")
	}
	h.shutdownSignal <- nil
}

// GetToken returns an access token obtained from either an existing session or new browser login.
func (g *BrowserGetter) GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {

	// Check for existing token in auth cache, if it exists.
	cache, readErr := Read()
	if readErr != nil {
		log.Printf("Failed to read cache from file: %s", readErr.Error())
	}

	var tok *oauth2.Token
	var err error

	// Check the session expiry of the retrieved token.
	// If session expiry has passed, then reauthenticate with browser login and reassign token.
	if readErr != nil || cache.SessionExpiry.Before(time.Now()) {

		// Login with browser.
		log.Print("No credentials found, proceeding with browser login.")

		tok, err = getTokenFromBrowser(ctx, conf)
		if err != nil {
			return nil, fmt.Errorf("failed to get access token: %w", err)
		}

		// Update cache with newly obtained token.
		newCache := Cache{
			AccessToken:       tok.AccessToken,
			RefreshToken:      tok.RefreshToken,
			AccessTokenExpiry: tok.Expiry,
			SessionExpiry:     time.Now().Add(SessionMaxAge),
		}

		err = Write(newCache)
		if err != nil {
			log.Printf("Failed to write cache to file: %s", err.Error())
		}

	} else { // Otherwise return existing, unexpired token to continue existing session.
		tok = &oauth2.Token{
			AccessToken:  cache.AccessToken,
			RefreshToken: cache.RefreshToken,
			Expiry:       cache.AccessTokenExpiry,
		}
	}

	return tok, nil
}

// getTokenFromBrwoser opens a browser window for the user to log in and handles the OAuth2 flow to obtain a token.
func getTokenFromBrowser(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	// Launch a request to Auth0's authorization endpoint.
	colorstring.Printf("[bold][yellow]The default web browser has been opened at %s. Please continue the login in the web browser.\n", conf.Endpoint.AuthURL)

	// Prepare the /authorize request with randomly generated state, offline access option, and audience
	aud := "https://api.hashicorp.cloud"
	opt := oauth2.SetAuthURLParam("audience", aud)
	authzURL := conf.AuthCodeURL(generateRandomString(32), oauth2.AccessTypeOffline, opt)

	// Handle ctrl-c while waiting for the callback
	sigintCh := make(chan os.Signal, 1)
	signal.Notify(sigintCh, os.Interrupt)
	defer signal.Stop(sigintCh)

	if err := open.Start(authzURL); err != nil {
		return nil, fmt.Errorf("failed to open browser at URL %q: %w", authzURL, err)
	}

	// Start callback server
	callbackEndpoint := &callbackEndpoint{}
	callbackEndpoint.shutdownSignal = make(chan error)
	server := &http.Server{
		Addr:           ":8443",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	callbackEndpoint.server = server
	http.Handle("/oidc/callback", callbackEndpoint)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			callbackEndpoint.shutdownSignal <- fmt.Errorf("failed to start callback server: %w", err)
		}
	}()

	// Wait for either the callback to finish, SIGINT to be received or up to 2 minutes
	select {
	case err := <-callbackEndpoint.shutdownSignal:

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

		// Save the token to config file.
		cache := Cache{
			AccessToken:       tok.AccessToken,
			RefreshToken:      tok.RefreshToken,
			AccessTokenExpiry: tok.Expiry,
			SessionExpiry:     time.Now().Add(SessionMaxAge),
		}
		err = Write(cache)
		if err != nil {
			return nil, fmt.Errorf("failed to write token to file: %w", err)
		}

		return tok, nil
	case <-sigintCh:
		return nil, errors.New("interrupted")
	case <-time.After(2 * time.Minute):
		return nil, errors.New("timed out waiting for response from provider")
	}
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
