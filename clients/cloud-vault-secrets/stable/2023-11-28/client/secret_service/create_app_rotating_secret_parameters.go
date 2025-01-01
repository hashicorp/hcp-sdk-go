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

// NewCreateAppRotatingSecretParams creates a new CreateAppRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppRotatingSecretParams() *CreateAppRotatingSecretParams {
	return &CreateAppRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppRotatingSecretParamsWithTimeout creates a new CreateAppRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewCreateAppRotatingSecretParamsWithTimeout(timeout time.Duration) *CreateAppRotatingSecretParams {
	return &CreateAppRotatingSecretParams{
		timeout: timeout,
	}
}

// NewCreateAppRotatingSecretParamsWithContext creates a new CreateAppRotatingSecretParams object
// with the ability to set a context for a request.
func NewCreateAppRotatingSecretParamsWithContext(ctx context.Context) *CreateAppRotatingSecretParams {
	return &CreateAppRotatingSecretParams{
		Context: ctx,
	}
}

// NewCreateAppRotatingSecretParamsWithHTTPClient creates a new CreateAppRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppRotatingSecretParamsWithHTTPClient(client *http.Client) *CreateAppRotatingSecretParams {
	return &CreateAppRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
CreateAppRotatingSecretParams contains all the parameters to send to the API endpoint

	for the create app rotating secret operation.

	Typically these are written to a http.Request.
*/
type CreateAppRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceCreateAppRotatingSecretBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppRotatingSecretParams) WithDefaults() *CreateAppRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithTimeout(timeout time.Duration) *CreateAppRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithContext(ctx context.Context) *CreateAppRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithHTTPClient(client *http.Client) *CreateAppRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithAppName(appName string) *CreateAppRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithBody(body *models.SecretServiceCreateAppRotatingSecretBody) *CreateAppRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetBody(body *models.SecretServiceCreateAppRotatingSecretBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithOrganizationID(organizationID string) *CreateAppRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) WithProjectID(projectID string) *CreateAppRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create app rotating secret params
func (o *CreateAppRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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