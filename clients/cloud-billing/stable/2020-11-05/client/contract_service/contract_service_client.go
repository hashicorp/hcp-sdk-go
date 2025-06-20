// Code generated by go-swagger; DO NOT EDIT.

package contract_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new contract service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new contract service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new contract service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for contract service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ContractServiceGetContract(params *ContractServiceGetContractParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContractServiceGetContractOK, error)

	ContractServiceListContracts(params *ContractServiceListContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContractServiceListContractsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ContractServiceGetContract gets contract returns a contract by its public ID
*/
func (a *Client) ContractServiceGetContract(params *ContractServiceGetContractParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContractServiceGetContractOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceGetContractParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_GetContract",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts/{contract_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractServiceGetContractReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceGetContractOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceGetContractDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContracts lists contracts returns a list of a billing account s contracts optionally filtered
*/
func (a *Client) ContractServiceListContracts(params *ContractServiceListContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ContractServiceListContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContracts",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts/{billing_account_id}/contracts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractServiceListContractsReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
