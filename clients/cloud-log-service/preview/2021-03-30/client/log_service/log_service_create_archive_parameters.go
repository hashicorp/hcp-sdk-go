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

// NewLogServiceCreateArchiveParams creates a new LogServiceCreateArchiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogServiceCreateArchiveParams() *LogServiceCreateArchiveParams {
	return &LogServiceCreateArchiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogServiceCreateArchiveParamsWithTimeout creates a new LogServiceCreateArchiveParams object
// with the ability to set a timeout on a request.
func NewLogServiceCreateArchiveParamsWithTimeout(timeout time.Duration) *LogServiceCreateArchiveParams {
	return &LogServiceCreateArchiveParams{
		timeout: timeout,
	}
}

// NewLogServiceCreateArchiveParamsWithContext creates a new LogServiceCreateArchiveParams object
// with the ability to set a context for a request.
func NewLogServiceCreateArchiveParamsWithContext(ctx context.Context) *LogServiceCreateArchiveParams {
	return &LogServiceCreateArchiveParams{
		Context: ctx,
	}
}

// NewLogServiceCreateArchiveParamsWithHTTPClient creates a new LogServiceCreateArchiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogServiceCreateArchiveParamsWithHTTPClient(client *http.Client) *LogServiceCreateArchiveParams {
	return &LogServiceCreateArchiveParams{
		HTTPClient: client,
	}
}

/*
LogServiceCreateArchiveParams contains all the parameters to send to the API endpoint

	for the log service create archive operation.

	Typically these are written to a http.Request.
*/
type LogServiceCreateArchiveParams struct {

	// Body.
	Body *models.LogService20210330CreateArchiveRequest

	/* ResourceID.

	   id is the identifier for this resource.
	*/
	ResourceID string

	/* ResourceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	ResourceLocationOrganizationID string

	/* ResourceLocationProjectID.

	   project_id is the projects id.
	*/
	ResourceLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log service create archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceCreateArchiveParams) WithDefaults() *LogServiceCreateArchiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log service create archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogServiceCreateArchiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithTimeout(timeout time.Duration) *LogServiceCreateArchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithContext(ctx context.Context) *LogServiceCreateArchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithHTTPClient(client *http.Client) *LogServiceCreateArchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithBody(body *models.LogService20210330CreateArchiveRequest) *LogServiceCreateArchiveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetBody(body *models.LogService20210330CreateArchiveRequest) {
	o.Body = body
}

// WithResourceID adds the resourceID to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithResourceID(resourceID string) *LogServiceCreateArchiveParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceLocationOrganizationID adds the resourceLocationOrganizationID to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithResourceLocationOrganizationID(resourceLocationOrganizationID string) *LogServiceCreateArchiveParams {
	o.SetResourceLocationOrganizationID(resourceLocationOrganizationID)
	return o
}

// SetResourceLocationOrganizationID adds the resourceLocationOrganizationId to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetResourceLocationOrganizationID(resourceLocationOrganizationID string) {
	o.ResourceLocationOrganizationID = resourceLocationOrganizationID
}

// WithResourceLocationProjectID adds the resourceLocationProjectID to the log service create archive params
func (o *LogServiceCreateArchiveParams) WithResourceLocationProjectID(resourceLocationProjectID string) *LogServiceCreateArchiveParams {
	o.SetResourceLocationProjectID(resourceLocationProjectID)
	return o
}

// SetResourceLocationProjectID adds the resourceLocationProjectId to the log service create archive params
func (o *LogServiceCreateArchiveParams) SetResourceLocationProjectID(resourceLocationProjectID string) {
	o.ResourceLocationProjectID = resourceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *LogServiceCreateArchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource.id
	if err := r.SetPathParam("resource.id", o.ResourceID); err != nil {
		return err
	}

	// path param resource.location.organization_id
	if err := r.SetPathParam("resource.location.organization_id", o.ResourceLocationOrganizationID); err != nil {
		return err
	}

	// path param resource.location.project_id
	if err := r.SetPathParam("resource.location.project_id", o.ResourceLocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
