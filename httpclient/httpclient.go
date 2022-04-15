package httpclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/hashicorp/go-cleanhttp"

	"github.com/hashicorp/hcp-sdk-go/version"
)

// Config contains the client's configuration options
type Config struct {
	// HostPath is the host name of the HCP API, without a scheme prefix.
	HostPath string `json:"host_path"`

	// AuthURL is the URL of the authentication provider, inclusive of the 'https' scheme prefix.
	AuthURL string `json:"auth_url"`

	// ClientID is the client ID of a Service Principal Key.
	ClientID string `json:"client_id"`

	// ClientSecret is the client secret of a Service Principal Key, only provided on creation.
	ClientSecret string `json:"client_secret"`

	// SourceChannel denotes the client (channel) that originiated the request.
	// This is synonymous to a user-agent.
	SourceChannel string `json:"source_channel"`
}
type roundTripperWithSourceChannel struct {
	OriginalRoundTripper http.RoundTripper
	SourceChannel        string
}

func (rt *roundTripperWithSourceChannel) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-HCP-Source-Channel", rt.SourceChannel)
	return rt.OriginalRoundTripper.RoundTrip(req)
}

// New creates a client with the right base path to connect to any HCP API
func New(cfg Config) (runtime *httptransport.Runtime, err error) {

	// Populate default values where possible.
	cfg.Canonicalize()

	// For local dev only
	// 1. skip client ID/Secret validation
	// if err := cfg.Validate(); err != nil {
	// 	return nil, fmt.Errorf("invalid config: %w", err)
	// }

	// Initialize an http client with a transport to be shared by the service-specific clients.
	client := cleanhttp.DefaultPooledClient()

	// For local dev only
	// 2. skip client ID/Secret authentication
	// // The oauth2 client initializer requires the http client to be passed in via context.
	// ctx := context.WithValue(context.Background(), oauth2.HTTPClient, client)

	// client, err = auth.WithClientCredentials(ctx, cfg.ClientID, cfg.ClientSecret, cfg.AuthURL)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to obtain credentials: %w", err)
	// }

	if cfg.SourceChannel != "" {
		// Using a custom transport in order to set the source channel header when it is present.
		sourceChannel := fmt.Sprintf("%s hcp-go-sdk/%s", cfg.SourceChannel, version.Version)
		client.Transport = &roundTripperWithSourceChannel{OriginalRoundTripper: client.Transport, SourceChannel: sourceChannel}
	}

	// This is the type of runtime that can be used by the generated clients.
	runtime = httptransport.NewWithClient(cfg.HostPath, "", []string{"https"}, client)

	// For local dev only
	// 2. attach bearer token string
	runtime.DefaultAuthentication = httptransport.BearerToken(os.Getenv("HCP_BEARER_TOKEN"))

	return runtime, nil
}

// Canonicalize populates default values for config fields that are otherwise unset.
func (c *Config) Canonicalize() {
	if c.HostPath == "" {
		c.HostPath = os.Getenv("HCP_API_HOST")
		if c.HostPath == "" {
			// prod default
			c.HostPath = "api.cloud.hashicorp.com"
		}
	}
	if c.AuthURL == "" {
		c.AuthURL = os.Getenv("HCP_AUTH_URL")
		if c.AuthURL == "" {
			// prod default
			c.AuthURL = "https://auth.hashicorp.com"
		}
	}
}

// Validate ensures the http client configuration is valid.
func (c *Config) Validate() error {

	if c.ClientID == "" {
		c.ClientID = os.Getenv("HCP_CLIENT_ID")
		if c.ClientID == "" {
			return errors.New("client ID not provided")
		}
	}

	if c.ClientSecret == "" {
		c.ClientSecret = os.Getenv("HCP_CLIENT_SECRET")
		if c.ClientSecret == "" {
			return errors.New("client secret not provided")
		}
	}

	return validation.ValidateStruct(c,
		// The scheme will be appended by the go-openapi client, so it must not be included in the host path.
		validation.Field(&c.HostPath, validation.Required, validation.By(excludesScheme)),
		// For security, the authentication provider's Token URL must use 'https'.
		validation.Field(&c.AuthURL, validation.Required, validation.By(includesHTTPSScheme)),
	)
}

func excludesScheme(value interface{}) error {
	u, _ := value.(string)

	p, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", u, err)
	}

	// This scheme is only allowed for testing.
	if p.Scheme == "localhost" {
		return nil
	}

	if p.Scheme != "" {
		return fmt.Errorf("%q must not include any scheme", u)
	}

	return nil
}

func includesHTTPSScheme(value interface{}) error {
	u, _ := value.(string)

	p, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", u, err)
	}

	if p.Scheme != "https" {
		return fmt.Errorf("%q must start with 'https' scheme", u)
	}

	return nil
}
