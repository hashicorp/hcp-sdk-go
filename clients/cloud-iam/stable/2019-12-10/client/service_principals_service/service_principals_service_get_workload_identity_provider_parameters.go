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

// NewServicePrincipalsServiceGetWorkloadIdentityProviderParams creates a new ServicePrincipalsServiceGetWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceGetWorkloadIdentityProviderParams() *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceGetWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithTimeout creates a new ServicePrincipalsServiceGetWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceGetWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithContext creates a new ServicePrincipalsServiceGetWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithContext(ctx context.Context) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceGetWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithHTTPClient creates a new ServicePrincipalsServiceGetWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceGetWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceGetWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceGetWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the service principals service get workload identity provider operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceGetWorkloadIdentityProviderParams struct {

	/* ResourceName2.

	     resource_name is the resource name of the workload identity provider to
	retrieve.
	*/
	ResourceName2 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service get workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WithDefaults() *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service get workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WithContext(ctx context.Context) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName2 adds the resourceName2 to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WithResourceName2(resourceName2 string) *ServicePrincipalsServiceGetWorkloadIdentityProviderParams {
	o.SetResourceName2(resourceName2)
	return o
}

// SetResourceName2 adds the resourceName2 to the service principals service get workload identity provider params
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) SetResourceName2(resourceName2 string) {
	o.ResourceName2 = resourceName2
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceGetWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name_2
	if err := r.SetPathParam("resource_name_2", o.ResourceName2); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
