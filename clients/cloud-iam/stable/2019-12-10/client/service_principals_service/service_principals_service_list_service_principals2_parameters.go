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

// NewServicePrincipalsServiceListServicePrincipals2Params creates a new ServicePrincipalsServiceListServicePrincipals2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceListServicePrincipals2Params() *ServicePrincipalsServiceListServicePrincipals2Params {
	return &ServicePrincipalsServiceListServicePrincipals2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceListServicePrincipals2ParamsWithTimeout creates a new ServicePrincipalsServiceListServicePrincipals2Params object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceListServicePrincipals2ParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceListServicePrincipals2Params {
	return &ServicePrincipalsServiceListServicePrincipals2Params{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceListServicePrincipals2ParamsWithContext creates a new ServicePrincipalsServiceListServicePrincipals2Params object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceListServicePrincipals2ParamsWithContext(ctx context.Context) *ServicePrincipalsServiceListServicePrincipals2Params {
	return &ServicePrincipalsServiceListServicePrincipals2Params{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceListServicePrincipals2ParamsWithHTTPClient creates a new ServicePrincipalsServiceListServicePrincipals2Params object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceListServicePrincipals2ParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceListServicePrincipals2Params {
	return &ServicePrincipalsServiceListServicePrincipals2Params{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceListServicePrincipals2Params contains all the parameters to send to the API endpoint

	for the service principals service list service principals2 operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceListServicePrincipals2Params struct {

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

	/* ParentResourceName1.

	     parent_resource_name is the resource name of the project or organization
	at which the service principal should be listed (e.g.
	"project/<project_id>" or "organization/<organization_id>")
	*/
	ParentResourceName1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service list service principals2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithDefaults() *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service list service principals2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithContext(ctx context.Context) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithPaginationNextPageToken(paginationNextPageToken *string) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithPaginationPageSize(paginationPageSize *int64) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithParentResourceName1 adds the parentResourceName1 to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WithParentResourceName1(parentResourceName1 string) *ServicePrincipalsServiceListServicePrincipals2Params {
	o.SetParentResourceName1(parentResourceName1)
	return o
}

// SetParentResourceName1 adds the parentResourceName1 to the service principals service list service principals2 params
func (o *ServicePrincipalsServiceListServicePrincipals2Params) SetParentResourceName1(parentResourceName1 string) {
	o.ParentResourceName1 = parentResourceName1
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceListServicePrincipals2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param parent_resource_name_1
	if err := r.SetPathParam("parent_resource_name_1", o.ParentResourceName1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
