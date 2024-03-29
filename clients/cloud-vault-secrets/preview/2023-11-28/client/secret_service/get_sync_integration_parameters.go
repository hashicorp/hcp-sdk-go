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

// NewGetSyncIntegrationParams creates a new GetSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSyncIntegrationParams() *GetSyncIntegrationParams {
	return &GetSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncIntegrationParamsWithTimeout creates a new GetSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewGetSyncIntegrationParamsWithTimeout(timeout time.Duration) *GetSyncIntegrationParams {
	return &GetSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewGetSyncIntegrationParamsWithContext creates a new GetSyncIntegrationParams object
// with the ability to set a context for a request.
func NewGetSyncIntegrationParamsWithContext(ctx context.Context) *GetSyncIntegrationParams {
	return &GetSyncIntegrationParams{
		Context: ctx,
	}
}

// NewGetSyncIntegrationParamsWithHTTPClient creates a new GetSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSyncIntegrationParamsWithHTTPClient(client *http.Client) *GetSyncIntegrationParams {
	return &GetSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
GetSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the get sync integration operation.

	Typically these are written to a http.Request.
*/
type GetSyncIntegrationParams struct {

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

// WithDefaults hydrates default values in the get sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncIntegrationParams) WithDefaults() *GetSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sync integration params
func (o *GetSyncIntegrationParams) WithTimeout(timeout time.Duration) *GetSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync integration params
func (o *GetSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync integration params
func (o *GetSyncIntegrationParams) WithContext(ctx context.Context) *GetSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync integration params
func (o *GetSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync integration params
func (o *GetSyncIntegrationParams) WithHTTPClient(client *http.Client) *GetSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync integration params
func (o *GetSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get sync integration params
func (o *GetSyncIntegrationParams) WithName(name string) *GetSyncIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get sync integration params
func (o *GetSyncIntegrationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get sync integration params
func (o *GetSyncIntegrationParams) WithOrganizationID(organizationID string) *GetSyncIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get sync integration params
func (o *GetSyncIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get sync integration params
func (o *GetSyncIntegrationParams) WithProjectID(projectID string) *GetSyncIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get sync integration params
func (o *GetSyncIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
