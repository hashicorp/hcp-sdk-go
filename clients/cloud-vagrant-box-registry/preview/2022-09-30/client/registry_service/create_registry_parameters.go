// Code generated by go-swagger; DO NOT EDIT.

package registry_service

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

// NewCreateRegistryParams creates a new CreateRegistryParams object
// with the default values initialized.
func NewCreateRegistryParams() *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRegistryParamsWithTimeout creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRegistryParamsWithTimeout(timeout time.Duration) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		timeout: timeout,
	}
}

// NewCreateRegistryParamsWithContext creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRegistryParamsWithContext(ctx context.Context) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		Context: ctx,
	}
}

// NewCreateRegistryParamsWithHTTPClient creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRegistryParamsWithHTTPClient(client *http.Client) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{
		HTTPClient: client,
	}
}

/*CreateRegistryParams contains all the parameters to send to the API endpoint
for the create registry operation typically these are written to a http.Request
*/
type CreateRegistryParams struct {

	/*Body*/
	Body *models.HashicorpCloudVagrant20220930CreateRegistryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) WithTimeout(timeout time.Duration) *CreateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create registry params
func (o *CreateRegistryParams) WithContext(ctx context.Context) *CreateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create registry params
func (o *CreateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) WithHTTPClient(client *http.Client) *CreateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create registry params
func (o *CreateRegistryParams) WithBody(body *models.HashicorpCloudVagrant20220930CreateRegistryRequest) *CreateRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create registry params
func (o *CreateRegistryParams) SetBody(body *models.HashicorpCloudVagrant20220930CreateRegistryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
