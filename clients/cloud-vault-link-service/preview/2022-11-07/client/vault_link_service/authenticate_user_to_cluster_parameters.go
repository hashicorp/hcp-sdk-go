// Code generated by go-swagger; DO NOT EDIT.

package vault_link_service

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

// NewAuthenticateUserToClusterParams creates a new AuthenticateUserToClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticateUserToClusterParams() *AuthenticateUserToClusterParams {
	return &AuthenticateUserToClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticateUserToClusterParamsWithTimeout creates a new AuthenticateUserToClusterParams object
// with the ability to set a timeout on a request.
func NewAuthenticateUserToClusterParamsWithTimeout(timeout time.Duration) *AuthenticateUserToClusterParams {
	return &AuthenticateUserToClusterParams{
		timeout: timeout,
	}
}

// NewAuthenticateUserToClusterParamsWithContext creates a new AuthenticateUserToClusterParams object
// with the ability to set a context for a request.
func NewAuthenticateUserToClusterParamsWithContext(ctx context.Context) *AuthenticateUserToClusterParams {
	return &AuthenticateUserToClusterParams{
		Context: ctx,
	}
}

// NewAuthenticateUserToClusterParamsWithHTTPClient creates a new AuthenticateUserToClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticateUserToClusterParamsWithHTTPClient(client *http.Client) *AuthenticateUserToClusterParams {
	return &AuthenticateUserToClusterParams{
		HTTPClient: client,
	}
}

/*
AuthenticateUserToClusterParams contains all the parameters to send to the API endpoint

	for the authenticate user to cluster operation.

	Typically these are written to a http.Request.
*/
type AuthenticateUserToClusterParams struct {

	// Body.
	Body AuthenticateUserToClusterBody

	// ClusterID.
	ClusterID string

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

// WithDefaults hydrates default values in the authenticate user to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateUserToClusterParams) WithDefaults() *AuthenticateUserToClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authenticate user to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateUserToClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithTimeout(timeout time.Duration) *AuthenticateUserToClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithContext(ctx context.Context) *AuthenticateUserToClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithHTTPClient(client *http.Client) *AuthenticateUserToClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithBody(body AuthenticateUserToClusterBody) *AuthenticateUserToClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetBody(body AuthenticateUserToClusterBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithClusterID(clusterID string) *AuthenticateUserToClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithLocationOrganizationID(locationOrganizationID string) *AuthenticateUserToClusterParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) WithLocationProjectID(locationProjectID string) *AuthenticateUserToClusterParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the authenticate user to cluster params
func (o *AuthenticateUserToClusterParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticateUserToClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
