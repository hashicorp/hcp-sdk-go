package workload

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironmentVariableCredentialSource_Validate(t *testing.T) {
	tests := []struct {
		name   string
		ec     *EnvironmentVariableCredentialSource
		errStr string
	}{
		{
			name:   "empty cred source",
			ec:     &EnvironmentVariableCredentialSource{},
			errStr: "environment variable must be specified",
		},
		{
			name: "valid cred source",
			ec: &EnvironmentVariableCredentialSource{
				Var:              "env-var",
				CredentialFormat: CredentialFormat{},
			},
			errStr: "",
		},
		{
			name: "invalid formatter",
			ec: &EnvironmentVariableCredentialSource{
				Var: "env-var",
				CredentialFormat: CredentialFormat{
					Type: "bad",
				},
			},
			errStr: "format type must either be",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ec.Validate()
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestEnvironmentVariableCredentialSource_token(t *testing.T) {
	tests := []struct {
		name        string
		ec          *EnvironmentVariableCredentialSource
		envVarName  string
		envVarValue string
		want        string
		errStr      string
	}{
		{
			name: "basic",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
			},
			envVarName:  "MY_VAR",
			envVarValue: "token",
			want:        "token",
			errStr:      "",
		},
		{
			name: "not found",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
			},
			envVarName:  "MY_OTHER_VAR",
			envVarValue: "token",
			want:        "",
			errStr:      "environment variable not found",
		},
		{
			name: "empty env var",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
			},
			envVarName:  "MY_VAR",
			envVarValue: "",
			want:        "",
			errStr:      "environment variable value is empty",
		},
		{
			name: "json",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			envVarName:  "MY_VAR",
			envVarValue: `{"access_token": "token"}`,
			want:        "token",
			errStr:      "",
		},
		{
			name: "bad json",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			envVarName:  "MY_VAR",
			envVarValue: `"access_token": "token"}`,
			want:        "",
			errStr:      "failed to unmarshal json value",
		},
		{
			name: "missing json key",
			ec: &EnvironmentVariableCredentialSource{
				Var: "MY_VAR",
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/missing",
				},
			},
			envVarName:  "MY_VAR",
			envVarValue: `{"access_token": "token"}`,
			want:        "",
			errStr:      "Object has no key 'missing'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the env var
			t.Setenv(tt.envVarName, tt.envVarValue)
			got, err := tt.ec.token()
			if tt.errStr == "" {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}
