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

// NewWaypointServiceDeleteApplicationTemplateParams creates a new WaypointServiceDeleteApplicationTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceDeleteApplicationTemplateParams() *WaypointServiceDeleteApplicationTemplateParams {
	return &WaypointServiceDeleteApplicationTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceDeleteApplicationTemplateParamsWithTimeout creates a new WaypointServiceDeleteApplicationTemplateParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceDeleteApplicationTemplateParamsWithTimeout(timeout time.Duration) *WaypointServiceDeleteApplicationTemplateParams {
	return &WaypointServiceDeleteApplicationTemplateParams{
		timeout: timeout,
	}
}

// NewWaypointServiceDeleteApplicationTemplateParamsWithContext creates a new WaypointServiceDeleteApplicationTemplateParams object
// with the ability to set a context for a request.
func NewWaypointServiceDeleteApplicationTemplateParamsWithContext(ctx context.Context) *WaypointServiceDeleteApplicationTemplateParams {
	return &WaypointServiceDeleteApplicationTemplateParams{
		Context: ctx,
	}
}

// NewWaypointServiceDeleteApplicationTemplateParamsWithHTTPClient creates a new WaypointServiceDeleteApplicationTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceDeleteApplicationTemplateParamsWithHTTPClient(client *http.Client) *WaypointServiceDeleteApplicationTemplateParams {
	return &WaypointServiceDeleteApplicationTemplateParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceDeleteApplicationTemplateParams contains all the parameters to send to the API endpoint

	for the waypoint service delete application template operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceDeleteApplicationTemplateParams struct {

	/* ApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ApplicationTemplateID string

	/* ApplicationTemplateName.

	   Name of the ApplicationTemplate
	*/
	ApplicationTemplateName *string

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

// WithDefaults hydrates default values in the waypoint service delete application template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDeleteApplicationTemplateParams) WithDefaults() *WaypointServiceDeleteApplicationTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service delete application template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceDeleteApplicationTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithTimeout(timeout time.Duration) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithContext(ctx context.Context) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithHTTPClient(client *http.Client) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationTemplateID adds the applicationTemplateID to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithApplicationTemplateID(applicationTemplateID string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetApplicationTemplateID(applicationTemplateID)
	return o
}

// SetApplicationTemplateID adds the applicationTemplateId to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetApplicationTemplateID(applicationTemplateID string) {
	o.ApplicationTemplateID = applicationTemplateID
}

// WithApplicationTemplateName adds the applicationTemplateName to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithApplicationTemplateName(applicationTemplateName *string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetApplicationTemplateName(applicationTemplateName)
	return o
}

// SetApplicationTemplateName adds the applicationTemplateName to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetApplicationTemplateName(applicationTemplateName *string) {
	o.ApplicationTemplateName = applicationTemplateName
}

// WithNamespaceID adds the namespaceID to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithNamespaceID(namespaceID *string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceDeleteApplicationTemplateParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service delete application template params
func (o *WaypointServiceDeleteApplicationTemplateParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceDeleteApplicationTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
