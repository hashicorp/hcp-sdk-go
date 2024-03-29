// Code generated by go-swagger; DO NOT EDIT.

package webhook_service

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

// NewWebhookServiceDeleteWebhookParams creates a new WebhookServiceDeleteWebhookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebhookServiceDeleteWebhookParams() *WebhookServiceDeleteWebhookParams {
	return &WebhookServiceDeleteWebhookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebhookServiceDeleteWebhookParamsWithTimeout creates a new WebhookServiceDeleteWebhookParams object
// with the ability to set a timeout on a request.
func NewWebhookServiceDeleteWebhookParamsWithTimeout(timeout time.Duration) *WebhookServiceDeleteWebhookParams {
	return &WebhookServiceDeleteWebhookParams{
		timeout: timeout,
	}
}

// NewWebhookServiceDeleteWebhookParamsWithContext creates a new WebhookServiceDeleteWebhookParams object
// with the ability to set a context for a request.
func NewWebhookServiceDeleteWebhookParamsWithContext(ctx context.Context) *WebhookServiceDeleteWebhookParams {
	return &WebhookServiceDeleteWebhookParams{
		Context: ctx,
	}
}

// NewWebhookServiceDeleteWebhookParamsWithHTTPClient creates a new WebhookServiceDeleteWebhookParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebhookServiceDeleteWebhookParamsWithHTTPClient(client *http.Client) *WebhookServiceDeleteWebhookParams {
	return &WebhookServiceDeleteWebhookParams{
		HTTPClient: client,
	}
}

/*
WebhookServiceDeleteWebhookParams contains all the parameters to send to the API endpoint

	for the webhook service delete webhook operation.

	Typically these are written to a http.Request.
*/
type WebhookServiceDeleteWebhookParams struct {

	/* ResourceName.

	     The webhook resource name.
	A webhook's resource name format name is `webhook/project/<project-id>/geo/<geography>/webhook/<webhook-name>`.
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webhook service delete webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookServiceDeleteWebhookParams) WithDefaults() *WebhookServiceDeleteWebhookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webhook service delete webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookServiceDeleteWebhookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) WithTimeout(timeout time.Duration) *WebhookServiceDeleteWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) WithContext(ctx context.Context) *WebhookServiceDeleteWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) WithHTTPClient(client *http.Client) *WebhookServiceDeleteWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) WithResourceName(resourceName string) *WebhookServiceDeleteWebhookParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the webhook service delete webhook params
func (o *WebhookServiceDeleteWebhookParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *WebhookServiceDeleteWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name
	if err := r.SetPathParam("resource_name", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
