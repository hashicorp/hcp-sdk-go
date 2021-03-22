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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-11-25/models"
)

// NewUpdatePublicIpsParams creates a new UpdatePublicIpsParams object
// with the default values initialized.
func NewUpdatePublicIpsParams() *UpdatePublicIpsParams {
	var ()
	return &UpdatePublicIpsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePublicIpsParamsWithTimeout creates a new UpdatePublicIpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePublicIpsParamsWithTimeout(timeout time.Duration) *UpdatePublicIpsParams {
	var ()
	return &UpdatePublicIpsParams{

		timeout: timeout,
	}
}

// NewUpdatePublicIpsParamsWithContext creates a new UpdatePublicIpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePublicIpsParamsWithContext(ctx context.Context) *UpdatePublicIpsParams {
	var ()
	return &UpdatePublicIpsParams{

		Context: ctx,
	}
}

// NewUpdatePublicIpsParamsWithHTTPClient creates a new UpdatePublicIpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePublicIpsParamsWithHTTPClient(client *http.Client) *UpdatePublicIpsParams {
	var ()
	return &UpdatePublicIpsParams{
		HTTPClient: client,
	}
}

/*UpdatePublicIpsParams contains all the parameters to send to the API endpoint
for the update public ips operation typically these are written to a http.Request
*/
type UpdatePublicIpsParams struct {

	/*Body*/
	Body *models.HashicorpCloudVault20201125UpdatePublicIpsRequest
	/*ClusterID*/
	ClusterID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update public ips params
func (o *UpdatePublicIpsParams) WithTimeout(timeout time.Duration) *UpdatePublicIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update public ips params
func (o *UpdatePublicIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update public ips params
func (o *UpdatePublicIpsParams) WithContext(ctx context.Context) *UpdatePublicIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update public ips params
func (o *UpdatePublicIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update public ips params
func (o *UpdatePublicIpsParams) WithHTTPClient(client *http.Client) *UpdatePublicIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update public ips params
func (o *UpdatePublicIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update public ips params
func (o *UpdatePublicIpsParams) WithBody(body *models.HashicorpCloudVault20201125UpdatePublicIpsRequest) *UpdatePublicIpsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update public ips params
func (o *UpdatePublicIpsParams) SetBody(body *models.HashicorpCloudVault20201125UpdatePublicIpsRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update public ips params
func (o *UpdatePublicIpsParams) WithClusterID(clusterID string) *UpdatePublicIpsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update public ips params
func (o *UpdatePublicIpsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update public ips params
func (o *UpdatePublicIpsParams) WithLocationOrganizationID(locationOrganizationID string) *UpdatePublicIpsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update public ips params
func (o *UpdatePublicIpsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update public ips params
func (o *UpdatePublicIpsParams) WithLocationProjectID(locationProjectID string) *UpdatePublicIpsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update public ips params
func (o *UpdatePublicIpsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePublicIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
