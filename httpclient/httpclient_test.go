package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	consul "github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/client/consul_service"
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

	numRequests := uint32(0)

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		atomic.AddUint32(&numRequests, 1)

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

	cfg := Config{
		HostPath:     ts.URL,
		AuthURL:      ts.URL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// Override the base http.Client so that we trust the test server's TLS
		Client: ts.Client(),
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
	_, err = consulClient.List(listParams, nil)
	require.NoError(t, err)

	// Make sure we actually handled both the auth and the GET request and didn't
	// just skip all the assertions!
	require.Equal(t, uint32(2), atomic.LoadUint32(&numRequests))
}
