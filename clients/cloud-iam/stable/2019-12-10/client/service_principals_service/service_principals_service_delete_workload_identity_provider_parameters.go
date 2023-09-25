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

// NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParams creates a new ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParams() *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithTimeout creates a new ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithContext creates a new ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithContext(ctx context.Context) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithHTTPClient creates a new ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceDeleteWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the service principals service delete workload identity provider operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams struct {

	/* ResourceName.

	     resource_name is the resource name of the workload identity provider to
	delete.
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service delete workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WithDefaults() *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service delete workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WithContext(ctx context.Context) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WithResourceName(resourceName string) *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the service principals service delete workload identity provider params
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceDeleteWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
