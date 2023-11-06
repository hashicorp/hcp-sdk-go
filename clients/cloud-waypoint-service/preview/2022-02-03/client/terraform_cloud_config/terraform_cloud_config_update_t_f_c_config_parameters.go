// Code generated by go-swagger; DO NOT EDIT.

package terraform_cloud_config

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

// NewTerraformCloudConfigUpdateTFCConfigParams creates a new TerraformCloudConfigUpdateTFCConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTerraformCloudConfigUpdateTFCConfigParams() *TerraformCloudConfigUpdateTFCConfigParams {
	return &TerraformCloudConfigUpdateTFCConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTerraformCloudConfigUpdateTFCConfigParamsWithTimeout creates a new TerraformCloudConfigUpdateTFCConfigParams object
// with the ability to set a timeout on a request.
func NewTerraformCloudConfigUpdateTFCConfigParamsWithTimeout(timeout time.Duration) *TerraformCloudConfigUpdateTFCConfigParams {
	return &TerraformCloudConfigUpdateTFCConfigParams{
		timeout: timeout,
	}
}

// NewTerraformCloudConfigUpdateTFCConfigParamsWithContext creates a new TerraformCloudConfigUpdateTFCConfigParams object
// with the ability to set a context for a request.
func NewTerraformCloudConfigUpdateTFCConfigParamsWithContext(ctx context.Context) *TerraformCloudConfigUpdateTFCConfigParams {
	return &TerraformCloudConfigUpdateTFCConfigParams{
		Context: ctx,
	}
}

// NewTerraformCloudConfigUpdateTFCConfigParamsWithHTTPClient creates a new TerraformCloudConfigUpdateTFCConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewTerraformCloudConfigUpdateTFCConfigParamsWithHTTPClient(client *http.Client) *TerraformCloudConfigUpdateTFCConfigParams {
	return &TerraformCloudConfigUpdateTFCConfigParams{
		HTTPClient: client,
	}
}

/*
TerraformCloudConfigUpdateTFCConfigParams contains all the parameters to send to the API endpoint

	for the terraform cloud config update t f c config operation.

	Typically these are written to a http.Request.
*/
type TerraformCloudConfigUpdateTFCConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointUpdateTFCConfigRequest

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the terraform cloud config update t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithDefaults() *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the terraform cloud config update t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithTimeout(timeout time.Duration) *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithContext(ctx context.Context) *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithHTTPClient(client *http.Client) *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithBody(body *models.HashicorpCloudWaypointUpdateTFCConfigRequest) *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetBody(body *models.HashicorpCloudWaypointUpdateTFCConfigRequest) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) WithNamespaceID(namespaceID string) *TerraformCloudConfigUpdateTFCConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the terraform cloud config update t f c config params
func (o *TerraformCloudConfigUpdateTFCConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *TerraformCloudConfigUpdateTFCConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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