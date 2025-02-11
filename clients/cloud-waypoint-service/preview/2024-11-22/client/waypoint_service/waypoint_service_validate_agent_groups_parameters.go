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

// NewWaypointServiceValidateAgentGroupsParams creates a new WaypointServiceValidateAgentGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceValidateAgentGroupsParams() *WaypointServiceValidateAgentGroupsParams {
	return &WaypointServiceValidateAgentGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceValidateAgentGroupsParamsWithTimeout creates a new WaypointServiceValidateAgentGroupsParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceValidateAgentGroupsParamsWithTimeout(timeout time.Duration) *WaypointServiceValidateAgentGroupsParams {
	return &WaypointServiceValidateAgentGroupsParams{
		timeout: timeout,
	}
}

// NewWaypointServiceValidateAgentGroupsParamsWithContext creates a new WaypointServiceValidateAgentGroupsParams object
// with the ability to set a context for a request.
func NewWaypointServiceValidateAgentGroupsParamsWithContext(ctx context.Context) *WaypointServiceValidateAgentGroupsParams {
	return &WaypointServiceValidateAgentGroupsParams{
		Context: ctx,
	}
}

// NewWaypointServiceValidateAgentGroupsParamsWithHTTPClient creates a new WaypointServiceValidateAgentGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceValidateAgentGroupsParamsWithHTTPClient(client *http.Client) *WaypointServiceValidateAgentGroupsParams {
	return &WaypointServiceValidateAgentGroupsParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceValidateAgentGroupsParams contains all the parameters to send to the API endpoint

	for the waypoint service validate agent groups operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceValidateAgentGroupsParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceValidateAgentGroupsBody

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

// WithDefaults hydrates default values in the waypoint service validate agent groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceValidateAgentGroupsParams) WithDefaults() *WaypointServiceValidateAgentGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service validate agent groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceValidateAgentGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithTimeout(timeout time.Duration) *WaypointServiceValidateAgentGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithContext(ctx context.Context) *WaypointServiceValidateAgentGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithHTTPClient(client *http.Client) *WaypointServiceValidateAgentGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceValidateAgentGroupsBody) *WaypointServiceValidateAgentGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceValidateAgentGroupsBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceValidateAgentGroupsParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceValidateAgentGroupsParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service validate agent groups params
func (o *WaypointServiceValidateAgentGroupsParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceValidateAgentGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
