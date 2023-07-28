// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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
	"github.com/go-openapi/swag"
)

// NewListOrganizationServicePrincipalsParams creates a new ListOrganizationServicePrincipalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOrganizationServicePrincipalsParams() *ListOrganizationServicePrincipalsParams {
	return &ListOrganizationServicePrincipalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOrganizationServicePrincipalsParamsWithTimeout creates a new ListOrganizationServicePrincipalsParams object
// with the ability to set a timeout on a request.
func NewListOrganizationServicePrincipalsParamsWithTimeout(timeout time.Duration) *ListOrganizationServicePrincipalsParams {
	return &ListOrganizationServicePrincipalsParams{
		timeout: timeout,
	}
}

// NewListOrganizationServicePrincipalsParamsWithContext creates a new ListOrganizationServicePrincipalsParams object
// with the ability to set a context for a request.
func NewListOrganizationServicePrincipalsParamsWithContext(ctx context.Context) *ListOrganizationServicePrincipalsParams {
	return &ListOrganizationServicePrincipalsParams{
		Context: ctx,
	}
}

// NewListOrganizationServicePrincipalsParamsWithHTTPClient creates a new ListOrganizationServicePrincipalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOrganizationServicePrincipalsParamsWithHTTPClient(client *http.Client) *ListOrganizationServicePrincipalsParams {
	return &ListOrganizationServicePrincipalsParams{
		HTTPClient: client,
	}
}

/*
ListOrganizationServicePrincipalsParams contains all the parameters to send to the API endpoint

	for the list organization service principals operation.

	Typically these are written to a http.Request.
*/
type ListOrganizationServicePrincipalsParams struct {

	/* Include.

	     include defines which service principals in the hirarchical level
	below an organization should be included in the returned list.
	Allowed values: UNSET, ALL, PROJECTS.

	 - UNSET: UNSET is the default value.
	UNSET will only include service principals at organization level.
	 - ALL: ALL will include all service principals on all hirarchical levels
	under the given organization.
	 - PROJECTS: PROJECTS includes service principals at project level
	in the list response.

	     Default: "UNSET"
	*/
	Include *string

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization that
	owns the service principals to be listed.
	*/
	OrganizationID string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list organization service principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrganizationServicePrincipalsParams) WithDefaults() *ListOrganizationServicePrincipalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list organization service principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrganizationServicePrincipalsParams) SetDefaults() {
	var (
		includeDefault = string("UNSET")
	)

	val := ListOrganizationServicePrincipalsParams{
		Include: &includeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithTimeout(timeout time.Duration) *ListOrganizationServicePrincipalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithContext(ctx context.Context) *ListOrganizationServicePrincipalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithHTTPClient(client *http.Client) *ListOrganizationServicePrincipalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithInclude(include *string) *ListOrganizationServicePrincipalsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetInclude(include *string) {
	o.Include = include
}

// WithOrganizationID adds the organizationID to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithOrganizationID(organizationID string) *ListOrganizationServicePrincipalsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListOrganizationServicePrincipalsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithPaginationPageSize(paginationPageSize *int64) *ListOrganizationServicePrincipalsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListOrganizationServicePrincipalsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list organization service principals params
func (o *ListOrganizationServicePrincipalsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrganizationServicePrincipalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Include != nil {

		// query param include
		var qrInclude string

		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {

			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
