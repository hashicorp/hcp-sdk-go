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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2020-08-26/models"
)

// NewCreateSnapshotParams creates a new CreateSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSnapshotParams() *CreateSnapshotParams {
	return &CreateSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotParamsWithTimeout creates a new CreateSnapshotParams object
// with the ability to set a timeout on a request.
func NewCreateSnapshotParamsWithTimeout(timeout time.Duration) *CreateSnapshotParams {
	return &CreateSnapshotParams{
		timeout: timeout,
	}
}

// NewCreateSnapshotParamsWithContext creates a new CreateSnapshotParams object
// with the ability to set a context for a request.
func NewCreateSnapshotParamsWithContext(ctx context.Context) *CreateSnapshotParams {
	return &CreateSnapshotParams{
		Context: ctx,
	}
}

// NewCreateSnapshotParamsWithHTTPClient creates a new CreateSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSnapshotParamsWithHTTPClient(client *http.Client) *CreateSnapshotParams {
	return &CreateSnapshotParams{
		HTTPClient: client,
	}
}

/* CreateSnapshotParams contains all the parameters to send to the API endpoint
   for the create snapshot operation.

   Typically these are written to a http.Request.
*/
type CreateSnapshotParams struct {

	// Body.
	Body *models.HashicorpCloudConsul20200826CreateSnapshotRequest

	/* ResourceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	ResourceLocationOrganizationID string

	/* ResourceLocationProjectID.

	   project_id is the projects id.
	*/
	ResourceLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSnapshotParams) WithDefaults() *CreateSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create snapshot params
func (o *CreateSnapshotParams) WithTimeout(timeout time.Duration) *CreateSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot params
func (o *CreateSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot params
func (o *CreateSnapshotParams) WithContext(ctx context.Context) *CreateSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot params
func (o *CreateSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot params
func (o *CreateSnapshotParams) WithHTTPClient(client *http.Client) *CreateSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot params
func (o *CreateSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create snapshot params
func (o *CreateSnapshotParams) WithBody(body *models.HashicorpCloudConsul20200826CreateSnapshotRequest) *CreateSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create snapshot params
func (o *CreateSnapshotParams) SetBody(body *models.HashicorpCloudConsul20200826CreateSnapshotRequest) {
	o.Body = body
}

// WithResourceLocationOrganizationID adds the resourceLocationOrganizationID to the create snapshot params
func (o *CreateSnapshotParams) WithResourceLocationOrganizationID(resourceLocationOrganizationID string) *CreateSnapshotParams {
	o.SetResourceLocationOrganizationID(resourceLocationOrganizationID)
	return o
}

// SetResourceLocationOrganizationID adds the resourceLocationOrganizationId to the create snapshot params
func (o *CreateSnapshotParams) SetResourceLocationOrganizationID(resourceLocationOrganizationID string) {
	o.ResourceLocationOrganizationID = resourceLocationOrganizationID
}

// WithResourceLocationProjectID adds the resourceLocationProjectID to the create snapshot params
func (o *CreateSnapshotParams) WithResourceLocationProjectID(resourceLocationProjectID string) *CreateSnapshotParams {
	o.SetResourceLocationProjectID(resourceLocationProjectID)
	return o
}

// SetResourceLocationProjectID adds the resourceLocationProjectId to the create snapshot params
func (o *CreateSnapshotParams) SetResourceLocationProjectID(resourceLocationProjectID string) {
	o.ResourceLocationProjectID = resourceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
