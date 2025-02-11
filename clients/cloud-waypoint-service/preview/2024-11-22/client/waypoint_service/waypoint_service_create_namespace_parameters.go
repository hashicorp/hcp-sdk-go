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

// NewWaypointServiceCreateNamespaceParams creates a new WaypointServiceCreateNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateNamespaceParams() *WaypointServiceCreateNamespaceParams {
	return &WaypointServiceCreateNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateNamespaceParamsWithTimeout creates a new WaypointServiceCreateNamespaceParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateNamespaceParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateNamespaceParams {
	return &WaypointServiceCreateNamespaceParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateNamespaceParamsWithContext creates a new WaypointServiceCreateNamespaceParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateNamespaceParamsWithContext(ctx context.Context) *WaypointServiceCreateNamespaceParams {
	return &WaypointServiceCreateNamespaceParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateNamespaceParamsWithHTTPClient creates a new WaypointServiceCreateNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateNamespaceParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateNamespaceParams {
	return &WaypointServiceCreateNamespaceParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateNamespaceParams contains all the parameters to send to the API endpoint

	for the waypoint service create namespace operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateNamespaceParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateNamespaceBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service create namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateNamespaceParams) WithDefaults() *WaypointServiceCreateNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateNamespaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithContext(ctx context.Context) *WaypointServiceCreateNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateNamespaceBody) *WaypointServiceCreateNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceCreateNamespaceBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithLocationOrganizationID(locationOrganizationID string) *WaypointServiceCreateNamespaceParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) WithLocationProjectID(locationProjectID string) *WaypointServiceCreateNamespaceParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the waypoint service create namespace params
func (o *WaypointServiceCreateNamespaceParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
