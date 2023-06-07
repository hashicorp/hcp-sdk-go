// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewUpdateAwsSmSyncIntegrationParams creates a new UpdateAwsSmSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAwsSmSyncIntegrationParams() *UpdateAwsSmSyncIntegrationParams {
	return &UpdateAwsSmSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAwsSmSyncIntegrationParamsWithTimeout creates a new UpdateAwsSmSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpdateAwsSmSyncIntegrationParamsWithTimeout(timeout time.Duration) *UpdateAwsSmSyncIntegrationParams {
	return &UpdateAwsSmSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewUpdateAwsSmSyncIntegrationParamsWithContext creates a new UpdateAwsSmSyncIntegrationParams object
// with the ability to set a context for a request.
func NewUpdateAwsSmSyncIntegrationParamsWithContext(ctx context.Context) *UpdateAwsSmSyncIntegrationParams {
	return &UpdateAwsSmSyncIntegrationParams{
		Context: ctx,
	}
}

// NewUpdateAwsSmSyncIntegrationParamsWithHTTPClient creates a new UpdateAwsSmSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAwsSmSyncIntegrationParamsWithHTTPClient(client *http.Client) *UpdateAwsSmSyncIntegrationParams {
	return &UpdateAwsSmSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpdateAwsSmSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the update aws sm sync integration operation.

	Typically these are written to a http.Request.
*/
type UpdateAwsSmSyncIntegrationParams struct {

	// Body.
	Body UpdateAwsSmSyncIntegrationBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update aws sm sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsSmSyncIntegrationParams) WithDefaults() *UpdateAwsSmSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update aws sm sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsSmSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithTimeout(timeout time.Duration) *UpdateAwsSmSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithContext(ctx context.Context) *UpdateAwsSmSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithHTTPClient(client *http.Client) *UpdateAwsSmSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithBody(body UpdateAwsSmSyncIntegrationBody) *UpdateAwsSmSyncIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetBody(body UpdateAwsSmSyncIntegrationBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateAwsSmSyncIntegrationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithLocationProjectID(locationProjectID string) *UpdateAwsSmSyncIntegrationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithName adds the name to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) WithName(name string) *UpdateAwsSmSyncIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update aws sm sync integration params
func (o *UpdateAwsSmSyncIntegrationParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAwsSmSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}