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

// NewGetGitHubEnvironmentsParams creates a new GetGitHubEnvironmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGitHubEnvironmentsParams() *GetGitHubEnvironmentsParams {
	return &GetGitHubEnvironmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitHubEnvironmentsParamsWithTimeout creates a new GetGitHubEnvironmentsParams object
// with the ability to set a timeout on a request.
func NewGetGitHubEnvironmentsParamsWithTimeout(timeout time.Duration) *GetGitHubEnvironmentsParams {
	return &GetGitHubEnvironmentsParams{
		timeout: timeout,
	}
}

// NewGetGitHubEnvironmentsParamsWithContext creates a new GetGitHubEnvironmentsParams object
// with the ability to set a context for a request.
func NewGetGitHubEnvironmentsParamsWithContext(ctx context.Context) *GetGitHubEnvironmentsParams {
	return &GetGitHubEnvironmentsParams{
		Context: ctx,
	}
}

// NewGetGitHubEnvironmentsParamsWithHTTPClient creates a new GetGitHubEnvironmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGitHubEnvironmentsParamsWithHTTPClient(client *http.Client) *GetGitHubEnvironmentsParams {
	return &GetGitHubEnvironmentsParams{
		HTTPClient: client,
	}
}

/*
GetGitHubEnvironmentsParams contains all the parameters to send to the API endpoint

	for the get git hub environments operation.

	Typically these are written to a http.Request.
*/
type GetGitHubEnvironmentsParams struct {

	// InstallationName.
	InstallationName *string

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

	// Repository.
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get git hub environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubEnvironmentsParams) WithDefaults() *GetGitHubEnvironmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get git hub environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubEnvironmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithTimeout(timeout time.Duration) *GetGitHubEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithContext(ctx context.Context) *GetGitHubEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithHTTPClient(client *http.Client) *GetGitHubEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallationName adds the installationName to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithInstallationName(installationName *string) *GetGitHubEnvironmentsParams {
	o.SetInstallationName(installationName)
	return o
}

// SetInstallationName adds the installationName to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetInstallationName(installationName *string) {
	o.InstallationName = installationName
}

// WithLocationOrganizationID adds the locationOrganizationID to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithLocationOrganizationID(locationOrganizationID string) *GetGitHubEnvironmentsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithLocationProjectID(locationProjectID string) *GetGitHubEnvironmentsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithLocationRegionProvider(locationRegionProvider *string) *GetGitHubEnvironmentsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithLocationRegionRegion(locationRegionRegion *string) *GetGitHubEnvironmentsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithRepository adds the repository to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) WithRepository(repository string) *GetGitHubEnvironmentsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get git hub environments params
func (o *GetGitHubEnvironmentsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitHubEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InstallationName != nil {

		// query param installation_name
		var qrInstallationName string

		if o.InstallationName != nil {
			qrInstallationName = *o.InstallationName
		}
		qInstallationName := qrInstallationName
		if qInstallationName != "" {

			if err := r.SetQueryParam("installation_name", qInstallationName); err != nil {
				return err
			}
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

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
