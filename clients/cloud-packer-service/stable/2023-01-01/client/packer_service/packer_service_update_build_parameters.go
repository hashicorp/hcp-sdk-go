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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
)

// NewPackerServiceUpdateBuildParams creates a new PackerServiceUpdateBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceUpdateBuildParams() *PackerServiceUpdateBuildParams {
	return &PackerServiceUpdateBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceUpdateBuildParamsWithTimeout creates a new PackerServiceUpdateBuildParams object
// with the ability to set a timeout on a request.
func NewPackerServiceUpdateBuildParamsWithTimeout(timeout time.Duration) *PackerServiceUpdateBuildParams {
	return &PackerServiceUpdateBuildParams{
		timeout: timeout,
	}
}

// NewPackerServiceUpdateBuildParamsWithContext creates a new PackerServiceUpdateBuildParams object
// with the ability to set a context for a request.
func NewPackerServiceUpdateBuildParamsWithContext(ctx context.Context) *PackerServiceUpdateBuildParams {
	return &PackerServiceUpdateBuildParams{
		Context: ctx,
	}
}

// NewPackerServiceUpdateBuildParamsWithHTTPClient creates a new PackerServiceUpdateBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceUpdateBuildParamsWithHTTPClient(client *http.Client) *PackerServiceUpdateBuildParams {
	return &PackerServiceUpdateBuildParams{
		HTTPClient: client,
	}
}

/*
PackerServiceUpdateBuildParams contains all the parameters to send to the API endpoint

	for the packer service update build operation.

	Typically these are written to a http.Request.
*/
type PackerServiceUpdateBuildParams struct {

	// Body.
	Body *models.HashicorpCloudPacker20230101UpdateBuildBody

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

// WithDefaults hydrates default values in the packer service update build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateBuildParams) WithDefaults() *PackerServiceUpdateBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service update build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithTimeout(timeout time.Duration) *PackerServiceUpdateBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithContext(ctx context.Context) *PackerServiceUpdateBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithHTTPClient(client *http.Client) *PackerServiceUpdateBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithBody(body *models.HashicorpCloudPacker20230101UpdateBuildBody) *PackerServiceUpdateBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetBody(body *models.HashicorpCloudPacker20230101UpdateBuildBody) {
	o.Body = body
}

// WithBucketName adds the bucketName to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithBucketName(bucketName string) *PackerServiceUpdateBuildParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithBuildID adds the buildID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithBuildID(buildID string) *PackerServiceUpdateBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithFingerprint adds the fingerprint to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithFingerprint(fingerprint string) *PackerServiceUpdateBuildParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceUpdateBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationProjectID(locationProjectID string) *PackerServiceUpdateBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceUpdateBuildParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceUpdateBuildParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceUpdateBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
