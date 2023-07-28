// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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

// NewGetOrganizationAuthMetadataParams creates a new GetOrganizationAuthMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAuthMetadataParams() *GetOrganizationAuthMetadataParams {
	return &GetOrganizationAuthMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAuthMetadataParamsWithTimeout creates a new GetOrganizationAuthMetadataParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAuthMetadataParamsWithTimeout(timeout time.Duration) *GetOrganizationAuthMetadataParams {
	return &GetOrganizationAuthMetadataParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAuthMetadataParamsWithContext creates a new GetOrganizationAuthMetadataParams object
// with the ability to set a context for a request.
func NewGetOrganizationAuthMetadataParamsWithContext(ctx context.Context) *GetOrganizationAuthMetadataParams {
	return &GetOrganizationAuthMetadataParams{
		Context: ctx,
	}
}

// NewGetOrganizationAuthMetadataParamsWithHTTPClient creates a new GetOrganizationAuthMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAuthMetadataParamsWithHTTPClient(client *http.Client) *GetOrganizationAuthMetadataParams {
	return &GetOrganizationAuthMetadataParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationAuthMetadataParams contains all the parameters to send to the API endpoint

	for the get organization auth metadata operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationAuthMetadataParams struct {

	/* OrganizationID.

	     organization_id is the UUID identifier of the organization for which
	authentication metadata will be retrieved.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization auth metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAuthMetadataParams) WithDefaults() *GetOrganizationAuthMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization auth metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAuthMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) WithTimeout(timeout time.Duration) *GetOrganizationAuthMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) WithContext(ctx context.Context) *GetOrganizationAuthMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) WithHTTPClient(client *http.Client) *GetOrganizationAuthMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) WithOrganizationID(organizationID string) *GetOrganizationAuthMetadataParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization auth metadata params
func (o *GetOrganizationAuthMetadataParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAuthMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
