// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

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

// NewWaypointServiceListNoCodeModulesParams creates a new WaypointServiceListNoCodeModulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceListNoCodeModulesParams() *WaypointServiceListNoCodeModulesParams {
	return &WaypointServiceListNoCodeModulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceListNoCodeModulesParamsWithTimeout creates a new WaypointServiceListNoCodeModulesParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceListNoCodeModulesParamsWithTimeout(timeout time.Duration) *WaypointServiceListNoCodeModulesParams {
	return &WaypointServiceListNoCodeModulesParams{
		timeout: timeout,
	}
}

// NewWaypointServiceListNoCodeModulesParamsWithContext creates a new WaypointServiceListNoCodeModulesParams object
// with the ability to set a context for a request.
func NewWaypointServiceListNoCodeModulesParamsWithContext(ctx context.Context) *WaypointServiceListNoCodeModulesParams {
	return &WaypointServiceListNoCodeModulesParams{
		Context: ctx,
	}
}

// NewWaypointServiceListNoCodeModulesParamsWithHTTPClient creates a new WaypointServiceListNoCodeModulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceListNoCodeModulesParamsWithHTTPClient(client *http.Client) *WaypointServiceListNoCodeModulesParams {
	return &WaypointServiceListNoCodeModulesParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceListNoCodeModulesParams contains all the parameters to send to the API endpoint

	for the waypoint service list no code modules operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceListNoCodeModulesParams struct {

	// NamespaceID.
	NamespaceID *string

	/* NamespaceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	NamespaceLocationOrganizationID string

	/* NamespaceLocationProjectID.

	   project_id is the projects id.
	*/
	NamespaceLocationProjectID string

	/* NamespaceLocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	NamespaceLocationRegionProvider *string

	/* NamespaceLocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	NamespaceLocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service list no code modules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListNoCodeModulesParams) WithDefaults() *WaypointServiceListNoCodeModulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service list no code modules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListNoCodeModulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithTimeout(timeout time.Duration) *WaypointServiceListNoCodeModulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithContext(ctx context.Context) *WaypointServiceListNoCodeModulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithHTTPClient(client *http.Client) *WaypointServiceListNoCodeModulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithNamespaceID(namespaceID *string) *WaypointServiceListNoCodeModulesParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceListNoCodeModulesParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceListNoCodeModulesParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceListNoCodeModulesParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceListNoCodeModulesParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service list no code modules params
func (o *WaypointServiceListNoCodeModulesParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceListNoCodeModulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NamespaceID != nil {

		// query param namespace.id
		var qrNamespaceID string

		if o.NamespaceID != nil {
			qrNamespaceID = *o.NamespaceID
		}
		qNamespaceID := qrNamespaceID
		if qNamespaceID != "" {

			if err := r.SetQueryParam("namespace.id", qNamespaceID); err != nil {
				return err
			}
		}
	}

	// path param namespace.location.organization_id
	if err := r.SetPathParam("namespace.location.organization_id", o.NamespaceLocationOrganizationID); err != nil {
		return err
	}

	// path param namespace.location.project_id
	if err := r.SetPathParam("namespace.location.project_id", o.NamespaceLocationProjectID); err != nil {
		return err
	}

	if o.NamespaceLocationRegionProvider != nil {

		// query param namespace.location.region.provider
		var qrNamespaceLocationRegionProvider string

		if o.NamespaceLocationRegionProvider != nil {
			qrNamespaceLocationRegionProvider = *o.NamespaceLocationRegionProvider
		}
		qNamespaceLocationRegionProvider := qrNamespaceLocationRegionProvider
		if qNamespaceLocationRegionProvider != "" {

			if err := r.SetQueryParam("namespace.location.region.provider", qNamespaceLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.NamespaceLocationRegionRegion != nil {

		// query param namespace.location.region.region
		var qrNamespaceLocationRegionRegion string

		if o.NamespaceLocationRegionRegion != nil {
			qrNamespaceLocationRegionRegion = *o.NamespaceLocationRegionRegion
		}
		qNamespaceLocationRegionRegion := qrNamespaceLocationRegionRegion
		if qNamespaceLocationRegionRegion != "" {

			if err := r.SetQueryParam("namespace.location.region.region", qNamespaceLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
