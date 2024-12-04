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

// NewCreateRandomRotatingSecretParams creates a new CreateRandomRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRandomRotatingSecretParams() *CreateRandomRotatingSecretParams {
	return &CreateRandomRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRandomRotatingSecretParamsWithTimeout creates a new CreateRandomRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewCreateRandomRotatingSecretParamsWithTimeout(timeout time.Duration) *CreateRandomRotatingSecretParams {
	return &CreateRandomRotatingSecretParams{
		timeout: timeout,
	}
}

// NewCreateRandomRotatingSecretParamsWithContext creates a new CreateRandomRotatingSecretParams object
// with the ability to set a context for a request.
func NewCreateRandomRotatingSecretParamsWithContext(ctx context.Context) *CreateRandomRotatingSecretParams {
	return &CreateRandomRotatingSecretParams{
		Context: ctx,
	}
}

// NewCreateRandomRotatingSecretParamsWithHTTPClient creates a new CreateRandomRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRandomRotatingSecretParamsWithHTTPClient(client *http.Client) *CreateRandomRotatingSecretParams {
	return &CreateRandomRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
CreateRandomRotatingSecretParams contains all the parameters to send to the API endpoint

	for the create random rotating secret operation.

	Typically these are written to a http.Request.
*/
type CreateRandomRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceCreateRandomRotatingSecretBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create random rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRandomRotatingSecretParams) WithDefaults() *CreateRandomRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create random rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRandomRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithTimeout(timeout time.Duration) *CreateRandomRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithContext(ctx context.Context) *CreateRandomRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithHTTPClient(client *http.Client) *CreateRandomRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithAppName(appName string) *CreateRandomRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithBody(body *models.SecretServiceCreateRandomRotatingSecretBody) *CreateRandomRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetBody(body *models.SecretServiceCreateRandomRotatingSecretBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithOrganizationID(organizationID string) *CreateRandomRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) WithProjectID(projectID string) *CreateRandomRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create random rotating secret params
func (o *CreateRandomRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRandomRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}
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
