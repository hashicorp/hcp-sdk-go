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

// NewServicePrincipalsServiceCreateServicePrincipalKey2Params creates a new ServicePrincipalsServiceCreateServicePrincipalKey2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceCreateServicePrincipalKey2Params() *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithTimeout creates a new ServicePrincipalsServiceCreateServicePrincipalKey2Params object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2Params{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithContext creates a new ServicePrincipalsServiceCreateServicePrincipalKey2Params object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithContext(ctx context.Context) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2Params{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithHTTPClient creates a new ServicePrincipalsServiceCreateServicePrincipalKey2Params object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceCreateServicePrincipalKey2ParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2Params{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceCreateServicePrincipalKey2Params contains all the parameters to send to the API endpoint

	for the service principals service create service principal key2 operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceCreateServicePrincipalKey2Params struct {

	// Body.
	Body interface{}

	/* ParentResourceName1.

	     parent_resource_name is the resource name of the service principal to
	generate a key for.
	*/
	ParentResourceName1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service create service principal key2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithDefaults() *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service create service principal key2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithContext(ctx context.Context) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithBody(body interface{}) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetBody(body interface{}) {
	o.Body = body
}

// WithParentResourceName1 adds the parentResourceName1 to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WithParentResourceName1(parentResourceName1 string) *ServicePrincipalsServiceCreateServicePrincipalKey2Params {
	o.SetParentResourceName1(parentResourceName1)
	return o
}

// SetParentResourceName1 adds the parentResourceName1 to the service principals service create service principal key2 params
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) SetParentResourceName1(parentResourceName1 string) {
	o.ParentResourceName1 = parentResourceName1
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param parent_resource_name_1
	if err := r.SetPathParam("parent_resource_name_1", o.ParentResourceName1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}