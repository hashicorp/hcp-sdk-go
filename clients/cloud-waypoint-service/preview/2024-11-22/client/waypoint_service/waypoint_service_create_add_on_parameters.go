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

// NewWaypointServiceCreateAddOnParams creates a new WaypointServiceCreateAddOnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateAddOnParams() *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateAddOnParamsWithTimeout creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateAddOnParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateAddOnParamsWithContext creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateAddOnParamsWithContext(ctx context.Context) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateAddOnParamsWithHTTPClient creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateAddOnParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateAddOnParams contains all the parameters to send to the API endpoint

	for the waypoint service create add on operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateAddOnParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAddOnBody

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

// WithDefaults hydrates default values in the waypoint service create add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAddOnParams) WithDefaults() *WaypointServiceCreateAddOnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAddOnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateAddOnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithContext(ctx context.Context) *WaypointServiceCreateAddOnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateAddOnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAddOnBody) *WaypointServiceCreateAddOnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAddOnBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceCreateAddOnParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceCreateAddOnParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateAddOnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
