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

// NewCreateAppKVSecretParams creates a new CreateAppKVSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppKVSecretParams() *CreateAppKVSecretParams {
	return &CreateAppKVSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppKVSecretParamsWithTimeout creates a new CreateAppKVSecretParams object
// with the ability to set a timeout on a request.
func NewCreateAppKVSecretParamsWithTimeout(timeout time.Duration) *CreateAppKVSecretParams {
	return &CreateAppKVSecretParams{
		timeout: timeout,
	}
}

// NewCreateAppKVSecretParamsWithContext creates a new CreateAppKVSecretParams object
// with the ability to set a context for a request.
func NewCreateAppKVSecretParamsWithContext(ctx context.Context) *CreateAppKVSecretParams {
	return &CreateAppKVSecretParams{
		Context: ctx,
	}
}

// NewCreateAppKVSecretParamsWithHTTPClient creates a new CreateAppKVSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppKVSecretParamsWithHTTPClient(client *http.Client) *CreateAppKVSecretParams {
	return &CreateAppKVSecretParams{
		HTTPClient: client,
	}
}

/*
CreateAppKVSecretParams contains all the parameters to send to the API endpoint

	for the create app k v secret operation.

	Typically these are written to a http.Request.
*/
type CreateAppKVSecretParams struct {

	// AppName.
	AppName string

	// Body.
	Body CreateAppKVSecretBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app k v secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppKVSecretParams) WithDefaults() *CreateAppKVSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app k v secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppKVSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app k v secret params
func (o *CreateAppKVSecretParams) WithTimeout(timeout time.Duration) *CreateAppKVSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app k v secret params
func (o *CreateAppKVSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app k v secret params
func (o *CreateAppKVSecretParams) WithContext(ctx context.Context) *CreateAppKVSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app k v secret params
func (o *CreateAppKVSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app k v secret params
func (o *CreateAppKVSecretParams) WithHTTPClient(client *http.Client) *CreateAppKVSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app k v secret params
func (o *CreateAppKVSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the create app k v secret params
func (o *CreateAppKVSecretParams) WithAppName(appName string) *CreateAppKVSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create app k v secret params
func (o *CreateAppKVSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create app k v secret params
func (o *CreateAppKVSecretParams) WithBody(body CreateAppKVSecretBody) *CreateAppKVSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create app k v secret params
func (o *CreateAppKVSecretParams) SetBody(body CreateAppKVSecretBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the create app k v secret params
func (o *CreateAppKVSecretParams) WithLocationOrganizationID(locationOrganizationID string) *CreateAppKVSecretParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the create app k v secret params
func (o *CreateAppKVSecretParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the create app k v secret params
func (o *CreateAppKVSecretParams) WithLocationProjectID(locationProjectID string) *CreateAppKVSecretParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the create app k v secret params
func (o *CreateAppKVSecretParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppKVSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
