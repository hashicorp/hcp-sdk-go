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

// NewLogServiceDeleteStreamingDestinationParams creates a new LogServiceDeleteStreamingDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceDeleteStreamingDestinationParams() *LogServiceDeleteStreamingDestinationParams {
	return &LogServiceDeleteStreamingDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceDeleteStreamingDestinationParamsWithTimeout creates a new LogServiceDeleteStreamingDestinationParams object
// with the ability to set a timeout on a request.
func NewLogServiceDeleteStreamingDestinationParamsWithTimeout(timeout time.Duration) *LogServiceDeleteStreamingDestinationParams {
	return &LogServiceDeleteStreamingDestinationParams{
		timeout: timeout,
	}
}

// NewLogServiceDeleteStreamingDestinationParamsWithContext creates a new LogServiceDeleteStreamingDestinationParams object
// with the ability to set a context for a request.
func NewLogServiceDeleteStreamingDestinationParamsWithContext(ctx context.Context) *LogServiceDeleteStreamingDestinationParams {
	return &LogServiceDeleteStreamingDestinationParams{
		Context: ctx,
	}
}

// NewLogServiceDeleteStreamingDestinationParamsWithHTTPClient creates a new LogServiceDeleteStreamingDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceDeleteStreamingDestinationParamsWithHTTPClient(client *http.Client) *LogServiceDeleteStreamingDestinationParams {
	return &LogServiceDeleteStreamingDestinationParams{
		HTTPClient: client,
	}
}

/*
LogServiceDeleteStreamingDestinationParams contains all the parameters to send to the API endpoint

	for the log service delete streaming destination operation.

	Typically these are written to a http.Request.
*/
type LogServiceDeleteStreamingDestinationParams struct {

	/* DestinationID.

	   destination_id represents the destination id to delete.
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

	/* SourceChannel.

	   source_channel is the source channel from which the delete request came from. For example terraform vs HCP Portal.
	*/
	SourceChannel *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log service delete streaming destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceDeleteStreamingDestinationParams) WithDefaults() *LogServiceDeleteStreamingDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service delete streaming destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceDeleteStreamingDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithTimeout(timeout time.Duration) *LogServiceDeleteStreamingDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithContext(ctx context.Context) *LogServiceDeleteStreamingDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithHTTPClient(client *http.Client) *LogServiceDeleteStreamingDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithDestinationID(destinationID string) *LogServiceDeleteStreamingDestinationParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithLocationOrganizationID(locationOrganizationID string) *LogServiceDeleteStreamingDestinationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithLocationProjectID(locationProjectID string) *LogServiceDeleteStreamingDestinationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithLocationRegionProvider(locationRegionProvider *string) *LogServiceDeleteStreamingDestinationParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithLocationRegionRegion(locationRegionRegion *string) *LogServiceDeleteStreamingDestinationParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithSourceChannel adds the sourceChannel to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) WithSourceChannel(sourceChannel *string) *LogServiceDeleteStreamingDestinationParams {
	o.SetSourceChannel(sourceChannel)
	return o
}

// SetSourceChannel adds the sourceChannel to the log service delete streaming destination params
func (o *LogServiceDeleteStreamingDestinationParams) SetSourceChannel(sourceChannel *string) {
	o.SourceChannel = sourceChannel
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceDeleteStreamingDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SourceChannel != nil {

		// query param source_channel
		var qrSourceChannel string

		if o.SourceChannel != nil {
			qrSourceChannel = *o.SourceChannel
		}
		qSourceChannel := qrSourceChannel
		if qSourceChannel != "" {

			if err := r.SetQueryParam("source_channel", qSourceChannel); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
