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

// NewWaypointServiceUILoadProductBannerParams creates a new WaypointServiceUILoadProductBannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUILoadProductBannerParams() *WaypointServiceUILoadProductBannerParams {
	return &WaypointServiceUILoadProductBannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUILoadProductBannerParamsWithTimeout creates a new WaypointServiceUILoadProductBannerParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUILoadProductBannerParamsWithTimeout(timeout time.Duration) *WaypointServiceUILoadProductBannerParams {
	return &WaypointServiceUILoadProductBannerParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUILoadProductBannerParamsWithContext creates a new WaypointServiceUILoadProductBannerParams object
// with the ability to set a context for a request.
func NewWaypointServiceUILoadProductBannerParamsWithContext(ctx context.Context) *WaypointServiceUILoadProductBannerParams {
	return &WaypointServiceUILoadProductBannerParams{
		Context: ctx,
	}
}

// NewWaypointServiceUILoadProductBannerParamsWithHTTPClient creates a new WaypointServiceUILoadProductBannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUILoadProductBannerParamsWithHTTPClient(client *http.Client) *WaypointServiceUILoadProductBannerParams {
	return &WaypointServiceUILoadProductBannerParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUILoadProductBannerParams contains all the parameters to send to the API endpoint

	for the waypoint service UI load product banner operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUILoadProductBannerParams struct {

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service UI load product banner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUILoadProductBannerParams) WithDefaults() *WaypointServiceUILoadProductBannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service UI load product banner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUILoadProductBannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) WithTimeout(timeout time.Duration) *WaypointServiceUILoadProductBannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) WithContext(ctx context.Context) *WaypointServiceUILoadProductBannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) WithHTTPClient(client *http.Client) *WaypointServiceUILoadProductBannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) WithNamespaceID(namespaceID string) *WaypointServiceUILoadProductBannerParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service UI load product banner params
func (o *WaypointServiceUILoadProductBannerParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUILoadProductBannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
