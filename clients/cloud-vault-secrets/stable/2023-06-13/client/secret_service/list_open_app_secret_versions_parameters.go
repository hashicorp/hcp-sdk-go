// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewListOpenAppSecretVersionsParams creates a new ListOpenAppSecretVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOpenAppSecretVersionsParams() *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenAppSecretVersionsParamsWithTimeout creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a timeout on a request.
func NewListOpenAppSecretVersionsParamsWithTimeout(timeout time.Duration) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		timeout: timeout,
	}
}

// NewListOpenAppSecretVersionsParamsWithContext creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a context for a request.
func NewListOpenAppSecretVersionsParamsWithContext(ctx context.Context) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		Context: ctx,
	}
}

// NewListOpenAppSecretVersionsParamsWithHTTPClient creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOpenAppSecretVersionsParamsWithHTTPClient(client *http.Client) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		HTTPClient: client,
	}
}

/*
ListOpenAppSecretVersionsParams contains all the parameters to send to the API endpoint

	for the list open app secret versions operation.

	Typically these are written to a http.Request.
*/
type ListOpenAppSecretVersionsParams struct {

	// AppName.
	AppName string

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

	// SecretName.
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list open app secret versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenAppSecretVersionsParams) WithDefaults() *ListOpenAppSecretVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list open app secret versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenAppSecretVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithTimeout(timeout time.Duration) *ListOpenAppSecretVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithContext(ctx context.Context) *ListOpenAppSecretVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithHTTPClient(client *http.Client) *ListOpenAppSecretVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithAppName(appName string) *ListOpenAppSecretVersionsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithLocationOrganizationID adds the locationOrganizationID to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithLocationOrganizationID(locationOrganizationID string) *ListOpenAppSecretVersionsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithLocationProjectID(locationProjectID string) *ListOpenAppSecretVersionsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithLocationRegionProvider(locationRegionProvider *string) *ListOpenAppSecretVersionsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithLocationRegionRegion(locationRegionRegion *string) *ListOpenAppSecretVersionsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithSecretName adds the secretName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithSecretName(secretName string) *ListOpenAppSecretVersionsParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenAppSecretVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
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

	// path param secret_name
	if err := r.SetPathParam("secret_name", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
