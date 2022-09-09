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

// NewUpdateMajorVersionUpgradeConfigParams creates a new UpdateMajorVersionUpgradeConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMajorVersionUpgradeConfigParams() *UpdateMajorVersionUpgradeConfigParams {
	return &UpdateMajorVersionUpgradeConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMajorVersionUpgradeConfigParamsWithTimeout creates a new UpdateMajorVersionUpgradeConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateMajorVersionUpgradeConfigParamsWithTimeout(timeout time.Duration) *UpdateMajorVersionUpgradeConfigParams {
	return &UpdateMajorVersionUpgradeConfigParams{
		timeout: timeout,
	}
}

// NewUpdateMajorVersionUpgradeConfigParamsWithContext creates a new UpdateMajorVersionUpgradeConfigParams object
// with the ability to set a context for a request.
func NewUpdateMajorVersionUpgradeConfigParamsWithContext(ctx context.Context) *UpdateMajorVersionUpgradeConfigParams {
	return &UpdateMajorVersionUpgradeConfigParams{
		Context: ctx,
	}
}

// NewUpdateMajorVersionUpgradeConfigParamsWithHTTPClient creates a new UpdateMajorVersionUpgradeConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMajorVersionUpgradeConfigParamsWithHTTPClient(client *http.Client) *UpdateMajorVersionUpgradeConfigParams {
	return &UpdateMajorVersionUpgradeConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateMajorVersionUpgradeConfigParams contains all the parameters to send to the API endpoint

	for the update major version upgrade config operation.

	Typically these are written to a http.Request.
*/
type UpdateMajorVersionUpgradeConfigParams struct {

	// Body.
	Body *models.HashicorpCloudVault20201125UpdateMajorVersionUpgradeConfigRequest

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

// WithDefaults hydrates default values in the update major version upgrade config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMajorVersionUpgradeConfigParams) WithDefaults() *UpdateMajorVersionUpgradeConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update major version upgrade config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMajorVersionUpgradeConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithTimeout(timeout time.Duration) *UpdateMajorVersionUpgradeConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithContext(ctx context.Context) *UpdateMajorVersionUpgradeConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithHTTPClient(client *http.Client) *UpdateMajorVersionUpgradeConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithBody(body *models.HashicorpCloudVault20201125UpdateMajorVersionUpgradeConfigRequest) *UpdateMajorVersionUpgradeConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetBody(body *models.HashicorpCloudVault20201125UpdateMajorVersionUpgradeConfigRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithClusterID(clusterID string) *UpdateMajorVersionUpgradeConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateMajorVersionUpgradeConfigParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) WithLocationProjectID(locationProjectID string) *UpdateMajorVersionUpgradeConfigParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update major version upgrade config params
func (o *UpdateMajorVersionUpgradeConfigParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMajorVersionUpgradeConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
