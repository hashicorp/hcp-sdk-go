package workload

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProvider_New(t *testing.T) {
	tests := []struct {
		name   string
		ipc    *IdentityProviderConfig
		errStr string
	}{
		{
			name:   "nil config",
			ipc:    nil,
			errStr: "workload identity provider config must not be nil",
		},
		{
			name:   "empty config",
			ipc:    &IdentityProviderConfig{},
			errStr: "workload identity provider resource name must be set",
		},
		{
			name:   "no credential source",
			ipc:    &IdentityProviderConfig{ProviderResourceName: "iam/test"},
			errStr: "a credential source must be configured",
		},
		{
			name: "more than one credential source",
			ipc: &IdentityProviderConfig{
				ProviderResourceName: "iam/test",
				File: &FileCredentialSource{
					Path: "test",
				},
				AWS: &AWSCredentialSource{
					IMDSv2: false,
				},
			},
			errStr: "only one credential source may be configured",
		},
		{
			name: "calls validate on cred source",
			ipc: &IdentityProviderConfig{
				ProviderResourceName: "iam/test",
				File: &FileCredentialSource{
					Path: "",
				},
			},
			errStr: "path must be set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.ipc)
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestProvider_Token_JWT(t *testing.T) {
	require := require.New(t)

	// Set the credential in an env var
	envVar := "test"
	cred := "example_cred"
	t.Setenv(envVar, cred)

	// Create the provider
	pname := "iam/project/123/service-principal/bar/workload-identity-provider/baz"
	c := &IdentityProviderConfig{
		ProviderResourceName: pname,
		EnvironmentVariable: &EnvironmentVariableCredentialSource{
			Var: envVar,
		},
	}
	p, err := New(c)
	require.NoError(err)

	// Create an HTTP test server
	hcpToken := "hcp_ftw"
	hcpTokenExpiration := 42
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(fmt.Sprintf("/2019-12-10/%s/exchange-token", pname), r.URL.Path)
		var req exchangeTokenRequest
		require.NoError(json.NewDecoder(r.Body).Decode(&req))
		require.Equal(cred, req.JwtToken)
		require.Nil(req.AwsGetCallerIDToken)

		resp := &exchangeTokenResponse{
			AccessToken: hcpToken,
			ExpiresIn:   hcpTokenExpiration,
		}
		require.NoError(json.NewEncoder(w).Encode(resp))
	}))
	defer ts.Close()

	// Create the API Config
	hcpAPI := &hcpAPI{
		apiAddress: ts.Listener.Addr().String(),
	}
	transport, ok := ts.Client().Transport.(*http.Transport)
	require.True(ok)
	hcpAPI.tlsConfig = transport.TLSClientConfig.Clone()
	p.SetAPI(hcpAPI)

	// Call token exchange
	token, err := p.Token()
	require.NoError(err)
	require.Equal(hcpToken, token.AccessToken)
	require.WithinDuration(time.Now().Add(time.Duration(hcpTokenExpiration)*time.Second), token.Expiry, 100*time.Millisecond)
}

func TestProvider_Token_AWS(t *testing.T) {
	require := require.New(t)

	// Create the provider
	pname := "iam/project/123/service-principal/bar/workload-identity-provider/baz"
	c := &IdentityProviderConfig{
		ProviderResourceName: pname,
		AWS:                  &AWSCredentialSource{},
	}
	p, err := New(c)
	require.NoError(err)

	// Swap the actual AWS credential source for a test one
	id := &callerIdentityRequest{
		Headers: map[string]string{"hello": "world!"},
		Method:  "PUT",
		URL:     "https://sts.us-east1.awsapi.com/",
	}
	p.awsProvider = &mockAWS{id: id}

	// Create an HTTP test server
	hcpToken := "hcp_ftw"
	hcpTokenExpiration := 42
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(fmt.Sprintf("/2019-12-10/%s/exchange-token", pname), r.URL.Path)
		var req exchangeTokenRequest
		require.NoError(json.NewDecoder(r.Body).Decode(&req))
		require.NotNil(req.AwsGetCallerIDToken)
		require.EqualValues(id.Headers, req.AwsGetCallerIDToken.Headers)
		require.Equal(id.Method, req.AwsGetCallerIDToken.Method)
		require.Equal(id.URL, req.AwsGetCallerIDToken.URL)
		resp := &exchangeTokenResponse{
			AccessToken: hcpToken,
			ExpiresIn:   hcpTokenExpiration,
		}
		require.NoError(json.NewEncoder(w).Encode(resp))
	}))
	defer ts.Close()

	// Create the API Config
	hcpAPI := &hcpAPI{
		apiAddress: ts.Listener.Addr().String(),
	}
	transport, ok := ts.Client().Transport.(*http.Transport)
	require.True(ok)
	hcpAPI.tlsConfig = transport.TLSClientConfig.Clone()
	p.SetAPI(hcpAPI)

	// Call token exchange
	token, err := p.Token()
	require.NoError(err)
	require.Equal(hcpToken, token.AccessToken)
	require.WithinDuration(time.Now().Add(time.Duration(hcpTokenExpiration)*time.Second), token.Expiry, 100*time.Millisecond)
}

type mockAWS struct {
	id *callerIdentityRequest
}

func (aws *mockAWS) getCallerIdentityReq(_ string) (*callerIdentityRequest, error) {
	return aws.id, nil
}

type hcpAPI struct {
	// apiAddress is the test HCP API address (<hostname>[:port]).
	apiAddress string

	// tlsConfig is the test tls config
	tlsConfig *tls.Config
}

func (a *hcpAPI) APIAddress() string        { return a.apiAddress }
func (a *hcpAPI) APITLSConfig() *tls.Config { return a.tlsConfig }
