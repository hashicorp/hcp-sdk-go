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

// NewWaypointServiceGetIntegrationParams creates a new WaypointServiceGetIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetIntegrationParams() *WaypointServiceGetIntegrationParams {
	return &WaypointServiceGetIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetIntegrationParamsWithTimeout creates a new WaypointServiceGetIntegrationParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetIntegrationParamsWithTimeout(timeout time.Duration) *WaypointServiceGetIntegrationParams {
	return &WaypointServiceGetIntegrationParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetIntegrationParamsWithContext creates a new WaypointServiceGetIntegrationParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetIntegrationParamsWithContext(ctx context.Context) *WaypointServiceGetIntegrationParams {
	return &WaypointServiceGetIntegrationParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetIntegrationParamsWithHTTPClient creates a new WaypointServiceGetIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetIntegrationParamsWithHTTPClient(client *http.Client) *WaypointServiceGetIntegrationParams {
	return &WaypointServiceGetIntegrationParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetIntegrationParams contains all the parameters to send to the API endpoint

	for the waypoint service get integration operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetIntegrationParams struct {

	// IntegrationID.
	IntegrationID string

	// IntegrationName.
	IntegrationName *string

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

// WithDefaults hydrates default values in the waypoint service get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetIntegrationParams) WithDefaults() *WaypointServiceGetIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithTimeout(timeout time.Duration) *WaypointServiceGetIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithContext(ctx context.Context) *WaypointServiceGetIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithHTTPClient(client *http.Client) *WaypointServiceGetIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithIntegrationID(integrationID string) *WaypointServiceGetIntegrationParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WithIntegrationName adds the integrationName to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithIntegrationName(integrationName *string) *WaypointServiceGetIntegrationParams {
	o.SetIntegrationName(integrationName)
	return o
}

// SetIntegrationName adds the integrationName to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetIntegrationName(integrationName *string) {
	o.IntegrationName = integrationName
}

// WithNamespaceID adds the namespaceID to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithNamespaceID(namespaceID *string) *WaypointServiceGetIntegrationParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceGetIntegrationParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceGetIntegrationParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceGetIntegrationParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceGetIntegrationParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get integration params
func (o *WaypointServiceGetIntegrationParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integration.id
	if err := r.SetPathParam("integration.id", o.IntegrationID); err != nil {
		return err
	}

	if o.IntegrationName != nil {

		// query param integration.name
		var qrIntegrationName string

		if o.IntegrationName != nil {
			qrIntegrationName = *o.IntegrationName
		}
		qIntegrationName := qrIntegrationName
		if qIntegrationName != "" {

			if err := r.SetQueryParam("integration.name", qIntegrationName); err != nil {
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
