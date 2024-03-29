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

// NewWaypointServiceUpdateApplicationsListParams creates a new WaypointServiceUpdateApplicationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationsListParams() *WaypointServiceUpdateApplicationsListParams {
	return &WaypointServiceUpdateApplicationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationsListParamsWithTimeout creates a new WaypointServiceUpdateApplicationsListParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationsListParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationsListParams {
	return &WaypointServiceUpdateApplicationsListParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationsListParamsWithContext creates a new WaypointServiceUpdateApplicationsListParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationsListParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationsListParams {
	return &WaypointServiceUpdateApplicationsListParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationsListParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationsListParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationsListParams {
	return &WaypointServiceUpdateApplicationsListParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationsListParams contains all the parameters to send to the API endpoint

	for the waypoint service update applications list operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationsListParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationsListBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update applications list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationsListParams) WithDefaults() *WaypointServiceUpdateApplicationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update applications list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationsListBody) *WaypointServiceUpdateApplicationsListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationsListBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) WithNamespaceID(namespaceID string) *WaypointServiceUpdateApplicationsListParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update applications list params
func (o *WaypointServiceUpdateApplicationsListParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
