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

// NewGetAppDynamicSecretParams creates a new GetAppDynamicSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppDynamicSecretParams() *GetAppDynamicSecretParams {
	return &GetAppDynamicSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppDynamicSecretParamsWithTimeout creates a new GetAppDynamicSecretParams object
// with the ability to set a timeout on a request.
func NewGetAppDynamicSecretParamsWithTimeout(timeout time.Duration) *GetAppDynamicSecretParams {
	return &GetAppDynamicSecretParams{
		timeout: timeout,
	}
}

// NewGetAppDynamicSecretParamsWithContext creates a new GetAppDynamicSecretParams object
// with the ability to set a context for a request.
func NewGetAppDynamicSecretParamsWithContext(ctx context.Context) *GetAppDynamicSecretParams {
	return &GetAppDynamicSecretParams{
		Context: ctx,
	}
}

// NewGetAppDynamicSecretParamsWithHTTPClient creates a new GetAppDynamicSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppDynamicSecretParamsWithHTTPClient(client *http.Client) *GetAppDynamicSecretParams {
	return &GetAppDynamicSecretParams{
		HTTPClient: client,
	}
}

/*
GetAppDynamicSecretParams contains all the parameters to send to the API endpoint

	for the get app dynamic secret operation.

	Typically these are written to a http.Request.
*/
type GetAppDynamicSecretParams struct {

	// AppName.
	AppName string

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

// WithDefaults hydrates default values in the get app dynamic secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppDynamicSecretParams) WithDefaults() *GetAppDynamicSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app dynamic secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppDynamicSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithTimeout(timeout time.Duration) *GetAppDynamicSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithContext(ctx context.Context) *GetAppDynamicSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithHTTPClient(client *http.Client) *GetAppDynamicSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithAppName(appName string) *GetAppDynamicSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithName adds the name to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithName(name string) *GetAppDynamicSecretParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithOrganizationID(organizationID string) *GetAppDynamicSecretParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) WithProjectID(projectID string) *GetAppDynamicSecretParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get app dynamic secret params
func (o *GetAppDynamicSecretParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppDynamicSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
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
