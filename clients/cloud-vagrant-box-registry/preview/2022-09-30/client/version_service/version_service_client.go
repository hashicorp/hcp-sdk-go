// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new version service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVersion(params *CreateVersionParams) (*CreateVersionOK, error)

	DeleteVersion(params *DeleteVersionParams) (*DeleteVersionOK, error)

	ListVersions(params *ListVersionsParams) (*ListVersionsOK, error)

	ReadVersion(params *ReadVersionParams) (*ReadVersionOK, error)

	ReleaseVersion(params *ReleaseVersionParams) (*ReleaseVersionOK, error)

	RevokeVersion(params *RevokeVersionParams) (*RevokeVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVersion creates version creates a new box version
*/
func (a *Client) CreateVersion(params *CreateVersionParams) (*CreateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVersion deletes version deletes a box version

  Deleting a Box Version removes all its Providers as well. This
operation cannot be undone.
*/
func (a *Client) DeleteVersion(params *DeleteVersionParams) (*DeleteVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVersion",
		Method:             "DELETE",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListVersions lists version lists all of the versions within a particular box
*/
func (a *Client) ListVersions(params *ListVersionsParams) (*ListVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVersions",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadVersion reads version reads a box version
*/
func (a *Client) ReadVersion(params *ReadVersionParams) (*ReadVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadVersion",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReleaseVersion releases releases the specified version the version must not already be released
*/
func (a *Client) ReleaseVersion(params *ReleaseVersionParams) (*ReleaseVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReleaseVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReleaseVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReleaseVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReleaseVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevokeVersion revokes revokes the specified version the version must be actively released
*/
func (a *Client) RevokeVersion(params *RevokeVersionParams) (*RevokeVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RevokeVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RevokeVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RevokeVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
