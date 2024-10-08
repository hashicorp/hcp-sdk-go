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

// NewWaypointServiceGetTFWorkspaceInfo2Params creates a new WaypointServiceGetTFWorkspaceInfo2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetTFWorkspaceInfo2Params() *WaypointServiceGetTFWorkspaceInfo2Params {
	return &WaypointServiceGetTFWorkspaceInfo2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetTFWorkspaceInfo2ParamsWithTimeout creates a new WaypointServiceGetTFWorkspaceInfo2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetTFWorkspaceInfo2ParamsWithTimeout(timeout time.Duration) *WaypointServiceGetTFWorkspaceInfo2Params {
	return &WaypointServiceGetTFWorkspaceInfo2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceGetTFWorkspaceInfo2ParamsWithContext creates a new WaypointServiceGetTFWorkspaceInfo2Params object
// with the ability to set a context for a request.
func NewWaypointServiceGetTFWorkspaceInfo2ParamsWithContext(ctx context.Context) *WaypointServiceGetTFWorkspaceInfo2Params {
	return &WaypointServiceGetTFWorkspaceInfo2Params{
		Context: ctx,
	}
}

// NewWaypointServiceGetTFWorkspaceInfo2ParamsWithHTTPClient creates a new WaypointServiceGetTFWorkspaceInfo2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetTFWorkspaceInfo2ParamsWithHTTPClient(client *http.Client) *WaypointServiceGetTFWorkspaceInfo2Params {
	return &WaypointServiceGetTFWorkspaceInfo2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetTFWorkspaceInfo2Params contains all the parameters to send to the API endpoint

	for the waypoint service get t f workspace info2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetTFWorkspaceInfo2Params struct {

	// AddOnID.
	AddOnID *string

	// AddOnName.
	AddOnName *string

	// ApplicationID.
	ApplicationID *string

	// ApplicationName.
	ApplicationName string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get t f workspace info2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithDefaults() *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get t f workspace info2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithTimeout(timeout time.Duration) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithContext(ctx context.Context) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithHTTPClient(client *http.Client) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnID adds the addOnID to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithAddOnID(addOnID *string) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetAddOnID(addOnID)
	return o
}

// SetAddOnID adds the addOnId to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetAddOnID(addOnID *string) {
	o.AddOnID = addOnID
}

// WithAddOnName adds the addOnName to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithAddOnName(addOnName *string) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetAddOnName(addOnName *string) {
	o.AddOnName = addOnName
}

// WithApplicationID adds the applicationID to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithApplicationID(applicationID *string) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetApplicationID(applicationID *string) {
	o.ApplicationID = applicationID
}

// WithApplicationName adds the applicationName to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithApplicationName(applicationName string) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetApplicationName(applicationName string) {
	o.ApplicationName = applicationName
}

// WithNamespaceID adds the namespaceID to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WithNamespaceID(namespaceID string) *WaypointServiceGetTFWorkspaceInfo2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get t f workspace info2 params
func (o *WaypointServiceGetTFWorkspaceInfo2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetTFWorkspaceInfo2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ApplicationID != nil {

		// query param application.id
		var qrApplicationID string

		if o.ApplicationID != nil {
			qrApplicationID = *o.ApplicationID
		}
		qApplicationID := qrApplicationID
		if qApplicationID != "" {

			if err := r.SetQueryParam("application.id", qApplicationID); err != nil {
				return err
			}
		}
	}

	// path param application.name
	if err := r.SetPathParam("application.name", o.ApplicationName); err != nil {
		return err
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
