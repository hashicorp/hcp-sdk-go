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

// NewGetAzureApplicationPasswordRotatingSecretConfigParams creates a new GetAzureApplicationPasswordRotatingSecretConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAzureApplicationPasswordRotatingSecretConfigParams() *GetAzureApplicationPasswordRotatingSecretConfigParams {
	return &GetAzureApplicationPasswordRotatingSecretConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithTimeout creates a new GetAzureApplicationPasswordRotatingSecretConfigParams object
// with the ability to set a timeout on a request.
func NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithTimeout(timeout time.Duration) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	return &GetAzureApplicationPasswordRotatingSecretConfigParams{
		timeout: timeout,
	}
}

// NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithContext creates a new GetAzureApplicationPasswordRotatingSecretConfigParams object
// with the ability to set a context for a request.
func NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithContext(ctx context.Context) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	return &GetAzureApplicationPasswordRotatingSecretConfigParams{
		Context: ctx,
	}
}

// NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithHTTPClient creates a new GetAzureApplicationPasswordRotatingSecretConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAzureApplicationPasswordRotatingSecretConfigParamsWithHTTPClient(client *http.Client) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	return &GetAzureApplicationPasswordRotatingSecretConfigParams{
		HTTPClient: client,
	}
}

/*
GetAzureApplicationPasswordRotatingSecretConfigParams contains all the parameters to send to the API endpoint

	for the get azure application password rotating secret config operation.

	Typically these are written to a http.Request.
*/
type GetAzureApplicationPasswordRotatingSecretConfigParams struct {

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

// WithDefaults hydrates default values in the get azure application password rotating secret config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithDefaults() *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get azure application password rotating secret config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithTimeout(timeout time.Duration) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithContext(ctx context.Context) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithHTTPClient(client *http.Client) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithAppName(appName string) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithName adds the name to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithName(name string) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithOrganizationID(organizationID string) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WithProjectID(projectID string) *GetAzureApplicationPasswordRotatingSecretConfigParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get azure application password rotating secret config params
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureApplicationPasswordRotatingSecretConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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