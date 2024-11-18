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

// NewCreateAzureIntegrationParams creates a new CreateAzureIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAzureIntegrationParams() *CreateAzureIntegrationParams {
	return &CreateAzureIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureIntegrationParamsWithTimeout creates a new CreateAzureIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateAzureIntegrationParamsWithTimeout(timeout time.Duration) *CreateAzureIntegrationParams {
	return &CreateAzureIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateAzureIntegrationParamsWithContext creates a new CreateAzureIntegrationParams object
// with the ability to set a context for a request.
func NewCreateAzureIntegrationParamsWithContext(ctx context.Context) *CreateAzureIntegrationParams {
	return &CreateAzureIntegrationParams{
		Context: ctx,
	}
}

// NewCreateAzureIntegrationParamsWithHTTPClient creates a new CreateAzureIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAzureIntegrationParamsWithHTTPClient(client *http.Client) *CreateAzureIntegrationParams {
	return &CreateAzureIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateAzureIntegrationParams contains all the parameters to send to the API endpoint

	for the create azure integration operation.

	Typically these are written to a http.Request.
*/
type CreateAzureIntegrationParams struct {

	// Body.
	Body *models.SecretServiceCreateAzureIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create azure integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureIntegrationParams) WithDefaults() *CreateAzureIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create azure integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create azure integration params
func (o *CreateAzureIntegrationParams) WithTimeout(timeout time.Duration) *CreateAzureIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure integration params
func (o *CreateAzureIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure integration params
func (o *CreateAzureIntegrationParams) WithContext(ctx context.Context) *CreateAzureIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure integration params
func (o *CreateAzureIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure integration params
func (o *CreateAzureIntegrationParams) WithHTTPClient(client *http.Client) *CreateAzureIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure integration params
func (o *CreateAzureIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create azure integration params
func (o *CreateAzureIntegrationParams) WithBody(body *models.SecretServiceCreateAzureIntegrationBody) *CreateAzureIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create azure integration params
func (o *CreateAzureIntegrationParams) SetBody(body *models.SecretServiceCreateAzureIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create azure integration params
func (o *CreateAzureIntegrationParams) WithOrganizationID(organizationID string) *CreateAzureIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create azure integration params
func (o *CreateAzureIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create azure integration params
func (o *CreateAzureIntegrationParams) WithProjectID(projectID string) *CreateAzureIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create azure integration params
func (o *CreateAzureIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
