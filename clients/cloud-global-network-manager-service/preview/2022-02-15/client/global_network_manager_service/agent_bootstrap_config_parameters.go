// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

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

// NewAgentBootstrapConfigParams creates a new AgentBootstrapConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAgentBootstrapConfigParams() *AgentBootstrapConfigParams {
	return &AgentBootstrapConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAgentBootstrapConfigParamsWithTimeout creates a new AgentBootstrapConfigParams object
// with the ability to set a timeout on a request.
func NewAgentBootstrapConfigParamsWithTimeout(timeout time.Duration) *AgentBootstrapConfigParams {
	return &AgentBootstrapConfigParams{
		timeout: timeout,
	}
}

// NewAgentBootstrapConfigParamsWithContext creates a new AgentBootstrapConfigParams object
// with the ability to set a context for a request.
func NewAgentBootstrapConfigParamsWithContext(ctx context.Context) *AgentBootstrapConfigParams {
	return &AgentBootstrapConfigParams{
		Context: ctx,
	}
}

// NewAgentBootstrapConfigParamsWithHTTPClient creates a new AgentBootstrapConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewAgentBootstrapConfigParamsWithHTTPClient(client *http.Client) *AgentBootstrapConfigParams {
	return &AgentBootstrapConfigParams{
		HTTPClient: client,
	}
}

/*
AgentBootstrapConfigParams contains all the parameters to send to the API endpoint

	for the agent bootstrap config operation.

	Typically these are written to a http.Request.
*/
type AgentBootstrapConfigParams struct {

	/* ConsulVersion.

	   Consul version of the cluster
	*/
	ConsulVersion *string

	// ID.
	ID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the agent bootstrap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentBootstrapConfigParams) WithDefaults() *AgentBootstrapConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the agent bootstrap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentBootstrapConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithTimeout(timeout time.Duration) *AgentBootstrapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithContext(ctx context.Context) *AgentBootstrapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithHTTPClient(client *http.Client) *AgentBootstrapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsulVersion adds the consulVersion to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithConsulVersion(consulVersion *string) *AgentBootstrapConfigParams {
	o.SetConsulVersion(consulVersion)
	return o
}

// SetConsulVersion adds the consulVersion to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetConsulVersion(consulVersion *string) {
	o.ConsulVersion = consulVersion
}

// WithID adds the id to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithID(id string) *AgentBootstrapConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithLocationOrganizationID(locationOrganizationID string) *AgentBootstrapConfigParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithLocationProjectID(locationProjectID string) *AgentBootstrapConfigParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithLocationRegionProvider(locationRegionProvider *string) *AgentBootstrapConfigParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) WithLocationRegionRegion(locationRegionRegion *string) *AgentBootstrapConfigParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the agent bootstrap config params
func (o *AgentBootstrapConfigParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *AgentBootstrapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConsulVersion != nil {

		// query param consul_version
		var qrConsulVersion string

		if o.ConsulVersion != nil {
			qrConsulVersion = *o.ConsulVersion
		}
		qConsulVersion := qrConsulVersion
		if qConsulVersion != "" {

			if err := r.SetQueryParam("consul_version", qConsulVersion); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
