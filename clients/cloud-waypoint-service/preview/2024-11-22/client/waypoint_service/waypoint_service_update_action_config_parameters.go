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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2024-11-22/models"
)

// NewWaypointServiceUpdateActionConfigParams creates a new WaypointServiceUpdateActionConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateActionConfigParams() *WaypointServiceUpdateActionConfigParams {
	return &WaypointServiceUpdateActionConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateActionConfigParamsWithTimeout creates a new WaypointServiceUpdateActionConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateActionConfigParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateActionConfigParams {
	return &WaypointServiceUpdateActionConfigParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateActionConfigParamsWithContext creates a new WaypointServiceUpdateActionConfigParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateActionConfigParamsWithContext(ctx context.Context) *WaypointServiceUpdateActionConfigParams {
	return &WaypointServiceUpdateActionConfigParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateActionConfigParamsWithHTTPClient creates a new WaypointServiceUpdateActionConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateActionConfigParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateActionConfigParams {
	return &WaypointServiceUpdateActionConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateActionConfigParams contains all the parameters to send to the API endpoint

	for the waypoint service update action config operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateActionConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateActionConfigBody

	/* NamespaceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	NamespaceLocationOrganizationID string

	/* NamespaceLocationProjectID.

	   project_id is the projects id.
	*/
	NamespaceLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update action config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateActionConfigParams) WithDefaults() *WaypointServiceUpdateActionConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update action config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateActionConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateActionConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithContext(ctx context.Context) *WaypointServiceUpdateActionConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateActionConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateActionConfigBody) *WaypointServiceUpdateActionConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateActionConfigBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceUpdateActionConfigParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceUpdateActionConfigParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service update action config params
func (o *WaypointServiceUpdateActionConfigParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateActionConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace.location.organization_id
	if err := r.SetPathParam("namespace.location.organization_id", o.NamespaceLocationOrganizationID); err != nil {
		return err
	}

	// path param namespace.location.project_id
	if err := r.SetPathParam("namespace.location.project_id", o.NamespaceLocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
