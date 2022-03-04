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

// NewPackerServiceCreateBuildParams creates a new PackerServiceCreateBuildParams object
// with the default values initialized.
func NewPackerServiceCreateBuildParams() *PackerServiceCreateBuildParams {
	var ()
	return &PackerServiceCreateBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceCreateBuildParamsWithTimeout creates a new PackerServiceCreateBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceCreateBuildParamsWithTimeout(timeout time.Duration) *PackerServiceCreateBuildParams {
	var ()
	return &PackerServiceCreateBuildParams{

		timeout: timeout,
	}
}

// NewPackerServiceCreateBuildParamsWithContext creates a new PackerServiceCreateBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceCreateBuildParamsWithContext(ctx context.Context) *PackerServiceCreateBuildParams {
	var ()
	return &PackerServiceCreateBuildParams{

		Context: ctx,
	}
}

// NewPackerServiceCreateBuildParamsWithHTTPClient creates a new PackerServiceCreateBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceCreateBuildParamsWithHTTPClient(client *http.Client) *PackerServiceCreateBuildParams {
	var ()
	return &PackerServiceCreateBuildParams{
		HTTPClient: client,
	}
}

/*PackerServiceCreateBuildParams contains all the parameters to send to the API endpoint
for the packer service create build operation typically these are written to a http.Request
*/
type PackerServiceCreateBuildParams struct {

	/*Body*/
	Body *models.HashicorpCloudPackerCreateBuildRequest
	/*BucketSlug
	  Human-readable name for the bucket.

	*/
	BucketSlug string
	/*IterationID
	  ULID of the iteration that this build should be associated with.

	*/
	IterationID string
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

// WithTimeout adds the timeout to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithTimeout(timeout time.Duration) *PackerServiceCreateBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithContext(ctx context.Context) *PackerServiceCreateBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithHTTPClient(client *http.Client) *PackerServiceCreateBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithBody(body *models.HashicorpCloudPackerCreateBuildRequest) *PackerServiceCreateBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetBody(body *models.HashicorpCloudPackerCreateBuildRequest) {
	o.Body = body
}

// WithBucketSlug adds the bucketSlug to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithBucketSlug(bucketSlug string) *PackerServiceCreateBuildParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithIterationID adds the iterationID to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithIterationID(iterationID string) *PackerServiceCreateBuildParams {
	o.SetIterationID(iterationID)
	return o
}

// SetIterationID adds the iterationId to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetIterationID(iterationID string) {
	o.IterationID = iterationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceCreateBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service create build params
func (o *PackerServiceCreateBuildParams) WithLocationProjectID(locationProjectID string) *PackerServiceCreateBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service create build params
func (o *PackerServiceCreateBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceCreateBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param iteration_id
	if err := r.SetPathParam("iteration_id", o.IterationID); err != nil {
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
