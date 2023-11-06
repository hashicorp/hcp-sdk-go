// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// NewWaypointCreateAddOnDefinitionParams creates a new WaypointCreateAddOnDefinitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointCreateAddOnDefinitionParams() *WaypointCreateAddOnDefinitionParams {
	return &WaypointCreateAddOnDefinitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointCreateAddOnDefinitionParamsWithTimeout creates a new WaypointCreateAddOnDefinitionParams object
// with the ability to set a timeout on a request.
func NewWaypointCreateAddOnDefinitionParamsWithTimeout(timeout time.Duration) *WaypointCreateAddOnDefinitionParams {
	return &WaypointCreateAddOnDefinitionParams{
		timeout: timeout,
	}
}

// NewWaypointCreateAddOnDefinitionParamsWithContext creates a new WaypointCreateAddOnDefinitionParams object
// with the ability to set a context for a request.
func NewWaypointCreateAddOnDefinitionParamsWithContext(ctx context.Context) *WaypointCreateAddOnDefinitionParams {
	return &WaypointCreateAddOnDefinitionParams{
		Context: ctx,
	}
}

// NewWaypointCreateAddOnDefinitionParamsWithHTTPClient creates a new WaypointCreateAddOnDefinitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointCreateAddOnDefinitionParamsWithHTTPClient(client *http.Client) *WaypointCreateAddOnDefinitionParams {
	return &WaypointCreateAddOnDefinitionParams{
		HTTPClient: client,
	}
}

/*
WaypointCreateAddOnDefinitionParams contains all the parameters to send to the API endpoint

	for the waypoint create add on definition operation.

	Typically these are written to a http.Request.
*/
type WaypointCreateAddOnDefinitionParams struct {

	// Body.
	Body *models.HashicorpWaypointCreateAddOnDefinitionRequest

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint create add on definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointCreateAddOnDefinitionParams) WithDefaults() *WaypointCreateAddOnDefinitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint create add on definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointCreateAddOnDefinitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) WithTimeout(timeout time.Duration) *WaypointCreateAddOnDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) WithContext(ctx context.Context) *WaypointCreateAddOnDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) WithHTTPClient(client *http.Client) *WaypointCreateAddOnDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) WithBody(body *models.HashicorpWaypointCreateAddOnDefinitionRequest) *WaypointCreateAddOnDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) SetBody(body *models.HashicorpWaypointCreateAddOnDefinitionRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) WithNamespaceID(namespaceID string) *WaypointCreateAddOnDefinitionParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint create add on definition params
func (o *WaypointCreateAddOnDefinitionParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointCreateAddOnDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}