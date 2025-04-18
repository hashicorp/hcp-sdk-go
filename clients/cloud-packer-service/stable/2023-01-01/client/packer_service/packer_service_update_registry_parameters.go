// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
)

// NewPackerServiceUpdateRegistryParams creates a new PackerServiceUpdateRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceUpdateRegistryParams() *PackerServiceUpdateRegistryParams {
	return &PackerServiceUpdateRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceUpdateRegistryParamsWithTimeout creates a new PackerServiceUpdateRegistryParams object
// with the ability to set a timeout on a request.
func NewPackerServiceUpdateRegistryParamsWithTimeout(timeout time.Duration) *PackerServiceUpdateRegistryParams {
	return &PackerServiceUpdateRegistryParams{
		timeout: timeout,
	}
}

// NewPackerServiceUpdateRegistryParamsWithContext creates a new PackerServiceUpdateRegistryParams object
// with the ability to set a context for a request.
func NewPackerServiceUpdateRegistryParamsWithContext(ctx context.Context) *PackerServiceUpdateRegistryParams {
	return &PackerServiceUpdateRegistryParams{
		Context: ctx,
	}
}

// NewPackerServiceUpdateRegistryParamsWithHTTPClient creates a new PackerServiceUpdateRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceUpdateRegistryParamsWithHTTPClient(client *http.Client) *PackerServiceUpdateRegistryParams {
	return &PackerServiceUpdateRegistryParams{
		HTTPClient: client,
	}
}

/*
PackerServiceUpdateRegistryParams contains all the parameters to send to the API endpoint

	for the packer service update registry operation.

	Typically these are written to a http.Request.
*/
type PackerServiceUpdateRegistryParams struct {

	// Body.
	Body *models.HashicorpCloudPacker20230101UpdateRegistryBody

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

// WithDefaults hydrates default values in the packer service update registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateRegistryParams) WithDefaults() *PackerServiceUpdateRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service update registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithTimeout(timeout time.Duration) *PackerServiceUpdateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithContext(ctx context.Context) *PackerServiceUpdateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithHTTPClient(client *http.Client) *PackerServiceUpdateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithBody(body *models.HashicorpCloudPacker20230101UpdateRegistryBody) *PackerServiceUpdateRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetBody(body *models.HashicorpCloudPacker20230101UpdateRegistryBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceUpdateRegistryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) WithLocationProjectID(locationProjectID string) *PackerServiceUpdateRegistryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service update registry params
func (o *PackerServiceUpdateRegistryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceUpdateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
