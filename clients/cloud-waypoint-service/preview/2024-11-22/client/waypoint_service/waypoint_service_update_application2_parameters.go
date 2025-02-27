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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2024-11-22/models"
)

// NewWaypointServiceUpdateApplication2Params creates a new WaypointServiceUpdateApplication2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplication2Params() *WaypointServiceUpdateApplication2Params {
	return &WaypointServiceUpdateApplication2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplication2ParamsWithTimeout creates a new WaypointServiceUpdateApplication2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplication2ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplication2Params {
	return &WaypointServiceUpdateApplication2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplication2ParamsWithContext creates a new WaypointServiceUpdateApplication2Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplication2ParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplication2Params {
	return &WaypointServiceUpdateApplication2Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplication2ParamsWithHTTPClient creates a new WaypointServiceUpdateApplication2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplication2ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplication2Params {
	return &WaypointServiceUpdateApplication2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplication2Params contains all the parameters to send to the API endpoint

	for the waypoint service update application2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplication2Params struct {

	// ApplicationName.
	ApplicationName string

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody

	/* NamespaceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	NamespaceLocationOrganizationID string

	/* NamespaceLocationProjectID.

	   project_id is the projects id.
	*/
	NamespaceLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update application2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplication2Params) WithDefaults() *WaypointServiceUpdateApplication2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplication2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplication2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithContext(ctx context.Context) *WaypointServiceUpdateApplication2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplication2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationName adds the applicationName to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithApplicationName(applicationName string) *WaypointServiceUpdateApplication2Params {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetApplicationName(applicationName string) {
	o.ApplicationName = applicationName
}

// WithBody adds the body to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody) *WaypointServiceUpdateApplication2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateApplicationBody) {
	o.Body = body
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceUpdateApplication2Params {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceUpdateApplication2Params {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service update application2 params
func (o *WaypointServiceUpdateApplication2Params) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplication2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.name
	if err := r.SetPathParam("application.name", o.ApplicationName); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
