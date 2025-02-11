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

// NewWaypointServiceUpdateApplicationParams creates a new WaypointServiceUpdateApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationParams() *WaypointServiceUpdateApplicationParams {
	return &WaypointServiceUpdateApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationParamsWithTimeout creates a new WaypointServiceUpdateApplicationParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationParams {
	return &WaypointServiceUpdateApplicationParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationParamsWithContext creates a new WaypointServiceUpdateApplicationParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationParams {
	return &WaypointServiceUpdateApplicationParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationParams {
	return &WaypointServiceUpdateApplicationParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationParams contains all the parameters to send to the API endpoint

	for the waypoint service update application operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationParams struct {

	// ApplicationID.
	ApplicationID string

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody

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

// WithDefaults hydrates default values in the waypoint service update application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationParams) WithDefaults() *WaypointServiceUpdateApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithApplicationID(applicationID string) *WaypointServiceUpdateApplicationParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithBody adds the body to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody) *WaypointServiceUpdateApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceUpdateApplicationParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceUpdateApplicationParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service update application params
func (o *WaypointServiceUpdateApplicationParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.id
	if err := r.SetPathParam("application.id", o.ApplicationID); err != nil {
		return err
	}
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
