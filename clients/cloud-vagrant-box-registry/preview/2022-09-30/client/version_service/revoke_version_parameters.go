// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// NewRevokeVersionParams creates a new RevokeVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeVersionParams() *RevokeVersionParams {
	return &RevokeVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeVersionParamsWithTimeout creates a new RevokeVersionParams object
// with the ability to set a timeout on a request.
func NewRevokeVersionParamsWithTimeout(timeout time.Duration) *RevokeVersionParams {
	return &RevokeVersionParams{
		timeout: timeout,
	}
}

// NewRevokeVersionParamsWithContext creates a new RevokeVersionParams object
// with the ability to set a context for a request.
func NewRevokeVersionParamsWithContext(ctx context.Context) *RevokeVersionParams {
	return &RevokeVersionParams{
		Context: ctx,
	}
}

// NewRevokeVersionParamsWithHTTPClient creates a new RevokeVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeVersionParamsWithHTTPClient(client *http.Client) *RevokeVersionParams {
	return &RevokeVersionParams{
		HTTPClient: client,
	}
}

/*
RevokeVersionParams contains all the parameters to send to the API endpoint

	for the revoke version operation.

	Typically these are written to a http.Request.
*/
type RevokeVersionParams struct {

	// Body.
	Body *models.HashicorpCloudVagrantRevokeVersionRequest

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	/* Version.

	   The name of the Version to revoke.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeVersionParams) WithDefaults() *RevokeVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke version params
func (o *RevokeVersionParams) WithTimeout(timeout time.Duration) *RevokeVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke version params
func (o *RevokeVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke version params
func (o *RevokeVersionParams) WithContext(ctx context.Context) *RevokeVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke version params
func (o *RevokeVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke version params
func (o *RevokeVersionParams) WithHTTPClient(client *http.Client) *RevokeVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke version params
func (o *RevokeVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the revoke version params
func (o *RevokeVersionParams) WithBody(body *models.HashicorpCloudVagrantRevokeVersionRequest) *RevokeVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the revoke version params
func (o *RevokeVersionParams) SetBody(body *models.HashicorpCloudVagrantRevokeVersionRequest) {
	o.Body = body
}

// WithBox adds the box to the revoke version params
func (o *RevokeVersionParams) WithBox(box string) *RevokeVersionParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the revoke version params
func (o *RevokeVersionParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the revoke version params
func (o *RevokeVersionParams) WithRegistry(registry string) *RevokeVersionParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the revoke version params
func (o *RevokeVersionParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the revoke version params
func (o *RevokeVersionParams) WithVersion(version string) *RevokeVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the revoke version params
func (o *RevokeVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param box
	if err := r.SetPathParam("box", o.Box); err != nil {
		return err
	}

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
