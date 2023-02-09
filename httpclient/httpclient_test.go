// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	consul "github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/client/consul_service"
	"github.com/hashicorp/hcp-sdk-go/config"
	"github.com/hashicorp/hcp-sdk-go/profile"
)

func TestNew(t *testing.T) {

	tokenPath := "/oauth2/token"
	token := "90d64460d14870c08c81352a05dedd3465940a7c"
	clientID := "CLIENT_ID"
	clientSecret := "CLIENT_SECRET"

	orgID := "test-org"
	projID := "test-project"
	apiPath := fmt.Sprintf("/consul/2021-02-04/organizations/%s/projects/%s/clusters", orgID, projID)

	numRequests := uint32(0)

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		atomic.AddUint32(&numRequests, 1)

		// The incoming test request.
		switch r.URL.String() {

		// This handler mocks responses from the auth provider.
		case tokenPath:
			// Before returning the access token, verify request payload includes expected params.
			body, err := io.ReadAll(r.Body)
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

	ts.EnableHTTP2 = true
	ts.StartTLS()
	defer ts.Close()

	t.Run("legacy configuration", func(t *testing.T) {
		numRequests = 0

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

		// This SDK request hits the mock Consul List API handler above ('/consul/2021-02-04/organizations...').
		// The handler verifies that the expected bearer token is provided.
		_, err = consulClient.List(listParams, nil)
		require.NoError(t, err)

		// Make sure we actually handled both the auth and the GET request and didn't
		// just skip all the assertions!
		require.Equal(t, uint32(2), atomic.LoadUint32(&numRequests))
	})

	t.Run("HCPConfig", func(t *testing.T) {
		numRequests = 0

		serverURL, err := url.Parse(ts.URL)
		require.NoError(t, err)

		// Compile a custom TLS config that trusts the test server's certificate
		certPool := x509.NewCertPool()
		certPool.AddCert(ts.Certificate())

		tlsConfig := &tls.Config{
			RootCAs: certPool,
		}

		// Create a config that calls the test server
		hcpConfig, err := config.NewHCPConfig(
			config.WithClientCredentials(clientID, clientSecret),
			config.WithAuth(ts.URL, tlsConfig),
			config.WithAPI(serverURL.Host, tlsConfig),
			config.WithProfile(&profile.UserProfile{OrganizationID: orgID, ProjectID: projID}),
		)
		require.NoError(t, err)

		cfg := Config{
			HCPConfig: hcpConfig,
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

		// This SDK request hits the mock Consul List API handler above ('/consul/2021-02-04/organizations...').
		// The handler verifies that the expected bearer token is provided.
		_, err = consulClient.List(listParams, nil)
		require.NoError(t, err)

		// Make sure we actually handled both the auth and the GET request and didn't
		// just skip all the assertions!
		require.Equal(t, uint32(2), atomic.LoadUint32(&numRequests))
	})

}

func TestMiddleware(t *testing.T) {

	// Start with a plain request.
	request, err := http.NewRequest("GET", "api.cloud.hashicorp.com/consul/2021-02-04/organizations//projects//clusters", httptest.NewRecorder().Body)
	require.NoError(t, err)

	// Prepare header is unset.
	require.Equal(t, request.Header.Get("X-HCP-Source-Channel"), "")

	// Prepare middleware function.
	expectedSourceChannel := "source_channel_foo"
	sourceChannelMiddleware := withSourceChannel(expectedSourceChannel)

	// Apply middleware function.
	err = sourceChannelMiddleware(request)
	require.NoError(t, err)

	// Assert request is modified as expected.
	assert.Equal(t, request.Header.Get("X-HCP-Source-Channel"), expectedSourceChannel)

	// Assert path is unmodified.
	expectedOrgID := "org_id_77"
	expectedProjID := "proj_id_123"
	assert.NotContains(t, request.URL.Path, expectedOrgID)
	assert.NotContains(t, request.URL.Path, expectedProjID)

	// Prepare middleware function.
	profileMiddleware := withOrgAndProjectIDs(expectedOrgID, expectedProjID)

	// Apply middleware function.
	err = profileMiddleware(request)
	require.NoError(t, err)

	// Assert request is modified as expected.
	assert.Contains(t, request.URL.Path, expectedOrgID)
	assert.Contains(t, request.URL.Path, expectedProjID)
	assert.Equal(t, request.Header.Get("X-HCP-Source-Channel"), expectedSourceChannel)
}
