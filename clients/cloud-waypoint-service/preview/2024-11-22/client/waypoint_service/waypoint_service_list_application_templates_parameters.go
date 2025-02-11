// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

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

// NewWaypointServiceListApplicationTemplatesParams creates a new WaypointServiceListApplicationTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceListApplicationTemplatesParams() *WaypointServiceListApplicationTemplatesParams {
	return &WaypointServiceListApplicationTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceListApplicationTemplatesParamsWithTimeout creates a new WaypointServiceListApplicationTemplatesParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceListApplicationTemplatesParamsWithTimeout(timeout time.Duration) *WaypointServiceListApplicationTemplatesParams {
	return &WaypointServiceListApplicationTemplatesParams{
		timeout: timeout,
	}
}

// NewWaypointServiceListApplicationTemplatesParamsWithContext creates a new WaypointServiceListApplicationTemplatesParams object
// with the ability to set a context for a request.
func NewWaypointServiceListApplicationTemplatesParamsWithContext(ctx context.Context) *WaypointServiceListApplicationTemplatesParams {
	return &WaypointServiceListApplicationTemplatesParams{
		Context: ctx,
	}
}

// NewWaypointServiceListApplicationTemplatesParamsWithHTTPClient creates a new WaypointServiceListApplicationTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceListApplicationTemplatesParamsWithHTTPClient(client *http.Client) *WaypointServiceListApplicationTemplatesParams {
	return &WaypointServiceListApplicationTemplatesParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceListApplicationTemplatesParams contains all the parameters to send to the API endpoint

	for the waypoint service list application templates operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceListApplicationTemplatesParams struct {

	// NamespaceID.
	NamespaceID *string

	/* NamespaceLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	NamespaceLocationOrganizationID string

	/* NamespaceLocationProjectID.

	   project_id is the projects id.
	*/
	NamespaceLocationProjectID string

	/* NamespaceLocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	NamespaceLocationRegionProvider *string

	/* NamespaceLocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	NamespaceLocationRegionRegion *string

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

	/* WithTotalCount.

	   If set to false or not provided, response will not include a total_count value
	*/
	WithTotalCount *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service list application templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListApplicationTemplatesParams) WithDefaults() *WaypointServiceListApplicationTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service list application templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceListApplicationTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithTimeout(timeout time.Duration) *WaypointServiceListApplicationTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithContext(ctx context.Context) *WaypointServiceListApplicationTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithHTTPClient(client *http.Client) *WaypointServiceListApplicationTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithNamespaceID(namespaceID *string) *WaypointServiceListApplicationTemplatesParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetNamespaceID(namespaceID *string) {
	o.NamespaceID = namespaceID
}

// WithNamespaceLocationOrganizationID adds the namespaceLocationOrganizationID to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) *WaypointServiceListApplicationTemplatesParams {
	o.SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID)
	return o
}

// SetNamespaceLocationOrganizationID adds the namespaceLocationOrganizationId to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetNamespaceLocationOrganizationID(namespaceLocationOrganizationID string) {
	o.NamespaceLocationOrganizationID = namespaceLocationOrganizationID
}

// WithNamespaceLocationProjectID adds the namespaceLocationProjectID to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithNamespaceLocationProjectID(namespaceLocationProjectID string) *WaypointServiceListApplicationTemplatesParams {
	o.SetNamespaceLocationProjectID(namespaceLocationProjectID)
	return o
}

// SetNamespaceLocationProjectID adds the namespaceLocationProjectId to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetNamespaceLocationProjectID(namespaceLocationProjectID string) {
	o.NamespaceLocationProjectID = namespaceLocationProjectID
}

// WithNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) *WaypointServiceListApplicationTemplatesParams {
	o.SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider)
	return o
}

// SetNamespaceLocationRegionProvider adds the namespaceLocationRegionProvider to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetNamespaceLocationRegionProvider(namespaceLocationRegionProvider *string) {
	o.NamespaceLocationRegionProvider = namespaceLocationRegionProvider
}

// WithNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) *WaypointServiceListApplicationTemplatesParams {
	o.SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion)
	return o
}

// SetNamespaceLocationRegionRegion adds the namespaceLocationRegionRegion to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetNamespaceLocationRegionRegion(namespaceLocationRegionRegion *string) {
	o.NamespaceLocationRegionRegion = namespaceLocationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointServiceListApplicationTemplatesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithPaginationPageSize(paginationPageSize *int64) *WaypointServiceListApplicationTemplatesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointServiceListApplicationTemplatesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithWithTotalCount adds the withTotalCount to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) WithWithTotalCount(withTotalCount *bool) *WaypointServiceListApplicationTemplatesParams {
	o.SetWithTotalCount(withTotalCount)
	return o
}

// SetWithTotalCount adds the withTotalCount to the waypoint service list application templates params
func (o *WaypointServiceListApplicationTemplatesParams) SetWithTotalCount(withTotalCount *bool) {
	o.WithTotalCount = withTotalCount
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceListApplicationTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NamespaceID != nil {

		// query param namespace.id
		var qrNamespaceID string

		if o.NamespaceID != nil {
			qrNamespaceID = *o.NamespaceID
		}
		qNamespaceID := qrNamespaceID
		if qNamespaceID != "" {

			if err := r.SetQueryParam("namespace.id", qNamespaceID); err != nil {
				return err
			}
		}
	}

	// path param namespace.location.organization_id
	if err := r.SetPathParam("namespace.location.organization_id", o.NamespaceLocationOrganizationID); err != nil {
		return err
	}

	// path param namespace.location.project_id
	if err := r.SetPathParam("namespace.location.project_id", o.NamespaceLocationProjectID); err != nil {
		return err
	}

	if o.NamespaceLocationRegionProvider != nil {

		// query param namespace.location.region.provider
		var qrNamespaceLocationRegionProvider string

		if o.NamespaceLocationRegionProvider != nil {
			qrNamespaceLocationRegionProvider = *o.NamespaceLocationRegionProvider
		}
		qNamespaceLocationRegionProvider := qrNamespaceLocationRegionProvider
		if qNamespaceLocationRegionProvider != "" {

			if err := r.SetQueryParam("namespace.location.region.provider", qNamespaceLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.NamespaceLocationRegionRegion != nil {

		// query param namespace.location.region.region
		var qrNamespaceLocationRegionRegion string

		if o.NamespaceLocationRegionRegion != nil {
			qrNamespaceLocationRegionRegion = *o.NamespaceLocationRegionRegion
		}
		qNamespaceLocationRegionRegion := qrNamespaceLocationRegionRegion
		if qNamespaceLocationRegionRegion != "" {

			if err := r.SetQueryParam("namespace.location.region.region", qNamespaceLocationRegionRegion); err != nil {
				return err
			}
		}
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

	if o.WithTotalCount != nil {

		// query param with_total_count
		var qrWithTotalCount bool

		if o.WithTotalCount != nil {
			qrWithTotalCount = *o.WithTotalCount
		}
		qWithTotalCount := swag.FormatBool(qrWithTotalCount)
		if qWithTotalCount != "" {

			if err := r.SetQueryParam("with_total_count", qWithTotalCount); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
