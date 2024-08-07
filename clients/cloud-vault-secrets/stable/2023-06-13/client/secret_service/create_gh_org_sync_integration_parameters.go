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

// NewCreateGhOrgSyncIntegrationParams creates a new CreateGhOrgSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGhOrgSyncIntegrationParams() *CreateGhOrgSyncIntegrationParams {
	return &CreateGhOrgSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGhOrgSyncIntegrationParamsWithTimeout creates a new CreateGhOrgSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateGhOrgSyncIntegrationParamsWithTimeout(timeout time.Duration) *CreateGhOrgSyncIntegrationParams {
	return &CreateGhOrgSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateGhOrgSyncIntegrationParamsWithContext creates a new CreateGhOrgSyncIntegrationParams object
// with the ability to set a context for a request.
func NewCreateGhOrgSyncIntegrationParamsWithContext(ctx context.Context) *CreateGhOrgSyncIntegrationParams {
	return &CreateGhOrgSyncIntegrationParams{
		Context: ctx,
	}
}

// NewCreateGhOrgSyncIntegrationParamsWithHTTPClient creates a new CreateGhOrgSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGhOrgSyncIntegrationParamsWithHTTPClient(client *http.Client) *CreateGhOrgSyncIntegrationParams {
	return &CreateGhOrgSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateGhOrgSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the create gh org sync integration operation.

	Typically these are written to a http.Request.
*/
type CreateGhOrgSyncIntegrationParams struct {

	// Body.
	Body CreateGhOrgSyncIntegrationBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create gh org sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGhOrgSyncIntegrationParams) WithDefaults() *CreateGhOrgSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create gh org sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGhOrgSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithTimeout(timeout time.Duration) *CreateGhOrgSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithContext(ctx context.Context) *CreateGhOrgSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithHTTPClient(client *http.Client) *CreateGhOrgSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithBody(body CreateGhOrgSyncIntegrationBody) *CreateGhOrgSyncIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetBody(body CreateGhOrgSyncIntegrationBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithLocationOrganizationID(locationOrganizationID string) *CreateGhOrgSyncIntegrationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) WithLocationProjectID(locationProjectID string) *CreateGhOrgSyncIntegrationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the create gh org sync integration params
func (o *CreateGhOrgSyncIntegrationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGhOrgSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
