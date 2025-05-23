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

// NewWaypointServiceUpdateApplicationTemplateParams creates a new WaypointServiceUpdateApplicationTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationTemplateParams() *WaypointServiceUpdateApplicationTemplateParams {
	return &WaypointServiceUpdateApplicationTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplateParamsWithTimeout creates a new WaypointServiceUpdateApplicationTemplateParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationTemplateParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplateParams {
	return &WaypointServiceUpdateApplicationTemplateParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplateParamsWithContext creates a new WaypointServiceUpdateApplicationTemplateParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationTemplateParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplateParams {
	return &WaypointServiceUpdateApplicationTemplateParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationTemplateParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationTemplateParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplateParams {
	return &WaypointServiceUpdateApplicationTemplateParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationTemplateParams contains all the parameters to send to the API endpoint

	for the waypoint service update application template operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationTemplateParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationTemplateBody

	/* ExistingApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ExistingApplicationTemplateID string

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

// WithDefaults hydrates default values in the waypoint service update application template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplateParams) WithDefaults() *WaypointServiceUpdateApplicationTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationTemplateBody) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationTemplateBody) {
	o.Body = body
}

// WithExistingApplicationTemplateID adds the existingApplicationTemplateID to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithExistingApplicationTemplateID(existingApplicationTemplateID string) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetExistingApplicationTemplateID(existingApplicationTemplateID)
	return o
}

// SetExistingApplicationTemplateID adds the existingApplicationTemplateId to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetExistingApplicationTemplateID(existingApplicationTemplateID string) {
	o.ExistingApplicationTemplateID = existingApplicationTemplateID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceUpdateApplicationTemplateParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service update application template params
func (o *WaypointServiceUpdateApplicationTemplateParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param existing_application_template.id
	if err := r.SetPathParam("existing_application_template.id", o.ExistingApplicationTemplateID); err != nil {
		return err
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
