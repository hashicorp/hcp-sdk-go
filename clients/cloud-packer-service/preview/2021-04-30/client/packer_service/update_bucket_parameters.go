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

// NewUpdateBucketParams creates a new UpdateBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBucketParams() *UpdateBucketParams {
	return &UpdateBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBucketParamsWithTimeout creates a new UpdateBucketParams object
// with the ability to set a timeout on a request.
func NewUpdateBucketParamsWithTimeout(timeout time.Duration) *UpdateBucketParams {
	return &UpdateBucketParams{
		timeout: timeout,
	}
}

// NewUpdateBucketParamsWithContext creates a new UpdateBucketParams object
// with the ability to set a context for a request.
func NewUpdateBucketParamsWithContext(ctx context.Context) *UpdateBucketParams {
	return &UpdateBucketParams{
		Context: ctx,
	}
}

// NewUpdateBucketParamsWithHTTPClient creates a new UpdateBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBucketParamsWithHTTPClient(client *http.Client) *UpdateBucketParams {
	return &UpdateBucketParams{
		HTTPClient: client,
	}
}

/* UpdateBucketParams contains all the parameters to send to the API endpoint
   for the update bucket operation.

   Typically these are written to a http.Request.
*/
type UpdateBucketParams struct {

	// Body.
	Body *models.HashicorpCloudPackerUpdateBucketRequest

	/* BucketSlug.

	   base-ubuntu-18-secure
	*/
	BucketSlug string

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

// WithDefaults hydrates default values in the update bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBucketParams) WithDefaults() *UpdateBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update bucket params
func (o *UpdateBucketParams) WithTimeout(timeout time.Duration) *UpdateBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update bucket params
func (o *UpdateBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update bucket params
func (o *UpdateBucketParams) WithContext(ctx context.Context) *UpdateBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update bucket params
func (o *UpdateBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update bucket params
func (o *UpdateBucketParams) WithHTTPClient(client *http.Client) *UpdateBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update bucket params
func (o *UpdateBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update bucket params
func (o *UpdateBucketParams) WithBody(body *models.HashicorpCloudPackerUpdateBucketRequest) *UpdateBucketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update bucket params
func (o *UpdateBucketParams) SetBody(body *models.HashicorpCloudPackerUpdateBucketRequest) {
	o.Body = body
}

// WithBucketSlug adds the bucketSlug to the update bucket params
func (o *UpdateBucketParams) WithBucketSlug(bucketSlug string) *UpdateBucketParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the update bucket params
func (o *UpdateBucketParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the update bucket params
func (o *UpdateBucketParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateBucketParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update bucket params
func (o *UpdateBucketParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update bucket params
func (o *UpdateBucketParams) WithLocationProjectID(locationProjectID string) *UpdateBucketParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update bucket params
func (o *UpdateBucketParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
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
