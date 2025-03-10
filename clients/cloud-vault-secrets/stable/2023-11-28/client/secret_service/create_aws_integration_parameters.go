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

// NewCreateAwsIntegrationParams creates a new CreateAwsIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAwsIntegrationParams() *CreateAwsIntegrationParams {
	return &CreateAwsIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAwsIntegrationParamsWithTimeout creates a new CreateAwsIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateAwsIntegrationParamsWithTimeout(timeout time.Duration) *CreateAwsIntegrationParams {
	return &CreateAwsIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateAwsIntegrationParamsWithContext creates a new CreateAwsIntegrationParams object
// with the ability to set a context for a request.
func NewCreateAwsIntegrationParamsWithContext(ctx context.Context) *CreateAwsIntegrationParams {
	return &CreateAwsIntegrationParams{
		Context: ctx,
	}
}

// NewCreateAwsIntegrationParamsWithHTTPClient creates a new CreateAwsIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAwsIntegrationParamsWithHTTPClient(client *http.Client) *CreateAwsIntegrationParams {
	return &CreateAwsIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateAwsIntegrationParams contains all the parameters to send to the API endpoint

	for the create aws integration operation.

	Typically these are written to a http.Request.
*/
type CreateAwsIntegrationParams struct {

	// Body.
	Body *models.SecretServiceCreateAwsIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create aws integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAwsIntegrationParams) WithDefaults() *CreateAwsIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create aws integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAwsIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create aws integration params
func (o *CreateAwsIntegrationParams) WithTimeout(timeout time.Duration) *CreateAwsIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create aws integration params
func (o *CreateAwsIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create aws integration params
func (o *CreateAwsIntegrationParams) WithContext(ctx context.Context) *CreateAwsIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create aws integration params
func (o *CreateAwsIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create aws integration params
func (o *CreateAwsIntegrationParams) WithHTTPClient(client *http.Client) *CreateAwsIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create aws integration params
func (o *CreateAwsIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create aws integration params
func (o *CreateAwsIntegrationParams) WithBody(body *models.SecretServiceCreateAwsIntegrationBody) *CreateAwsIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create aws integration params
func (o *CreateAwsIntegrationParams) SetBody(body *models.SecretServiceCreateAwsIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create aws integration params
func (o *CreateAwsIntegrationParams) WithOrganizationID(organizationID string) *CreateAwsIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create aws integration params
func (o *CreateAwsIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create aws integration params
func (o *CreateAwsIntegrationParams) WithProjectID(projectID string) *CreateAwsIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create aws integration params
func (o *CreateAwsIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAwsIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
