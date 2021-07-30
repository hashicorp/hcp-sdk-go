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
	"github.com/go-openapi/swag"
)

// NewGetIterationParams creates a new GetIterationParams object
// with the default values initialized.
func NewGetIterationParams() *GetIterationParams {
	var ()
	return &GetIterationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIterationParamsWithTimeout creates a new GetIterationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIterationParamsWithTimeout(timeout time.Duration) *GetIterationParams {
	var ()
	return &GetIterationParams{

		timeout: timeout,
	}
}

// NewGetIterationParamsWithContext creates a new GetIterationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIterationParamsWithContext(ctx context.Context) *GetIterationParams {
	var ()
	return &GetIterationParams{

		Context: ctx,
	}
}

// NewGetIterationParamsWithHTTPClient creates a new GetIterationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIterationParamsWithHTTPClient(client *http.Client) *GetIterationParams {
	var ()
	return &GetIterationParams{
		HTTPClient: client,
	}
}

/*GetIterationParams contains all the parameters to send to the API endpoint
for the get iteration operation typically these are written to a http.Request
*/
type GetIterationParams struct {

	/*BucketSlug*/
	BucketSlug string
	/*Fingerprint*/
	Fingerprint *string
	/*IncrementalVersion*/
	IncrementalVersion *int32
	/*IterationID*/
	IterationID *string
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

// WithTimeout adds the timeout to the get iteration params
func (o *GetIterationParams) WithTimeout(timeout time.Duration) *GetIterationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iteration params
func (o *GetIterationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iteration params
func (o *GetIterationParams) WithContext(ctx context.Context) *GetIterationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iteration params
func (o *GetIterationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iteration params
func (o *GetIterationParams) WithHTTPClient(client *http.Client) *GetIterationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iteration params
func (o *GetIterationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the get iteration params
func (o *GetIterationParams) WithBucketSlug(bucketSlug string) *GetIterationParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the get iteration params
func (o *GetIterationParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithFingerprint adds the fingerprint to the get iteration params
func (o *GetIterationParams) WithFingerprint(fingerprint *string) *GetIterationParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the get iteration params
func (o *GetIterationParams) SetFingerprint(fingerprint *string) {
	o.Fingerprint = fingerprint
}

// WithIncrementalVersion adds the incrementalVersion to the get iteration params
func (o *GetIterationParams) WithIncrementalVersion(incrementalVersion *int32) *GetIterationParams {
	o.SetIncrementalVersion(incrementalVersion)
	return o
}

// SetIncrementalVersion adds the incrementalVersion to the get iteration params
func (o *GetIterationParams) SetIncrementalVersion(incrementalVersion *int32) {
	o.IncrementalVersion = incrementalVersion
}

// WithIterationID adds the iterationID to the get iteration params
func (o *GetIterationParams) WithIterationID(iterationID *string) *GetIterationParams {
	o.SetIterationID(iterationID)
	return o
}

// SetIterationID adds the iterationId to the get iteration params
func (o *GetIterationParams) SetIterationID(iterationID *string) {
	o.IterationID = iterationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the get iteration params
func (o *GetIterationParams) WithLocationOrganizationID(locationOrganizationID string) *GetIterationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get iteration params
func (o *GetIterationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get iteration params
func (o *GetIterationParams) WithLocationProjectID(locationProjectID string) *GetIterationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get iteration params
func (o *GetIterationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get iteration params
func (o *GetIterationParams) WithLocationRegionProvider(locationRegionProvider *string) *GetIterationParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get iteration params
func (o *GetIterationParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get iteration params
func (o *GetIterationParams) WithLocationRegionRegion(locationRegionRegion *string) *GetIterationParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get iteration params
func (o *GetIterationParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetIterationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
		return err
	}

	if o.Fingerprint != nil {

		// query param fingerprint
		var qrFingerprint string
		if o.Fingerprint != nil {
			qrFingerprint = *o.Fingerprint
		}
		qFingerprint := qrFingerprint
		if qFingerprint != "" {
			if err := r.SetQueryParam("fingerprint", qFingerprint); err != nil {
				return err
			}
		}

	}

	if o.IncrementalVersion != nil {

		// query param incremental_version
		var qrIncrementalVersion int32
		if o.IncrementalVersion != nil {
			qrIncrementalVersion = *o.IncrementalVersion
		}
		qIncrementalVersion := swag.FormatInt32(qrIncrementalVersion)
		if qIncrementalVersion != "" {
			if err := r.SetQueryParam("incremental_version", qIncrementalVersion); err != nil {
				return err
			}
		}

	}

	if o.IterationID != nil {

		// query param iteration_id
		var qrIterationID string
		if o.IterationID != nil {
			qrIterationID = *o.IterationID
		}
		qIterationID := qrIterationID
		if qIterationID != "" {
			if err := r.SetQueryParam("iteration_id", qIterationID); err != nil {
				return err
			}
		}

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