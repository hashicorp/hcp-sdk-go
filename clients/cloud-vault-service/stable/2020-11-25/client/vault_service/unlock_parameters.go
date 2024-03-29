// Code generated by go-swagger; DO NOT EDIT.

package vault_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// NewUnlockParams creates a new UnlockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnlockParams() *UnlockParams {
	return &UnlockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnlockParamsWithTimeout creates a new UnlockParams object
// with the ability to set a timeout on a request.
func NewUnlockParamsWithTimeout(timeout time.Duration) *UnlockParams {
	return &UnlockParams{
		timeout: timeout,
	}
}

// NewUnlockParamsWithContext creates a new UnlockParams object
// with the ability to set a context for a request.
func NewUnlockParamsWithContext(ctx context.Context) *UnlockParams {
	return &UnlockParams{
		Context: ctx,
	}
}

// NewUnlockParamsWithHTTPClient creates a new UnlockParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnlockParamsWithHTTPClient(client *http.Client) *UnlockParams {
	return &UnlockParams{
		HTTPClient: client,
	}
}

/*
UnlockParams contains all the parameters to send to the API endpoint

	for the unlock operation.

	Typically these are written to a http.Request.
*/
type UnlockParams struct {

	// Body.
	Body *models.HashicorpCloudVault20201125UnlockRequest

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

// WithDefaults hydrates default values in the unlock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnlockParams) WithDefaults() *UnlockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unlock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnlockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unlock params
func (o *UnlockParams) WithTimeout(timeout time.Duration) *UnlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unlock params
func (o *UnlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unlock params
func (o *UnlockParams) WithContext(ctx context.Context) *UnlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unlock params
func (o *UnlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unlock params
func (o *UnlockParams) WithHTTPClient(client *http.Client) *UnlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unlock params
func (o *UnlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unlock params
func (o *UnlockParams) WithBody(body *models.HashicorpCloudVault20201125UnlockRequest) *UnlockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unlock params
func (o *UnlockParams) SetBody(body *models.HashicorpCloudVault20201125UnlockRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the unlock params
func (o *UnlockParams) WithClusterID(clusterID string) *UnlockParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the unlock params
func (o *UnlockParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the unlock params
func (o *UnlockParams) WithLocationOrganizationID(locationOrganizationID string) *UnlockParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the unlock params
func (o *UnlockParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the unlock params
func (o *UnlockParams) WithLocationProjectID(locationProjectID string) *UnlockParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the unlock params
func (o *UnlockParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UnlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
