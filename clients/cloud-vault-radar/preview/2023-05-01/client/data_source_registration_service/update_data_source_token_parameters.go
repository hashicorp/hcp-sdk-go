// Code generated by go-swagger; DO NOT EDIT.

package data_source_registration_service

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

// NewUpdateDataSourceTokenParams creates a new UpdateDataSourceTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDataSourceTokenParams() *UpdateDataSourceTokenParams {
	return &UpdateDataSourceTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataSourceTokenParamsWithTimeout creates a new UpdateDataSourceTokenParams object
// with the ability to set a timeout on a request.
func NewUpdateDataSourceTokenParamsWithTimeout(timeout time.Duration) *UpdateDataSourceTokenParams {
	return &UpdateDataSourceTokenParams{
		timeout: timeout,
	}
}

// NewUpdateDataSourceTokenParamsWithContext creates a new UpdateDataSourceTokenParams object
// with the ability to set a context for a request.
func NewUpdateDataSourceTokenParamsWithContext(ctx context.Context) *UpdateDataSourceTokenParams {
	return &UpdateDataSourceTokenParams{
		Context: ctx,
	}
}

// NewUpdateDataSourceTokenParamsWithHTTPClient creates a new UpdateDataSourceTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDataSourceTokenParamsWithHTTPClient(client *http.Client) *UpdateDataSourceTokenParams {
	return &UpdateDataSourceTokenParams{
		HTTPClient: client,
	}
}

/*
UpdateDataSourceTokenParams contains all the parameters to send to the API endpoint

	for the update data source token operation.

	Typically these are written to a http.Request.
*/
type UpdateDataSourceTokenParams struct {

	// Body.
	Body UpdateDataSourceTokenBody

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update data source token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataSourceTokenParams) WithDefaults() *UpdateDataSourceTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update data source token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataSourceTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update data source token params
func (o *UpdateDataSourceTokenParams) WithTimeout(timeout time.Duration) *UpdateDataSourceTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data source token params
func (o *UpdateDataSourceTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data source token params
func (o *UpdateDataSourceTokenParams) WithContext(ctx context.Context) *UpdateDataSourceTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data source token params
func (o *UpdateDataSourceTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data source token params
func (o *UpdateDataSourceTokenParams) WithHTTPClient(client *http.Client) *UpdateDataSourceTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data source token params
func (o *UpdateDataSourceTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update data source token params
func (o *UpdateDataSourceTokenParams) WithBody(body UpdateDataSourceTokenBody) *UpdateDataSourceTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update data source token params
func (o *UpdateDataSourceTokenParams) SetBody(body UpdateDataSourceTokenBody) {
	o.Body = body
}

// WithLocationProjectID adds the locationProjectID to the update data source token params
func (o *UpdateDataSourceTokenParams) WithLocationProjectID(locationProjectID string) *UpdateDataSourceTokenParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update data source token params
func (o *UpdateDataSourceTokenParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataSourceTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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