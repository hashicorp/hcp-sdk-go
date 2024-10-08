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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/stable/2022-09-30/models"
)

// NewUpdateProviderParams creates a new UpdateProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProviderParams() *UpdateProviderParams {
	return &UpdateProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProviderParamsWithTimeout creates a new UpdateProviderParams object
// with the ability to set a timeout on a request.
func NewUpdateProviderParamsWithTimeout(timeout time.Duration) *UpdateProviderParams {
	return &UpdateProviderParams{
		timeout: timeout,
	}
}

// NewUpdateProviderParamsWithContext creates a new UpdateProviderParams object
// with the ability to set a context for a request.
func NewUpdateProviderParamsWithContext(ctx context.Context) *UpdateProviderParams {
	return &UpdateProviderParams{
		Context: ctx,
	}
}

// NewUpdateProviderParamsWithHTTPClient creates a new UpdateProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProviderParamsWithHTTPClient(client *http.Client) *UpdateProviderParams {
	return &UpdateProviderParams{
		HTTPClient: client,
	}
}

/*
UpdateProviderParams contains all the parameters to send to the API endpoint

	for the update provider operation.

	Typically these are written to a http.Request.
*/
type UpdateProviderParams struct {

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Data.

	   Details of the Provider to update.
	*/
	Data *models.HashicorpCloudVagrant20220930Provider

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

// WithDefaults hydrates default values in the update provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProviderParams) WithDefaults() *UpdateProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update provider params
func (o *UpdateProviderParams) WithTimeout(timeout time.Duration) *UpdateProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update provider params
func (o *UpdateProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update provider params
func (o *UpdateProviderParams) WithContext(ctx context.Context) *UpdateProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update provider params
func (o *UpdateProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update provider params
func (o *UpdateProviderParams) WithHTTPClient(client *http.Client) *UpdateProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update provider params
func (o *UpdateProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBox adds the box to the update provider params
func (o *UpdateProviderParams) WithBox(box string) *UpdateProviderParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the update provider params
func (o *UpdateProviderParams) SetBox(box string) {
	o.Box = box
}

// WithData adds the data to the update provider params
func (o *UpdateProviderParams) WithData(data *models.HashicorpCloudVagrant20220930Provider) *UpdateProviderParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the update provider params
func (o *UpdateProviderParams) SetData(data *models.HashicorpCloudVagrant20220930Provider) {
	o.Data = data
}

// WithProvider adds the provider to the update provider params
func (o *UpdateProviderParams) WithProvider(provider string) *UpdateProviderParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the update provider params
func (o *UpdateProviderParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegistry adds the registry to the update provider params
func (o *UpdateProviderParams) WithRegistry(registry string) *UpdateProviderParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the update provider params
func (o *UpdateProviderParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the update provider params
func (o *UpdateProviderParams) WithVersion(version string) *UpdateProviderParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update provider params
func (o *UpdateProviderParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param box
	if err := r.SetPathParam("box", o.Box); err != nil {
		return err
	}
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
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
