// Code generated by go-swagger; DO NOT EDIT.

package log_service

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

// NewLogServiceListStreamingDestinations2Params creates a new LogServiceListStreamingDestinations2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceListStreamingDestinations2Params() *LogServiceListStreamingDestinations2Params {
	return &LogServiceListStreamingDestinations2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceListStreamingDestinations2ParamsWithTimeout creates a new LogServiceListStreamingDestinations2Params object
// with the ability to set a timeout on a request.
func NewLogServiceListStreamingDestinations2ParamsWithTimeout(timeout time.Duration) *LogServiceListStreamingDestinations2Params {
	return &LogServiceListStreamingDestinations2Params{
		timeout: timeout,
	}
}

// NewLogServiceListStreamingDestinations2ParamsWithContext creates a new LogServiceListStreamingDestinations2Params object
// with the ability to set a context for a request.
func NewLogServiceListStreamingDestinations2ParamsWithContext(ctx context.Context) *LogServiceListStreamingDestinations2Params {
	return &LogServiceListStreamingDestinations2Params{
		Context: ctx,
	}
}

// NewLogServiceListStreamingDestinations2ParamsWithHTTPClient creates a new LogServiceListStreamingDestinations2Params object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceListStreamingDestinations2ParamsWithHTTPClient(client *http.Client) *LogServiceListStreamingDestinations2Params {
	return &LogServiceListStreamingDestinations2Params{
		HTTPClient: client,
	}
}

/*
LogServiceListStreamingDestinations2Params contains all the parameters to send to the API endpoint

	for the log service list streaming destinations2 operation.

	Typically these are written to a http.Request.
*/
type LogServiceListStreamingDestinations2Params struct {

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID *string

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

// WithDefaults hydrates default values in the log service list streaming destinations2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceListStreamingDestinations2Params) WithDefaults() *LogServiceListStreamingDestinations2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service list streaming destinations2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceListStreamingDestinations2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithTimeout(timeout time.Duration) *LogServiceListStreamingDestinations2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithContext(ctx context.Context) *LogServiceListStreamingDestinations2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithHTTPClient(client *http.Client) *LogServiceListStreamingDestinations2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithLocationOrganizationID(locationOrganizationID string) *LogServiceListStreamingDestinations2Params {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithLocationProjectID(locationProjectID *string) *LogServiceListStreamingDestinations2Params {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetLocationProjectID(locationProjectID *string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithLocationRegionProvider(locationRegionProvider *string) *LogServiceListStreamingDestinations2Params {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) WithLocationRegionRegion(locationRegionRegion *string) *LogServiceListStreamingDestinations2Params {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the log service list streaming destinations2 params
func (o *LogServiceListStreamingDestinations2Params) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceListStreamingDestinations2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	if o.LocationProjectID != nil {

		// query param location.project_id
		var qrLocationProjectID string

		if o.LocationProjectID != nil {
			qrLocationProjectID = *o.LocationProjectID
		}
		qLocationProjectID := qrLocationProjectID
		if qLocationProjectID != "" {

			if err := r.SetQueryParam("location.project_id", qLocationProjectID); err != nil {
				return err
			}
		}
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
