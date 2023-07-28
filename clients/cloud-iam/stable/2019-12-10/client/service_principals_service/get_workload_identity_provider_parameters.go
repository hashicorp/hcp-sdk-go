// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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

// NewGetWorkloadIdentityProviderParams creates a new GetWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkloadIdentityProviderParams() *GetWorkloadIdentityProviderParams {
	return &GetWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkloadIdentityProviderParamsWithTimeout creates a new GetWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewGetWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *GetWorkloadIdentityProviderParams {
	return &GetWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewGetWorkloadIdentityProviderParamsWithContext creates a new GetWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewGetWorkloadIdentityProviderParamsWithContext(ctx context.Context) *GetWorkloadIdentityProviderParams {
	return &GetWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewGetWorkloadIdentityProviderParamsWithHTTPClient creates a new GetWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *GetWorkloadIdentityProviderParams {
	return &GetWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
GetWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the get workload identity provider operation.

	Typically these are written to a http.Request.
*/
type GetWorkloadIdentityProviderParams struct {

	/* ResourceName.

	     resource_name is the resource name of the workload identity provider to
	retrieve.
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkloadIdentityProviderParams) WithDefaults() *GetWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *GetWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) WithContext(ctx context.Context) *GetWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *GetWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) WithResourceName(resourceName string) *GetWorkloadIdentityProviderParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get workload identity provider params
func (o *GetWorkloadIdentityProviderParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name
	if err := r.SetPathParam("resource_name", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
