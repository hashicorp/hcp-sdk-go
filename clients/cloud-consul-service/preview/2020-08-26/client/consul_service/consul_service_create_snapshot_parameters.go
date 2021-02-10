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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/models"
)

// NewConsulServiceCreateSnapshotParams creates a new ConsulServiceCreateSnapshotParams object
// with the default values initialized.
func NewConsulServiceCreateSnapshotParams() *ConsulServiceCreateSnapshotParams {
	var ()
	return &ConsulServiceCreateSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsulServiceCreateSnapshotParamsWithTimeout creates a new ConsulServiceCreateSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsulServiceCreateSnapshotParamsWithTimeout(timeout time.Duration) *ConsulServiceCreateSnapshotParams {
	var ()
	return &ConsulServiceCreateSnapshotParams{

		timeout: timeout,
	}
}

// NewConsulServiceCreateSnapshotParamsWithContext creates a new ConsulServiceCreateSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsulServiceCreateSnapshotParamsWithContext(ctx context.Context) *ConsulServiceCreateSnapshotParams {
	var ()
	return &ConsulServiceCreateSnapshotParams{

		Context: ctx,
	}
}

// NewConsulServiceCreateSnapshotParamsWithHTTPClient creates a new ConsulServiceCreateSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsulServiceCreateSnapshotParamsWithHTTPClient(client *http.Client) *ConsulServiceCreateSnapshotParams {
	var ()
	return &ConsulServiceCreateSnapshotParams{
		HTTPClient: client,
	}
}

/*ConsulServiceCreateSnapshotParams contains all the parameters to send to the API endpoint
for the consul service create snapshot operation typically these are written to a http.Request
*/
type ConsulServiceCreateSnapshotParams struct {

	/*Body*/
	Body *models.HashicorpCloudConsul20200826CreateSnapshotRequest
	/*ResourceLocationOrganizationID
	  organization_id is the id of the organization.

	*/
	ResourceLocationOrganizationID string
	/*ResourceLocationProjectID
	  project_id is the projects id.

	*/
	ResourceLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithTimeout(timeout time.Duration) *ConsulServiceCreateSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithContext(ctx context.Context) *ConsulServiceCreateSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithHTTPClient(client *http.Client) *ConsulServiceCreateSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithBody(body *models.HashicorpCloudConsul20200826CreateSnapshotRequest) *ConsulServiceCreateSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetBody(body *models.HashicorpCloudConsul20200826CreateSnapshotRequest) {
	o.Body = body
}

// WithResourceLocationOrganizationID adds the resourceLocationOrganizationID to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithResourceLocationOrganizationID(resourceLocationOrganizationID string) *ConsulServiceCreateSnapshotParams {
	o.SetResourceLocationOrganizationID(resourceLocationOrganizationID)
	return o
}

// SetResourceLocationOrganizationID adds the resourceLocationOrganizationId to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetResourceLocationOrganizationID(resourceLocationOrganizationID string) {
	o.ResourceLocationOrganizationID = resourceLocationOrganizationID
}

// WithResourceLocationProjectID adds the resourceLocationProjectID to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) WithResourceLocationProjectID(resourceLocationProjectID string) *ConsulServiceCreateSnapshotParams {
	o.SetResourceLocationProjectID(resourceLocationProjectID)
	return o
}

// SetResourceLocationProjectID adds the resourceLocationProjectId to the consul service create snapshot params
func (o *ConsulServiceCreateSnapshotParams) SetResourceLocationProjectID(resourceLocationProjectID string) {
	o.ResourceLocationProjectID = resourceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsulServiceCreateSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource.location.organization_id
	if err := r.SetPathParam("resource.location.organization_id", o.ResourceLocationOrganizationID); err != nil {
		return err
	}

	// path param resource.location.project_id
	if err := r.SetPathParam("resource.location.project_id", o.ResourceLocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}