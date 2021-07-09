// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2021-02-04/models"
)

// NewUpdateParams creates a new UpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateParams() *UpdateParams {
	return &UpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the ability to set a timeout on a request.
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	return &UpdateParams{
		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the ability to set a context for a request.
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	return &UpdateParams{
		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	return &UpdateParams{
		HTTPClient: client,
	}
}

/* UpdateParams contains all the parameters to send to the API endpoint
   for the update operation.

   Typically these are written to a http.Request.
*/
type UpdateParams struct {

	/* Body.

	   cluster is the complete cluster resource to be updated.
	*/
	Body *models.HashicorpCloudConsul20210204Cluster

	/* ClusterID.

	   id is ID of the Consul cluster.
	*/
	ClusterID string

	/* ClusterLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	ClusterLocationOrganizationID string

	/* ClusterLocationProjectID.

	   project_id is the projects id.
	*/
	ClusterLocationProjectID string

	/* UpdateMaskPaths.

	   The set of field mask paths.
	*/
	UpdateMaskPaths []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParams) WithDefaults() *UpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update params
func (o *UpdateParams) WithBody(body *models.HashicorpCloudConsul20210204Cluster) *UpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update params
func (o *UpdateParams) SetBody(body *models.HashicorpCloudConsul20210204Cluster) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update params
func (o *UpdateParams) WithClusterID(clusterID string) *UpdateParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update params
func (o *UpdateParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithClusterLocationOrganizationID adds the clusterLocationOrganizationID to the update params
func (o *UpdateParams) WithClusterLocationOrganizationID(clusterLocationOrganizationID string) *UpdateParams {
	o.SetClusterLocationOrganizationID(clusterLocationOrganizationID)
	return o
}

// SetClusterLocationOrganizationID adds the clusterLocationOrganizationId to the update params
func (o *UpdateParams) SetClusterLocationOrganizationID(clusterLocationOrganizationID string) {
	o.ClusterLocationOrganizationID = clusterLocationOrganizationID
}

// WithClusterLocationProjectID adds the clusterLocationProjectID to the update params
func (o *UpdateParams) WithClusterLocationProjectID(clusterLocationProjectID string) *UpdateParams {
	o.SetClusterLocationProjectID(clusterLocationProjectID)
	return o
}

// SetClusterLocationProjectID adds the clusterLocationProjectId to the update params
func (o *UpdateParams) SetClusterLocationProjectID(clusterLocationProjectID string) {
	o.ClusterLocationProjectID = clusterLocationProjectID
}

// WithUpdateMaskPaths adds the updateMaskPaths to the update params
func (o *UpdateParams) WithUpdateMaskPaths(updateMaskPaths []string) *UpdateParams {
	o.SetUpdateMaskPaths(updateMaskPaths)
	return o
}

// SetUpdateMaskPaths adds the updateMaskPaths to the update params
func (o *UpdateParams) SetUpdateMaskPaths(updateMaskPaths []string) {
	o.UpdateMaskPaths = updateMaskPaths
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster.id
	if err := r.SetPathParam("cluster.id", o.ClusterID); err != nil {
		return err
	}

	// path param cluster.location.organization_id
	if err := r.SetPathParam("cluster.location.organization_id", o.ClusterLocationOrganizationID); err != nil {
		return err
	}

	// path param cluster.location.project_id
	if err := r.SetPathParam("cluster.location.project_id", o.ClusterLocationProjectID); err != nil {
		return err
	}

	if o.UpdateMaskPaths != nil {

		// binding items for update_mask.paths
		joinedUpdateMaskPaths := o.bindParamUpdateMaskPaths(reg)

		// query array param update_mask.paths
		if err := r.SetQueryParam("update_mask.paths", joinedUpdateMaskPaths...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUpdate binds the parameter update_mask.paths
func (o *UpdateParams) bindParamUpdateMaskPaths(formats strfmt.Registry) []string {
	updateMaskPathsIR := o.UpdateMaskPaths

	var updateMaskPathsIC []string
	for _, updateMaskPathsIIR := range updateMaskPathsIR { // explode []string

		updateMaskPathsIIV := updateMaskPathsIIR // string as string
		updateMaskPathsIC = append(updateMaskPathsIC, updateMaskPathsIIV)
	}

	// items.CollectionFormat: "multi"
	updateMaskPathsIS := swag.JoinByFormat(updateMaskPathsIC, "multi")

	return updateMaskPathsIS
}
