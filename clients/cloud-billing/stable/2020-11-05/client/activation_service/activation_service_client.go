// Code generated by go-swagger; DO NOT EDIT.

package activation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new activation service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new activation service API client with basic auth credentials.
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

// New creates a new activation service API client with a bearer token for authentication.
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
Client for activation service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActivationServiceActivate(params *ActivationServiceActivateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivationServiceActivateOK, error)

	ActivationServiceGetActivationDetails(params *ActivationServiceGetActivationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivationServiceGetActivationDetailsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ActivationServiceActivate activates is called by the front end to redeem the given activation code
*/
func (a *Client) ActivationServiceActivate(params *ActivationServiceActivateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivationServiceActivateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivationServiceActivateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ActivationService_Activate",
		Method:             "POST",
		PathPattern:        "/billing/2020-11-05/activations/{activation_code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ActivationServiceActivateReader{formats: a.formats},
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
	success, ok := result.(*ActivationServiceActivateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ActivationServiceActivateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ActivationServiceGetActivationDetails gets activation details returns contract activation details associated with an activation code
*/
func (a *Client) ActivationServiceGetActivationDetails(params *ActivationServiceGetActivationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivationServiceGetActivationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivationServiceGetActivationDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ActivationService_GetActivationDetails",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/activations/{activation_code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ActivationServiceGetActivationDetailsReader{formats: a.formats},
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
	success, ok := result.(*ActivationServiceGetActivationDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ActivationServiceGetActivationDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
