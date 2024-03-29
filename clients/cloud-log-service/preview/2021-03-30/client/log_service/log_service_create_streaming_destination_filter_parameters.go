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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-log-service/preview/2021-03-30/models"
)

// NewLogServiceCreateStreamingDestinationFilterParams creates a new LogServiceCreateStreamingDestinationFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceCreateStreamingDestinationFilterParams() *LogServiceCreateStreamingDestinationFilterParams {
	return &LogServiceCreateStreamingDestinationFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceCreateStreamingDestinationFilterParamsWithTimeout creates a new LogServiceCreateStreamingDestinationFilterParams object
// with the ability to set a timeout on a request.
func NewLogServiceCreateStreamingDestinationFilterParamsWithTimeout(timeout time.Duration) *LogServiceCreateStreamingDestinationFilterParams {
	return &LogServiceCreateStreamingDestinationFilterParams{
		timeout: timeout,
	}
}

// NewLogServiceCreateStreamingDestinationFilterParamsWithContext creates a new LogServiceCreateStreamingDestinationFilterParams object
// with the ability to set a context for a request.
func NewLogServiceCreateStreamingDestinationFilterParamsWithContext(ctx context.Context) *LogServiceCreateStreamingDestinationFilterParams {
	return &LogServiceCreateStreamingDestinationFilterParams{
		Context: ctx,
	}
}

// NewLogServiceCreateStreamingDestinationFilterParamsWithHTTPClient creates a new LogServiceCreateStreamingDestinationFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceCreateStreamingDestinationFilterParamsWithHTTPClient(client *http.Client) *LogServiceCreateStreamingDestinationFilterParams {
	return &LogServiceCreateStreamingDestinationFilterParams{
		HTTPClient: client,
	}
}

/*
LogServiceCreateStreamingDestinationFilterParams contains all the parameters to send to the API endpoint

	for the log service create streaming destination filter operation.

	Typically these are written to a http.Request.
*/
type LogServiceCreateStreamingDestinationFilterParams struct {

	// Body.
	Body *models.LogService20210330CreateStreamingDestinationFilterRequest

	/* DestinationID.

	   destination_id is represents the ID of destination we want to create the filter for.
	*/
	DestinationID string

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log service create streaming destination filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceCreateStreamingDestinationFilterParams) WithDefaults() *LogServiceCreateStreamingDestinationFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service create streaming destination filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceCreateStreamingDestinationFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithTimeout(timeout time.Duration) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithContext(ctx context.Context) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithHTTPClient(client *http.Client) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithBody(body *models.LogService20210330CreateStreamingDestinationFilterRequest) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetBody(body *models.LogService20210330CreateStreamingDestinationFilterRequest) {
	o.Body = body
}

// WithDestinationID adds the destinationID to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithDestinationID(destinationID string) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithOrganizationID adds the organizationID to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) WithOrganizationID(organizationID string) *LogServiceCreateStreamingDestinationFilterParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the log service create streaming destination filter params
func (o *LogServiceCreateStreamingDestinationFilterParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceCreateStreamingDestinationFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param destination_id
	if err := r.SetPathParam("destination_id", o.DestinationID); err != nil {
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
