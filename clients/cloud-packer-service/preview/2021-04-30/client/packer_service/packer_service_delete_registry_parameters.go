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

// NewPackerServiceDeleteRegistryParams creates a new PackerServiceDeleteRegistryParams object
// with the default values initialized.
func NewPackerServiceDeleteRegistryParams() *PackerServiceDeleteRegistryParams {
	var ()
	return &PackerServiceDeleteRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceDeleteRegistryParamsWithTimeout creates a new PackerServiceDeleteRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceDeleteRegistryParamsWithTimeout(timeout time.Duration) *PackerServiceDeleteRegistryParams {
	var ()
	return &PackerServiceDeleteRegistryParams{

		timeout: timeout,
	}
}

// NewPackerServiceDeleteRegistryParamsWithContext creates a new PackerServiceDeleteRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceDeleteRegistryParamsWithContext(ctx context.Context) *PackerServiceDeleteRegistryParams {
	var ()
	return &PackerServiceDeleteRegistryParams{

		Context: ctx,
	}
}

// NewPackerServiceDeleteRegistryParamsWithHTTPClient creates a new PackerServiceDeleteRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceDeleteRegistryParamsWithHTTPClient(client *http.Client) *PackerServiceDeleteRegistryParams {
	var ()
	return &PackerServiceDeleteRegistryParams{
		HTTPClient: client,
	}
}

/*PackerServiceDeleteRegistryParams contains all the parameters to send to the API endpoint
for the packer service delete registry operation typically these are written to a http.Request
*/
type PackerServiceDeleteRegistryParams struct {

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