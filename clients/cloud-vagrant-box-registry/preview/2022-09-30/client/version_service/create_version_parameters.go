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

// NewCreateVersionParams creates a new CreateVersionParams object
// with the default values initialized.
func NewCreateVersionParams() *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVersionParamsWithTimeout creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVersionParamsWithTimeout(timeout time.Duration) *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		timeout: timeout,
	}
}

// NewCreateVersionParamsWithContext creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVersionParamsWithContext(ctx context.Context) *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		Context: ctx,
	}
}

// NewCreateVersionParamsWithHTTPClient creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVersionParamsWithHTTPClient(client *http.Client) *CreateVersionParams {
	var ()
	return &CreateVersionParams{
		HTTPClient: client,
	}
}

/*CreateVersionParams contains all the parameters to send to the API endpoint
for the create version operation typically these are written to a http.Request
*/
type CreateVersionParams struct {

	/*Body*/
	Body *models.HashicorpCloudVagrant20220930CreateVersionRequest
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create version params
func (o *CreateVersionParams) WithTimeout(timeout time.Duration) *CreateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create version params
func (o *CreateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create version params
func (o *CreateVersionParams) WithContext(ctx context.Context) *CreateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create version params
func (o *CreateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create version params
func (o *CreateVersionParams) WithHTTPClient(client *http.Client) *CreateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create version params
func (o *CreateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create version params
func (o *CreateVersionParams) WithBody(body *models.HashicorpCloudVagrant20220930CreateVersionRequest) *CreateVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create version params
func (o *CreateVersionParams) SetBody(body *models.HashicorpCloudVagrant20220930CreateVersionRequest) {
	o.Body = body
}

// WithBox adds the box to the create version params
func (o *CreateVersionParams) WithBox(box string) *CreateVersionParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the create version params
func (o *CreateVersionParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the create version params
func (o *CreateVersionParams) WithRegistry(registry string) *CreateVersionParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the create version params
func (o *CreateVersionParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}