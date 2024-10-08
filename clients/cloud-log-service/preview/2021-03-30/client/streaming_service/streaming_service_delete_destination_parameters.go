// Code generated by go-swagger; DO NOT EDIT.

package streaming_service

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

// NewStreamingServiceDeleteDestinationParams creates a new StreamingServiceDeleteDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamingServiceDeleteDestinationParams() *StreamingServiceDeleteDestinationParams {
	return &StreamingServiceDeleteDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamingServiceDeleteDestinationParamsWithTimeout creates a new StreamingServiceDeleteDestinationParams object
// with the ability to set a timeout on a request.
func NewStreamingServiceDeleteDestinationParamsWithTimeout(timeout time.Duration) *StreamingServiceDeleteDestinationParams {
	return &StreamingServiceDeleteDestinationParams{
		timeout: timeout,
	}
}

// NewStreamingServiceDeleteDestinationParamsWithContext creates a new StreamingServiceDeleteDestinationParams object
// with the ability to set a context for a request.
func NewStreamingServiceDeleteDestinationParamsWithContext(ctx context.Context) *StreamingServiceDeleteDestinationParams {
	return &StreamingServiceDeleteDestinationParams{
		Context: ctx,
	}
}

// NewStreamingServiceDeleteDestinationParamsWithHTTPClient creates a new StreamingServiceDeleteDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamingServiceDeleteDestinationParamsWithHTTPClient(client *http.Client) *StreamingServiceDeleteDestinationParams {
	return &StreamingServiceDeleteDestinationParams{
		HTTPClient: client,
	}
}

/*
StreamingServiceDeleteDestinationParams contains all the parameters to send to the API endpoint

	for the streaming service delete destination operation.

	Typically these are written to a http.Request.
*/
type StreamingServiceDeleteDestinationParams struct {

	/* DestinationID.

	   destination_id represents the destination id to delete.
	*/
	DestinationID string

	// OrganizationID.
	OrganizationID string

	/* SourceChannel.

	   source_channel is the source channel from which the delete request came from. For example terraform vs HCP Portal.
	*/
	SourceChannel *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the streaming service delete destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceDeleteDestinationParams) WithDefaults() *StreamingServiceDeleteDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the streaming service delete destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceDeleteDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithTimeout(timeout time.Duration) *StreamingServiceDeleteDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithContext(ctx context.Context) *StreamingServiceDeleteDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithHTTPClient(client *http.Client) *StreamingServiceDeleteDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithDestinationID(destinationID string) *StreamingServiceDeleteDestinationParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithOrganizationID adds the organizationID to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithOrganizationID(organizationID string) *StreamingServiceDeleteDestinationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithSourceChannel adds the sourceChannel to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) WithSourceChannel(sourceChannel *string) *StreamingServiceDeleteDestinationParams {
	o.SetSourceChannel(sourceChannel)
	return o
}

// SetSourceChannel adds the sourceChannel to the streaming service delete destination params
func (o *StreamingServiceDeleteDestinationParams) SetSourceChannel(sourceChannel *string) {
	o.SourceChannel = sourceChannel
}

// WriteToRequest writes these params to a swagger request
func (o *StreamingServiceDeleteDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param destination_id
	if err := r.SetPathParam("destination_id", o.DestinationID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
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
