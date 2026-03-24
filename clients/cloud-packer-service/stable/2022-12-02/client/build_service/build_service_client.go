// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package build_service

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new build service API client (2022-12-02).
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// Client for Packer Build Service (2022-12-02) API.
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for the build service client.
type ClientService interface {
	GetImageByBuildLabels(params *GetImageByBuildLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImageByBuildLabelsOK, error)
	SetTransport(transport runtime.ClientTransport)
}

// GetImageByBuildLabels calls the 2022-12-02 GetImageByBuildLabels API.
func (a *Client) GetImageByBuildLabels(params *GetImageByBuildLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImageByBuildLabelsOK, error) {
	if params == nil {
		params = NewGetImageByBuildLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BuildService_GetImageByBuildLabels",
		Method:             "POST",
		PathPattern:        "/packer/2022-12-02/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_name}/channels/{channel_name}/_search/by_labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https", "http"},
		Params:             params,
		Reader:             &GetImageByBuildLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}
	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageByBuildLabelsOK)
	if ok {
		return success, nil
	}
	unexpectedSuccess := result.(*GetImageByBuildLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client.
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
