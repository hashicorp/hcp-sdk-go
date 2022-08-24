// Code generated by go-swagger; DO NOT EDIT.

package organization_service

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

// NewOrganizationServiceListRolesParams creates a new OrganizationServiceListRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationServiceListRolesParams() *OrganizationServiceListRolesParams {
	return &OrganizationServiceListRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationServiceListRolesParamsWithTimeout creates a new OrganizationServiceListRolesParams object
// with the ability to set a timeout on a request.
func NewOrganizationServiceListRolesParamsWithTimeout(timeout time.Duration) *OrganizationServiceListRolesParams {
	return &OrganizationServiceListRolesParams{
		timeout: timeout,
	}
}

// NewOrganizationServiceListRolesParamsWithContext creates a new OrganizationServiceListRolesParams object
// with the ability to set a context for a request.
func NewOrganizationServiceListRolesParamsWithContext(ctx context.Context) *OrganizationServiceListRolesParams {
	return &OrganizationServiceListRolesParams{
		Context: ctx,
	}
}

// NewOrganizationServiceListRolesParamsWithHTTPClient creates a new OrganizationServiceListRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationServiceListRolesParamsWithHTTPClient(client *http.Client) *OrganizationServiceListRolesParams {
	return &OrganizationServiceListRolesParams{
		HTTPClient: client,
	}
}

/* OrganizationServiceListRolesParams contains all the parameters to send to the API endpoint
   for the organization service list roles operation.

   Typically these are written to a http.Request.
*/
type OrganizationServiceListRolesParams struct {

	/* ID.

	   ID is the identifier of the organization.
	*/
	ID string

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

// WithDefaults hydrates default values in the organization service list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceListRolesParams) WithDefaults() *OrganizationServiceListRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization service list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceListRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithTimeout(timeout time.Duration) *OrganizationServiceListRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithContext(ctx context.Context) *OrganizationServiceListRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithHTTPClient(client *http.Client) *OrganizationServiceListRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithID(id string) *OrganizationServiceListRolesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetID(id string) {
	o.ID = id
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *OrganizationServiceListRolesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithPaginationPageSize(paginationPageSize *int64) *OrganizationServiceListRolesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the organization service list roles params
func (o *OrganizationServiceListRolesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *OrganizationServiceListRolesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the organization service list roles params
func (o *OrganizationServiceListRolesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationServiceListRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
