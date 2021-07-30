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

// NewDeleteChannelParams creates a new DeleteChannelParams object
// with the default values initialized.
func NewDeleteChannelParams() *DeleteChannelParams {
	var ()
	return &DeleteChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteChannelParamsWithTimeout creates a new DeleteChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteChannelParamsWithTimeout(timeout time.Duration) *DeleteChannelParams {
	var ()
	return &DeleteChannelParams{

		timeout: timeout,
	}
}

// NewDeleteChannelParamsWithContext creates a new DeleteChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteChannelParamsWithContext(ctx context.Context) *DeleteChannelParams {
	var ()
	return &DeleteChannelParams{

		Context: ctx,
	}
}

// NewDeleteChannelParamsWithHTTPClient creates a new DeleteChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteChannelParamsWithHTTPClient(client *http.Client) *DeleteChannelParams {
	var ()
	return &DeleteChannelParams{
		HTTPClient: client,
	}
}

/*DeleteChannelParams contains all the parameters to send to the API endpoint
for the delete channel operation typically these are written to a http.Request
*/
type DeleteChannelParams struct {

	/*BucketSlug
	  The bucket slug

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
	  The channel slug. e.g. production-stable

	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete channel params
func (o *DeleteChannelParams) WithTimeout(timeout time.Duration) *DeleteChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete channel params
func (o *DeleteChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete channel params
func (o *DeleteChannelParams) WithContext(ctx context.Context) *DeleteChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete channel params
func (o *DeleteChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete channel params
func (o *DeleteChannelParams) WithHTTPClient(client *http.Client) *DeleteChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete channel params
func (o *DeleteChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the delete channel params
func (o *DeleteChannelParams) WithBucketSlug(bucketSlug string) *DeleteChannelParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the delete channel params
func (o *DeleteChannelParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the delete channel params
func (o *DeleteChannelParams) WithLocationOrganizationID(locationOrganizationID string) *DeleteChannelParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the delete channel params
func (o *DeleteChannelParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the delete channel params
func (o *DeleteChannelParams) WithLocationProjectID(locationProjectID string) *DeleteChannelParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the delete channel params
func (o *DeleteChannelParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the delete channel params
func (o *DeleteChannelParams) WithLocationRegionProvider(locationRegionProvider *string) *DeleteChannelParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the delete channel params
func (o *DeleteChannelParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the delete channel params
func (o *DeleteChannelParams) WithLocationRegionRegion(locationRegionRegion *string) *DeleteChannelParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the delete channel params
func (o *DeleteChannelParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithRevocationMessage adds the revocationMessage to the delete channel params
func (o *DeleteChannelParams) WithRevocationMessage(revocationMessage *string) *DeleteChannelParams {
	o.SetRevocationMessage(revocationMessage)
	return o
}

// SetRevocationMessage adds the revocationMessage to the delete channel params
func (o *DeleteChannelParams) SetRevocationMessage(revocationMessage *string) {
	o.RevocationMessage = revocationMessage
}

// WithSlug adds the slug to the delete channel params
func (o *DeleteChannelParams) WithSlug(slug string) *DeleteChannelParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete channel params
func (o *DeleteChannelParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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