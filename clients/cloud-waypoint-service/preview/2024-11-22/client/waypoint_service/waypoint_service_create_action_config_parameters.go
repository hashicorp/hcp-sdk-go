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

// NewWaypointServiceCreateActionConfigParams creates a new WaypointServiceCreateActionConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateActionConfigParams() *WaypointServiceCreateActionConfigParams {
	return &WaypointServiceCreateActionConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateActionConfigParamsWithTimeout creates a new WaypointServiceCreateActionConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateActionConfigParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateActionConfigParams {
	return &WaypointServiceCreateActionConfigParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateActionConfigParamsWithContext creates a new WaypointServiceCreateActionConfigParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateActionConfigParamsWithContext(ctx context.Context) *WaypointServiceCreateActionConfigParams {
	return &WaypointServiceCreateActionConfigParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateActionConfigParamsWithHTTPClient creates a new WaypointServiceCreateActionConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateActionConfigParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateActionConfigParams {
	return &WaypointServiceCreateActionConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateActionConfigParams contains all the parameters to send to the API endpoint

	for the waypoint service create action config operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateActionConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateActionConfigBody

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

// WithDefaults hydrates default values in the waypoint service create action config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateActionConfigParams) WithDefaults() *WaypointServiceCreateActionConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create action config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateActionConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateActionConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithContext(ctx context.Context) *WaypointServiceCreateActionConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateActionConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateActionConfigBody) *WaypointServiceCreateActionConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateActionConfigBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceCreateActionConfigParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceCreateActionConfigParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service create action config params
func (o *WaypointServiceCreateActionConfigParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateActionConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
