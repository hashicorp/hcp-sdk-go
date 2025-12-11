// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredentialTokenSource_Validate(t *testing.T) {
	tests := []struct {
		name   string
		ct     *CredentialTokenSource
		errStr string
	}{
		{
			name:   "empty token",
			ct:     &CredentialTokenSource{},
			errStr: "token must be set",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.ct.Validate()
			if test.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.errStr)
			}
		})
	}
}
