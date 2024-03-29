// Code generated by go-swagger; DO NOT EDIT.

package vault_link_service

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

// NewGetLinkedClusterPolicyParams creates a new GetLinkedClusterPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLinkedClusterPolicyParams() *GetLinkedClusterPolicyParams {
	return &GetLinkedClusterPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLinkedClusterPolicyParamsWithTimeout creates a new GetLinkedClusterPolicyParams object
// with the ability to set a timeout on a request.
func NewGetLinkedClusterPolicyParamsWithTimeout(timeout time.Duration) *GetLinkedClusterPolicyParams {
	return &GetLinkedClusterPolicyParams{
		timeout: timeout,
	}
}

// NewGetLinkedClusterPolicyParamsWithContext creates a new GetLinkedClusterPolicyParams object
// with the ability to set a context for a request.
func NewGetLinkedClusterPolicyParamsWithContext(ctx context.Context) *GetLinkedClusterPolicyParams {
	return &GetLinkedClusterPolicyParams{
		Context: ctx,
	}
}

// NewGetLinkedClusterPolicyParamsWithHTTPClient creates a new GetLinkedClusterPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLinkedClusterPolicyParamsWithHTTPClient(client *http.Client) *GetLinkedClusterPolicyParams {
	return &GetLinkedClusterPolicyParams{
		HTTPClient: client,
	}
}

/*
GetLinkedClusterPolicyParams contains all the parameters to send to the API endpoint

	for the get linked cluster policy operation.

	Typically these are written to a http.Request.
*/
type GetLinkedClusterPolicyParams struct {

	// ClusterID.
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

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get linked cluster policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkedClusterPolicyParams) WithDefaults() *GetLinkedClusterPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get linked cluster policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkedClusterPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithTimeout(timeout time.Duration) *GetLinkedClusterPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithContext(ctx context.Context) *GetLinkedClusterPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithHTTPClient(client *http.Client) *GetLinkedClusterPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithClusterID(clusterID string) *GetLinkedClusterPolicyParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithLocationOrganizationID(locationOrganizationID string) *GetLinkedClusterPolicyParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithLocationProjectID(locationProjectID string) *GetLinkedClusterPolicyParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithLocationRegionProvider(locationRegionProvider *string) *GetLinkedClusterPolicyParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) WithLocationRegionRegion(locationRegionRegion *string) *GetLinkedClusterPolicyParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get linked cluster policy params
func (o *GetLinkedClusterPolicyParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetLinkedClusterPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
