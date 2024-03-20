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

// NewCreateTwilioRotationIntegrationParams creates a new CreateTwilioRotationIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTwilioRotationIntegrationParams() *CreateTwilioRotationIntegrationParams {
	return &CreateTwilioRotationIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTwilioRotationIntegrationParamsWithTimeout creates a new CreateTwilioRotationIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateTwilioRotationIntegrationParamsWithTimeout(timeout time.Duration) *CreateTwilioRotationIntegrationParams {
	return &CreateTwilioRotationIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateTwilioRotationIntegrationParamsWithContext creates a new CreateTwilioRotationIntegrationParams object
// with the ability to set a context for a request.
func NewCreateTwilioRotationIntegrationParamsWithContext(ctx context.Context) *CreateTwilioRotationIntegrationParams {
	return &CreateTwilioRotationIntegrationParams{
		Context: ctx,
	}
}

// NewCreateTwilioRotationIntegrationParamsWithHTTPClient creates a new CreateTwilioRotationIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTwilioRotationIntegrationParamsWithHTTPClient(client *http.Client) *CreateTwilioRotationIntegrationParams {
	return &CreateTwilioRotationIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateTwilioRotationIntegrationParams contains all the parameters to send to the API endpoint

	for the create twilio rotation integration operation.

	Typically these are written to a http.Request.
*/
type CreateTwilioRotationIntegrationParams struct {

	// Body.
	Body CreateTwilioRotationIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create twilio rotation integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTwilioRotationIntegrationParams) WithDefaults() *CreateTwilioRotationIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create twilio rotation integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTwilioRotationIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithTimeout(timeout time.Duration) *CreateTwilioRotationIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithContext(ctx context.Context) *CreateTwilioRotationIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithHTTPClient(client *http.Client) *CreateTwilioRotationIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithBody(body CreateTwilioRotationIntegrationBody) *CreateTwilioRotationIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetBody(body CreateTwilioRotationIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithOrganizationID(organizationID string) *CreateTwilioRotationIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) WithProjectID(projectID string) *CreateTwilioRotationIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create twilio rotation integration params
func (o *CreateTwilioRotationIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTwilioRotationIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}