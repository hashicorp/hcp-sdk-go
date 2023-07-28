// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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
)

// NewGetProjectServicePrincipalParams creates a new GetProjectServicePrincipalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectServicePrincipalParams() *GetProjectServicePrincipalParams {
	return &GetProjectServicePrincipalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectServicePrincipalParamsWithTimeout creates a new GetProjectServicePrincipalParams object
// with the ability to set a timeout on a request.
func NewGetProjectServicePrincipalParamsWithTimeout(timeout time.Duration) *GetProjectServicePrincipalParams {
	return &GetProjectServicePrincipalParams{
		timeout: timeout,
	}
}

// NewGetProjectServicePrincipalParamsWithContext creates a new GetProjectServicePrincipalParams object
// with the ability to set a context for a request.
func NewGetProjectServicePrincipalParamsWithContext(ctx context.Context) *GetProjectServicePrincipalParams {
	return &GetProjectServicePrincipalParams{
		Context: ctx,
	}
}

// NewGetProjectServicePrincipalParamsWithHTTPClient creates a new GetProjectServicePrincipalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectServicePrincipalParamsWithHTTPClient(client *http.Client) *GetProjectServicePrincipalParams {
	return &GetProjectServicePrincipalParams{
		HTTPClient: client,
	}
}

/*
GetProjectServicePrincipalParams contains all the parameters to send to the API endpoint

	for the get project service principal operation.

	Typically these are written to a http.Request.
*/
type GetProjectServicePrincipalParams struct {

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization that
	owns the requested service principal or project.
	*/
	OrganizationID string

	/* PrincipalID.

	   principal_id is the ID of the service principal being requested.
	*/
	PrincipalID string

	/* ProjectID.

	     project_id is the unique identifier (UUID) of the project that
	under which the service principal should be created.
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectServicePrincipalParams) WithDefaults() *GetProjectServicePrincipalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectServicePrincipalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithTimeout(timeout time.Duration) *GetProjectServicePrincipalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithContext(ctx context.Context) *GetProjectServicePrincipalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithHTTPClient(client *http.Client) *GetProjectServicePrincipalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithOrganizationID(organizationID string) *GetProjectServicePrincipalParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPrincipalID adds the principalID to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithPrincipalID(principalID string) *GetProjectServicePrincipalParams {
	o.SetPrincipalID(principalID)
	return o
}

// SetPrincipalID adds the principalId to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetPrincipalID(principalID string) {
	o.PrincipalID = principalID
}

// WithProjectID adds the projectID to the get project service principal params
func (o *GetProjectServicePrincipalParams) WithProjectID(projectID string) *GetProjectServicePrincipalParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project service principal params
func (o *GetProjectServicePrincipalParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectServicePrincipalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param principal_id
	if err := r.SetPathParam("principal_id", o.PrincipalID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
