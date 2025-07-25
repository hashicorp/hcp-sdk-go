// Code generated by go-swagger; DO NOT EDIT.

package operation_service

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

// NewGetParams creates a new GetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetParams() *GetParams {
	return &GetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetParamsWithTimeout creates a new GetParams object
// with the ability to set a timeout on a request.
func NewGetParamsWithTimeout(timeout time.Duration) *GetParams {
	return &GetParams{
		timeout: timeout,
	}
}

// NewGetParamsWithContext creates a new GetParams object
// with the ability to set a context for a request.
func NewGetParamsWithContext(ctx context.Context) *GetParams {
	return &GetParams{
		Context: ctx,
	}
}

// NewGetParamsWithHTTPClient creates a new GetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetParamsWithHTTPClient(client *http.Client) *GetParams {
	return &GetParams{
		HTTPClient: client,
	}
}

/*
GetParams contains all the parameters to send to the API endpoint

	for the get operation.

	Typically these are written to a http.Request.
*/
type GetParams struct {

	/* ID.

	   id is the UUID of the operation.
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

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParams) WithDefaults() *GetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get params
func (o *GetParams) WithTimeout(timeout time.Duration) *GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get params
func (o *GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get params
func (o *GetParams) WithContext(ctx context.Context) *GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get params
func (o *GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get params
func (o *GetParams) WithHTTPClient(client *http.Client) *GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get params
func (o *GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get params
func (o *GetParams) WithID(id string) *GetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get params
func (o *GetParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the get params
func (o *GetParams) WithLocationOrganizationID(locationOrganizationID string) *GetParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get params
func (o *GetParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get params
func (o *GetParams) WithLocationProjectID(locationProjectID string) *GetParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get params
func (o *GetParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get params
func (o *GetParams) WithLocationRegionProvider(locationRegionProvider *string) *GetParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get params
func (o *GetParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get params
func (o *GetParams) WithLocationRegionRegion(locationRegionRegion *string) *GetParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get params
func (o *GetParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
