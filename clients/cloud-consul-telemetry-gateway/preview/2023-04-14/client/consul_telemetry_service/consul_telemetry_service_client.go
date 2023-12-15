// Code generated by go-swagger; DO NOT EDIT.

package consul_telemetry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new consul telemetry service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for consul telemetry service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AgentTelemetryConfig(params *AgentTelemetryConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentTelemetryConfigOK, error)

	GetLabelValues(params *GetLabelValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLabelValuesOK, error)

	GetServiceTopology(params *GetServiceTopologyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceTopologyOK, error)

	QueryRangeBatch(params *QueryRangeBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryRangeBatchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AgentTelemetryConfig agent telemetry config API
*/
func (a *Client) AgentTelemetryConfig(params *AgentTelemetryConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentTelemetryConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAgentTelemetryConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AgentTelemetryConfig",
		Method:             "POST",
		PathPattern:        "/ctgw/2023-04-14/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/agent/telemetry_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AgentTelemetryConfigReader{formats: a.formats},
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
	success, ok := result.(*AgentTelemetryConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AgentTelemetryConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLabelValues returns the label values for a given cluster
*/
func (a *Client) GetLabelValues(params *GetLabelValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLabelValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLabelValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLabelValues",
		Method:             "POST",
		PathPattern:        "/ctgw/2023-04-14/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/label/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLabelValuesReader{formats: a.formats},
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
	success, ok := result.(*GetLabelValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLabelValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServiceTopology creates a service topology graph based on recent metrics from the mesh
*/
func (a *Client) GetServiceTopology(params *GetServiceTopologyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceTopologyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceTopologyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServiceTopology",
		Method:             "GET",
		PathPattern:        "/ctgw/2023-04-14/organizations/{location.organization_id}/projects/{location.project_id}/topologies/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceTopologyReader{formats: a.formats},
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
	success, ok := result.(*GetServiceTopologyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServiceTopologyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryRangeBatch customs endpoints for observability that batches query range requests
*/
func (a *Client) QueryRangeBatch(params *QueryRangeBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryRangeBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRangeBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryRangeBatch",
		Method:             "POST",
		PathPattern:        "/ctgw/2023-04-14/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/query_range_batch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryRangeBatchReader{formats: a.formats},
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
	success, ok := result.(*QueryRangeBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryRangeBatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
