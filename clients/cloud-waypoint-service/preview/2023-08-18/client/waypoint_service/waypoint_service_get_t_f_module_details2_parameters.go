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

// NewWaypointServiceGetTFModuleDetails2Params creates a new WaypointServiceGetTFModuleDetails2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetTFModuleDetails2Params() *WaypointServiceGetTFModuleDetails2Params {
	return &WaypointServiceGetTFModuleDetails2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetTFModuleDetails2ParamsWithTimeout creates a new WaypointServiceGetTFModuleDetails2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetTFModuleDetails2ParamsWithTimeout(timeout time.Duration) *WaypointServiceGetTFModuleDetails2Params {
	return &WaypointServiceGetTFModuleDetails2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceGetTFModuleDetails2ParamsWithContext creates a new WaypointServiceGetTFModuleDetails2Params object
// with the ability to set a context for a request.
func NewWaypointServiceGetTFModuleDetails2ParamsWithContext(ctx context.Context) *WaypointServiceGetTFModuleDetails2Params {
	return &WaypointServiceGetTFModuleDetails2Params{
		Context: ctx,
	}
}

// NewWaypointServiceGetTFModuleDetails2ParamsWithHTTPClient creates a new WaypointServiceGetTFModuleDetails2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetTFModuleDetails2ParamsWithHTTPClient(client *http.Client) *WaypointServiceGetTFModuleDetails2Params {
	return &WaypointServiceGetTFModuleDetails2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetTFModuleDetails2Params contains all the parameters to send to the API endpoint

	for the waypoint service get t f module details2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetTFModuleDetails2Params struct {

	/* Name.

	   name is the name of the Terraform module.
	*/
	Name string

	// NamespaceID.
	NamespaceID string

	/* Provider.

	   provider is the name of the provider for the Terraform module.
	*/
	Provider string

	/* TfcNamespace.

	   tfc_namespace is the Terraform user who owns the Terraform module.
	*/
	TfcNamespace string

	/* Version.

	   DEPRECATED: Do not use.
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get t f module details2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFModuleDetails2Params) WithDefaults() *WaypointServiceGetTFModuleDetails2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get t f module details2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFModuleDetails2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithTimeout(timeout time.Duration) *WaypointServiceGetTFModuleDetails2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithContext(ctx context.Context) *WaypointServiceGetTFModuleDetails2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithHTTPClient(client *http.Client) *WaypointServiceGetTFModuleDetails2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithName(name string) *WaypointServiceGetTFModuleDetails2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetName(name string) {
	o.Name = name
}

// WithNamespaceID adds the namespaceID to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithNamespaceID(namespaceID string) *WaypointServiceGetTFModuleDetails2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProvider adds the provider to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithProvider(provider string) *WaypointServiceGetTFModuleDetails2Params {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetProvider(provider string) {
	o.Provider = provider
}

// WithTfcNamespace adds the tfcNamespace to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithTfcNamespace(tfcNamespace string) *WaypointServiceGetTFModuleDetails2Params {
	o.SetTfcNamespace(tfcNamespace)
	return o
}

// SetTfcNamespace adds the tfcNamespace to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetTfcNamespace(tfcNamespace string) {
	o.TfcNamespace = tfcNamespace
}

// WithVersion adds the version to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) WithVersion(version *string) *WaypointServiceGetTFModuleDetails2Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the waypoint service get t f module details2 params
func (o *WaypointServiceGetTFModuleDetails2Params) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetTFModuleDetails2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	// path param tfc_namespace
	if err := r.SetPathParam("tfc_namespace", o.TfcNamespace); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
