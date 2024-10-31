// Code generated by go-swagger; DO NOT EDIT.

package integration_connection_service

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

// NewCreateIntegrationConnectionParams creates a new CreateIntegrationConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIntegrationConnectionParams() *CreateIntegrationConnectionParams {
	return &CreateIntegrationConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIntegrationConnectionParamsWithTimeout creates a new CreateIntegrationConnectionParams object
// with the ability to set a timeout on a request.
func NewCreateIntegrationConnectionParamsWithTimeout(timeout time.Duration) *CreateIntegrationConnectionParams {
	return &CreateIntegrationConnectionParams{
		timeout: timeout,
	}
}

// NewCreateIntegrationConnectionParamsWithContext creates a new CreateIntegrationConnectionParams object
// with the ability to set a context for a request.
func NewCreateIntegrationConnectionParamsWithContext(ctx context.Context) *CreateIntegrationConnectionParams {
	return &CreateIntegrationConnectionParams{
		Context: ctx,
	}
}

// NewCreateIntegrationConnectionParamsWithHTTPClient creates a new CreateIntegrationConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIntegrationConnectionParamsWithHTTPClient(client *http.Client) *CreateIntegrationConnectionParams {
	return &CreateIntegrationConnectionParams{
		HTTPClient: client,
	}
}

/*
CreateIntegrationConnectionParams contains all the parameters to send to the API endpoint

	for the create integration connection operation.

	Typically these are written to a http.Request.
*/
type CreateIntegrationConnectionParams struct {

	// Body.
	Body CreateIntegrationConnectionBody

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create integration connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIntegrationConnectionParams) WithDefaults() *CreateIntegrationConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create integration connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIntegrationConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create integration connection params
func (o *CreateIntegrationConnectionParams) WithTimeout(timeout time.Duration) *CreateIntegrationConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create integration connection params
func (o *CreateIntegrationConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create integration connection params
func (o *CreateIntegrationConnectionParams) WithContext(ctx context.Context) *CreateIntegrationConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create integration connection params
func (o *CreateIntegrationConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create integration connection params
func (o *CreateIntegrationConnectionParams) WithHTTPClient(client *http.Client) *CreateIntegrationConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create integration connection params
func (o *CreateIntegrationConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create integration connection params
func (o *CreateIntegrationConnectionParams) WithBody(body CreateIntegrationConnectionBody) *CreateIntegrationConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create integration connection params
func (o *CreateIntegrationConnectionParams) SetBody(body CreateIntegrationConnectionBody) {
	o.Body = body
}

// WithLocationProjectID adds the locationProjectID to the create integration connection params
func (o *CreateIntegrationConnectionParams) WithLocationProjectID(locationProjectID string) *CreateIntegrationConnectionParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the create integration connection params
func (o *CreateIntegrationConnectionParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIntegrationConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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