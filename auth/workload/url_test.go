package workload

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLCredentialSource_Validate(t *testing.T) {
	tests := []struct {
		name   string
		uc     *URLCredentialSource
		errStr string
	}{
		{
			name:   "empty cred source",
			uc:     &URLCredentialSource{},
			errStr: "non-empty URL is required",
		},
		{
			name: "valid url",
			uc: &URLCredentialSource{
				URL:              "https://localhost:4200/token",
				CredentialFormat: CredentialFormat{},
			},
			errStr: "",
		},
		{
			name: "invalid url",
			uc: &URLCredentialSource{
				URL:              "badUrl/123",
				CredentialFormat: CredentialFormat{},
			},
			errStr: "",
		},
		{
			name: "invalid formatter",
			uc: &URLCredentialSource{
				URL: "/folder/token.json",
				CredentialFormat: CredentialFormat{
					Type: "bad",
				},
			},
			errStr: "format type must either be",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.uc.Validate()
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestURLCredentialSource_token(t *testing.T) {
	tests := []struct {
		name     string
		uc       *URLCredentialSource
		respBody string
		want     string
		errStr   string
	}{
		{
			name:     "basic",
			uc:       &URLCredentialSource{},
			respBody: "token",
			want:     "token",
			errStr:   "",
		},
		{
			name: "not found",
			uc: &URLCredentialSource{
				URL: "http://unknown-url-path.localhost.com",
			},
			respBody: "token",
			want:     "",
			errStr:   "invalid response retrieving token",
		},
		{
			name:     "empty response",
			uc:       &URLCredentialSource{},
			respBody: "",
			want:     "",
			errStr:   "response body is empty",
		},
		{
			name: "json response",
			uc: &URLCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			respBody: `{"access_token": "token"}`,
			want:     "token",
			errStr:   "",
		},
		{
			name: "bad json",
			uc: &URLCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			respBody: `"access_token": "token"}`,
			want:     "",
			errStr:   "failed to unmarshal json value",
		},
		{
			name: "missing json key",
			uc: &URLCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/missing",
				},
			},
			respBody: `{"access_token": "token"}`,
			want:     "",
			errStr:   "Object has no key 'missing'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			// Create an HTTP test server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte(tt.respBody))
				require.NoError(err)
			}))
			defer ts.Close()

			// Configure the credential source to use the test server url
			if tt.uc.URL == "" {
				tt.uc.URL = ts.URL
			}

			got, err := tt.uc.token()
			if tt.errStr == "" {
				require.NoError(err)
				require.Equal(tt.want, got)
			} else {
				require.ErrorContains(err, tt.errStr)
			}
		})
	}
}
