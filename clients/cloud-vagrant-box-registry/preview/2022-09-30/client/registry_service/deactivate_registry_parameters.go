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

// NewDeactivateRegistryParams creates a new DeactivateRegistryParams object
// with the default values initialized.
func NewDeactivateRegistryParams() *DeactivateRegistryParams {
	var ()
	return &DeactivateRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeactivateRegistryParamsWithTimeout creates a new DeactivateRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeactivateRegistryParamsWithTimeout(timeout time.Duration) *DeactivateRegistryParams {
	var ()
	return &DeactivateRegistryParams{

		timeout: timeout,
	}
}

// NewDeactivateRegistryParamsWithContext creates a new DeactivateRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeactivateRegistryParamsWithContext(ctx context.Context) *DeactivateRegistryParams {
	var ()
	return &DeactivateRegistryParams{

		Context: ctx,
	}
}

// NewDeactivateRegistryParamsWithHTTPClient creates a new DeactivateRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeactivateRegistryParamsWithHTTPClient(client *http.Client) *DeactivateRegistryParams {
	var ()
	return &DeactivateRegistryParams{
		HTTPClient: client,
	}
}

/*DeactivateRegistryParams contains all the parameters to send to the API endpoint
for the deactivate registry operation typically these are written to a http.Request
*/
type DeactivateRegistryParams struct {

	/*Body*/
	Body *models.HashicorpCloudVagrant20220930DeactivateRegistryRequest
	/*Registry
	  The name of the Registry to deactivate.

	*/
	Registry string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deactivate registry params
func (o *DeactivateRegistryParams) WithTimeout(timeout time.Duration) *DeactivateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deactivate registry params
func (o *DeactivateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deactivate registry params
func (o *DeactivateRegistryParams) WithContext(ctx context.Context) *DeactivateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deactivate registry params
func (o *DeactivateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deactivate registry params
func (o *DeactivateRegistryParams) WithHTTPClient(client *http.Client) *DeactivateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deactivate registry params
func (o *DeactivateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deactivate registry params
func (o *DeactivateRegistryParams) WithBody(body *models.HashicorpCloudVagrant20220930DeactivateRegistryRequest) *DeactivateRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deactivate registry params
func (o *DeactivateRegistryParams) SetBody(body *models.HashicorpCloudVagrant20220930DeactivateRegistryRequest) {
	o.Body = body
}

// WithRegistry adds the registry to the deactivate registry params
func (o *DeactivateRegistryParams) WithRegistry(registry string) *DeactivateRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the deactivate registry params
func (o *DeactivateRegistryParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *DeactivateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
