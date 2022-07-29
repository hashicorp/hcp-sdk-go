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
)

// NewReadVersionParams creates a new ReadVersionParams object
// with the default values initialized.
func NewReadVersionParams() *ReadVersionParams {
	var ()
	return &ReadVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadVersionParamsWithTimeout creates a new ReadVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadVersionParamsWithTimeout(timeout time.Duration) *ReadVersionParams {
	var ()
	return &ReadVersionParams{

		timeout: timeout,
	}
}

// NewReadVersionParamsWithContext creates a new ReadVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadVersionParamsWithContext(ctx context.Context) *ReadVersionParams {
	var ()
	return &ReadVersionParams{

		Context: ctx,
	}
}

// NewReadVersionParamsWithHTTPClient creates a new ReadVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadVersionParamsWithHTTPClient(client *http.Client) *ReadVersionParams {
	var ()
	return &ReadVersionParams{
		HTTPClient: client,
	}
}

/*ReadVersionParams contains all the parameters to send to the API endpoint
for the read version operation typically these are written to a http.Request
*/
type ReadVersionParams struct {

	/*Box
	  The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".

	*/
	Box string
	/*Registry
	  The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".

	*/
	Registry string
	/*Version
	  The name of the version to look up.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read version params
func (o *ReadVersionParams) WithTimeout(timeout time.Duration) *ReadVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read version params
func (o *ReadVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read version params
func (o *ReadVersionParams) WithContext(ctx context.Context) *ReadVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read version params
func (o *ReadVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read version params
func (o *ReadVersionParams) WithHTTPClient(client *http.Client) *ReadVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read version params
func (o *ReadVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBox adds the box to the read version params
func (o *ReadVersionParams) WithBox(box string) *ReadVersionParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the read version params
func (o *ReadVersionParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the read version params
func (o *ReadVersionParams) WithRegistry(registry string) *ReadVersionParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the read version params
func (o *ReadVersionParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the read version params
func (o *ReadVersionParams) WithVersion(version string) *ReadVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the read version params
func (o *ReadVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReadVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
