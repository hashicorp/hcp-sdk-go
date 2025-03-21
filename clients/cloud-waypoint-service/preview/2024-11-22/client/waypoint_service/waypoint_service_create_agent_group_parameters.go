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

// NewWaypointServiceCreateAgentGroupParams creates a new WaypointServiceCreateAgentGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateAgentGroupParams() *WaypointServiceCreateAgentGroupParams {
	return &WaypointServiceCreateAgentGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateAgentGroupParamsWithTimeout creates a new WaypointServiceCreateAgentGroupParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateAgentGroupParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateAgentGroupParams {
	return &WaypointServiceCreateAgentGroupParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateAgentGroupParamsWithContext creates a new WaypointServiceCreateAgentGroupParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateAgentGroupParamsWithContext(ctx context.Context) *WaypointServiceCreateAgentGroupParams {
	return &WaypointServiceCreateAgentGroupParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateAgentGroupParamsWithHTTPClient creates a new WaypointServiceCreateAgentGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateAgentGroupParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateAgentGroupParams {
	return &WaypointServiceCreateAgentGroupParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateAgentGroupParams contains all the parameters to send to the API endpoint

	for the waypoint service create agent group operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateAgentGroupParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAgentGroupBody

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

// WithDefaults hydrates default values in the waypoint service create agent group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAgentGroupParams) WithDefaults() *WaypointServiceCreateAgentGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create agent group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAgentGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateAgentGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithContext(ctx context.Context) *WaypointServiceCreateAgentGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateAgentGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAgentGroupBody) *WaypointServiceCreateAgentGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateAgentGroupBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceCreateAgentGroupParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceCreateAgentGroupParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service create agent group params
func (o *WaypointServiceCreateAgentGroupParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateAgentGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
