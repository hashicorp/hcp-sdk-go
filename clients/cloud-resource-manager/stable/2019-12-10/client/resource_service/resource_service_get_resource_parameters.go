// Code generated by go-swagger; DO NOT EDIT.

package resource_service

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

// NewResourceServiceGetResourceParams creates a new ResourceServiceGetResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceServiceGetResourceParams() *ResourceServiceGetResourceParams {
	return &ResourceServiceGetResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceServiceGetResourceParamsWithTimeout creates a new ResourceServiceGetResourceParams object
// with the ability to set a timeout on a request.
func NewResourceServiceGetResourceParamsWithTimeout(timeout time.Duration) *ResourceServiceGetResourceParams {
	return &ResourceServiceGetResourceParams{
		timeout: timeout,
	}
}

// NewResourceServiceGetResourceParamsWithContext creates a new ResourceServiceGetResourceParams object
// with the ability to set a context for a request.
func NewResourceServiceGetResourceParamsWithContext(ctx context.Context) *ResourceServiceGetResourceParams {
	return &ResourceServiceGetResourceParams{
		Context: ctx,
	}
}

// NewResourceServiceGetResourceParamsWithHTTPClient creates a new ResourceServiceGetResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceServiceGetResourceParamsWithHTTPClient(client *http.Client) *ResourceServiceGetResourceParams {
	return &ResourceServiceGetResourceParams{
		HTTPClient: client,
	}
}

/*
ResourceServiceGetResourceParams contains all the parameters to send to the API endpoint

	for the resource service get resource operation.

	Typically these are written to a http.Request.
*/
type ResourceServiceGetResourceParams struct {

	/* ResourceID.

	   ResourceId is the unique identifier of the resource.
	*/
	ResourceID *string

	/* ResourceName.

	   ResourceName is the name of the resource.
	*/
	ResourceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource service get resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceGetResourceParams) WithDefaults() *ResourceServiceGetResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource service get resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceServiceGetResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource service get resource params
func (o *ResourceServiceGetResourceParams) WithTimeout(timeout time.Duration) *ResourceServiceGetResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource service get resource params
func (o *ResourceServiceGetResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource service get resource params
func (o *ResourceServiceGetResourceParams) WithContext(ctx context.Context) *ResourceServiceGetResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource service get resource params
func (o *ResourceServiceGetResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource service get resource params
func (o *ResourceServiceGetResourceParams) WithHTTPClient(client *http.Client) *ResourceServiceGetResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource service get resource params
func (o *ResourceServiceGetResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the resource service get resource params
func (o *ResourceServiceGetResourceParams) WithResourceID(resourceID *string) *ResourceServiceGetResourceParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the resource service get resource params
func (o *ResourceServiceGetResourceParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceName adds the resourceName to the resource service get resource params
func (o *ResourceServiceGetResourceParams) WithResourceName(resourceName *string) *ResourceServiceGetResourceParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the resource service get resource params
func (o *ResourceServiceGetResourceParams) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceServiceGetResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ResourceID != nil {

		// query param resource_id
		var qrResourceID string

		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {

			if err := r.SetQueryParam("resource_id", qResourceID); err != nil {
				return err
			}
		}
	}

	if o.ResourceName != nil {

		// query param resource_name
		var qrResourceName string

		if o.ResourceName != nil {
			qrResourceName = *o.ResourceName
		}
		qResourceName := qrResourceName
		if qResourceName != "" {

			if err := r.SetQueryParam("resource_name", qResourceName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
