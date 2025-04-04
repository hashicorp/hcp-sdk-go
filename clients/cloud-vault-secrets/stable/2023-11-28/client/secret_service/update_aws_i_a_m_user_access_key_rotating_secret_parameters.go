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

// NewUpdateAwsIAMUserAccessKeyRotatingSecretParams creates a new UpdateAwsIAMUserAccessKeyRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAwsIAMUserAccessKeyRotatingSecretParams() *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithTimeout creates a new UpdateAwsIAMUserAccessKeyRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithTimeout(timeout time.Duration) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretParams{
		timeout: timeout,
	}
}

// NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithContext creates a new UpdateAwsIAMUserAccessKeyRotatingSecretParams object
// with the ability to set a context for a request.
func NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithContext(ctx context.Context) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretParams{
		Context: ctx,
	}
}

// NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithHTTPClient creates a new UpdateAwsIAMUserAccessKeyRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAwsIAMUserAccessKeyRotatingSecretParamsWithHTTPClient(client *http.Client) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateAwsIAMUserAccessKeyRotatingSecretParams contains all the parameters to send to the API endpoint

	for the update aws i a m user access key rotating secret operation.

	Typically these are written to a http.Request.
*/
type UpdateAwsIAMUserAccessKeyRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceUpdateAwsIAMUserAccessKeyRotatingSecretBody

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

// WithDefaults hydrates default values in the update aws i a m user access key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithDefaults() *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update aws i a m user access key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithTimeout(timeout time.Duration) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithContext(ctx context.Context) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithHTTPClient(client *http.Client) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithAppName(appName string) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithBody(body *models.SecretServiceUpdateAwsIAMUserAccessKeyRotatingSecretBody) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetBody(body *models.SecretServiceUpdateAwsIAMUserAccessKeyRotatingSecretBody) {
	o.Body = body
}

// WithName adds the name to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithName(name string) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithOrganizationID(organizationID string) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WithProjectID(projectID string) *UpdateAwsIAMUserAccessKeyRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update aws i a m user access key rotating secret params
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
