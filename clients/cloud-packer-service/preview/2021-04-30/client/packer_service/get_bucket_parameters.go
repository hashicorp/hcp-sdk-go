// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

// NewGetBucketParams creates a new GetBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBucketParams() *GetBucketParams {
	return &GetBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBucketParamsWithTimeout creates a new GetBucketParams object
// with the ability to set a timeout on a request.
func NewGetBucketParamsWithTimeout(timeout time.Duration) *GetBucketParams {
	return &GetBucketParams{
		timeout: timeout,
	}
}

// NewGetBucketParamsWithContext creates a new GetBucketParams object
// with the ability to set a context for a request.
func NewGetBucketParamsWithContext(ctx context.Context) *GetBucketParams {
	return &GetBucketParams{
		Context: ctx,
	}
}

// NewGetBucketParamsWithHTTPClient creates a new GetBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBucketParamsWithHTTPClient(client *http.Client) *GetBucketParams {
	return &GetBucketParams{
		HTTPClient: client,
	}
}

/* GetBucketParams contains all the parameters to send to the API endpoint
   for the get bucket operation.

   Typically these are written to a http.Request.
*/
type GetBucketParams struct {

	// BucketID.
	BucketID *string

	// BucketSlug.
	BucketSlug string

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

// WithDefaults hydrates default values in the get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketParams) WithDefaults() *GetBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bucket params
func (o *GetBucketParams) WithTimeout(timeout time.Duration) *GetBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bucket params
func (o *GetBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bucket params
func (o *GetBucketParams) WithContext(ctx context.Context) *GetBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bucket params
func (o *GetBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bucket params
func (o *GetBucketParams) WithHTTPClient(client *http.Client) *GetBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bucket params
func (o *GetBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the get bucket params
func (o *GetBucketParams) WithBucketID(bucketID *string) *GetBucketParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the get bucket params
func (o *GetBucketParams) SetBucketID(bucketID *string) {
	o.BucketID = bucketID
}

// WithBucketSlug adds the bucketSlug to the get bucket params
func (o *GetBucketParams) WithBucketSlug(bucketSlug string) *GetBucketParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the get bucket params
func (o *GetBucketParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the get bucket params
func (o *GetBucketParams) WithLocationOrganizationID(locationOrganizationID string) *GetBucketParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get bucket params
func (o *GetBucketParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get bucket params
func (o *GetBucketParams) WithLocationProjectID(locationProjectID string) *GetBucketParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get bucket params
func (o *GetBucketParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get bucket params
func (o *GetBucketParams) WithLocationRegionProvider(locationRegionProvider *string) *GetBucketParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get bucket params
func (o *GetBucketParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get bucket params
func (o *GetBucketParams) WithLocationRegionRegion(locationRegionRegion *string) *GetBucketParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get bucket params
func (o *GetBucketParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketID != nil {

		// query param bucket_id
		var qrBucketID string

		if o.BucketID != nil {
			qrBucketID = *o.BucketID
		}
		qBucketID := qrBucketID
		if qBucketID != "" {

			if err := r.SetQueryParam("bucket_id", qBucketID); err != nil {
				return err
			}
		}
	}

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
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
