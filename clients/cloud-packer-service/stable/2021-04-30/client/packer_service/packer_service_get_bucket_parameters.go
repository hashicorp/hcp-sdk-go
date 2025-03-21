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

// NewPackerServiceGetBucketParams creates a new PackerServiceGetBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceGetBucketParams() *PackerServiceGetBucketParams {
	return &PackerServiceGetBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceGetBucketParamsWithTimeout creates a new PackerServiceGetBucketParams object
// with the ability to set a timeout on a request.
func NewPackerServiceGetBucketParamsWithTimeout(timeout time.Duration) *PackerServiceGetBucketParams {
	return &PackerServiceGetBucketParams{
		timeout: timeout,
	}
}

// NewPackerServiceGetBucketParamsWithContext creates a new PackerServiceGetBucketParams object
// with the ability to set a context for a request.
func NewPackerServiceGetBucketParamsWithContext(ctx context.Context) *PackerServiceGetBucketParams {
	return &PackerServiceGetBucketParams{
		Context: ctx,
	}
}

// NewPackerServiceGetBucketParamsWithHTTPClient creates a new PackerServiceGetBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceGetBucketParamsWithHTTPClient(client *http.Client) *PackerServiceGetBucketParams {
	return &PackerServiceGetBucketParams{
		HTTPClient: client,
	}
}

/*
PackerServiceGetBucketParams contains all the parameters to send to the API endpoint

	for the packer service get bucket operation.

	Typically these are written to a http.Request.
*/
type PackerServiceGetBucketParams struct {

	/* BucketID.

	   ULID of the bucket.
	*/
	BucketID *string

	/* BucketSlug.

	   Human-readable name for the bucket.
	*/
	BucketSlug string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceGetBucketParams) WithDefaults() *PackerServiceGetBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceGetBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithTimeout(timeout time.Duration) *PackerServiceGetBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithContext(ctx context.Context) *PackerServiceGetBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithHTTPClient(client *http.Client) *PackerServiceGetBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithBucketID(bucketID *string) *PackerServiceGetBucketParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetBucketID(bucketID *string) {
	o.BucketID = bucketID
}

// WithBucketSlug adds the bucketSlug to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithBucketSlug(bucketSlug string) *PackerServiceGetBucketParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceGetBucketParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service get bucket params
func (o *PackerServiceGetBucketParams) WithLocationProjectID(locationProjectID string) *PackerServiceGetBucketParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service get bucket params
func (o *PackerServiceGetBucketParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceGetBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
