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
)

// NewConsulServiceDeleteSnapshotParams creates a new ConsulServiceDeleteSnapshotParams object
// with the default values initialized.
func NewConsulServiceDeleteSnapshotParams() *ConsulServiceDeleteSnapshotParams {
	var ()
	return &ConsulServiceDeleteSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsulServiceDeleteSnapshotParamsWithTimeout creates a new ConsulServiceDeleteSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsulServiceDeleteSnapshotParamsWithTimeout(timeout time.Duration) *ConsulServiceDeleteSnapshotParams {
	var ()
	return &ConsulServiceDeleteSnapshotParams{

		timeout: timeout,
	}
}

// NewConsulServiceDeleteSnapshotParamsWithContext creates a new ConsulServiceDeleteSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsulServiceDeleteSnapshotParamsWithContext(ctx context.Context) *ConsulServiceDeleteSnapshotParams {
	var ()
	return &ConsulServiceDeleteSnapshotParams{

		Context: ctx,
	}
}

// NewConsulServiceDeleteSnapshotParamsWithHTTPClient creates a new ConsulServiceDeleteSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsulServiceDeleteSnapshotParamsWithHTTPClient(client *http.Client) *ConsulServiceDeleteSnapshotParams {
	var ()
	return &ConsulServiceDeleteSnapshotParams{
		HTTPClient: client,
	}
}

/*ConsulServiceDeleteSnapshotParams contains all the parameters to send to the API endpoint
for the consul service delete snapshot operation typically these are written to a http.Request
*/
type ConsulServiceDeleteSnapshotParams struct {

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
	/*SnapshotID
	  snapshot_id represents the snapshot to delete.

	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithTimeout(timeout time.Duration) *ConsulServiceDeleteSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithContext(ctx context.Context) *ConsulServiceDeleteSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithHTTPClient(client *http.Client) *ConsulServiceDeleteSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithLocationOrganizationID(locationOrganizationID string) *ConsulServiceDeleteSnapshotParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithLocationProjectID(locationProjectID string) *ConsulServiceDeleteSnapshotParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithLocationRegionProvider(locationRegionProvider *string) *ConsulServiceDeleteSnapshotParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithLocationRegionRegion(locationRegionRegion *string) *ConsulServiceDeleteSnapshotParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithSnapshotID adds the snapshotID to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) WithSnapshotID(snapshotID string) *ConsulServiceDeleteSnapshotParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the consul service delete snapshot params
func (o *ConsulServiceDeleteSnapshotParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsulServiceDeleteSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}