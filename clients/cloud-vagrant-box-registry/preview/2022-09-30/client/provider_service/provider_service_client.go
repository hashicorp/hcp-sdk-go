// Code generated by go-swagger; DO NOT EDIT.

package provider_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new provider service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for provider service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompleteUpload(params *CompleteUploadParams) (*CompleteUploadOK, error)

	CreateProvider(params *CreateProviderParams) (*CreateProviderOK, error)

	DeleteProvider(params *DeleteProviderParams) (*DeleteProviderOK, error)

	DownloadProvider(params *DownloadProviderParams) (*DownloadProviderOK, error)

	GetProvider(params *GetProviderParams) (*GetProviderOK, error)

	ListProviders(params *ListProvidersParams) (*ListProvidersOK, error)

	UpdateProvider(params *UpdateProviderParams) (*UpdateProviderOK, error)

	UploadProvider(params *UploadProviderParams) (*UploadProviderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompleteUpload completes upload signals that the upload for a provider is complete
*/
func (a *Client) CompleteUpload(params *CompleteUploadParams) (*CompleteUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteUploadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompleteUpload",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteUploadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompleteUpload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateProvider creates provider creates a provider in the specified version
*/
func (a *Client) CreateProvider(params *CreateProviderParams) (*CreateProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateProvider",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProvider deletes provider deletes a provider note that this deletes any data associated with a hosted provider so use with care
*/
func (a *Client) DeleteProvider(params *DeleteProviderParams) (*DeleteProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProvider",
		Method:             "DELETE",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DownloadProvider downloads provider initiates a provider download

  For Hosted Providers, this is an atomic call that prepares the Provider
for download from the object storage. For External Providers, this is
simply a pass-through to the external download data supplied in the
External Provider record.
*/
func (a *Client) DownloadProvider(params *DownloadProviderParams) (*DownloadProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadProvider",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DownloadProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProvider gets provider fetches a provider for the specified version
*/
func (a *Client) GetProvider(params *GetProviderParams) (*GetProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProvider",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProviders lists providers lists all providers within a version
*/
func (a *Client) ListProviders(params *ListProvidersParams) (*ListProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProvidersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListProviders",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListProvidersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProvider updates provider updates details about a specified provider

  Note that you don't use this to upload data to Hosted Providers, to do
that, use the Upload method.
*/
func (a *Client) UpdateProvider(params *UpdateProviderParams) (*UpdateProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateProvider",
		Method:             "PATCH",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadProvider uploads provider initiates the upload of a hosted provider the service will return details that can be used to upload the data over HTTP after complete the caller should call complete upload

  Note that a Version needs to be unreleased in order to upload Providers to
it.

Overwrite is permitted; old/existing data for a particular Provider will
be replaced with data from a new successful upload for that same
Provider.
*/
func (a *Client) UploadProvider(params *UploadProviderParams) (*UploadProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadProvider",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UploadProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
