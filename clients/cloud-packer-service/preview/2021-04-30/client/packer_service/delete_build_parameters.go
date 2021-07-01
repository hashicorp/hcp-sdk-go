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

// NewDeleteBuildParams creates a new DeleteBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildParams() *DeleteBuildParams {
	return &DeleteBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildParamsWithTimeout creates a new DeleteBuildParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildParamsWithTimeout(timeout time.Duration) *DeleteBuildParams {
	return &DeleteBuildParams{
		timeout: timeout,
	}
}

// NewDeleteBuildParamsWithContext creates a new DeleteBuildParams object
// with the ability to set a context for a request.
func NewDeleteBuildParamsWithContext(ctx context.Context) *DeleteBuildParams {
	return &DeleteBuildParams{
		Context: ctx,
	}
}

// NewDeleteBuildParamsWithHTTPClient creates a new DeleteBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildParamsWithHTTPClient(client *http.Client) *DeleteBuildParams {
	return &DeleteBuildParams{
		HTTPClient: client,
	}
}

/* DeleteBuildParams contains all the parameters to send to the API endpoint
   for the delete build operation.

   Typically these are written to a http.Request.
*/
type DeleteBuildParams struct {

	/* BuildID.

	   build ULID
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

// WithDefaults hydrates default values in the delete build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParams) WithDefaults() *DeleteBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) WithTimeout(timeout time.Duration) *DeleteBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build params
func (o *DeleteBuildParams) WithContext(ctx context.Context) *DeleteBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build params
func (o *DeleteBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) WithHTTPClient(client *http.Client) *DeleteBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the delete build params
func (o *DeleteBuildParams) WithBuildID(buildID string) *DeleteBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the delete build params
func (o *DeleteBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithLocationOrganizationID adds the locationOrganizationID to the delete build params
func (o *DeleteBuildParams) WithLocationOrganizationID(locationOrganizationID string) *DeleteBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the delete build params
func (o *DeleteBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the delete build params
func (o *DeleteBuildParams) WithLocationProjectID(locationProjectID string) *DeleteBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the delete build params
func (o *DeleteBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the delete build params
func (o *DeleteBuildParams) WithLocationRegionProvider(locationRegionProvider *string) *DeleteBuildParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the delete build params
func (o *DeleteBuildParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the delete build params
func (o *DeleteBuildParams) WithLocationRegionRegion(locationRegionRegion *string) *DeleteBuildParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the delete build params
func (o *DeleteBuildParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
