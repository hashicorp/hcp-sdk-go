// Code generated by go-swagger; DO NOT EDIT.

package provider_service

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

// NewCompleteUploadParams creates a new CompleteUploadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCompleteUploadParams() *CompleteUploadParams {
	return &CompleteUploadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteUploadParamsWithTimeout creates a new CompleteUploadParams object
// with the ability to set a timeout on a request.
func NewCompleteUploadParamsWithTimeout(timeout time.Duration) *CompleteUploadParams {
	return &CompleteUploadParams{
		timeout: timeout,
	}
}

// NewCompleteUploadParamsWithContext creates a new CompleteUploadParams object
// with the ability to set a context for a request.
func NewCompleteUploadParamsWithContext(ctx context.Context) *CompleteUploadParams {
	return &CompleteUploadParams{
		Context: ctx,
	}
}

// NewCompleteUploadParamsWithHTTPClient creates a new CompleteUploadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCompleteUploadParamsWithHTTPClient(client *http.Client) *CompleteUploadParams {
	return &CompleteUploadParams{
		HTTPClient: client,
	}
}

/* CompleteUploadParams contains all the parameters to send to the API endpoint
   for the complete upload operation.

   Typically these are written to a http.Request.
*/
type CompleteUploadParams struct {

	// Body.
	Body *models.HashicorpCloudVagrantCompleteUploadRequest

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Provider.

	   The name of the Provider.
	*/
	Provider string

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	/* Version.

	   The name of the Version for the Provider.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the complete upload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CompleteUploadParams) WithDefaults() *CompleteUploadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the complete upload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CompleteUploadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the complete upload params
func (o *CompleteUploadParams) WithTimeout(timeout time.Duration) *CompleteUploadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete upload params
func (o *CompleteUploadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete upload params
func (o *CompleteUploadParams) WithContext(ctx context.Context) *CompleteUploadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete upload params
func (o *CompleteUploadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete upload params
func (o *CompleteUploadParams) WithHTTPClient(client *http.Client) *CompleteUploadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete upload params
func (o *CompleteUploadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the complete upload params
func (o *CompleteUploadParams) WithBody(body *models.HashicorpCloudVagrantCompleteUploadRequest) *CompleteUploadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete upload params
func (o *CompleteUploadParams) SetBody(body *models.HashicorpCloudVagrantCompleteUploadRequest) {
	o.Body = body
}

// WithBox adds the box to the complete upload params
func (o *CompleteUploadParams) WithBox(box string) *CompleteUploadParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the complete upload params
func (o *CompleteUploadParams) SetBox(box string) {
	o.Box = box
}

// WithProvider adds the provider to the complete upload params
func (o *CompleteUploadParams) WithProvider(provider string) *CompleteUploadParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the complete upload params
func (o *CompleteUploadParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegistry adds the registry to the complete upload params
func (o *CompleteUploadParams) WithRegistry(registry string) *CompleteUploadParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the complete upload params
func (o *CompleteUploadParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the complete upload params
func (o *CompleteUploadParams) WithVersion(version string) *CompleteUploadParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the complete upload params
func (o *CompleteUploadParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteUploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
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
