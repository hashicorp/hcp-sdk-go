// Code generated by go-swagger; DO NOT EDIT.

package tenant_service

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

// NewProvisionTenantParams creates a new ProvisionTenantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionTenantParams() *ProvisionTenantParams {
	return &ProvisionTenantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionTenantParamsWithTimeout creates a new ProvisionTenantParams object
// with the ability to set a timeout on a request.
func NewProvisionTenantParamsWithTimeout(timeout time.Duration) *ProvisionTenantParams {
	return &ProvisionTenantParams{
		timeout: timeout,
	}
}

// NewProvisionTenantParamsWithContext creates a new ProvisionTenantParams object
// with the ability to set a context for a request.
func NewProvisionTenantParamsWithContext(ctx context.Context) *ProvisionTenantParams {
	return &ProvisionTenantParams{
		Context: ctx,
	}
}

// NewProvisionTenantParamsWithHTTPClient creates a new ProvisionTenantParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionTenantParamsWithHTTPClient(client *http.Client) *ProvisionTenantParams {
	return &ProvisionTenantParams{
		HTTPClient: client,
	}
}

/*
ProvisionTenantParams contains all the parameters to send to the API endpoint

	for the provision tenant operation.

	Typically these are written to a http.Request.
*/
type ProvisionTenantParams struct {

	// Body.
	Body ProvisionTenantBody

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provision tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionTenantParams) WithDefaults() *ProvisionTenantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provision tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionTenantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provision tenant params
func (o *ProvisionTenantParams) WithTimeout(timeout time.Duration) *ProvisionTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provision tenant params
func (o *ProvisionTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provision tenant params
func (o *ProvisionTenantParams) WithContext(ctx context.Context) *ProvisionTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provision tenant params
func (o *ProvisionTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provision tenant params
func (o *ProvisionTenantParams) WithHTTPClient(client *http.Client) *ProvisionTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provision tenant params
func (o *ProvisionTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the provision tenant params
func (o *ProvisionTenantParams) WithBody(body ProvisionTenantBody) *ProvisionTenantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the provision tenant params
func (o *ProvisionTenantParams) SetBody(body ProvisionTenantBody) {
	o.Body = body
}

// WithLocationProjectID adds the locationProjectID to the provision tenant params
func (o *ProvisionTenantParams) WithLocationProjectID(locationProjectID string) *ProvisionTenantParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the provision tenant params
func (o *ProvisionTenantParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
