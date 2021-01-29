package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	consul "github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/client/consul_service"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	tokenPath := "/oauth/token"
	token := "90d64460d14870c08c81352a05dedd3465940a7c"
	clientID := "CLIENT_ID"
	clientSecret := "CLIENT_SECRET"

	orgID := "test-org"
	projID := "test-project"
	apiPath := fmt.Sprintf("/consul/2020-08-26/organizations/%s/projects/%s/clusters", orgID, projID)

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// The incoming test request.
		switch r.URL.String() {

		// This handler mocks responses from the auth provider.
		case tokenPath:
			// Before returning the access token, verify request payload includes expected params.
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			require.NoError(t, err)

			expected := `audience=https%3A%2F%2Fapi.hashicorp.cloud&grant_type=client_credentials`
			require.Equal(t, expected, string(body))

			// Once params are verified, the test auth server returns an bearer token.
			w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
			w.Write([]byte(fmt.Sprintf("access_token=%s&token_type=bearer", token))) // nolint:errcheck

		// This handler mocks responses from the Consul List API.
		case apiPath:

			// Verify the SDK request provides the access token granted by the mocked auth provider above.
			tok := r.Header.Get("Authorization")
			expected := fmt.Sprintf("Bearer %s", token)
			require.Equal(t, expected, tok, fmt.Sprintf("provided token %q does not match expected token %q", tok, token))

			// Simulate a successful empty response.
			w.Write([]byte{}) // nolint:errcheck
		default:
			t.Errorf("invalid URL %q requested", r.URL)
		}
	}))
	// Enable once we're on Go 1.14
	// ts.EnableHTTP2 = true
	ts.StartTLS()
	defer ts.Close()

	// Extract port from the URL.
	parts := strings.SplitN(ts.URL, ":", 3)
	port := parts[2]

	cfg := Config{
		HostPath:     fmt.Sprintf("localhost:%s", port),
		AuthURL:      fmt.Sprintf("https://localhost:%s", port),
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	// The test's important assertions occur in the test server handler above.
	// When an http client is initialized, it tries to obtain an access token using the configured token URL (from an auth provider), client ID, and client secret.
	// The token request embedded in this client initializer hits the mock auth provider handler above ('/oauth/token').
	cl, err := New(cfg)
	require.NoError(t, err)

	consulClient := consul.New(cl, nil)
	listParams := consul.NewListParams()
	listParams.LocationOrganizationID = orgID
	listParams.LocationProjectID = projID

	// This SDK request hits the mock Consul List API handler above ('/consul/2020-08-26/organizations...').
	// The handler verifies that the expected bearer token is provided.
	consulClient.List(listParams, nil) // nolint:errcheck
}
