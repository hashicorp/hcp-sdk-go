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

// NewLogServiceGetStreamingDestination2Params creates a new LogServiceGetStreamingDestination2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceGetStreamingDestination2Params() *LogServiceGetStreamingDestination2Params {
	return &LogServiceGetStreamingDestination2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceGetStreamingDestination2ParamsWithTimeout creates a new LogServiceGetStreamingDestination2Params object
// with the ability to set a timeout on a request.
func NewLogServiceGetStreamingDestination2ParamsWithTimeout(timeout time.Duration) *LogServiceGetStreamingDestination2Params {
	return &LogServiceGetStreamingDestination2Params{
		timeout: timeout,
	}
}

// NewLogServiceGetStreamingDestination2ParamsWithContext creates a new LogServiceGetStreamingDestination2Params object
// with the ability to set a context for a request.
func NewLogServiceGetStreamingDestination2ParamsWithContext(ctx context.Context) *LogServiceGetStreamingDestination2Params {
	return &LogServiceGetStreamingDestination2Params{
		Context: ctx,
	}
}

// NewLogServiceGetStreamingDestination2ParamsWithHTTPClient creates a new LogServiceGetStreamingDestination2Params object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceGetStreamingDestination2ParamsWithHTTPClient(client *http.Client) *LogServiceGetStreamingDestination2Params {
	return &LogServiceGetStreamingDestination2Params{
		HTTPClient: client,
	}
}

/*
LogServiceGetStreamingDestination2Params contains all the parameters to send to the API endpoint

	for the log service get streaming destination2 operation.

	Typically these are written to a http.Request.
*/
type LogServiceGetStreamingDestination2Params struct {

	/* DestinationID.

	   destination_id is represents the ID of destination we want to create the filter for.
	*/
	DestinationID string

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

// WithDefaults hydrates default values in the log service get streaming destination2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetStreamingDestination2Params) WithDefaults() *LogServiceGetStreamingDestination2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service get streaming destination2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetStreamingDestination2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithTimeout(timeout time.Duration) *LogServiceGetStreamingDestination2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithContext(ctx context.Context) *LogServiceGetStreamingDestination2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithHTTPClient(client *http.Client) *LogServiceGetStreamingDestination2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithDestinationID(destinationID string) *LogServiceGetStreamingDestination2Params {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithLocationOrganizationID(locationOrganizationID string) *LogServiceGetStreamingDestination2Params {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithLocationProjectID(locationProjectID *string) *LogServiceGetStreamingDestination2Params {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetLocationProjectID(locationProjectID *string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithLocationRegionProvider(locationRegionProvider *string) *LogServiceGetStreamingDestination2Params {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) WithLocationRegionRegion(locationRegionRegion *string) *LogServiceGetStreamingDestination2Params {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the log service get streaming destination2 params
func (o *LogServiceGetStreamingDestination2Params) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceGetStreamingDestination2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param destination_id
	if err := r.SetPathParam("destination_id", o.DestinationID); err != nil {
		return err
	}

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
