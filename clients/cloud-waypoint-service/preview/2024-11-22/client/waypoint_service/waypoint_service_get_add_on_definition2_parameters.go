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

// NewWaypointServiceGetAddOnDefinition2Params creates a new WaypointServiceGetAddOnDefinition2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetAddOnDefinition2Params() *WaypointServiceGetAddOnDefinition2Params {
	return &WaypointServiceGetAddOnDefinition2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetAddOnDefinition2ParamsWithTimeout creates a new WaypointServiceGetAddOnDefinition2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetAddOnDefinition2ParamsWithTimeout(timeout time.Duration) *WaypointServiceGetAddOnDefinition2Params {
	return &WaypointServiceGetAddOnDefinition2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceGetAddOnDefinition2ParamsWithContext creates a new WaypointServiceGetAddOnDefinition2Params object
// with the ability to set a context for a request.
func NewWaypointServiceGetAddOnDefinition2ParamsWithContext(ctx context.Context) *WaypointServiceGetAddOnDefinition2Params {
	return &WaypointServiceGetAddOnDefinition2Params{
		Context: ctx,
	}
}

// NewWaypointServiceGetAddOnDefinition2ParamsWithHTTPClient creates a new WaypointServiceGetAddOnDefinition2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetAddOnDefinition2ParamsWithHTTPClient(client *http.Client) *WaypointServiceGetAddOnDefinition2Params {
	return &WaypointServiceGetAddOnDefinition2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetAddOnDefinition2Params contains all the parameters to send to the API endpoint

	for the waypoint service get add on definition2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetAddOnDefinition2Params struct {

	// AddOnDefinitionID.
	AddOnDefinitionID *string

	// AddOnDefinitionName.
	AddOnDefinitionName string

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

// WithDefaults hydrates default values in the waypoint service get add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetAddOnDefinition2Params) WithDefaults() *WaypointServiceGetAddOnDefinition2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get add on definition2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetAddOnDefinition2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithTimeout(timeout time.Duration) *WaypointServiceGetAddOnDefinition2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithContext(ctx context.Context) *WaypointServiceGetAddOnDefinition2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithHTTPClient(client *http.Client) *WaypointServiceGetAddOnDefinition2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnDefinitionID adds the addOnDefinitionID to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithAddOnDefinitionID(addOnDefinitionID *string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetAddOnDefinitionID(addOnDefinitionID)
	return o
}

// SetAddOnDefinitionID adds the addOnDefinitionId to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetAddOnDefinitionID(addOnDefinitionID *string) {
	o.AddOnDefinitionID = addOnDefinitionID
}

// WithAddOnDefinitionName adds the addOnDefinitionName to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithAddOnDefinitionName(addOnDefinitionName string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetAddOnDefinitionName(addOnDefinitionName)
	return o
}

// SetAddOnDefinitionName adds the addOnDefinitionName to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetAddOnDefinitionName(addOnDefinitionName string) {
	o.AddOnDefinitionName = addOnDefinitionName
}

// WithNamespaceID adds the namespaceID to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithNamespaceID(namespaceID *string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceGetAddOnDefinition2Params {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get add on definition2 params
func (o *WaypointServiceGetAddOnDefinition2Params) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetAddOnDefinition2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddOnDefinitionID != nil {

		// query param add_on_definition.id
		var qrAddOnDefinitionID string

		if o.AddOnDefinitionID != nil {
			qrAddOnDefinitionID = *o.AddOnDefinitionID
		}
		qAddOnDefinitionID := qrAddOnDefinitionID
		if qAddOnDefinitionID != "" {

			if err := r.SetQueryParam("add_on_definition.id", qAddOnDefinitionID); err != nil {
				return err
			}
		}
	}

	// path param add_on_definition.name
	if err := r.SetPathParam("add_on_definition.name", o.AddOnDefinitionName); err != nil {
		return err
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
