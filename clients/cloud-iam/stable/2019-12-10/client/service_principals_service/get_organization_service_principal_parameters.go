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

// NewGetOrganizationServicePrincipalParams creates a new GetOrganizationServicePrincipalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationServicePrincipalParams() *GetOrganizationServicePrincipalParams {
	return &GetOrganizationServicePrincipalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationServicePrincipalParamsWithTimeout creates a new GetOrganizationServicePrincipalParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationServicePrincipalParamsWithTimeout(timeout time.Duration) *GetOrganizationServicePrincipalParams {
	return &GetOrganizationServicePrincipalParams{
		timeout: timeout,
	}
}

// NewGetOrganizationServicePrincipalParamsWithContext creates a new GetOrganizationServicePrincipalParams object
// with the ability to set a context for a request.
func NewGetOrganizationServicePrincipalParamsWithContext(ctx context.Context) *GetOrganizationServicePrincipalParams {
	return &GetOrganizationServicePrincipalParams{
		Context: ctx,
	}
}

// NewGetOrganizationServicePrincipalParamsWithHTTPClient creates a new GetOrganizationServicePrincipalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationServicePrincipalParamsWithHTTPClient(client *http.Client) *GetOrganizationServicePrincipalParams {
	return &GetOrganizationServicePrincipalParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationServicePrincipalParams contains all the parameters to send to the API endpoint

	for the get organization service principal operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationServicePrincipalParams struct {

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization that
	owns the requested service principal or project.
	*/
	OrganizationID string

	/* PrincipalID.

	   principal_id is the ID of the service principal being requested.
	*/
	PrincipalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationServicePrincipalParams) WithDefaults() *GetOrganizationServicePrincipalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization service principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationServicePrincipalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) WithTimeout(timeout time.Duration) *GetOrganizationServicePrincipalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) WithContext(ctx context.Context) *GetOrganizationServicePrincipalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) WithHTTPClient(client *http.Client) *GetOrganizationServicePrincipalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) WithOrganizationID(organizationID string) *GetOrganizationServicePrincipalParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPrincipalID adds the principalID to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) WithPrincipalID(principalID string) *GetOrganizationServicePrincipalParams {
	o.SetPrincipalID(principalID)
	return o
}

// SetPrincipalID adds the principalId to the get organization service principal params
func (o *GetOrganizationServicePrincipalParams) SetPrincipalID(principalID string) {
	o.PrincipalID = principalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationServicePrincipalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param principal_id
	if err := r.SetPathParam("principal_id", o.PrincipalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
