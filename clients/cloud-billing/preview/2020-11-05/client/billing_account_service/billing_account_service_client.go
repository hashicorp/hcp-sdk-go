// Code generated by go-swagger; DO NOT EDIT.

package billing_account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing account service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing account service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BillingAccountServiceCreate(params *BillingAccountServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceCreateOK, error)

	BillingAccountServiceCreateSetupIntent(params *BillingAccountServiceCreateSetupIntentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceCreateSetupIntentOK, error)

	BillingAccountServiceDeleteBillingAccount(params *BillingAccountServiceDeleteBillingAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceDeleteBillingAccountOK, error)

	BillingAccountServiceGet(params *BillingAccountServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceGetOK, error)

	BillingAccountServiceGetByProject(params *BillingAccountServiceGetByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceGetByProjectOK, error)

	BillingAccountServiceList(params *BillingAccountServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceListOK, error)

	BillingAccountServiceRemoveOnDemandPaymentMethod(params *BillingAccountServiceRemoveOnDemandPaymentMethodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceRemoveOnDemandPaymentMethodOK, error)

	BillingAccountServiceUpdate(params *BillingAccountServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	BillingAccountServiceCreate creates a new billing account for the organization

	This endpoint is idempotent and it is expected that clients will retry

their requests on server-side failure.
*/
func (a *Client) BillingAccountServiceCreate(params *BillingAccountServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_Create",
		Method:             "POST",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceCreateSetupIntent creates setup intent creates a setup intent used to collect payment details with stripe js and perform any necessary checks e g 3 d secure

https://stripe.com/docs/payments/setup-intents
*/
func (a *Client) BillingAccountServiceCreateSetupIntent(params *BillingAccountServiceCreateSetupIntentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceCreateSetupIntentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceCreateSetupIntentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_CreateSetupIntent",
		Method:             "POST",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/setup-intents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceCreateSetupIntentReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceCreateSetupIntentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceCreateSetupIntentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceDeleteBillingAccount deletes deletes an organization s billing account this method does not follow the convention of the file due to swagger conflicts
*/
func (a *Client) BillingAccountServiceDeleteBillingAccount(params *BillingAccountServiceDeleteBillingAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceDeleteBillingAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceDeleteBillingAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_DeleteBillingAccount",
		Method:             "DELETE",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceDeleteBillingAccountReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceDeleteBillingAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceDeleteBillingAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceGet gets a billing account by ID
*/
func (a *Client) BillingAccountServiceGet(params *BillingAccountServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_Get",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceGetReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceGetByProject gets by project returns the billing account associated with the given project if no billing account is set for the project not found is returned
*/
func (a *Client) BillingAccountServiceGetByProject(params *BillingAccountServiceGetByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceGetByProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceGetByProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_GetByProject",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/projects/{project_id}/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceGetByProjectReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceGetByProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceGetByProjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceList lists an organization s billing accounts
*/
func (a *Client) BillingAccountServiceList(params *BillingAccountServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_List",
		Method:             "GET",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceListReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceRemoveOnDemandPaymentMethod billing account service remove on demand payment method API
*/
func (a *Client) BillingAccountServiceRemoveOnDemandPaymentMethod(params *BillingAccountServiceRemoveOnDemandPaymentMethodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceRemoveOnDemandPaymentMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceRemoveOnDemandPaymentMethodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_RemoveOnDemandPaymentMethod",
		Method:             "DELETE",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/billing_accounts/{billing_account_id}/on_demand_payment_method",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceRemoveOnDemandPaymentMethodReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceRemoveOnDemandPaymentMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceRemoveOnDemandPaymentMethodDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BillingAccountServiceUpdate updates a billing account s settings
*/
func (a *Client) BillingAccountServiceUpdate(params *BillingAccountServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingAccountServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingAccountServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BillingAccountService_Update",
		Method:             "PUT",
		PathPattern:        "/billing/2020-11-05/organizations/{organization_id}/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BillingAccountServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*BillingAccountServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BillingAccountServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}