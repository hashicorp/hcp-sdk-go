// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// NewWaypointServiceUpdateAddOnDefinition2Params creates a new WaypointServiceUpdateAddOnDefinition2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateAddOnDefinition2Params() *WaypointServiceUpdateAddOnDefinition2Params {
	return &WaypointServiceUpdateAddOnDefinition2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateAddOnDefinition2ParamsWithTimeout creates a new WaypointServiceUpdateAddOnDefinition2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateAddOnDefinition2ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateAddOnDefinition2Params {
	return &WaypointServiceUpdateAddOnDefinition2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateAddOnDefinition2ParamsWithContext creates a new WaypointServiceUpdateAddOnDefinition2Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateAddOnDefinition2ParamsWithContext(ctx context.Context) *WaypointServiceUpdateAddOnDefinition2Params {
	return &WaypointServiceUpdateAddOnDefinition2Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateAddOnDefinition2ParamsWithHTTPClient creates a new WaypointServiceUpdateAddOnDefinition2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateAddOnDefinition2ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateAddOnDefinition2Params {
	return &WaypointServiceUpdateAddOnDefinition2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateAddOnDefinition2Params contains all the parameters to send to the API endpoint

	for the waypoint service update add on definition2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateAddOnDefinition2Params struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateAddOnDefinitionBody

	// ExistingAddOnDefinitionName.
	ExistingAddOnDefinitionName string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithDefaults() *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithContext(ctx context.Context) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateAddOnDefinitionBody) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateAddOnDefinitionBody) {
	o.Body = body
}

// WithExistingAddOnDefinitionName adds the existingAddOnDefinitionName to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithExistingAddOnDefinitionName(existingAddOnDefinitionName string) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetExistingAddOnDefinitionName(existingAddOnDefinitionName)
	return o
}

// SetExistingAddOnDefinitionName adds the existingAddOnDefinitionName to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetExistingAddOnDefinitionName(existingAddOnDefinitionName string) {
	o.ExistingAddOnDefinitionName = existingAddOnDefinitionName
}

// WithNamespaceID adds the namespaceID to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) WithNamespaceID(namespaceID string) *WaypointServiceUpdateAddOnDefinition2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update add on definition2 params
func (o *WaypointServiceUpdateAddOnDefinition2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateAddOnDefinition2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param existing_add_on_definition.name
	if err := r.SetPathParam("existing_add_on_definition.name", o.ExistingAddOnDefinitionName); err != nil {
		return err
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
