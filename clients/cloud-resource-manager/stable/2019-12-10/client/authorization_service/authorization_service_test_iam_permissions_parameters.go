// Code generated by go-swagger; DO NOT EDIT.

package authorization_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
)

// NewAuthorizationServiceTestIamPermissionsParams creates a new AuthorizationServiceTestIamPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthorizationServiceTestIamPermissionsParams() *AuthorizationServiceTestIamPermissionsParams {
	return &AuthorizationServiceTestIamPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizationServiceTestIamPermissionsParamsWithTimeout creates a new AuthorizationServiceTestIamPermissionsParams object
// with the ability to set a timeout on a request.
func NewAuthorizationServiceTestIamPermissionsParamsWithTimeout(timeout time.Duration) *AuthorizationServiceTestIamPermissionsParams {
	return &AuthorizationServiceTestIamPermissionsParams{
		timeout: timeout,
	}
}

// NewAuthorizationServiceTestIamPermissionsParamsWithContext creates a new AuthorizationServiceTestIamPermissionsParams object
// with the ability to set a context for a request.
func NewAuthorizationServiceTestIamPermissionsParamsWithContext(ctx context.Context) *AuthorizationServiceTestIamPermissionsParams {
	return &AuthorizationServiceTestIamPermissionsParams{
		Context: ctx,
	}
}

// NewAuthorizationServiceTestIamPermissionsParamsWithHTTPClient creates a new AuthorizationServiceTestIamPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthorizationServiceTestIamPermissionsParamsWithHTTPClient(client *http.Client) *AuthorizationServiceTestIamPermissionsParams {
	return &AuthorizationServiceTestIamPermissionsParams{
		HTTPClient: client,
	}
}

/*
AuthorizationServiceTestIamPermissionsParams contains all the parameters to send to the API endpoint

	for the authorization service test iam permissions operation.

	Typically these are written to a http.Request.
*/
type AuthorizationServiceTestIamPermissionsParams struct {

	/* Body.

	   AuthorizationTestIamPermissionsRequest is a request to test a set of permissions.
	*/
	Body *models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authorization service test iam permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizationServiceTestIamPermissionsParams) WithDefaults() *AuthorizationServiceTestIamPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authorization service test iam permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizationServiceTestIamPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) WithTimeout(timeout time.Duration) *AuthorizationServiceTestIamPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) WithContext(ctx context.Context) *AuthorizationServiceTestIamPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) WithHTTPClient(client *http.Client) *AuthorizationServiceTestIamPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) WithBody(body *models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsRequest) *AuthorizationServiceTestIamPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authorization service test iam permissions params
func (o *AuthorizationServiceTestIamPermissionsParams) SetBody(body *models.HashicorpCloudResourcemanagerAuthorizationTestIamPermissionsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizationServiceTestIamPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
