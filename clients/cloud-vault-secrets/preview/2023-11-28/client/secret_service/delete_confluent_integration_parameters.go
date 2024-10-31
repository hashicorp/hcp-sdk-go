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

// NewDeleteConfluentIntegrationParams creates a new DeleteConfluentIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConfluentIntegrationParams() *DeleteConfluentIntegrationParams {
	return &DeleteConfluentIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfluentIntegrationParamsWithTimeout creates a new DeleteConfluentIntegrationParams object
// with the ability to set a timeout on a request.
func NewDeleteConfluentIntegrationParamsWithTimeout(timeout time.Duration) *DeleteConfluentIntegrationParams {
	return &DeleteConfluentIntegrationParams{
		timeout: timeout,
	}
}

// NewDeleteConfluentIntegrationParamsWithContext creates a new DeleteConfluentIntegrationParams object
// with the ability to set a context for a request.
func NewDeleteConfluentIntegrationParamsWithContext(ctx context.Context) *DeleteConfluentIntegrationParams {
	return &DeleteConfluentIntegrationParams{
		Context: ctx,
	}
}

// NewDeleteConfluentIntegrationParamsWithHTTPClient creates a new DeleteConfluentIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConfluentIntegrationParamsWithHTTPClient(client *http.Client) *DeleteConfluentIntegrationParams {
	return &DeleteConfluentIntegrationParams{
		HTTPClient: client,
	}
}

/*
DeleteConfluentIntegrationParams contains all the parameters to send to the API endpoint

	for the delete confluent integration operation.

	Typically these are written to a http.Request.
*/
type DeleteConfluentIntegrationParams struct {

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

// WithDefaults hydrates default values in the delete confluent integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfluentIntegrationParams) WithDefaults() *DeleteConfluentIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete confluent integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfluentIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithTimeout(timeout time.Duration) *DeleteConfluentIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithContext(ctx context.Context) *DeleteConfluentIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithHTTPClient(client *http.Client) *DeleteConfluentIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithName(name string) *DeleteConfluentIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithOrganizationID(organizationID string) *DeleteConfluentIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) WithProjectID(projectID string) *DeleteConfluentIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete confluent integration params
func (o *DeleteConfluentIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfluentIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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