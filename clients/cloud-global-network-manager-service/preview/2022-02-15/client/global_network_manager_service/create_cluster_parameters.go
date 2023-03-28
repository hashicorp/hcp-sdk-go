// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

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

// NewCreateClusterParams creates a new CreateClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterParams() *CreateClusterParams {
	return &CreateClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterParamsWithTimeout creates a new CreateClusterParams object
// with the ability to set a timeout on a request.
func NewCreateClusterParamsWithTimeout(timeout time.Duration) *CreateClusterParams {
	return &CreateClusterParams{
		timeout: timeout,
	}
}

// NewCreateClusterParamsWithContext creates a new CreateClusterParams object
// with the ability to set a context for a request.
func NewCreateClusterParamsWithContext(ctx context.Context) *CreateClusterParams {
	return &CreateClusterParams{
		Context: ctx,
	}
}

// NewCreateClusterParamsWithHTTPClient creates a new CreateClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterParamsWithHTTPClient(client *http.Client) *CreateClusterParams {
	return &CreateClusterParams{
		HTTPClient: client,
	}
}

/*
CreateClusterParams contains all the parameters to send to the API endpoint

	for the create cluster operation.

	Typically these are written to a http.Request.
*/
type CreateClusterParams struct {

	// Body.
	Body CreateClusterBody

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

// WithDefaults hydrates default values in the create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterParams) WithDefaults() *CreateClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) WithTimeout(timeout time.Duration) *CreateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster params
func (o *CreateClusterParams) WithContext(ctx context.Context) *CreateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster params
func (o *CreateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) WithHTTPClient(client *http.Client) *CreateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster params
func (o *CreateClusterParams) WithBody(body CreateClusterBody) *CreateClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster params
func (o *CreateClusterParams) SetBody(body CreateClusterBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the create cluster params
func (o *CreateClusterParams) WithLocationOrganizationID(locationOrganizationID string) *CreateClusterParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the create cluster params
func (o *CreateClusterParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the create cluster params
func (o *CreateClusterParams) WithLocationProjectID(locationProjectID string) *CreateClusterParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the create cluster params
func (o *CreateClusterParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
