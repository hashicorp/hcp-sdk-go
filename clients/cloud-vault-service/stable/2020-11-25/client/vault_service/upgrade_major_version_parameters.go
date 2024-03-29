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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// NewUpgradeMajorVersionParams creates a new UpgradeMajorVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeMajorVersionParams() *UpgradeMajorVersionParams {
	return &UpgradeMajorVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeMajorVersionParamsWithTimeout creates a new UpgradeMajorVersionParams object
// with the ability to set a timeout on a request.
func NewUpgradeMajorVersionParamsWithTimeout(timeout time.Duration) *UpgradeMajorVersionParams {
	return &UpgradeMajorVersionParams{
		timeout: timeout,
	}
}

// NewUpgradeMajorVersionParamsWithContext creates a new UpgradeMajorVersionParams object
// with the ability to set a context for a request.
func NewUpgradeMajorVersionParamsWithContext(ctx context.Context) *UpgradeMajorVersionParams {
	return &UpgradeMajorVersionParams{
		Context: ctx,
	}
}

// NewUpgradeMajorVersionParamsWithHTTPClient creates a new UpgradeMajorVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeMajorVersionParamsWithHTTPClient(client *http.Client) *UpgradeMajorVersionParams {
	return &UpgradeMajorVersionParams{
		HTTPClient: client,
	}
}

/*
UpgradeMajorVersionParams contains all the parameters to send to the API endpoint

	for the upgrade major version operation.

	Typically these are written to a http.Request.
*/
type UpgradeMajorVersionParams struct {

	// Body.
	Body *models.HashicorpCloudVault20201125UpgradeMajorVersionRequest

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

// WithDefaults hydrates default values in the upgrade major version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeMajorVersionParams) WithDefaults() *UpgradeMajorVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade major version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeMajorVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithTimeout(timeout time.Duration) *UpgradeMajorVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithContext(ctx context.Context) *UpgradeMajorVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithHTTPClient(client *http.Client) *UpgradeMajorVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithBody(body *models.HashicorpCloudVault20201125UpgradeMajorVersionRequest) *UpgradeMajorVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetBody(body *models.HashicorpCloudVault20201125UpgradeMajorVersionRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithClusterID(clusterID string) *UpgradeMajorVersionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithLocationOrganizationID(locationOrganizationID string) *UpgradeMajorVersionParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the upgrade major version params
func (o *UpgradeMajorVersionParams) WithLocationProjectID(locationProjectID string) *UpgradeMajorVersionParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the upgrade major version params
func (o *UpgradeMajorVersionParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeMajorVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
