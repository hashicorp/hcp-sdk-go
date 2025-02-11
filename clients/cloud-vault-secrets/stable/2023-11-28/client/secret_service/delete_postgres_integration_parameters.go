// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewDeletePostgresIntegrationParams creates a new DeletePostgresIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePostgresIntegrationParams() *DeletePostgresIntegrationParams {
	return &DeletePostgresIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePostgresIntegrationParamsWithTimeout creates a new DeletePostgresIntegrationParams object
// with the ability to set a timeout on a request.
func NewDeletePostgresIntegrationParamsWithTimeout(timeout time.Duration) *DeletePostgresIntegrationParams {
	return &DeletePostgresIntegrationParams{
		timeout: timeout,
	}
}

// NewDeletePostgresIntegrationParamsWithContext creates a new DeletePostgresIntegrationParams object
// with the ability to set a context for a request.
func NewDeletePostgresIntegrationParamsWithContext(ctx context.Context) *DeletePostgresIntegrationParams {
	return &DeletePostgresIntegrationParams{
		Context: ctx,
	}
}

// NewDeletePostgresIntegrationParamsWithHTTPClient creates a new DeletePostgresIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePostgresIntegrationParamsWithHTTPClient(client *http.Client) *DeletePostgresIntegrationParams {
	return &DeletePostgresIntegrationParams{
		HTTPClient: client,
	}
}

/*
DeletePostgresIntegrationParams contains all the parameters to send to the API endpoint

	for the delete postgres integration operation.

	Typically these are written to a http.Request.
*/
type DeletePostgresIntegrationParams struct {

	// Name.
	Name string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePostgresIntegrationParams) WithDefaults() *DeletePostgresIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePostgresIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithTimeout(timeout time.Duration) *DeletePostgresIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithContext(ctx context.Context) *DeletePostgresIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithHTTPClient(client *http.Client) *DeletePostgresIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithName(name string) *DeletePostgresIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithOrganizationID(organizationID string) *DeletePostgresIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) WithProjectID(projectID string) *DeletePostgresIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete postgres integration params
func (o *DeletePostgresIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePostgresIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
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
