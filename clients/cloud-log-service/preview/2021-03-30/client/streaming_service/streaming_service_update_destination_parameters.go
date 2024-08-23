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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// NewStreamingServiceUpdateDestinationParams creates a new StreamingServiceUpdateDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamingServiceUpdateDestinationParams() *StreamingServiceUpdateDestinationParams {
	return &StreamingServiceUpdateDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamingServiceUpdateDestinationParamsWithTimeout creates a new StreamingServiceUpdateDestinationParams object
// with the ability to set a timeout on a request.
func NewStreamingServiceUpdateDestinationParamsWithTimeout(timeout time.Duration) *StreamingServiceUpdateDestinationParams {
	return &StreamingServiceUpdateDestinationParams{
		timeout: timeout,
	}
}

// NewStreamingServiceUpdateDestinationParamsWithContext creates a new StreamingServiceUpdateDestinationParams object
// with the ability to set a context for a request.
func NewStreamingServiceUpdateDestinationParamsWithContext(ctx context.Context) *StreamingServiceUpdateDestinationParams {
	return &StreamingServiceUpdateDestinationParams{
		Context: ctx,
	}
}

// NewStreamingServiceUpdateDestinationParamsWithHTTPClient creates a new StreamingServiceUpdateDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamingServiceUpdateDestinationParamsWithHTTPClient(client *http.Client) *StreamingServiceUpdateDestinationParams {
	return &StreamingServiceUpdateDestinationParams{
		HTTPClient: client,
	}
}

/*
StreamingServiceUpdateDestinationParams contains all the parameters to send to the API endpoint

	for the streaming service update destination operation.

	Typically these are written to a http.Request.
*/
type StreamingServiceUpdateDestinationParams struct {

	// Body.
	Body *models.LogService20210330UpdateDestinationRequest

	/* ID.

	   id - the destination's id, required.
	*/
	ID string

	/* OrganizationID.

	   organization_id - the organization holding the destination, required.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the streaming service update destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceUpdateDestinationParams) WithDefaults() *StreamingServiceUpdateDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the streaming service update destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamingServiceUpdateDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithTimeout(timeout time.Duration) *StreamingServiceUpdateDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithContext(ctx context.Context) *StreamingServiceUpdateDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithHTTPClient(client *http.Client) *StreamingServiceUpdateDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithBody(body *models.LogService20210330UpdateDestinationRequest) *StreamingServiceUpdateDestinationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetBody(body *models.LogService20210330UpdateDestinationRequest) {
	o.Body = body
}

// WithID adds the id to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithID(id string) *StreamingServiceUpdateDestinationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) WithOrganizationID(organizationID string) *StreamingServiceUpdateDestinationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the streaming service update destination params
func (o *StreamingServiceUpdateDestinationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *StreamingServiceUpdateDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}