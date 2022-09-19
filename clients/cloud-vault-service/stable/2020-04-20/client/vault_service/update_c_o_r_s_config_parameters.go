// Code generated by go-swagger; DO NOT EDIT.

package vault_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-04-20/models"
)

// NewUpdateCORSConfigParams creates a new UpdateCORSConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCORSConfigParams() *UpdateCORSConfigParams {
	return &UpdateCORSConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCORSConfigParamsWithTimeout creates a new UpdateCORSConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateCORSConfigParamsWithTimeout(timeout time.Duration) *UpdateCORSConfigParams {
	return &UpdateCORSConfigParams{
		timeout: timeout,
	}
}

// NewUpdateCORSConfigParamsWithContext creates a new UpdateCORSConfigParams object
// with the ability to set a context for a request.
func NewUpdateCORSConfigParamsWithContext(ctx context.Context) *UpdateCORSConfigParams {
	return &UpdateCORSConfigParams{
		Context: ctx,
	}
}

// NewUpdateCORSConfigParamsWithHTTPClient creates a new UpdateCORSConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCORSConfigParamsWithHTTPClient(client *http.Client) *UpdateCORSConfigParams {
	return &UpdateCORSConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateCORSConfigParams contains all the parameters to send to the API endpoint

	for the update c o r s config operation.

	Typically these are written to a http.Request.
*/
type UpdateCORSConfigParams struct {

	// Body.
	Body *models.HashicorpCloudVault20200420UpdateCORSConfigRequest

	// ClusterID.
	ClusterID string

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

// WithDefaults hydrates default values in the update c o r s config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCORSConfigParams) WithDefaults() *UpdateCORSConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update c o r s config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCORSConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update c o r s config params
func (o *UpdateCORSConfigParams) WithTimeout(timeout time.Duration) *UpdateCORSConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update c o r s config params
func (o *UpdateCORSConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update c o r s config params
func (o *UpdateCORSConfigParams) WithContext(ctx context.Context) *UpdateCORSConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update c o r s config params
func (o *UpdateCORSConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update c o r s config params
func (o *UpdateCORSConfigParams) WithHTTPClient(client *http.Client) *UpdateCORSConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update c o r s config params
func (o *UpdateCORSConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update c o r s config params
func (o *UpdateCORSConfigParams) WithBody(body *models.HashicorpCloudVault20200420UpdateCORSConfigRequest) *UpdateCORSConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update c o r s config params
func (o *UpdateCORSConfigParams) SetBody(body *models.HashicorpCloudVault20200420UpdateCORSConfigRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update c o r s config params
func (o *UpdateCORSConfigParams) WithClusterID(clusterID string) *UpdateCORSConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update c o r s config params
func (o *UpdateCORSConfigParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update c o r s config params
func (o *UpdateCORSConfigParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateCORSConfigParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update c o r s config params
func (o *UpdateCORSConfigParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update c o r s config params
func (o *UpdateCORSConfigParams) WithLocationProjectID(locationProjectID string) *UpdateCORSConfigParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update c o r s config params
func (o *UpdateCORSConfigParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCORSConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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
