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

// NewWaypointServiceUpdateAddOnDefinition3Params creates a new WaypointServiceUpdateAddOnDefinition3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateAddOnDefinition3Params() *WaypointServiceUpdateAddOnDefinition3Params {
	return &WaypointServiceUpdateAddOnDefinition3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateAddOnDefinition3ParamsWithTimeout creates a new WaypointServiceUpdateAddOnDefinition3Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateAddOnDefinition3ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateAddOnDefinition3Params {
	return &WaypointServiceUpdateAddOnDefinition3Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateAddOnDefinition3ParamsWithContext creates a new WaypointServiceUpdateAddOnDefinition3Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateAddOnDefinition3ParamsWithContext(ctx context.Context) *WaypointServiceUpdateAddOnDefinition3Params {
	return &WaypointServiceUpdateAddOnDefinition3Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateAddOnDefinition3ParamsWithHTTPClient creates a new WaypointServiceUpdateAddOnDefinition3Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateAddOnDefinition3ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateAddOnDefinition3Params {
	return &WaypointServiceUpdateAddOnDefinition3Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateAddOnDefinition3Params contains all the parameters to send to the API endpoint

	for the waypoint service update add on definition3 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateAddOnDefinition3Params struct {

	// Body.
	Body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnDefinitionBody

	// ExistingAddOnDefinitionID.
	ExistingAddOnDefinitionID string

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

// WithDefaults hydrates default values in the waypoint service update add on definition3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithDefaults() *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update add on definition3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithContext(ctx context.Context) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnDefinitionBody) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetBody(body *models.HashicorpCloudWaypointV20241122WaypointServiceUpdateAddOnDefinitionBody) {
	o.Body = body
}

// WithExistingAddOnDefinitionID adds the existingAddOnDefinitionID to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithExistingAddOnDefinitionID(existingAddOnDefinitionID string) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetExistingAddOnDefinitionID(existingAddOnDefinitionID)
	return o
}

// SetExistingAddOnDefinitionID adds the existingAddOnDefinitionId to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetExistingAddOnDefinitionID(existingAddOnDefinitionID string) {
	o.ExistingAddOnDefinitionID = existingAddOnDefinitionID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceUpdateAddOnDefinition3Params {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service update add on definition3 params
func (o *WaypointServiceUpdateAddOnDefinition3Params) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateAddOnDefinition3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param existing_add_on_definition.id
	if err := r.SetPathParam("existing_add_on_definition.id", o.ExistingAddOnDefinitionID); err != nil {
		return err
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
