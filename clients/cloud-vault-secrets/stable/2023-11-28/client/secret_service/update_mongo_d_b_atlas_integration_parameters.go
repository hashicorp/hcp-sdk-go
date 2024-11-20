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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// NewUpdateMongoDBAtlasIntegrationParams creates a new UpdateMongoDBAtlasIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMongoDBAtlasIntegrationParams() *UpdateMongoDBAtlasIntegrationParams {
	return &UpdateMongoDBAtlasIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMongoDBAtlasIntegrationParamsWithTimeout creates a new UpdateMongoDBAtlasIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpdateMongoDBAtlasIntegrationParamsWithTimeout(timeout time.Duration) *UpdateMongoDBAtlasIntegrationParams {
	return &UpdateMongoDBAtlasIntegrationParams{
		timeout: timeout,
	}
}

// NewUpdateMongoDBAtlasIntegrationParamsWithContext creates a new UpdateMongoDBAtlasIntegrationParams object
// with the ability to set a context for a request.
func NewUpdateMongoDBAtlasIntegrationParamsWithContext(ctx context.Context) *UpdateMongoDBAtlasIntegrationParams {
	return &UpdateMongoDBAtlasIntegrationParams{
		Context: ctx,
	}
}

// NewUpdateMongoDBAtlasIntegrationParamsWithHTTPClient creates a new UpdateMongoDBAtlasIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMongoDBAtlasIntegrationParamsWithHTTPClient(client *http.Client) *UpdateMongoDBAtlasIntegrationParams {
	return &UpdateMongoDBAtlasIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpdateMongoDBAtlasIntegrationParams contains all the parameters to send to the API endpoint

	for the update mongo d b atlas integration operation.

	Typically these are written to a http.Request.
*/
type UpdateMongoDBAtlasIntegrationParams struct {

	// Body.
	Body *models.SecretServiceUpdateMongoDBAtlasIntegrationBody

	// Name.
	Name string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mongo d b atlas integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMongoDBAtlasIntegrationParams) WithDefaults() *UpdateMongoDBAtlasIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mongo d b atlas integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMongoDBAtlasIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithTimeout(timeout time.Duration) *UpdateMongoDBAtlasIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithContext(ctx context.Context) *UpdateMongoDBAtlasIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithHTTPClient(client *http.Client) *UpdateMongoDBAtlasIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithBody(body *models.SecretServiceUpdateMongoDBAtlasIntegrationBody) *UpdateMongoDBAtlasIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetBody(body *models.SecretServiceUpdateMongoDBAtlasIntegrationBody) {
	o.Body = body
}

// WithName adds the name to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithName(name string) *UpdateMongoDBAtlasIntegrationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithOrganizationID(organizationID string) *UpdateMongoDBAtlasIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) WithProjectID(projectID string) *UpdateMongoDBAtlasIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update mongo d b atlas integration params
func (o *UpdateMongoDBAtlasIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMongoDBAtlasIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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