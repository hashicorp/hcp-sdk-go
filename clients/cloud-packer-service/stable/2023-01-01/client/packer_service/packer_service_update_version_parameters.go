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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
)

// NewPackerServiceUpdateVersionParams creates a new PackerServiceUpdateVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceUpdateVersionParams() *PackerServiceUpdateVersionParams {
	return &PackerServiceUpdateVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceUpdateVersionParamsWithTimeout creates a new PackerServiceUpdateVersionParams object
// with the ability to set a timeout on a request.
func NewPackerServiceUpdateVersionParamsWithTimeout(timeout time.Duration) *PackerServiceUpdateVersionParams {
	return &PackerServiceUpdateVersionParams{
		timeout: timeout,
	}
}

// NewPackerServiceUpdateVersionParamsWithContext creates a new PackerServiceUpdateVersionParams object
// with the ability to set a context for a request.
func NewPackerServiceUpdateVersionParamsWithContext(ctx context.Context) *PackerServiceUpdateVersionParams {
	return &PackerServiceUpdateVersionParams{
		Context: ctx,
	}
}

// NewPackerServiceUpdateVersionParamsWithHTTPClient creates a new PackerServiceUpdateVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceUpdateVersionParamsWithHTTPClient(client *http.Client) *PackerServiceUpdateVersionParams {
	return &PackerServiceUpdateVersionParams{
		HTTPClient: client,
	}
}

/*
PackerServiceUpdateVersionParams contains all the parameters to send to the API endpoint

	for the packer service update version operation.

	Typically these are written to a http.Request.
*/
type PackerServiceUpdateVersionParams struct {

	// Body.
	Body *models.HashicorpCloudPacker20230101UpdateVersionBody

	// BucketName.
	BucketName string

	// Fingerprint.
	Fingerprint string

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

// WithDefaults hydrates default values in the packer service update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateVersionParams) WithDefaults() *PackerServiceUpdateVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceUpdateVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithTimeout(timeout time.Duration) *PackerServiceUpdateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithContext(ctx context.Context) *PackerServiceUpdateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithHTTPClient(client *http.Client) *PackerServiceUpdateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithBody(body *models.HashicorpCloudPacker20230101UpdateVersionBody) *PackerServiceUpdateVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetBody(body *models.HashicorpCloudPacker20230101UpdateVersionBody) {
	o.Body = body
}

// WithBucketName adds the bucketName to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithBucketName(bucketName string) *PackerServiceUpdateVersionParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithFingerprint adds the fingerprint to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithFingerprint(fingerprint string) *PackerServiceUpdateVersionParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceUpdateVersionParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service update version params
func (o *PackerServiceUpdateVersionParams) WithLocationProjectID(locationProjectID string) *PackerServiceUpdateVersionParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service update version params
func (o *PackerServiceUpdateVersionParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceUpdateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_name
	if err := r.SetPathParam("bucket_name", o.BucketName); err != nil {
		return err
	}

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
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
