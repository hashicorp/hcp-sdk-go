// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

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

// NewSessionsParams creates a new SessionsParams object
// with the default values initialized.
func NewSessionsParams() *SessionsParams {
	var ()
	return &SessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSessionsParamsWithTimeout creates a new SessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSessionsParamsWithTimeout(timeout time.Duration) *SessionsParams {
	var ()
	return &SessionsParams{

		timeout: timeout,
	}
}

// NewSessionsParamsWithContext creates a new SessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSessionsParamsWithContext(ctx context.Context) *SessionsParams {
	var ()
	return &SessionsParams{

		Context: ctx,
	}
}

// NewSessionsParamsWithHTTPClient creates a new SessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSessionsParamsWithHTTPClient(client *http.Client) *SessionsParams {
	var ()
	return &SessionsParams{
		HTTPClient: client,
	}
}

/*SessionsParams contains all the parameters to send to the API endpoint
for the sessions operation typically these are written to a http.Request
*/
type SessionsParams struct {

	/*ClusterID
	  cluster_id is the id of the cluster set by user on creation.

	*/
	ClusterID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sessions params
func (o *SessionsParams) WithTimeout(timeout time.Duration) *SessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sessions params
func (o *SessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sessions params
func (o *SessionsParams) WithContext(ctx context.Context) *SessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sessions params
func (o *SessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sessions params
func (o *SessionsParams) WithHTTPClient(client *http.Client) *SessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sessions params
func (o *SessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the sessions params
func (o *SessionsParams) WithClusterID(clusterID string) *SessionsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the sessions params
func (o *SessionsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the sessions params
func (o *SessionsParams) WithLocationOrganizationID(locationOrganizationID string) *SessionsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the sessions params
func (o *SessionsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the sessions params
func (o *SessionsParams) WithLocationProjectID(locationProjectID string) *SessionsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the sessions params
func (o *SessionsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the sessions params
func (o *SessionsParams) WithLocationRegionProvider(locationRegionProvider *string) *SessionsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the sessions params
func (o *SessionsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the sessions params
func (o *SessionsParams) WithLocationRegionRegion(locationRegionRegion *string) *SessionsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the sessions params
func (o *SessionsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *SessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
