// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new boundary service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for boundary service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOK, error)

	Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error)

	Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error)

	List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error)

	Sessions(params *SessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SessionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Create creates creates a new h c p boundary cluster
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
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
	success, ok := result.(*CreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Delete deletes deletes the specified h c p boundary cluster
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Get gets returns a single existing h c p boundary cluster
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
List lists returns all existing h c p boundary clusters
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
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
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Sessions sessions returns all session information for existing h c p boundary cluster
*/
func (a *Client) Sessions(params *SessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSessionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Sessions",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SessionsReader{formats: a.formats},
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
	success, ok := result.(*SessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
