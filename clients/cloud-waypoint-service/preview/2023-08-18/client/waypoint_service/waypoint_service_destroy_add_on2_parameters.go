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

// NewWaypointServiceDestroyAddOn2Params creates a new WaypointServiceDestroyAddOn2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceDestroyAddOn2Params() *WaypointServiceDestroyAddOn2Params {
	return &WaypointServiceDestroyAddOn2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceDestroyAddOn2ParamsWithTimeout creates a new WaypointServiceDestroyAddOn2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceDestroyAddOn2ParamsWithTimeout(timeout time.Duration) *WaypointServiceDestroyAddOn2Params {
	return &WaypointServiceDestroyAddOn2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceDestroyAddOn2ParamsWithContext creates a new WaypointServiceDestroyAddOn2Params object
// with the ability to set a context for a request.
func NewWaypointServiceDestroyAddOn2ParamsWithContext(ctx context.Context) *WaypointServiceDestroyAddOn2Params {
	return &WaypointServiceDestroyAddOn2Params{
		Context: ctx,
	}
}

// NewWaypointServiceDestroyAddOn2ParamsWithHTTPClient creates a new WaypointServiceDestroyAddOn2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceDestroyAddOn2ParamsWithHTTPClient(client *http.Client) *WaypointServiceDestroyAddOn2Params {
	return &WaypointServiceDestroyAddOn2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceDestroyAddOn2Params contains all the parameters to send to the API endpoint

	for the waypoint service destroy add on2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceDestroyAddOn2Params struct {

	// AddOnID.
	AddOnID *string

	// AddOnName.
	AddOnName string

	// NamespaceID.
	NamespaceID string

	/* NamespaceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	NamespaceLocationOrganizationID *string

	/* NamespaceLocationProjectID.

	   project_id is the projects id.
	*/
	NamespaceLocationProjectID *string

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

// WithDefaults hydrates default values in the waypoint service destroy add on2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDestroyAddOn2Params) WithDefaults() *WaypointServiceDestroyAddOn2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service destroy add on2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDestroyAddOn2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithTimeout(timeout time.Duration) *WaypointServiceDestroyAddOn2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithContext(ctx context.Context) *WaypointServiceDestroyAddOn2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithHTTPClient(client *http.Client) *WaypointServiceDestroyAddOn2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnID adds the addOnID to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithAddOnID(addOnID *string) *WaypointServiceDestroyAddOn2Params {
	o.SetAddOnID(addOnID)
	return o
}

// SetAddOnID adds the addOnId to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetAddOnID(addOnID *string) {
	o.AddOnID = addOnID
}

// WithAddOnName adds the addOnName to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithAddOnName(addOnName string) *WaypointServiceDestroyAddOn2Params {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetAddOnName(addOnName string) {
	o.AddOnName = addOnName
}

// WithNamespaceID adds the namespaceID to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithNamespaceID(namespaceID string) *WaypointServiceDestroyAddOn2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) *WaypointServiceDestroyAddOn2Params {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithNamespaceLocationProjectID(namespaceLocationProjectID *string) *WaypointServiceDestroyAddOn2Params {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetNamespaceLocationProjectID(namespaceLocationProjectID *string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceDestroyAddOn2Params {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceDestroyAddOn2Params {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service destroy add on2 params
func (o *WaypointServiceDestroyAddOn2Params) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceDestroyAddOn2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddOnID != nil {

		// query param add_on.id
		var qrAddOnID string

		if o.AddOnID != nil {
			qrAddOnID = *o.AddOnID
		}
		qAddOnID := qrAddOnID
		if qAddOnID != "" {

			if err := r.SetQueryParam("add_on.id", qAddOnID); err != nil {
				return err
			}
		}
	}

	// path param add_on.name
	if err := r.SetPathParam("add_on.name", o.AddOnName); err != nil {
		return err
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if o.NamespaceLocationOrganizationID != nil {

		// query param namespace.location.organization_id
		var qrNamespaceLocationOrganizationID string

		if o.NamespaceLocationOrganizationID != nil {
			qrNamespaceLocationOrganizationID = *o.NamespaceLocationOrganizationID
		}
		qNamespaceLocationOrganizationID := qrNamespaceLocationOrganizationID
		if qNamespaceLocationOrganizationID != "" {

			if err := r.SetQueryParam("namespace.location.organization_id", qNamespaceLocationOrganizationID); err != nil {
				return err
			}
		}
	}

	if o.NamespaceLocationProjectID != nil {

		// query param namespace.location.project_id
		var qrNamespaceLocationProjectID string

		if o.NamespaceLocationProjectID != nil {
			qrNamespaceLocationProjectID = *o.NamespaceLocationProjectID
		}
		qNamespaceLocationProjectID := qrNamespaceLocationProjectID
		if qNamespaceLocationProjectID != "" {

			if err := r.SetQueryParam("namespace.location.project_id", qNamespaceLocationProjectID); err != nil {
				return err
			}
		}
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
