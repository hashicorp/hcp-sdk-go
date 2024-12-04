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
)

// NewWaypointServiceListTFAgentPoolsParams creates a new WaypointServiceListTFAgentPoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceListTFAgentPoolsParams() *WaypointServiceListTFAgentPoolsParams {
	return &WaypointServiceListTFAgentPoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceListTFAgentPoolsParamsWithTimeout creates a new WaypointServiceListTFAgentPoolsParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceListTFAgentPoolsParamsWithTimeout(timeout time.Duration) *WaypointServiceListTFAgentPoolsParams {
	return &WaypointServiceListTFAgentPoolsParams{
		timeout: timeout,
	}
}

// NewWaypointServiceListTFAgentPoolsParamsWithContext creates a new WaypointServiceListTFAgentPoolsParams object
// with the ability to set a context for a request.
func NewWaypointServiceListTFAgentPoolsParamsWithContext(ctx context.Context) *WaypointServiceListTFAgentPoolsParams {
	return &WaypointServiceListTFAgentPoolsParams{
		Context: ctx,
	}
}

// NewWaypointServiceListTFAgentPoolsParamsWithHTTPClient creates a new WaypointServiceListTFAgentPoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceListTFAgentPoolsParamsWithHTTPClient(client *http.Client) *WaypointServiceListTFAgentPoolsParams {
	return &WaypointServiceListTFAgentPoolsParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceListTFAgentPoolsParams contains all the parameters to send to the API endpoint

	for the waypoint service list t f agent pools operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceListTFAgentPoolsParams struct {

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service list t f agent pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListTFAgentPoolsParams) WithDefaults() *WaypointServiceListTFAgentPoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service list t f agent pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListTFAgentPoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) WithTimeout(timeout time.Duration) *WaypointServiceListTFAgentPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) WithContext(ctx context.Context) *WaypointServiceListTFAgentPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) WithHTTPClient(client *http.Client) *WaypointServiceListTFAgentPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) WithNamespaceID(namespaceID string) *WaypointServiceListTFAgentPoolsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service list t f agent pools params
func (o *WaypointServiceListTFAgentPoolsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceListTFAgentPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}