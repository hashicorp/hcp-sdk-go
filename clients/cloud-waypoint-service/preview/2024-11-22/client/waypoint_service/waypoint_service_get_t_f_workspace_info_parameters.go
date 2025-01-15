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

// NewWaypointServiceGetTFWorkspaceInfoParams creates a new WaypointServiceGetTFWorkspaceInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetTFWorkspaceInfoParams() *WaypointServiceGetTFWorkspaceInfoParams {
	return &WaypointServiceGetTFWorkspaceInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetTFWorkspaceInfoParamsWithTimeout creates a new WaypointServiceGetTFWorkspaceInfoParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetTFWorkspaceInfoParamsWithTimeout(timeout time.Duration) *WaypointServiceGetTFWorkspaceInfoParams {
	return &WaypointServiceGetTFWorkspaceInfoParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetTFWorkspaceInfoParamsWithContext creates a new WaypointServiceGetTFWorkspaceInfoParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetTFWorkspaceInfoParamsWithContext(ctx context.Context) *WaypointServiceGetTFWorkspaceInfoParams {
	return &WaypointServiceGetTFWorkspaceInfoParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetTFWorkspaceInfoParamsWithHTTPClient creates a new WaypointServiceGetTFWorkspaceInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetTFWorkspaceInfoParamsWithHTTPClient(client *http.Client) *WaypointServiceGetTFWorkspaceInfoParams {
	return &WaypointServiceGetTFWorkspaceInfoParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetTFWorkspaceInfoParams contains all the parameters to send to the API endpoint

	for the waypoint service get t f workspace info operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetTFWorkspaceInfoParams struct {

	// AddOnID.
	AddOnID *string

	// AddOnName.
	AddOnName *string

	// ApplicationID.
	ApplicationID string

	// ApplicationName.
	ApplicationName *string

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

// WithDefaults hydrates default values in the waypoint service get t f workspace info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithDefaults() *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get t f workspace info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithTimeout(timeout time.Duration) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithContext(ctx context.Context) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithHTTPClient(client *http.Client) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnID adds the addOnID to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithAddOnID(addOnID *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetAddOnID(addOnID)
	return o
}

// SetAddOnID adds the addOnId to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetAddOnID(addOnID *string) {
	o.AddOnID = addOnID
}

// WithAddOnName adds the addOnName to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithAddOnName(addOnName *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetAddOnName(addOnName *string) {
	o.AddOnName = addOnName
}

// WithApplicationID adds the applicationID to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithApplicationID(applicationID string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithApplicationName adds the applicationName to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithApplicationName(applicationName *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetApplicationName(applicationName *string) {
	o.ApplicationName = applicationName
}

// WithNamespaceID adds the namespaceID to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithNamespaceID(namespaceID *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceGetTFWorkspaceInfoParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get t f workspace info params
func (o *WaypointServiceGetTFWorkspaceInfoParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetTFWorkspaceInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param application.id
	if err := r.SetPathParam("application.id", o.ApplicationID); err != nil {
		return err
	}

	if o.ApplicationName != nil {

		// query param application.name
		var qrApplicationName string

		if o.ApplicationName != nil {
			qrApplicationName = *o.ApplicationName
		}
		qApplicationName := qrApplicationName
		if qApplicationName != "" {

			if err := r.SetQueryParam("application.name", qApplicationName); err != nil {
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
