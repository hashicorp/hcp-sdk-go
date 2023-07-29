package workload

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredentialFormat_Validate(t *testing.T) {
	tests := []struct {
		name                     string
		Type                     string
		SubjectCredentialPointer string
		errStr                   string
	}{
		{
			name:   "default",
			errStr: "",
		},
		{
			name:                     "default type with subject token",
			SubjectCredentialPointer: "/field",
			errStr:                   "subject credential pointer must not be set with format type",
		},
		{
			name:   "value valid",
			Type:   FormatTypeValue,
			errStr: "",
		},
		{
			name:                     "value type with subject token",
			Type:                     FormatTypeValue,
			SubjectCredentialPointer: "/field",
			errStr:                   "subject credential pointer must not be set with format type",
		},
		{
			name:                     "json valid",
			Type:                     FormatTypeJSON,
			SubjectCredentialPointer: "/token",
			errStr:                   "",
		},
		{
			name:                     "json and no subject token field name",
			Type:                     FormatTypeJSON,
			SubjectCredentialPointer: "",
			errStr:                   "subject credential pointer must be set",
		},
		{
			name:                     "json invalid pointer",
			Type:                     FormatTypeJSON,
			SubjectCredentialPointer: "token!",
			errStr:                   "subject credential pointer is invalid: JSON pointer must be empty or start with a \"/\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := CredentialFormat{
				Type:                     tt.Type,
				SubjectCredentialPointer: tt.SubjectCredentialPointer,
			}
			err := cf.Validate()
			if tt.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}

func TestCredentialFormat_get(t *testing.T) {
	tests := []struct {
		name   string
		format CredentialFormat
		value  []byte
		want   string
		errStr string
	}{
		{
			name:   "default value",
			format: CredentialFormat{},
			value:  []byte("hello, world!"),
			want:   "hello, world!",
			errStr: "",
		},
		{
			name: "value",
			format: CredentialFormat{
				Type: FormatTypeValue,
			},
			value:  []byte("hello, world!"),
			want:   "hello, world!",
			errStr: "",
		},
		{
			name: "invalid json",
			format: CredentialFormat{
				Type:                     FormatTypeJSON,
				SubjectCredentialPointer: "/good",
			},
			value:  []byte(`{"good": "token"`),
			want:   "",
			errStr: "failed to unmarshal json value: unexpected end of JSON input",
		},
		{
			name: "json top level",
			format: CredentialFormat{
				Type:                     FormatTypeJSON,
				SubjectCredentialPointer: "/good",
			},
			value:  []byte(`{"good": "token"}`),
			want:   "token",
			errStr: "",
		},
		{
			name: "json not found key",
			format: CredentialFormat{
				Type:                     FormatTypeJSON,
				SubjectCredentialPointer: "/unknown",
			},
			value:  []byte(`{"good": "token"}`),
			want:   "",
			errStr: "Object has no key 'unknown'",
		},
		{
			name: "json deep pointer",
			format: CredentialFormat{
				Type:                     FormatTypeJSON,
				SubjectCredentialPointer: "/top/inner/deeper",
			},
			value:  []byte(`{"top": {"inner": {"deeper": "token"}}}`),
			want:   "token",
			errStr: "",
		},
		{
			name: "json non-string value",
			format: CredentialFormat{
				Type:                     FormatTypeJSON,
				SubjectCredentialPointer: "/top",
			},
			value:  []byte(`{"top": true}`),
			want:   "",
			errStr: "credential must be a string; got bool",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.format.get(tt.value)
			if tt.errStr == "" {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.ErrorContains(t, err, tt.errStr)
			}
		})
	}
}
