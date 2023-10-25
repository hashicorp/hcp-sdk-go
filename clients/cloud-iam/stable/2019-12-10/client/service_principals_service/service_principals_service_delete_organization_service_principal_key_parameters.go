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

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams creates a new ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams() *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithTimeout creates a new ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithContext creates a new ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithContext(ctx context.Context) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithHTTPClient creates a new ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams contains all the parameters to send to the API endpoint

	for the service principals service delete organization service principal key operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams struct {

	/* ClientID.

	   client_id is client ID of the service principal key.
	*/
	ClientID string

	/* OrganizationID.

	   organization_id is the organization the service principal key belongs to.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service delete organization service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithDefaults() *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service delete organization service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithContext(ctx context.Context) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithClientID(clientID string) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithOrganizationID adds the organizationID to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WithOrganizationID(organizationID string) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the service principals service delete organization service principal key params
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}