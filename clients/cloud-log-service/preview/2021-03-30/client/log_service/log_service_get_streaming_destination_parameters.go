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

// NewLogServiceGetStreamingDestinationParams creates a new LogServiceGetStreamingDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceGetStreamingDestinationParams() *LogServiceGetStreamingDestinationParams {
	return &LogServiceGetStreamingDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceGetStreamingDestinationParamsWithTimeout creates a new LogServiceGetStreamingDestinationParams object
// with the ability to set a timeout on a request.
func NewLogServiceGetStreamingDestinationParamsWithTimeout(timeout time.Duration) *LogServiceGetStreamingDestinationParams {
	return &LogServiceGetStreamingDestinationParams{
		timeout: timeout,
	}
}

// NewLogServiceGetStreamingDestinationParamsWithContext creates a new LogServiceGetStreamingDestinationParams object
// with the ability to set a context for a request.
func NewLogServiceGetStreamingDestinationParamsWithContext(ctx context.Context) *LogServiceGetStreamingDestinationParams {
	return &LogServiceGetStreamingDestinationParams{
		Context: ctx,
	}
}

// NewLogServiceGetStreamingDestinationParamsWithHTTPClient creates a new LogServiceGetStreamingDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceGetStreamingDestinationParamsWithHTTPClient(client *http.Client) *LogServiceGetStreamingDestinationParams {
	return &LogServiceGetStreamingDestinationParams{
		HTTPClient: client,
	}
}

/*
LogServiceGetStreamingDestinationParams contains all the parameters to send to the API endpoint

	for the log service get streaming destination operation.

	Typically these are written to a http.Request.
*/
type LogServiceGetStreamingDestinationParams struct {

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

// WithDefaults hydrates default values in the log service get streaming destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetStreamingDestinationParams) WithDefaults() *LogServiceGetStreamingDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service get streaming destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceGetStreamingDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithTimeout(timeout time.Duration) *LogServiceGetStreamingDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithContext(ctx context.Context) *LogServiceGetStreamingDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithHTTPClient(client *http.Client) *LogServiceGetStreamingDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithDestinationID(destinationID string) *LogServiceGetStreamingDestinationParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithLocationOrganizationID(locationOrganizationID string) *LogServiceGetStreamingDestinationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithLocationProjectID(locationProjectID string) *LogServiceGetStreamingDestinationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithLocationRegionProvider(locationRegionProvider *string) *LogServiceGetStreamingDestinationParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) WithLocationRegionRegion(locationRegionRegion *string) *LogServiceGetStreamingDestinationParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the log service get streaming destination params
func (o *LogServiceGetStreamingDestinationParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceGetStreamingDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
