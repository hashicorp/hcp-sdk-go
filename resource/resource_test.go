// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"encoding"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/stretchr/testify/require"
)

var (
	_ encoding.TextMarshaler   = &Resource{}
	_ encoding.TextUnmarshaler = &Resource{}
	_ json.Marshaler           = &Resource{}
	_ json.Unmarshaler         = &Resource{}
	_ fmt.Stringer             = &Resource{}
)

func testResource() Resource {
	return Resource{
		ID:           "foo",
		Type:         "resource.example",
		Organization: "209ffeb2-542e-4e81-a99d-ce9374c6ea7c",
		Project:      "262305a9-76e1-4a3b-b940-df55d493edb8",
	}
}

func TestFromString(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected Resource
		errStr   string
	}{
		{
			name:     "normal case",
			input:    "organization/209ffeb2-542e-4e81-a99d-ce9374c6ea7c/project/262305a9-76e1-4a3b-b940-df55d493edb8/resource.example/foo",
			expected: testResource(),
		},
		{
			name:   "bad organization token",
			input:  "org/209ffeb2-542e-4e81-a99d-ce9374c6ea7c/project/262305a9-76e1-4a3b-b940-df55d493edb8/resource.example/foo",
			errStr: "unexpected token",
		},
		{
			name:   "bad organization id",
			input:  "organization/abc123/project/262305a9-76e1-4a3b-b940-df55d493edb8/resource.example/foo",
			errStr: "must be a valid UUID",
		},
		{
			name:   "bad project token",
			input:  "organization/209ffeb2-542e-4e81-a99d-ce9374c6ea7c/proj/262305a9-76e1-4a3b-b940-df55d493edb8/resource.example/foo",
			errStr: "unexpected token",
		},
		{
			name:   "bad project id",
			input:  "organization/262305a9-76e1-4a3b-b940-df55d493edb8/project/abc123/resource.example/foo",
			errStr: "must be a valid UUID",
		},
		{
			name:   "bad ID format",
			input:  "foobar",
			errStr: "unexpected number of tokens",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tc *testing.T) {
			res, err := FromString(c.input)
			if len(c.errStr) == 0 {
				require.NoError(tc, err)
				require.Equal(tc, c.expected.String(), res.String())
			} else {
				require.Error(tc, err)
				require.Contains(tc, err.Error(), c.errStr)
			}
		})
	}
}

func TestFromLink(t *testing.T) {
	cases := []struct {
		name     string
		input    *models.HashicorpCloudLocationLink
		expected Resource
		errStr   string
	}{
		{
			name: "normal case",
			input: &models.HashicorpCloudLocationLink{
				ID:   "foo",
				Type: "resource.example",
				Location: &models.HashicorpCloudLocationLocation{
					ProjectID:      "262305a9-76e1-4a3b-b940-df55d493edb8",
					OrganizationID: "209ffeb2-542e-4e81-a99d-ce9374c6ea7c",
				},
			},
			expected: testResource(),
		},
		{
			name:   "nil link",
			errStr: "must not be nil",
		},
		{
			name:   "nil link.Location",
			input:  &models.HashicorpCloudLocationLink{},
			errStr: "must not be nil",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tc *testing.T) {
			res, err := FromLink(c.input)
			if len(c.errStr) == 0 {
				require.NoError(tc, err)
				require.Equal(tc, c.expected.String(), res.String())
			} else {
				require.Error(tc, err)
				require.Contains(tc, err.Error(), c.errStr)
			}
		})
	}
}

type testResourceStruct struct {
	ID Resource
}

func TestResourceEncoding(t *testing.T) {
	expectedStr := "organization/209ffeb2-542e-4e81-a99d-ce9374c6ea7c/project/262305a9-76e1-4a3b-b940-df55d493edb8/resource.example/foo"
	require.Equal(t, expectedStr, testResource().String())

	obj := testResourceStruct{
		ID: testResource(),
	}
	data, err := json.Marshal(&obj)
	require.NoError(t, err)
	dataMap := map[string]string{}
	require.NoError(t, json.Unmarshal(data, &dataMap))
	require.Equal(t, expectedStr, dataMap["ID"])

	var obj2 testResourceStruct
	require.NoError(t, json.Unmarshal(data, &obj2))
	require.Equal(t, obj, obj2)
}

func TestMust(t *testing.T) {
	require.NotPanics(t, func() {
		Must(FromString(testResource().String()))
	})
	require.Panics(t, func() {
		Must(FromString("abc"))
	})
}
