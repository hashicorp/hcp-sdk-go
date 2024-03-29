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

// NewWaypointServiceEndingActionParams creates a new WaypointServiceEndingActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceEndingActionParams() *WaypointServiceEndingActionParams {
	return &WaypointServiceEndingActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceEndingActionParamsWithTimeout creates a new WaypointServiceEndingActionParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceEndingActionParamsWithTimeout(timeout time.Duration) *WaypointServiceEndingActionParams {
	return &WaypointServiceEndingActionParams{
		timeout: timeout,
	}
}

// NewWaypointServiceEndingActionParamsWithContext creates a new WaypointServiceEndingActionParams object
// with the ability to set a context for a request.
func NewWaypointServiceEndingActionParamsWithContext(ctx context.Context) *WaypointServiceEndingActionParams {
	return &WaypointServiceEndingActionParams{
		Context: ctx,
	}
}

// NewWaypointServiceEndingActionParamsWithHTTPClient creates a new WaypointServiceEndingActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceEndingActionParamsWithHTTPClient(client *http.Client) *WaypointServiceEndingActionParams {
	return &WaypointServiceEndingActionParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceEndingActionParams contains all the parameters to send to the API endpoint

	for the waypoint service ending action operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceEndingActionParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceEndingActionBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service ending action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceEndingActionParams) WithDefaults() *WaypointServiceEndingActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service ending action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceEndingActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) WithTimeout(timeout time.Duration) *WaypointServiceEndingActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) WithContext(ctx context.Context) *WaypointServiceEndingActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) WithHTTPClient(client *http.Client) *WaypointServiceEndingActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceEndingActionBody) *WaypointServiceEndingActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceEndingActionBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) WithNamespaceID(namespaceID string) *WaypointServiceEndingActionParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service ending action params
func (o *WaypointServiceEndingActionParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceEndingActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
