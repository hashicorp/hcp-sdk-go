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

// NewPackerServiceDeleteRegistryParams creates a new PackerServiceDeleteRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceDeleteRegistryParams() *PackerServiceDeleteRegistryParams {
	return &PackerServiceDeleteRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceDeleteRegistryParamsWithTimeout creates a new PackerServiceDeleteRegistryParams object
// with the ability to set a timeout on a request.
func NewPackerServiceDeleteRegistryParamsWithTimeout(timeout time.Duration) *PackerServiceDeleteRegistryParams {
	return &PackerServiceDeleteRegistryParams{
		timeout: timeout,
	}
}

// NewPackerServiceDeleteRegistryParamsWithContext creates a new PackerServiceDeleteRegistryParams object
// with the ability to set a context for a request.
func NewPackerServiceDeleteRegistryParamsWithContext(ctx context.Context) *PackerServiceDeleteRegistryParams {
	return &PackerServiceDeleteRegistryParams{
		Context: ctx,
	}
}

// NewPackerServiceDeleteRegistryParamsWithHTTPClient creates a new PackerServiceDeleteRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceDeleteRegistryParamsWithHTTPClient(client *http.Client) *PackerServiceDeleteRegistryParams {
	return &PackerServiceDeleteRegistryParams{
		HTTPClient: client,
	}
}

/*
PackerServiceDeleteRegistryParams contains all the parameters to send to the API endpoint

	for the packer service delete registry operation.

	Typically these are written to a http.Request.
*/
type PackerServiceDeleteRegistryParams struct {

	/* HardDelete.

	     When set to true, the registry will be deleted from database
	and recovery will no longer be possible.
	*/
	HardDelete *bool

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

// WithDefaults hydrates default values in the packer service delete registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteRegistryParams) WithDefaults() *PackerServiceDeleteRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service delete registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithTimeout(timeout time.Duration) *PackerServiceDeleteRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithContext(ctx context.Context) *PackerServiceDeleteRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithHTTPClient(client *http.Client) *PackerServiceDeleteRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHardDelete adds the hardDelete to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithHardDelete(hardDelete *bool) *PackerServiceDeleteRegistryParams {
	o.SetHardDelete(hardDelete)
	return o
}

// SetHardDelete adds the hardDelete to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetHardDelete(hardDelete *bool) {
	o.HardDelete = hardDelete
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceDeleteRegistryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithLocationProjectID(locationProjectID string) *PackerServiceDeleteRegistryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceDeleteRegistryParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceDeleteRegistryParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service delete registry params
func (o *PackerServiceDeleteRegistryParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceDeleteRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HardDelete != nil {

		// query param hard_delete
		var qrHardDelete bool

		if o.HardDelete != nil {
			qrHardDelete = *o.HardDelete
		}
		qHardDelete := swag.FormatBool(qrHardDelete)
		if qHardDelete != "" {

			if err := r.SetQueryParam("hard_delete", qHardDelete); err != nil {
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
