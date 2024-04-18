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

// NewWaypointServiceUIBulkUpdateActionAssignForAppParams creates a new WaypointServiceUIBulkUpdateActionAssignForAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUIBulkUpdateActionAssignForAppParams() *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	return &WaypointServiceUIBulkUpdateActionAssignForAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithTimeout creates a new WaypointServiceUIBulkUpdateActionAssignForAppParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithTimeout(timeout time.Duration) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	return &WaypointServiceUIBulkUpdateActionAssignForAppParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithContext creates a new WaypointServiceUIBulkUpdateActionAssignForAppParams object
// with the ability to set a context for a request.
func NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithContext(ctx context.Context) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	return &WaypointServiceUIBulkUpdateActionAssignForAppParams{
		Context: ctx,
	}
}

// NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithHTTPClient creates a new WaypointServiceUIBulkUpdateActionAssignForAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUIBulkUpdateActionAssignForAppParamsWithHTTPClient(client *http.Client) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	return &WaypointServiceUIBulkUpdateActionAssignForAppParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUIBulkUpdateActionAssignForAppParams contains all the parameters to send to the API endpoint

	for the waypoint service UI bulk update action assign for app operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUIBulkUpdateActionAssignForAppParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUIBulkUpdateActionAssignForAppBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service UI bulk update action assign for app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithDefaults() *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service UI bulk update action assign for app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithTimeout(timeout time.Duration) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithContext(ctx context.Context) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithHTTPClient(client *http.Client) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUIBulkUpdateActionAssignForAppBody) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUIBulkUpdateActionAssignForAppBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WithNamespaceID(namespaceID string) *WaypointServiceUIBulkUpdateActionAssignForAppParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service UI bulk update action assign for app params
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUIBulkUpdateActionAssignForAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
