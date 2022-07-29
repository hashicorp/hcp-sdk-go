// Code generated by go-swagger; DO NOT EDIT.

package box_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new box service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for box service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBox(params *CreateBoxParams) (*CreateBoxOK, error)

	DeleteBox(params *DeleteBoxParams) (*DeleteBoxOK, error)

	ListBoxes(params *ListBoxesParams) (*ListBoxesOK, error)

	ReadBox(params *ReadBoxParams) (*ReadBoxOK, error)

	UpdateBox(params *UpdateBoxParams) (*UpdateBoxOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBox creates box creates a new vagrant box
*/
func (a *Client) CreateBox(params *CreateBoxParams) (*CreateBoxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBoxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBox",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBoxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBoxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBox deletes box deletes a vagrant box

  Deleting a Box removes all its Versions and Providers as
well. This operation cannot be undone.
*/
func (a *Client) DeleteBox(params *DeleteBoxParams) (*DeleteBoxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBoxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteBox",
		Method:             "DELETE",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBoxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBoxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteBox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBoxes lists boxes lists all of the boxes within a particular registry
*/
func (a *Client) ListBoxes(params *ListBoxesParams) (*ListBoxesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBoxesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListBoxes",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListBoxesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBoxesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListBoxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBox reads box reads a vagrant box
*/
func (a *Client) ReadBox(params *ReadBoxParams) (*ReadBoxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBoxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadBox",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadBoxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBoxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadBox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBox updates box updates the details of a vagrant box

  Note that this only updates details about the Box itself. To
work with a Box's Versions or Providers, use those respective
services.
*/
func (a *Client) UpdateBox(params *UpdateBoxParams) (*UpdateBoxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBoxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateBox",
		Method:             "PATCH",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBoxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBoxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
