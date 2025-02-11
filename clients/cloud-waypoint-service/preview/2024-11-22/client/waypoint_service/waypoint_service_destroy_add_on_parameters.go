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

// NewWaypointServiceDestroyAddOnParams creates a new WaypointServiceDestroyAddOnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceDestroyAddOnParams() *WaypointServiceDestroyAddOnParams {
	return &WaypointServiceDestroyAddOnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceDestroyAddOnParamsWithTimeout creates a new WaypointServiceDestroyAddOnParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceDestroyAddOnParamsWithTimeout(timeout time.Duration) *WaypointServiceDestroyAddOnParams {
	return &WaypointServiceDestroyAddOnParams{
		timeout: timeout,
	}
}

// NewWaypointServiceDestroyAddOnParamsWithContext creates a new WaypointServiceDestroyAddOnParams object
// with the ability to set a context for a request.
func NewWaypointServiceDestroyAddOnParamsWithContext(ctx context.Context) *WaypointServiceDestroyAddOnParams {
	return &WaypointServiceDestroyAddOnParams{
		Context: ctx,
	}
}

// NewWaypointServiceDestroyAddOnParamsWithHTTPClient creates a new WaypointServiceDestroyAddOnParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceDestroyAddOnParamsWithHTTPClient(client *http.Client) *WaypointServiceDestroyAddOnParams {
	return &WaypointServiceDestroyAddOnParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceDestroyAddOnParams contains all the parameters to send to the API endpoint

	for the waypoint service destroy add on operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceDestroyAddOnParams struct {

	// AddOnID.
	AddOnID string

	// AddOnName.
	AddOnName *string

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

// WithDefaults hydrates default values in the waypoint service destroy add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDestroyAddOnParams) WithDefaults() *WaypointServiceDestroyAddOnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service destroy add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDestroyAddOnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithTimeout(timeout time.Duration) *WaypointServiceDestroyAddOnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithContext(ctx context.Context) *WaypointServiceDestroyAddOnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithHTTPClient(client *http.Client) *WaypointServiceDestroyAddOnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnID adds the addOnID to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithAddOnID(addOnID string) *WaypointServiceDestroyAddOnParams {
	o.SetAddOnID(addOnID)
	return o
}

// SetAddOnID adds the addOnId to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetAddOnID(addOnID string) {
	o.AddOnID = addOnID
}

// WithAddOnName adds the addOnName to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithAddOnName(addOnName *string) *WaypointServiceDestroyAddOnParams {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetAddOnName(addOnName *string) {
	o.AddOnName = addOnName
}

// WithNamespaceID adds the namespaceID to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithNamespaceID(namespaceID *string) *WaypointServiceDestroyAddOnParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceDestroyAddOnParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceDestroyAddOnParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceDestroyAddOnParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceDestroyAddOnParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service destroy add on params
func (o *WaypointServiceDestroyAddOnParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceDestroyAddOnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param add_on.id
	if err := r.SetPathParam("add_on.id", o.AddOnID); err != nil {
		return err
	}

	if o.AddOnName != nil {

		// query param add_on.name
		var qrAddOnName string

		if o.AddOnName != nil {
			qrAddOnName = *o.AddOnName
		}
		qAddOnName := qrAddOnName
		if qAddOnName != "" {

			if err := r.SetQueryParam("add_on.name", qAddOnName); err != nil {
				return err
			}
		}
	}

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
