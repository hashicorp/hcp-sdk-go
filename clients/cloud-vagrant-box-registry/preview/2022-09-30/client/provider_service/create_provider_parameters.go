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

// NewCreateProviderParams creates a new CreateProviderParams object
// with the default values initialized.
func NewCreateProviderParams() *CreateProviderParams {
	var ()
	return &CreateProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProviderParamsWithTimeout creates a new CreateProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProviderParamsWithTimeout(timeout time.Duration) *CreateProviderParams {
	var ()
	return &CreateProviderParams{

		timeout: timeout,
	}
}

// NewCreateProviderParamsWithContext creates a new CreateProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProviderParamsWithContext(ctx context.Context) *CreateProviderParams {
	var ()
	return &CreateProviderParams{

		Context: ctx,
	}
}

// NewCreateProviderParamsWithHTTPClient creates a new CreateProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProviderParamsWithHTTPClient(client *http.Client) *CreateProviderParams {
	var ()
	return &CreateProviderParams{
		HTTPClient: client,
	}
}

/*CreateProviderParams contains all the parameters to send to the API endpoint
for the create provider operation typically these are written to a http.Request
*/
type CreateProviderParams struct {

	/*Body*/
	Body *models.HashicorpCloudVagrantCreateProviderRequest
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
	  The name of the Version to create the Provider in.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create provider params
func (o *CreateProviderParams) WithTimeout(timeout time.Duration) *CreateProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create provider params
func (o *CreateProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create provider params
func (o *CreateProviderParams) WithContext(ctx context.Context) *CreateProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create provider params
func (o *CreateProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create provider params
func (o *CreateProviderParams) WithHTTPClient(client *http.Client) *CreateProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create provider params
func (o *CreateProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create provider params
func (o *CreateProviderParams) WithBody(body *models.HashicorpCloudVagrantCreateProviderRequest) *CreateProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create provider params
func (o *CreateProviderParams) SetBody(body *models.HashicorpCloudVagrantCreateProviderRequest) {
	o.Body = body
}

// WithBox adds the box to the create provider params
func (o *CreateProviderParams) WithBox(box string) *CreateProviderParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the create provider params
func (o *CreateProviderParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the create provider params
func (o *CreateProviderParams) WithRegistry(registry string) *CreateProviderParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the create provider params
func (o *CreateProviderParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the create provider params
func (o *CreateProviderParams) WithVersion(version string) *CreateProviderParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create provider params
func (o *CreateProviderParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
