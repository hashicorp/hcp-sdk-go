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
)

// NewListUpgradeVersionsParams creates a new ListUpgradeVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUpgradeVersionsParams() *ListUpgradeVersionsParams {
	return &ListUpgradeVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUpgradeVersionsParamsWithTimeout creates a new ListUpgradeVersionsParams object
// with the ability to set a timeout on a request.
func NewListUpgradeVersionsParamsWithTimeout(timeout time.Duration) *ListUpgradeVersionsParams {
	return &ListUpgradeVersionsParams{
		timeout: timeout,
	}
}

// NewListUpgradeVersionsParamsWithContext creates a new ListUpgradeVersionsParams object
// with the ability to set a context for a request.
func NewListUpgradeVersionsParamsWithContext(ctx context.Context) *ListUpgradeVersionsParams {
	return &ListUpgradeVersionsParams{
		Context: ctx,
	}
}

// NewListUpgradeVersionsParamsWithHTTPClient creates a new ListUpgradeVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUpgradeVersionsParamsWithHTTPClient(client *http.Client) *ListUpgradeVersionsParams {
	return &ListUpgradeVersionsParams{
		HTTPClient: client,
	}
}

/*
ListUpgradeVersionsParams contains all the parameters to send to the API endpoint

	for the list upgrade versions operation.

	Typically these are written to a http.Request.
*/
type ListUpgradeVersionsParams struct {

	/* ID.

	   id of the Consul cluster to list versions for.
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

// WithDefaults hydrates default values in the list upgrade versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUpgradeVersionsParams) WithDefaults() *ListUpgradeVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list upgrade versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUpgradeVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithTimeout(timeout time.Duration) *ListUpgradeVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithContext(ctx context.Context) *ListUpgradeVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithHTTPClient(client *http.Client) *ListUpgradeVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithID(id string) *ListUpgradeVersionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithLocationOrganizationID(locationOrganizationID string) *ListUpgradeVersionsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithLocationProjectID(locationProjectID string) *ListUpgradeVersionsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithLocationRegionProvider(locationRegionProvider *string) *ListUpgradeVersionsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list upgrade versions params
func (o *ListUpgradeVersionsParams) WithLocationRegionRegion(locationRegionRegion *string) *ListUpgradeVersionsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list upgrade versions params
func (o *ListUpgradeVersionsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *ListUpgradeVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
