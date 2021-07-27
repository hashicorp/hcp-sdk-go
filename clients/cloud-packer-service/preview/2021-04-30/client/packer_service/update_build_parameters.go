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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
)

// NewUpdateBuildParams creates a new UpdateBuildParams object
// with the default values initialized.
func NewUpdateBuildParams() *UpdateBuildParams {
	var ()
	return &UpdateBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBuildParamsWithTimeout creates a new UpdateBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBuildParamsWithTimeout(timeout time.Duration) *UpdateBuildParams {
	var ()
	return &UpdateBuildParams{

		timeout: timeout,
	}
}

// NewUpdateBuildParamsWithContext creates a new UpdateBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBuildParamsWithContext(ctx context.Context) *UpdateBuildParams {
	var ()
	return &UpdateBuildParams{

		Context: ctx,
	}
}

// NewUpdateBuildParamsWithHTTPClient creates a new UpdateBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBuildParamsWithHTTPClient(client *http.Client) *UpdateBuildParams {
	var ()
	return &UpdateBuildParams{
		HTTPClient: client,
	}
}

/*UpdateBuildParams contains all the parameters to send to the API endpoint
for the update build operation typically these are written to a http.Request
*/
type UpdateBuildParams struct {

	/*Body*/
	Body *models.HashicorpCloudPackerUpdateBuildRequest
	/*BuildID
	  build ULID

	*/
	BuildID string
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

// WithTimeout adds the timeout to the update build params
func (o *UpdateBuildParams) WithTimeout(timeout time.Duration) *UpdateBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update build params
func (o *UpdateBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update build params
func (o *UpdateBuildParams) WithContext(ctx context.Context) *UpdateBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update build params
func (o *UpdateBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update build params
func (o *UpdateBuildParams) WithHTTPClient(client *http.Client) *UpdateBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update build params
func (o *UpdateBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update build params
func (o *UpdateBuildParams) WithBody(body *models.HashicorpCloudPackerUpdateBuildRequest) *UpdateBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update build params
func (o *UpdateBuildParams) SetBody(body *models.HashicorpCloudPackerUpdateBuildRequest) {
	o.Body = body
}

// WithBuildID adds the buildID to the update build params
func (o *UpdateBuildParams) WithBuildID(buildID string) *UpdateBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the update build params
func (o *UpdateBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update build params
func (o *UpdateBuildParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update build params
func (o *UpdateBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update build params
func (o *UpdateBuildParams) WithLocationProjectID(locationProjectID string) *UpdateBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update build params
func (o *UpdateBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
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
