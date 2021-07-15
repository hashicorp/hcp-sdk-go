package main

import (
	"encoding/json"
	"testing"

	"github.com/go-openapi/loads"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/require"
)

var swagger = spec.Swagger{
	SwaggerProps: spec.SwaggerProps{
		ID:       "http://localhost:3849/hcp-api",
		Swagger:  "2.0",
		Host:     "some.api.out.there",
		BasePath: "/",
		Definitions: map[string]spec.Schema{
			"hashicorp.cloud.common.PaginationRequest": {SchemaProps: spec.SchemaProps{Type: []string{"object"}}},
			"hashicorp.cloud.location.Link":            {SchemaProps: spec.SchemaProps{Type: []string{"object"}}},
			"grpc.gateway.runtime.Error":               {SchemaProps: spec.SchemaProps{Type: []string{"object"}}},
		},
	},
}

const specJSON = `{
	"id": "http://localhost:3849/hcp-api",
	"swagger": "2.0",
	"host": "some.api.out.there",
	"basePath": "/",
	"definitions": {
		"hashicorp.cloud.common.PaginationRequest": {
            "type": "object"
		},
		"hashicorp.cloud.location.Link": {
            "type": "object"
        },
		"grpc.gateway.runtime.Error": {
            "type": "object"
        }
	}
}`

// For reference only, since we do more granular assertions in the test below.
//nolint
const transformedSpecJSON = `{
	"id": "http://localhost:3849/hcp-api",
	"swagger": "2.0",
	"host": "some.api.out.there",
	"basePath": "/",
	"definitions": {
		"hashicorp.cloud.common.PaginationRequest": {
            "type": "object",
            "x-go-type": {
                "import": {
                    "package": "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
                    "alias": "cloud"
                },
                "type": "HashicorpCloudCommonPaginationRequest"
            }
		},
		"hashicorp.cloud.location.Link": {
            "type": "object",
            "x-go-type": {
                "import": {
                    "package": "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
                    "alias": "cloud"
                },
                "type": "HashicorpCloudLocationLink"
            }
        },
		"grpc.gateway.runtime.Error": {
            "type": "object",
            "x-go-type": {
                "import": {
                    "package": "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
                    "alias": "cloud"
                },
                "type": "GrpcGatewayRuntimeError"
            }
        }
	}
}`

func TestAddSharedExtensions(t *testing.T) {

	var actual spec.Swagger
	err := json.Unmarshal([]byte(specJSON), &actual)
	require.NoError(t, err)
	require.EqualValues(t, actual, swagger)

	var raw json.RawMessage
	err = json.Unmarshal([]byte(specJSON), &raw)
	require.NoError(t, err)

	svcDoc, err := loads.Analyzed(raw, "")
	require.NoError(t, err)

	svcDocs := map[string]*loads.Document{"temp": svcDoc}

	sharedDefs := map[string]bool{
		"hashicorp.cloud.common.PaginationRequest": true,
		"hashicorp.cloud.location.Link":            true,
		"grpc.gateway.runtime.Error":               true,
	}

	sp := svcDocs["temp"].Spec()

	updatedSpec, err := addSharedExtension(sp, sharedDefs)
	require.NoError(t, err)

	require.Equal(t, spec.VendorExtensible{
		Extensions: spec.Extensions{
			"x-go-type": XGoType{
				Import: Import{
					Package: "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
					Alias:   "cloud",
				},
				Type: "HashicorpCloudCommonPaginationRequest",
			},
		},
	}, updatedSpec.SwaggerProps.Definitions["hashicorp.cloud.common.PaginationRequest"].VendorExtensible)

	require.Equal(t, spec.VendorExtensible{
		Extensions: spec.Extensions{
			"x-go-type": XGoType{
				Import: Import{
					Package: "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
					Alias:   "cloud",
				},
				Type: "HashicorpCloudLocationLink",
			},
		},
	}, updatedSpec.SwaggerProps.Definitions["hashicorp.cloud.location.Link"].VendorExtensible)

	require.Equal(t, spec.VendorExtensible{
		Extensions: spec.Extensions{
			"x-go-type": XGoType{
				Import: Import{
					Package: "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
					Alias:   "cloud",
				},
				Type: "GrpcGatewayRuntimeError",
			},
		},
	}, updatedSpec.SwaggerProps.Definitions["grpc.gateway.runtime.Error"].VendorExtensible)

}
