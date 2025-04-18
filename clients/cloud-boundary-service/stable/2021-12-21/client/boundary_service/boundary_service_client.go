// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new boundary service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new boundary service API client with basic auth credentials.
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

// New creates a new boundary service API client with a bearer token for authentication.
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
Client for boundary service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BoundaryServiceClusterUpdateDeadlineGet(params *BoundaryServiceClusterUpdateDeadlineGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceClusterUpdateDeadlineGetOK, error)

	BoundaryServiceCreate(params *BoundaryServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceCreateOK, error)

	BoundaryServiceDelete(params *BoundaryServiceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceDeleteOK, error)

	BoundaryServiceGet(params *BoundaryServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceGetOK, error)

	BoundaryServiceGetControllerConfiguration(params *BoundaryServiceGetControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceGetControllerConfigurationOK, error)

	BoundaryServiceList(params *BoundaryServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceListOK, error)

	BoundaryServiceMaintenanceWindowGet(params *BoundaryServiceMaintenanceWindowGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceMaintenanceWindowGetOK, error)

	BoundaryServiceMaintenanceWindowUpdate(params *BoundaryServiceMaintenanceWindowUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceMaintenanceWindowUpdateOK, error)

	BoundaryServiceResetControllerConfiguration(params *BoundaryServiceResetControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceResetControllerConfigurationOK, error)

	BoundaryServiceSessions(params *BoundaryServiceSessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceSessionsOK, error)

	BoundaryServiceUpdate(params *BoundaryServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateOK, error)

	BoundaryServiceUpdateApply(params *BoundaryServiceUpdateApplyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateApplyOK, error)

	BoundaryServiceUpdateControllerConfiguration(params *BoundaryServiceUpdateControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateControllerConfigurationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BoundaryServiceClusterUpdateDeadlineGet clusters update deadline get is used to get the deadline for a cluster update
*/
func (a *Client) BoundaryServiceClusterUpdateDeadlineGet(params *BoundaryServiceClusterUpdateDeadlineGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceClusterUpdateDeadlineGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceClusterUpdateDeadlineGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_ClusterUpdateDeadlineGet",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/update-deadline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceClusterUpdateDeadlineGetReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceClusterUpdateDeadlineGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceClusterUpdateDeadlineGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceCreate creates creates a new h c p boundary cluster
*/
func (a *Client) BoundaryServiceCreate(params *BoundaryServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_Create",
		Method:             "POST",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceDelete deletes deletes the specified h c p boundary cluster
*/
func (a *Client) BoundaryServiceDelete(params *BoundaryServiceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_Delete",
		Method:             "DELETE",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceGet gets returns a single existing h c p boundary cluster
*/
func (a *Client) BoundaryServiceGet(params *BoundaryServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_Get",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceGetReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceGetControllerConfiguration gets controller configuration is used to retrieve the controller configuration values of a specified boundary cluster
*/
func (a *Client) BoundaryServiceGetControllerConfiguration(params *BoundaryServiceGetControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceGetControllerConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceGetControllerConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_GetControllerConfiguration",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceGetControllerConfigurationReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceGetControllerConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceGetControllerConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceList lists returns all existing h c p boundary clusters
*/
func (a *Client) BoundaryServiceList(params *BoundaryServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_List",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceListReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceMaintenanceWindowGet maintenances window get is used to fetch the type of maintenance for the given cluster maintenance can either be automatic or scheduled if scheduled the response returns the window set by the user
*/
func (a *Client) BoundaryServiceMaintenanceWindowGet(params *BoundaryServiceMaintenanceWindowGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceMaintenanceWindowGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceMaintenanceWindowGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_MaintenanceWindowGet",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/maintenance-window",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceMaintenanceWindowGetReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceMaintenanceWindowGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceMaintenanceWindowGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceMaintenanceWindowUpdate maintenances window update sets maintenance window for update version
*/
func (a *Client) BoundaryServiceMaintenanceWindowUpdate(params *BoundaryServiceMaintenanceWindowUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceMaintenanceWindowUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceMaintenanceWindowUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_MaintenanceWindowUpdate",
		Method:             "POST",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/maintenance-window",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceMaintenanceWindowUpdateReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceMaintenanceWindowUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceMaintenanceWindowUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceResetControllerConfiguration resets controller configuration is used to reset a cluster s controller configuration to its default values
*/
func (a *Client) BoundaryServiceResetControllerConfiguration(params *BoundaryServiceResetControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceResetControllerConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceResetControllerConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_ResetControllerConfiguration",
		Method:             "DELETE",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceResetControllerConfigurationReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceResetControllerConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceResetControllerConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceSessions sessions returns all session information for existing h c p boundary cluster
*/
func (a *Client) BoundaryServiceSessions(params *BoundaryServiceSessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceSessionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_Sessions",
		Method:             "GET",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceSessionsReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceSessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceUpdate updates triggers checks if update for the specific cluster is needed and forces an update
*/
func (a *Client) BoundaryServiceUpdate(params *BoundaryServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_Update",
		Method:             "POST",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceUpdateApply updates apply is used to manually update a cluster to a specific version
*/
func (a *Client) BoundaryServiceUpdateApply(params *BoundaryServiceUpdateApplyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateApplyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceUpdateApplyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_UpdateApply",
		Method:             "POST",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceUpdateApplyReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceUpdateApplyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceUpdateApplyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BoundaryServiceUpdateControllerConfiguration updates controller configuration is used to modify the controller configuration values of a specified boundary cluster
*/
func (a *Client) BoundaryServiceUpdateControllerConfiguration(params *BoundaryServiceUpdateControllerConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BoundaryServiceUpdateControllerConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBoundaryServiceUpdateControllerConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BoundaryService_UpdateControllerConfiguration",
		Method:             "PUT",
		PathPattern:        "/boundary/2021-12-21/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/controller-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BoundaryServiceUpdateControllerConfigurationReader{formats: a.formats},
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
	success, ok := result.(*BoundaryServiceUpdateControllerConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BoundaryServiceUpdateControllerConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
