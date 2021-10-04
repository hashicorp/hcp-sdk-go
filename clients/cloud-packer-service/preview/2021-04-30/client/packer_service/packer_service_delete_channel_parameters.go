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
)

// NewPackerServiceDeleteChannelParams creates a new PackerServiceDeleteChannelParams object
// with the default values initialized.
func NewPackerServiceDeleteChannelParams() *PackerServiceDeleteChannelParams {
	var ()
	return &PackerServiceDeleteChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceDeleteChannelParamsWithTimeout creates a new PackerServiceDeleteChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceDeleteChannelParamsWithTimeout(timeout time.Duration) *PackerServiceDeleteChannelParams {
	var ()
	return &PackerServiceDeleteChannelParams{

		timeout: timeout,
	}
}

// NewPackerServiceDeleteChannelParamsWithContext creates a new PackerServiceDeleteChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceDeleteChannelParamsWithContext(ctx context.Context) *PackerServiceDeleteChannelParams {
	var ()
	return &PackerServiceDeleteChannelParams{

		Context: ctx,
	}
}

// NewPackerServiceDeleteChannelParamsWithHTTPClient creates a new PackerServiceDeleteChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceDeleteChannelParamsWithHTTPClient(client *http.Client) *PackerServiceDeleteChannelParams {
	var ()
	return &PackerServiceDeleteChannelParams{
		HTTPClient: client,
	}
}

/*PackerServiceDeleteChannelParams contains all the parameters to send to the API endpoint
for the packer service delete channel operation typically these are written to a http.Request
*/
type PackerServiceDeleteChannelParams struct {

	/*BucketSlug
	  Human-readable name for the bucket that the channel is associated with.

	*/
	BucketSlug string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*LocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	LocationRegionProvider *string
	/*LocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	LocationRegionRegion *string
	/*RevocationMessage
	  Optional field to provide the reason for why this channel is being revoked.
	Only useful for a channel that is assigned to an iteration.

	*/
	RevocationMessage *string
	/*Slug
	  Human-readable name for the channel.

	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithTimeout(timeout time.Duration) *PackerServiceDeleteChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithContext(ctx context.Context) *PackerServiceDeleteChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithHTTPClient(client *http.Client) *PackerServiceDeleteChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithBucketSlug(bucketSlug string) *PackerServiceDeleteChannelParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceDeleteChannelParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithLocationProjectID(locationProjectID string) *PackerServiceDeleteChannelParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceDeleteChannelParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceDeleteChannelParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithRevocationMessage adds the revocationMessage to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithRevocationMessage(revocationMessage *string) *PackerServiceDeleteChannelParams {
	o.SetRevocationMessage(revocationMessage)
	return o
}

// SetRevocationMessage adds the revocationMessage to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetRevocationMessage(revocationMessage *string) {
	o.RevocationMessage = revocationMessage
}

// WithSlug adds the slug to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) WithSlug(slug string) *PackerServiceDeleteChannelParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the packer service delete channel params
func (o *PackerServiceDeleteChannelParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceDeleteChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string
		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {
			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}

	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string
		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {
			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}

	}

	if o.RevocationMessage != nil {

		// query param revocation_message
		var qrRevocationMessage string
		if o.RevocationMessage != nil {
			qrRevocationMessage = *o.RevocationMessage
		}
		qRevocationMessage := qrRevocationMessage
		if qRevocationMessage != "" {
			if err := r.SetQueryParam("revocation_message", qRevocationMessage); err != nil {
				return err
			}
		}

	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
