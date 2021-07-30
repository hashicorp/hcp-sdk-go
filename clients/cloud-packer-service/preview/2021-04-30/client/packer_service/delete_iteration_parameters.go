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

// NewDeleteIterationParams creates a new DeleteIterationParams object
// with the default values initialized.
func NewDeleteIterationParams() *DeleteIterationParams {
	var ()
	return &DeleteIterationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIterationParamsWithTimeout creates a new DeleteIterationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIterationParamsWithTimeout(timeout time.Duration) *DeleteIterationParams {
	var ()
	return &DeleteIterationParams{

		timeout: timeout,
	}
}

// NewDeleteIterationParamsWithContext creates a new DeleteIterationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIterationParamsWithContext(ctx context.Context) *DeleteIterationParams {
	var ()
	return &DeleteIterationParams{

		Context: ctx,
	}
}

// NewDeleteIterationParamsWithHTTPClient creates a new DeleteIterationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIterationParamsWithHTTPClient(client *http.Client) *DeleteIterationParams {
	var ()
	return &DeleteIterationParams{
		HTTPClient: client,
	}
}

/*DeleteIterationParams contains all the parameters to send to the API endpoint
for the delete iteration operation typically these are written to a http.Request
*/
type DeleteIterationParams struct {

	/*BucketSlug*/
	BucketSlug *string
	/*IterationID
	  id of the iteration you would like to delete

	*/
	IterationID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*LocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	LocationRegionProvider *string
	/*LocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iteration params
func (o *DeleteIterationParams) WithTimeout(timeout time.Duration) *DeleteIterationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iteration params
func (o *DeleteIterationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iteration params
func (o *DeleteIterationParams) WithContext(ctx context.Context) *DeleteIterationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iteration params
func (o *DeleteIterationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iteration params
func (o *DeleteIterationParams) WithHTTPClient(client *http.Client) *DeleteIterationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iteration params
func (o *DeleteIterationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the delete iteration params
func (o *DeleteIterationParams) WithBucketSlug(bucketSlug *string) *DeleteIterationParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the delete iteration params
func (o *DeleteIterationParams) SetBucketSlug(bucketSlug *string) {
	o.BucketSlug = bucketSlug
}

// WithIterationID adds the iterationID to the delete iteration params
func (o *DeleteIterationParams) WithIterationID(iterationID string) *DeleteIterationParams {
	o.SetIterationID(iterationID)
	return o
}

// SetIterationID adds the iterationId to the delete iteration params
func (o *DeleteIterationParams) SetIterationID(iterationID string) {
	o.IterationID = iterationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the delete iteration params
func (o *DeleteIterationParams) WithLocationOrganizationID(locationOrganizationID string) *DeleteIterationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the delete iteration params
func (o *DeleteIterationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the delete iteration params
func (o *DeleteIterationParams) WithLocationProjectID(locationProjectID string) *DeleteIterationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the delete iteration params
func (o *DeleteIterationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the delete iteration params
func (o *DeleteIterationParams) WithLocationRegionProvider(locationRegionProvider *string) *DeleteIterationParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the delete iteration params
func (o *DeleteIterationParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the delete iteration params
func (o *DeleteIterationParams) WithLocationRegionRegion(locationRegionRegion *string) *DeleteIterationParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the delete iteration params
func (o *DeleteIterationParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIterationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketSlug != nil {

		// query param bucket_slug
		var qrBucketSlug string
		if o.BucketSlug != nil {
			qrBucketSlug = *o.BucketSlug
		}
		qBucketSlug := qrBucketSlug
		if qBucketSlug != "" {
			if err := r.SetQueryParam("bucket_slug", qBucketSlug); err != nil {
				return err
			}
		}

	}

	// path param iteration_id
	if err := r.SetPathParam("iteration_id", o.IterationID); err != nil {
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