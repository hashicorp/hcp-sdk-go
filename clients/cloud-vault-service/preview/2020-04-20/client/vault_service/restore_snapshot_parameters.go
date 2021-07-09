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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-04-20/models"
)

// NewRestoreSnapshotParams creates a new RestoreSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreSnapshotParams() *RestoreSnapshotParams {
	return &RestoreSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreSnapshotParamsWithTimeout creates a new RestoreSnapshotParams object
// with the ability to set a timeout on a request.
func NewRestoreSnapshotParamsWithTimeout(timeout time.Duration) *RestoreSnapshotParams {
	return &RestoreSnapshotParams{
		timeout: timeout,
	}
}

// NewRestoreSnapshotParamsWithContext creates a new RestoreSnapshotParams object
// with the ability to set a context for a request.
func NewRestoreSnapshotParamsWithContext(ctx context.Context) *RestoreSnapshotParams {
	return &RestoreSnapshotParams{
		Context: ctx,
	}
}

// NewRestoreSnapshotParamsWithHTTPClient creates a new RestoreSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreSnapshotParamsWithHTTPClient(client *http.Client) *RestoreSnapshotParams {
	return &RestoreSnapshotParams{
		HTTPClient: client,
	}
}

/* RestoreSnapshotParams contains all the parameters to send to the API endpoint
   for the restore snapshot operation.

   Typically these are written to a http.Request.
*/
type RestoreSnapshotParams struct {

	// Body.
	Body *models.HashicorpCloudVault20200420RestoreSnapshotRequest

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

// WithDefaults hydrates default values in the restore snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreSnapshotParams) WithDefaults() *RestoreSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore snapshot params
func (o *RestoreSnapshotParams) WithTimeout(timeout time.Duration) *RestoreSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore snapshot params
func (o *RestoreSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore snapshot params
func (o *RestoreSnapshotParams) WithContext(ctx context.Context) *RestoreSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore snapshot params
func (o *RestoreSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore snapshot params
func (o *RestoreSnapshotParams) WithHTTPClient(client *http.Client) *RestoreSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore snapshot params
func (o *RestoreSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restore snapshot params
func (o *RestoreSnapshotParams) WithBody(body *models.HashicorpCloudVault20200420RestoreSnapshotRequest) *RestoreSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restore snapshot params
func (o *RestoreSnapshotParams) SetBody(body *models.HashicorpCloudVault20200420RestoreSnapshotRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the restore snapshot params
func (o *RestoreSnapshotParams) WithClusterID(clusterID string) *RestoreSnapshotParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the restore snapshot params
func (o *RestoreSnapshotParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the restore snapshot params
func (o *RestoreSnapshotParams) WithLocationOrganizationID(locationOrganizationID string) *RestoreSnapshotParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the restore snapshot params
func (o *RestoreSnapshotParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the restore snapshot params
func (o *RestoreSnapshotParams) WithLocationProjectID(locationProjectID string) *RestoreSnapshotParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the restore snapshot params
func (o *RestoreSnapshotParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
