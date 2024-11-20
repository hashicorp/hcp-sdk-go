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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// NewCreatePostgresIntegrationParams creates a new CreatePostgresIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePostgresIntegrationParams() *CreatePostgresIntegrationParams {
	return &CreatePostgresIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePostgresIntegrationParamsWithTimeout creates a new CreatePostgresIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreatePostgresIntegrationParamsWithTimeout(timeout time.Duration) *CreatePostgresIntegrationParams {
	return &CreatePostgresIntegrationParams{
		timeout: timeout,
	}
}

// NewCreatePostgresIntegrationParamsWithContext creates a new CreatePostgresIntegrationParams object
// with the ability to set a context for a request.
func NewCreatePostgresIntegrationParamsWithContext(ctx context.Context) *CreatePostgresIntegrationParams {
	return &CreatePostgresIntegrationParams{
		Context: ctx,
	}
}

// NewCreatePostgresIntegrationParamsWithHTTPClient creates a new CreatePostgresIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePostgresIntegrationParamsWithHTTPClient(client *http.Client) *CreatePostgresIntegrationParams {
	return &CreatePostgresIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreatePostgresIntegrationParams contains all the parameters to send to the API endpoint

	for the create postgres integration operation.

	Typically these are written to a http.Request.
*/
type CreatePostgresIntegrationParams struct {

	// Body.
	Body *models.SecretServiceCreatePostgresIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePostgresIntegrationParams) WithDefaults() *CreatePostgresIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create postgres integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePostgresIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithTimeout(timeout time.Duration) *CreatePostgresIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithContext(ctx context.Context) *CreatePostgresIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithHTTPClient(client *http.Client) *CreatePostgresIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithBody(body *models.SecretServiceCreatePostgresIntegrationBody) *CreatePostgresIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetBody(body *models.SecretServiceCreatePostgresIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithOrganizationID(organizationID string) *CreatePostgresIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create postgres integration params
func (o *CreatePostgresIntegrationParams) WithProjectID(projectID string) *CreatePostgresIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create postgres integration params
func (o *CreatePostgresIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePostgresIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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