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

// NewOpenAppSecretParams creates a new OpenAppSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenAppSecretParams() *OpenAppSecretParams {
	return &OpenAppSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenAppSecretParamsWithTimeout creates a new OpenAppSecretParams object
// with the ability to set a timeout on a request.
func NewOpenAppSecretParamsWithTimeout(timeout time.Duration) *OpenAppSecretParams {
	return &OpenAppSecretParams{
		timeout: timeout,
	}
}

// NewOpenAppSecretParamsWithContext creates a new OpenAppSecretParams object
// with the ability to set a context for a request.
func NewOpenAppSecretParamsWithContext(ctx context.Context) *OpenAppSecretParams {
	return &OpenAppSecretParams{
		Context: ctx,
	}
}

// NewOpenAppSecretParamsWithHTTPClient creates a new OpenAppSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenAppSecretParamsWithHTTPClient(client *http.Client) *OpenAppSecretParams {
	return &OpenAppSecretParams{
		HTTPClient: client,
	}
}

/*
OpenAppSecretParams contains all the parameters to send to the API endpoint

	for the open app secret operation.

	Typically these are written to a http.Request.
*/
type OpenAppSecretParams struct {

	// AppName.
	AppName string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	// SecretName.
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the open app secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAppSecretParams) WithDefaults() *OpenAppSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the open app secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAppSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the open app secret params
func (o *OpenAppSecretParams) WithTimeout(timeout time.Duration) *OpenAppSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open app secret params
func (o *OpenAppSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open app secret params
func (o *OpenAppSecretParams) WithContext(ctx context.Context) *OpenAppSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open app secret params
func (o *OpenAppSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open app secret params
func (o *OpenAppSecretParams) WithHTTPClient(client *http.Client) *OpenAppSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open app secret params
func (o *OpenAppSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the open app secret params
func (o *OpenAppSecretParams) WithAppName(appName string) *OpenAppSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the open app secret params
func (o *OpenAppSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithOrganizationID adds the organizationID to the open app secret params
func (o *OpenAppSecretParams) WithOrganizationID(organizationID string) *OpenAppSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the open app secret params
func (o *OpenAppSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the open app secret params
func (o *OpenAppSecretParams) WithProjectID(projectID string) *OpenAppSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the open app secret params
func (o *OpenAppSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSecretName adds the secretName to the open app secret params
func (o *OpenAppSecretParams) WithSecretName(secretName string) *OpenAppSecretParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the open app secret params
func (o *OpenAppSecretParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *OpenAppSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
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

	// path param secret_name
	if err := r.SetPathParam("secret_name", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
