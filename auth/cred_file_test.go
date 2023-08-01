package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/auth/workload"
	"github.com/stretchr/testify/require"
)

func TestCredentialFile_Validate(t *testing.T) {
	tests := []struct {
		name   string
		cf     *CredentialFile
		errStr string
	}{
		{
			name: "valid oauth",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeServicePrincipal,
				Oauth: &OauthConfig{
					ClientID: "abc",
					SecretID: "def",
				},
			},
			errStr: "",
		},
		{
			name: "valid workload",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeWorkload,
				Workload: &workload.IdentityProviderConfig{
					ProviderResourceName: "iam/test",
					AWS:                  &workload.AWSCredentialSource{},
				},
			},
			errStr: "",
		},
		{
			name: "mismatched scheme",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeServicePrincipal,
				Workload: &workload.IdentityProviderConfig{
					ProviderResourceName: "iam/test",
					AWS:                  &workload.AWSCredentialSource{},
				},
			},
			errStr: "oauth config must be set when scheme is \"service_principal_creds\"",
		},
		{
			name: "invalid oauth",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeServicePrincipal,
				Oauth: &OauthConfig{
					ClientID: "abc",
					SecretID: "",
				},
			},
			errStr: "oauth: both client_id and secret_id must be set",
		},
		{
			name: "invalid workload",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeWorkload,
				Workload: &workload.IdentityProviderConfig{
					AWS: &workload.AWSCredentialSource{},
				},
			},
			errStr: "workload: workload identity provider resource name must be set",
		},
		{
			name: "both set",
			cf: &CredentialFile{
				ProjectID: "123",
				Scheme:    CredentialFileSchemeWorkload,
				Workload: &workload.IdentityProviderConfig{
					ProviderResourceName: "iam/test",
					AWS:                  &workload.AWSCredentialSource{},
				},
				Oauth: &OauthConfig{
					ClientID: "abc",
					SecretID: "def",
				},
			},
			errStr: "only one of oauth or workload may be set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cf.Validate()
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestReadCredentialFile(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		r := require.New(t)
		cf := &CredentialFile{
			Scheme: CredentialFileSchemeServicePrincipal,
			Oauth: &OauthConfig{
				ClientID: "abc",
				SecretID: "123",
			},
		}

		f, err := ioutil.TempFile("", "")
		r.NoError(err)
		r.NoError(WriteCredentialFile(f.Name(), cf))

		out, err := ReadCredentialFile(f.Name())
		r.NoError(err)
		r.EqualValues(cf, out)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := ReadCredentialFile(fmt.Sprintf("random-%d", rand.Int()))
		require.Error(t, err)
		require.ErrorContains(t, err, "failed to read credential file: open random")
	})

	t.Run("invalid", func(t *testing.T) {
		r := require.New(t)
		cf := &CredentialFile{
			Scheme: CredentialFileSchemeServicePrincipal,
			Oauth: &OauthConfig{
				ClientID: "abc",
			},
		}

		f, err := ioutil.TempFile("", "")
		r.NoError(err)
		r.NoError(WriteCredentialFile(f.Name(), cf))

		_, err = ReadCredentialFile(f.Name())
		r.ErrorContains(err, "oauth: both client_id and secret_id must be set")
	})

	t.Run("bad format", func(t *testing.T) {
		r := require.New(t)
		data, err := json.Marshal("hello, world!")
		r.NoError(err)

		f, err := ioutil.TempFile("", "")
		r.NoError(err)

		_, err = io.Copy(f, bytes.NewBuffer(data))
		r.NoError(err)

		_, err = ReadCredentialFile(f.Name())
		r.ErrorContains(err, "failed to unmarshal credential file")
	})
}

func TestGetDefaultCredentialFile(t *testing.T) {
	t.Run("with env", func(t *testing.T) {
		r := require.New(t)

		cf := &CredentialFile{
			Scheme: CredentialFileSchemeServicePrincipal,
			Oauth: &OauthConfig{
				ClientID: "abc",
				SecretID: "123",
			},
		}

		f, err := ioutil.TempFile("", "")
		r.NoError(err)
		r.NoError(WriteCredentialFile(f.Name(), cf))

		t.Setenv(EnvHCPCredFile, f.Name())
		out, err := GetDefaultCredentialFile()
		r.NoError(err)
		r.EqualValues(cf, out)
	})

	t.Run("without env", func(t *testing.T) {
		r := require.New(t)

		cf := &CredentialFile{
			Scheme: CredentialFileSchemeServicePrincipal,
			Oauth: &OauthConfig{
				ClientID: "abc",
				SecretID: "123",
			},
		}

		f, err := ioutil.TempFile("", "")
		r.NoError(err)
		r.NoError(WriteCredentialFile(f.Name(), cf))

		testDefaultHCPCredFilePath = f.Name()
		out, err := GetDefaultCredentialFile()
		testDefaultHCPCredFilePath = ""
		r.NoError(err)
		r.EqualValues(cf, out)
	})
}

func Test_getCredentialFilePath(t *testing.T) {
	t.Run("with env", func(t *testing.T) {
		r := require.New(t)
		cf := "test-path"
		t.Setenv(EnvHCPCredFile, cf)
		p, err := GetCredentialFilePath()
		r.NoError(err)
		r.Equal(cf, p)
	})

	t.Run("without  env", func(t *testing.T) {
		r := require.New(t)
		p, err := GetCredentialFilePath()
		r.NoError(err)
		r.Contains(p, CredentialFileName)
	})
}

func Test_WriteCredentialFile(t *testing.T) {
	r := require.New(t)
	cf := &CredentialFile{
		Scheme: CredentialFileSchemeServicePrincipal,
		Oauth: &OauthConfig{
			ClientID: "abc",
			SecretID: "123",
		},
	}

	f, err := ioutil.TempFile("", "")
	r.NoError(err)
	r.NoError(WriteCredentialFile(f.Name(), cf))

	back, err := ReadCredentialFile(f.Name())
	r.NoError(err)
	r.EqualValues(cf, back)
}
