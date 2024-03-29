// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

// NewPackerServiceRegenerateTFCRunTaskHmacKeyParams creates a new PackerServiceRegenerateTFCRunTaskHmacKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceRegenerateTFCRunTaskHmacKeyParams() *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithTimeout creates a new PackerServiceRegenerateTFCRunTaskHmacKeyParams object
// with the ability to set a timeout on a request.
func NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithTimeout(timeout time.Duration) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyParams{
		timeout: timeout,
	}
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithContext creates a new PackerServiceRegenerateTFCRunTaskHmacKeyParams object
// with the ability to set a context for a request.
func NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithContext(ctx context.Context) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyParams{
		Context: ctx,
	}
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithHTTPClient creates a new PackerServiceRegenerateTFCRunTaskHmacKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceRegenerateTFCRunTaskHmacKeyParamsWithHTTPClient(client *http.Client) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyParams{
		HTTPClient: client,
	}
}

/*
PackerServiceRegenerateTFCRunTaskHmacKeyParams contains all the parameters to send to the API endpoint

	for the packer service regenerate t f c run task hmac key operation.

	Typically these are written to a http.Request.
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyParams struct {

	// Body.
	Body PackerServiceRegenerateTFCRunTaskHmacKeyBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service regenerate t f c run task hmac key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithDefaults() *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service regenerate t f c run task hmac key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithTimeout(timeout time.Duration) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithContext(ctx context.Context) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithHTTPClient(client *http.Client) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithBody(body PackerServiceRegenerateTFCRunTaskHmacKeyBody) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetBody(body PackerServiceRegenerateTFCRunTaskHmacKeyBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WithLocationProjectID(locationProjectID string) *PackerServiceRegenerateTFCRunTaskHmacKeyParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service regenerate t f c run task hmac key params
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
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
