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

// NewWaypointServiceGetTFRunStatusParams creates a new WaypointServiceGetTFRunStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetTFRunStatusParams() *WaypointServiceGetTFRunStatusParams {
	return &WaypointServiceGetTFRunStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetTFRunStatusParamsWithTimeout creates a new WaypointServiceGetTFRunStatusParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetTFRunStatusParamsWithTimeout(timeout time.Duration) *WaypointServiceGetTFRunStatusParams {
	return &WaypointServiceGetTFRunStatusParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetTFRunStatusParamsWithContext creates a new WaypointServiceGetTFRunStatusParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetTFRunStatusParamsWithContext(ctx context.Context) *WaypointServiceGetTFRunStatusParams {
	return &WaypointServiceGetTFRunStatusParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetTFRunStatusParamsWithHTTPClient creates a new WaypointServiceGetTFRunStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetTFRunStatusParamsWithHTTPClient(client *http.Client) *WaypointServiceGetTFRunStatusParams {
	return &WaypointServiceGetTFRunStatusParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetTFRunStatusParams contains all the parameters to send to the API endpoint

	for the waypoint service get t f run status operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetTFRunStatusParams struct {

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

	// WorkspaceName.
	WorkspaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get t f run status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFRunStatusParams) WithDefaults() *WaypointServiceGetTFRunStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get t f run status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFRunStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithTimeout(timeout time.Duration) *WaypointServiceGetTFRunStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithContext(ctx context.Context) *WaypointServiceGetTFRunStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithHTTPClient(client *http.Client) *WaypointServiceGetTFRunStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithNamespaceID(namespaceID string) *WaypointServiceGetTFRunStatusParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) *WaypointServiceGetTFRunStatusParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithNamespaceLocationProjectID(namespaceLocationProjectID *string) *WaypointServiceGetTFRunStatusParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetNamespaceLocationProjectID(namespaceLocationProjectID *string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceGetTFRunStatusParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceGetTFRunStatusParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WithWorkspaceName adds the workspaceName to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) WithWorkspaceName(workspaceName string) *WaypointServiceGetTFRunStatusParams {
	o.SetWorkspaceName(workspaceName)
	return o
}

// SetWorkspaceName adds the workspaceName to the waypoint service get t f run status params
func (o *WaypointServiceGetTFRunStatusParams) SetWorkspaceName(workspaceName string) {
	o.WorkspaceName = workspaceName
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetTFRunStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param workspace_name
	if err := r.SetPathParam("workspace_name", o.WorkspaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
