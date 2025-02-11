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

// NewCreateRandomIntegrationParams creates a new CreateRandomIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRandomIntegrationParams() *CreateRandomIntegrationParams {
	return &CreateRandomIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRandomIntegrationParamsWithTimeout creates a new CreateRandomIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateRandomIntegrationParamsWithTimeout(timeout time.Duration) *CreateRandomIntegrationParams {
	return &CreateRandomIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateRandomIntegrationParamsWithContext creates a new CreateRandomIntegrationParams object
// with the ability to set a context for a request.
func NewCreateRandomIntegrationParamsWithContext(ctx context.Context) *CreateRandomIntegrationParams {
	return &CreateRandomIntegrationParams{
		Context: ctx,
	}
}

// NewCreateRandomIntegrationParamsWithHTTPClient creates a new CreateRandomIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRandomIntegrationParamsWithHTTPClient(client *http.Client) *CreateRandomIntegrationParams {
	return &CreateRandomIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateRandomIntegrationParams contains all the parameters to send to the API endpoint

	for the create random integration operation.

	Typically these are written to a http.Request.
*/
type CreateRandomIntegrationParams struct {

	// Body.
	Body *models.SecretServiceCreateRandomIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create random integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRandomIntegrationParams) WithDefaults() *CreateRandomIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create random integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRandomIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create random integration params
func (o *CreateRandomIntegrationParams) WithTimeout(timeout time.Duration) *CreateRandomIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create random integration params
func (o *CreateRandomIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create random integration params
func (o *CreateRandomIntegrationParams) WithContext(ctx context.Context) *CreateRandomIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create random integration params
func (o *CreateRandomIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create random integration params
func (o *CreateRandomIntegrationParams) WithHTTPClient(client *http.Client) *CreateRandomIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create random integration params
func (o *CreateRandomIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create random integration params
func (o *CreateRandomIntegrationParams) WithBody(body *models.SecretServiceCreateRandomIntegrationBody) *CreateRandomIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create random integration params
func (o *CreateRandomIntegrationParams) SetBody(body *models.SecretServiceCreateRandomIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create random integration params
func (o *CreateRandomIntegrationParams) WithOrganizationID(organizationID string) *CreateRandomIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create random integration params
func (o *CreateRandomIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create random integration params
func (o *CreateRandomIntegrationParams) WithProjectID(projectID string) *CreateRandomIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create random integration params
func (o *CreateRandomIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRandomIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
