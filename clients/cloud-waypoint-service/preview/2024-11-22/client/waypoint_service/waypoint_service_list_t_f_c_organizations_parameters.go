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

// NewWaypointServiceListTFCOrganizationsParams creates a new WaypointServiceListTFCOrganizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceListTFCOrganizationsParams() *WaypointServiceListTFCOrganizationsParams {
	return &WaypointServiceListTFCOrganizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceListTFCOrganizationsParamsWithTimeout creates a new WaypointServiceListTFCOrganizationsParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceListTFCOrganizationsParamsWithTimeout(timeout time.Duration) *WaypointServiceListTFCOrganizationsParams {
	return &WaypointServiceListTFCOrganizationsParams{
		timeout: timeout,
	}
}

// NewWaypointServiceListTFCOrganizationsParamsWithContext creates a new WaypointServiceListTFCOrganizationsParams object
// with the ability to set a context for a request.
func NewWaypointServiceListTFCOrganizationsParamsWithContext(ctx context.Context) *WaypointServiceListTFCOrganizationsParams {
	return &WaypointServiceListTFCOrganizationsParams{
		Context: ctx,
	}
}

// NewWaypointServiceListTFCOrganizationsParamsWithHTTPClient creates a new WaypointServiceListTFCOrganizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceListTFCOrganizationsParamsWithHTTPClient(client *http.Client) *WaypointServiceListTFCOrganizationsParams {
	return &WaypointServiceListTFCOrganizationsParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceListTFCOrganizationsParams contains all the parameters to send to the API endpoint

	for the waypoint service list t f c organizations operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceListTFCOrganizationsParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceListTFCOrganizationsBody

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

// WithDefaults hydrates default values in the waypoint service list t f c organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListTFCOrganizationsParams) WithDefaults() *WaypointServiceListTFCOrganizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service list t f c organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListTFCOrganizationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithTimeout(timeout time.Duration) *WaypointServiceListTFCOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithContext(ctx context.Context) *WaypointServiceListTFCOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithHTTPClient(client *http.Client) *WaypointServiceListTFCOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceListTFCOrganizationsBody) *WaypointServiceListTFCOrganizationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceListTFCOrganizationsBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceListTFCOrganizationsParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceListTFCOrganizationsParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service list t f c organizations params
func (o *WaypointServiceListTFCOrganizationsParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceListTFCOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
