// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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
	"github.com/go-openapi/swag"
)

// NewGetClientConfigParams creates a new GetClientConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClientConfigParams() *GetClientConfigParams {
	return &GetClientConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClientConfigParamsWithTimeout creates a new GetClientConfigParams object
// with the ability to set a timeout on a request.
func NewGetClientConfigParamsWithTimeout(timeout time.Duration) *GetClientConfigParams {
	return &GetClientConfigParams{
		timeout: timeout,
	}
}

// NewGetClientConfigParamsWithContext creates a new GetClientConfigParams object
// with the ability to set a context for a request.
func NewGetClientConfigParamsWithContext(ctx context.Context) *GetClientConfigParams {
	return &GetClientConfigParams{
		Context: ctx,
	}
}

// NewGetClientConfigParamsWithHTTPClient creates a new GetClientConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClientConfigParamsWithHTTPClient(client *http.Client) *GetClientConfigParams {
	return &GetClientConfigParams{
		HTTPClient: client,
	}
}

/*
GetClientConfigParams contains all the parameters to send to the API endpoint

	for the get client config operation.

	Typically these are written to a http.Request.
*/
type GetClientConfigParams struct {

	/* Bundle.

	   bundle asks the service to respond with a file bundle as a zip archive.
	*/
	Bundle *bool

	/* ID.

	   id is the unique ID of the HCP Consul cluster to get
	*/
	ID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get client config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClientConfigParams) WithDefaults() *GetClientConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get client config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClientConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get client config params
func (o *GetClientConfigParams) WithTimeout(timeout time.Duration) *GetClientConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get client config params
func (o *GetClientConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get client config params
func (o *GetClientConfigParams) WithContext(ctx context.Context) *GetClientConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get client config params
func (o *GetClientConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get client config params
func (o *GetClientConfigParams) WithHTTPClient(client *http.Client) *GetClientConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get client config params
func (o *GetClientConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundle adds the bundle to the get client config params
func (o *GetClientConfigParams) WithBundle(bundle *bool) *GetClientConfigParams {
	o.SetBundle(bundle)
	return o
}

// SetBundle adds the bundle to the get client config params
func (o *GetClientConfigParams) SetBundle(bundle *bool) {
	o.Bundle = bundle
}

// WithID adds the id to the get client config params
func (o *GetClientConfigParams) WithID(id string) *GetClientConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get client config params
func (o *GetClientConfigParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the get client config params
func (o *GetClientConfigParams) WithLocationOrganizationID(locationOrganizationID string) *GetClientConfigParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get client config params
func (o *GetClientConfigParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get client config params
func (o *GetClientConfigParams) WithLocationProjectID(locationProjectID string) *GetClientConfigParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get client config params
func (o *GetClientConfigParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get client config params
func (o *GetClientConfigParams) WithLocationRegionProvider(locationRegionProvider *string) *GetClientConfigParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get client config params
func (o *GetClientConfigParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get client config params
func (o *GetClientConfigParams) WithLocationRegionRegion(locationRegionRegion *string) *GetClientConfigParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get client config params
func (o *GetClientConfigParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetClientConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bundle != nil {

		// query param bundle
		var qrBundle bool

		if o.Bundle != nil {
			qrBundle = *o.Bundle
		}
		qBundle := swag.FormatBool(qrBundle)
		if qBundle != "" {

			if err := r.SetQueryParam("bundle", qBundle); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
