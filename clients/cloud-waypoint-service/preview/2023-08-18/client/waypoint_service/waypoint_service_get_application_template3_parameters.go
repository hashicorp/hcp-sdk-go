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

// NewWaypointServiceGetApplicationTemplate3Params creates a new WaypointServiceGetApplicationTemplate3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetApplicationTemplate3Params() *WaypointServiceGetApplicationTemplate3Params {
	return &WaypointServiceGetApplicationTemplate3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetApplicationTemplate3ParamsWithTimeout creates a new WaypointServiceGetApplicationTemplate3Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetApplicationTemplate3ParamsWithTimeout(timeout time.Duration) *WaypointServiceGetApplicationTemplate3Params {
	return &WaypointServiceGetApplicationTemplate3Params{
		timeout: timeout,
	}
}

// NewWaypointServiceGetApplicationTemplate3ParamsWithContext creates a new WaypointServiceGetApplicationTemplate3Params object
// with the ability to set a context for a request.
func NewWaypointServiceGetApplicationTemplate3ParamsWithContext(ctx context.Context) *WaypointServiceGetApplicationTemplate3Params {
	return &WaypointServiceGetApplicationTemplate3Params{
		Context: ctx,
	}
}

// NewWaypointServiceGetApplicationTemplate3ParamsWithHTTPClient creates a new WaypointServiceGetApplicationTemplate3Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetApplicationTemplate3ParamsWithHTTPClient(client *http.Client) *WaypointServiceGetApplicationTemplate3Params {
	return &WaypointServiceGetApplicationTemplate3Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetApplicationTemplate3Params contains all the parameters to send to the API endpoint

	for the waypoint service get application template3 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetApplicationTemplate3Params struct {

	/* ApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ApplicationTemplateID string

	/* ApplicationTemplateName.

	   Name of the ApplicationTemplate
	*/
	ApplicationTemplateName *string

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

// WithDefaults hydrates default values in the waypoint service get application template3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetApplicationTemplate3Params) WithDefaults() *WaypointServiceGetApplicationTemplate3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get application template3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetApplicationTemplate3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithTimeout(timeout time.Duration) *WaypointServiceGetApplicationTemplate3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithContext(ctx context.Context) *WaypointServiceGetApplicationTemplate3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithHTTPClient(client *http.Client) *WaypointServiceGetApplicationTemplate3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationTemplateID adds the applicationTemplateID to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithApplicationTemplateID(applicationTemplateID string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetApplicationTemplateID(applicationTemplateID)
	return o
}

// SetApplicationTemplateID adds the applicationTemplateId to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetApplicationTemplateID(applicationTemplateID string) {
	o.ApplicationTemplateID = applicationTemplateID
}

// WithApplicationTemplateName adds the applicationTemplateName to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithApplicationTemplateName(applicationTemplateName *string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetApplicationTemplateName(applicationTemplateName)
	return o
}

// SetApplicationTemplateName adds the applicationTemplateName to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetApplicationTemplateName(applicationTemplateName *string) {
	o.ApplicationTemplateName = applicationTemplateName
}

// WithNamespaceID adds the namespaceID to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithNamespaceID(namespaceID string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID *string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithNamespaceLocationProjectID(namespaceLocationProjectID *string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetNamespaceLocationProjectID(namespaceLocationProjectID *string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceGetApplicationTemplate3Params {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service get application template3 params
func (o *WaypointServiceGetApplicationTemplate3Params) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetApplicationTemplate3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application_template.id
	if err := r.SetPathParam("application_template.id", o.ApplicationTemplateID); err != nil {
		return err
	}

	if o.ApplicationTemplateName != nil {

		// query param application_template.name
		var qrApplicationTemplateName string

		if o.ApplicationTemplateName != nil {
			qrApplicationTemplateName = *o.ApplicationTemplateName
		}
		qApplicationTemplateName := qrApplicationTemplateName
		if qApplicationTemplateName != "" {

			if err := r.SetQueryParam("application_template.name", qApplicationTemplateName); err != nil {
				return err
			}
		}
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
