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

// NewPackerServiceGetBuildParams creates a new PackerServiceGetBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceGetBuildParams() *PackerServiceGetBuildParams {
	return &PackerServiceGetBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceGetBuildParamsWithTimeout creates a new PackerServiceGetBuildParams object
// with the ability to set a timeout on a request.
func NewPackerServiceGetBuildParamsWithTimeout(timeout time.Duration) *PackerServiceGetBuildParams {
	return &PackerServiceGetBuildParams{
		timeout: timeout,
	}
}

// NewPackerServiceGetBuildParamsWithContext creates a new PackerServiceGetBuildParams object
// with the ability to set a context for a request.
func NewPackerServiceGetBuildParamsWithContext(ctx context.Context) *PackerServiceGetBuildParams {
	return &PackerServiceGetBuildParams{
		Context: ctx,
	}
}

// NewPackerServiceGetBuildParamsWithHTTPClient creates a new PackerServiceGetBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceGetBuildParamsWithHTTPClient(client *http.Client) *PackerServiceGetBuildParams {
	return &PackerServiceGetBuildParams{
		HTTPClient: client,
	}
}

/*
PackerServiceGetBuildParams contains all the parameters to send to the API endpoint

	for the packer service get build operation.

	Typically these are written to a http.Request.
*/
type PackerServiceGetBuildParams struct {

	/* BuildID.

	   ULID of the build that should be retrieved.
	*/
	BuildID string

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

// WithDefaults hydrates default values in the packer service get build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceGetBuildParams) WithDefaults() *PackerServiceGetBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service get build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceGetBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service get build params
func (o *PackerServiceGetBuildParams) WithTimeout(timeout time.Duration) *PackerServiceGetBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service get build params
func (o *PackerServiceGetBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service get build params
func (o *PackerServiceGetBuildParams) WithContext(ctx context.Context) *PackerServiceGetBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service get build params
func (o *PackerServiceGetBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service get build params
func (o *PackerServiceGetBuildParams) WithHTTPClient(client *http.Client) *PackerServiceGetBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service get build params
func (o *PackerServiceGetBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the packer service get build params
func (o *PackerServiceGetBuildParams) WithBuildID(buildID string) *PackerServiceGetBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the packer service get build params
func (o *PackerServiceGetBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service get build params
func (o *PackerServiceGetBuildParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceGetBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service get build params
func (o *PackerServiceGetBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service get build params
func (o *PackerServiceGetBuildParams) WithLocationProjectID(locationProjectID string) *PackerServiceGetBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service get build params
func (o *PackerServiceGetBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service get build params
func (o *PackerServiceGetBuildParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceGetBuildParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service get build params
func (o *PackerServiceGetBuildParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service get build params
func (o *PackerServiceGetBuildParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceGetBuildParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service get build params
func (o *PackerServiceGetBuildParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceGetBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
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
