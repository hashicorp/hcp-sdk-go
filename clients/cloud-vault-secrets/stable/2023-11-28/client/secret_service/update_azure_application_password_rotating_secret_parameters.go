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

// NewUpdateAzureApplicationPasswordRotatingSecretParams creates a new UpdateAzureApplicationPasswordRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAzureApplicationPasswordRotatingSecretParams() *UpdateAzureApplicationPasswordRotatingSecretParams {
	return &UpdateAzureApplicationPasswordRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAzureApplicationPasswordRotatingSecretParamsWithTimeout creates a new UpdateAzureApplicationPasswordRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateAzureApplicationPasswordRotatingSecretParamsWithTimeout(timeout time.Duration) *UpdateAzureApplicationPasswordRotatingSecretParams {
	return &UpdateAzureApplicationPasswordRotatingSecretParams{
		timeout: timeout,
	}
}

// NewUpdateAzureApplicationPasswordRotatingSecretParamsWithContext creates a new UpdateAzureApplicationPasswordRotatingSecretParams object
// with the ability to set a context for a request.
func NewUpdateAzureApplicationPasswordRotatingSecretParamsWithContext(ctx context.Context) *UpdateAzureApplicationPasswordRotatingSecretParams {
	return &UpdateAzureApplicationPasswordRotatingSecretParams{
		Context: ctx,
	}
}

// NewUpdateAzureApplicationPasswordRotatingSecretParamsWithHTTPClient creates a new UpdateAzureApplicationPasswordRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAzureApplicationPasswordRotatingSecretParamsWithHTTPClient(client *http.Client) *UpdateAzureApplicationPasswordRotatingSecretParams {
	return &UpdateAzureApplicationPasswordRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateAzureApplicationPasswordRotatingSecretParams contains all the parameters to send to the API endpoint

	for the update azure application password rotating secret operation.

	Typically these are written to a http.Request.
*/
type UpdateAzureApplicationPasswordRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceUpdateAzureApplicationPasswordRotatingSecretBody

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

// WithDefaults hydrates default values in the update azure application password rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithDefaults() *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update azure application password rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithTimeout(timeout time.Duration) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithContext(ctx context.Context) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithHTTPClient(client *http.Client) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithAppName(appName string) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithBody(body *models.SecretServiceUpdateAzureApplicationPasswordRotatingSecretBody) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetBody(body *models.SecretServiceUpdateAzureApplicationPasswordRotatingSecretBody) {
	o.Body = body
}

// WithName adds the name to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithName(name string) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithOrganizationID(organizationID string) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WithProjectID(projectID string) *UpdateAzureApplicationPasswordRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update azure application password rotating secret params
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAzureApplicationPasswordRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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