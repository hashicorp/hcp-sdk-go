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

// NewGetClusterTelemetryParams creates a new GetClusterTelemetryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterTelemetryParams() *GetClusterTelemetryParams {
	return &GetClusterTelemetryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTelemetryParamsWithTimeout creates a new GetClusterTelemetryParams object
// with the ability to set a timeout on a request.
func NewGetClusterTelemetryParamsWithTimeout(timeout time.Duration) *GetClusterTelemetryParams {
	return &GetClusterTelemetryParams{
		timeout: timeout,
	}
}

// NewGetClusterTelemetryParamsWithContext creates a new GetClusterTelemetryParams object
// with the ability to set a context for a request.
func NewGetClusterTelemetryParamsWithContext(ctx context.Context) *GetClusterTelemetryParams {
	return &GetClusterTelemetryParams{
		Context: ctx,
	}
}

// NewGetClusterTelemetryParamsWithHTTPClient creates a new GetClusterTelemetryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterTelemetryParamsWithHTTPClient(client *http.Client) *GetClusterTelemetryParams {
	return &GetClusterTelemetryParams{
		HTTPClient: client,
	}
}

/*
GetClusterTelemetryParams contains all the parameters to send to the API endpoint

	for the get cluster telemetry operation.

	Typically these are written to a http.Request.
*/
type GetClusterTelemetryParams struct {

	/* ClusterID.

	   cluster_id represents the cluster to retrieve telemetry from
	*/
	ClusterID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster telemetry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTelemetryParams) WithDefaults() *GetClusterTelemetryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster telemetry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTelemetryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithTimeout(timeout time.Duration) *GetClusterTelemetryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithContext(ctx context.Context) *GetClusterTelemetryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithHTTPClient(client *http.Client) *GetClusterTelemetryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithClusterID(clusterID string) *GetClusterTelemetryParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithLocationOrganizationID(locationOrganizationID string) *GetClusterTelemetryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithLocationProjectID(locationProjectID string) *GetClusterTelemetryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithLocationRegionProvider(locationRegionProvider *string) *GetClusterTelemetryParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get cluster telemetry params
func (o *GetClusterTelemetryParams) WithLocationRegionRegion(locationRegionRegion *string) *GetClusterTelemetryParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get cluster telemetry params
func (o *GetClusterTelemetryParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTelemetryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
