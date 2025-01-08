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

// NewPackerServiceListSbomsParams creates a new PackerServiceListSbomsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceListSbomsParams() *PackerServiceListSbomsParams {
	return &PackerServiceListSbomsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceListSbomsParamsWithTimeout creates a new PackerServiceListSbomsParams object
// with the ability to set a timeout on a request.
func NewPackerServiceListSbomsParamsWithTimeout(timeout time.Duration) *PackerServiceListSbomsParams {
	return &PackerServiceListSbomsParams{
		timeout: timeout,
	}
}

// NewPackerServiceListSbomsParamsWithContext creates a new PackerServiceListSbomsParams object
// with the ability to set a context for a request.
func NewPackerServiceListSbomsParamsWithContext(ctx context.Context) *PackerServiceListSbomsParams {
	return &PackerServiceListSbomsParams{
		Context: ctx,
	}
}

// NewPackerServiceListSbomsParamsWithHTTPClient creates a new PackerServiceListSbomsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceListSbomsParamsWithHTTPClient(client *http.Client) *PackerServiceListSbomsParams {
	return &PackerServiceListSbomsParams{
		HTTPClient: client,
	}
}

/*
PackerServiceListSbomsParams contains all the parameters to send to the API endpoint

	for the packer service list sboms operation.

	Typically these are written to a http.Request.
*/
type PackerServiceListSbomsParams struct {

	// BucketName.
	BucketName string

	// BuildID.
	BuildID string

	// Fingerprint.
	Fingerprint string

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

// WithDefaults hydrates default values in the packer service list sboms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceListSbomsParams) WithDefaults() *PackerServiceListSbomsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service list sboms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceListSbomsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithTimeout(timeout time.Duration) *PackerServiceListSbomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithContext(ctx context.Context) *PackerServiceListSbomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithHTTPClient(client *http.Client) *PackerServiceListSbomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketName adds the bucketName to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithBucketName(bucketName string) *PackerServiceListSbomsParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithBuildID adds the buildID to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithBuildID(buildID string) *PackerServiceListSbomsParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithFingerprint adds the fingerprint to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithFingerprint(fingerprint string) *PackerServiceListSbomsParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceListSbomsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithLocationProjectID(locationProjectID string) *PackerServiceListSbomsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceListSbomsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service list sboms params
func (o *PackerServiceListSbomsParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceListSbomsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service list sboms params
func (o *PackerServiceListSbomsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceListSbomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_name
	if err := r.SetPathParam("bucket_name", o.BucketName); err != nil {
		return err
	}

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
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
