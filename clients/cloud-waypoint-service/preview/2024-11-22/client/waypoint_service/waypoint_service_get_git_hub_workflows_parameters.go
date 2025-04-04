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

// NewWaypointServiceGetGitHubWorkflowsParams creates a new WaypointServiceGetGitHubWorkflowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetGitHubWorkflowsParams() *WaypointServiceGetGitHubWorkflowsParams {
	return &WaypointServiceGetGitHubWorkflowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetGitHubWorkflowsParamsWithTimeout creates a new WaypointServiceGetGitHubWorkflowsParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetGitHubWorkflowsParamsWithTimeout(timeout time.Duration) *WaypointServiceGetGitHubWorkflowsParams {
	return &WaypointServiceGetGitHubWorkflowsParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetGitHubWorkflowsParamsWithContext creates a new WaypointServiceGetGitHubWorkflowsParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetGitHubWorkflowsParamsWithContext(ctx context.Context) *WaypointServiceGetGitHubWorkflowsParams {
	return &WaypointServiceGetGitHubWorkflowsParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetGitHubWorkflowsParamsWithHTTPClient creates a new WaypointServiceGetGitHubWorkflowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetGitHubWorkflowsParamsWithHTTPClient(client *http.Client) *WaypointServiceGetGitHubWorkflowsParams {
	return &WaypointServiceGetGitHubWorkflowsParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetGitHubWorkflowsParams contains all the parameters to send to the API endpoint

	for the waypoint service get git hub workflows operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetGitHubWorkflowsParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceGetGitHubWorkflowsBody

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

// WithDefaults hydrates default values in the waypoint service get git hub workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetGitHubWorkflowsParams) WithDefaults() *WaypointServiceGetGitHubWorkflowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get git hub workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetGitHubWorkflowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithTimeout(timeout time.Duration) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithContext(ctx context.Context) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithHTTPClient(client *http.Client) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceGetGitHubWorkflowsBody) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceGetGitHubWorkflowsBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceGetGitHubWorkflowsParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get git hub workflows params
func (o *WaypointServiceGetGitHubWorkflowsParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetGitHubWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
