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

// NewServicePrincipalsServiceCreateServicePrincipalKeyParams creates a new ServicePrincipalsServiceCreateServicePrincipalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceCreateServicePrincipalKeyParams() *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	return &ServicePrincipalsServiceCreateServicePrincipalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithTimeout creates a new ServicePrincipalsServiceCreateServicePrincipalKeyParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	return &ServicePrincipalsServiceCreateServicePrincipalKeyParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithContext creates a new ServicePrincipalsServiceCreateServicePrincipalKeyParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithContext(ctx context.Context) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	return &ServicePrincipalsServiceCreateServicePrincipalKeyParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithHTTPClient creates a new ServicePrincipalsServiceCreateServicePrincipalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceCreateServicePrincipalKeyParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	return &ServicePrincipalsServiceCreateServicePrincipalKeyParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceCreateServicePrincipalKeyParams contains all the parameters to send to the API endpoint

	for the service principals service create service principal key operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceCreateServicePrincipalKeyParams struct {

	// Body.
	Body interface{}

	/* ParentResourceName.

	     parent_resource_name is the resource name of the service principal to
	generate a key for.
	*/
	ParentResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service create service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithDefaults() *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service create service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithContext(ctx context.Context) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithBody(body interface{}) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetBody(body interface{}) {
	o.Body = body
}

// WithParentResourceName adds the parentResourceName to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WithParentResourceName(parentResourceName string) *ServicePrincipalsServiceCreateServicePrincipalKeyParams {
	o.SetParentResourceName(parentResourceName)
	return o
}

// SetParentResourceName adds the parentResourceName to the service principals service create service principal key params
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) SetParentResourceName(parentResourceName string) {
	o.ParentResourceName = parentResourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceCreateServicePrincipalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param parent_resource_name
	if err := r.SetPathParam("parent_resource_name", o.ParentResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
