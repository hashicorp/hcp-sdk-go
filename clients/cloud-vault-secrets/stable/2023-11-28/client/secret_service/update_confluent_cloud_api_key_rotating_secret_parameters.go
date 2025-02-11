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

// NewUpdateConfluentCloudAPIKeyRotatingSecretParams creates a new UpdateConfluentCloudAPIKeyRotatingSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConfluentCloudAPIKeyRotatingSecretParams() *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	return &UpdateConfluentCloudAPIKeyRotatingSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithTimeout creates a new UpdateConfluentCloudAPIKeyRotatingSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithTimeout(timeout time.Duration) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	return &UpdateConfluentCloudAPIKeyRotatingSecretParams{
		timeout: timeout,
	}
}

// NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithContext creates a new UpdateConfluentCloudAPIKeyRotatingSecretParams object
// with the ability to set a context for a request.
func NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithContext(ctx context.Context) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	return &UpdateConfluentCloudAPIKeyRotatingSecretParams{
		Context: ctx,
	}
}

// NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithHTTPClient creates a new UpdateConfluentCloudAPIKeyRotatingSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConfluentCloudAPIKeyRotatingSecretParamsWithHTTPClient(client *http.Client) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	return &UpdateConfluentCloudAPIKeyRotatingSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateConfluentCloudAPIKeyRotatingSecretParams contains all the parameters to send to the API endpoint

	for the update confluent cloud Api key rotating secret operation.

	Typically these are written to a http.Request.
*/
type UpdateConfluentCloudAPIKeyRotatingSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body *models.SecretServiceUpdateConfluentCloudAPIKeyRotatingSecretBody

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

// WithDefaults hydrates default values in the update confluent cloud Api key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithDefaults() *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update confluent cloud Api key rotating secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithTimeout(timeout time.Duration) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithContext(ctx context.Context) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithHTTPClient(client *http.Client) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithAppName(appName string) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithBody(body *models.SecretServiceUpdateConfluentCloudAPIKeyRotatingSecretBody) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetBody(body *models.SecretServiceUpdateConfluentCloudAPIKeyRotatingSecretBody) {
	o.Body = body
}

// WithName adds the name to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithName(name string) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithOrganizationID(organizationID string) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WithProjectID(projectID string) *UpdateConfluentCloudAPIKeyRotatingSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update confluent cloud Api key rotating secret params
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfluentCloudAPIKeyRotatingSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
