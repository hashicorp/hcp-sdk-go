package workload

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileCredentialSource_Validate(t *testing.T) {
	tests := []struct {
		name   string
		fc     *FileCredentialSource
		errStr string
	}{
		{
			name:   "empty cred source",
			fc:     &FileCredentialSource{},
			errStr: "path must be set",
		},
		{
			name: "valid cred source",
			fc: &FileCredentialSource{
				Path:             "~/token.json",
				CredentialFormat: CredentialFormat{},
			},
			errStr: "",
		},
		{
			name: "invalid formatter",
			fc: &FileCredentialSource{
				Path: "/folder/token.json",
				CredentialFormat: CredentialFormat{
					Type: "bad",
				},
			},
			errStr: "format type must either be",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fc.Validate()
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestFileCredentialSource_token(t *testing.T) {
	tests := []struct {
		name        string
		fc          *FileCredentialSource
		fileContent string
		want        string
		errStr      string
	}{
		{
			name:        "basic",
			fc:          &FileCredentialSource{},
			fileContent: "token",
			want:        "token",
			errStr:      "",
		},
		{
			name: "not found",
			fc: &FileCredentialSource{
				Path: "unknown_file",
			},
			fileContent: "token",
			want:        "",
			errStr:      "failed to open credential file",
		},
		{
			name:        "empty file",
			fc:          &FileCredentialSource{},
			fileContent: "",
			want:        "",
			errStr:      "credential file is empty",
		},
		{
			name: "json",
			fc: &FileCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			fileContent: `{"access_token": "token"}`,
			want:        "token",
			errStr:      "",
		},
		{
			name: "bad json",
			fc: &FileCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/access_token",
				},
			},
			fileContent: `"access_token": "token"}`,
			want:        "",
			errStr:      "failed to unmarshal json value",
		},
		{
			name: "missing json key",
			fc: &FileCredentialSource{
				CredentialFormat: CredentialFormat{
					Type:                     FormatTypeJSON,
					SubjectCredentialPointer: "/missing",
				},
			},
			fileContent: `{"access_token": "token"}`,
			want:        "",
			errStr:      "Object has no key 'missing'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			// Create a temp file with the given value
			f, err := ioutil.TempFile("", "file_cred_source")
			require.NoError(err)
			_, err = io.Copy(f, strings.NewReader(tt.fileContent))
			require.NoError(err)

			// Configure the credential source to use the temp file
			if tt.fc.Path == "" {
				tt.fc.Path = f.Name()
			}

			got, err := tt.fc.token()
			if tt.errStr == "" {
				require.NoError(err)
				require.Equal(tt.want, got)
			} else {
				require.ErrorContains(err, tt.errStr)
			}
		})
	}
}
