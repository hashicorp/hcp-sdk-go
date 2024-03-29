// Code generated by go-swagger; DO NOT EDIT.

package invitations_service

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

// NewInvitationsServiceListOrganizationInvitationsParams creates a new InvitationsServiceListOrganizationInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvitationsServiceListOrganizationInvitationsParams() *InvitationsServiceListOrganizationInvitationsParams {
	return &InvitationsServiceListOrganizationInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvitationsServiceListOrganizationInvitationsParamsWithTimeout creates a new InvitationsServiceListOrganizationInvitationsParams object
// with the ability to set a timeout on a request.
func NewInvitationsServiceListOrganizationInvitationsParamsWithTimeout(timeout time.Duration) *InvitationsServiceListOrganizationInvitationsParams {
	return &InvitationsServiceListOrganizationInvitationsParams{
		timeout: timeout,
	}
}

// NewInvitationsServiceListOrganizationInvitationsParamsWithContext creates a new InvitationsServiceListOrganizationInvitationsParams object
// with the ability to set a context for a request.
func NewInvitationsServiceListOrganizationInvitationsParamsWithContext(ctx context.Context) *InvitationsServiceListOrganizationInvitationsParams {
	return &InvitationsServiceListOrganizationInvitationsParams{
		Context: ctx,
	}
}

// NewInvitationsServiceListOrganizationInvitationsParamsWithHTTPClient creates a new InvitationsServiceListOrganizationInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvitationsServiceListOrganizationInvitationsParamsWithHTTPClient(client *http.Client) *InvitationsServiceListOrganizationInvitationsParams {
	return &InvitationsServiceListOrganizationInvitationsParams{
		HTTPClient: client,
	}
}

/*
InvitationsServiceListOrganizationInvitationsParams contains all the parameters to send to the API endpoint

	for the invitations service list organization invitations operation.

	Typically these are written to a http.Request.
*/
type InvitationsServiceListOrganizationInvitationsParams struct {

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization
	for which invites are being requested.
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

	/* States.

	     States limits the organization invitations to be returned by one or more states.

	 - PENDING: PENDING is the state of an organization invitation that has been created successfully but
	has not been accepted yet.
	 - ACCEPTED: ACCEPTED is the state of an organization invitation that has been accepted.
	*/
	States []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invitations service list organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceListOrganizationInvitationsParams) WithDefaults() *InvitationsServiceListOrganizationInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invitations service list organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceListOrganizationInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithTimeout(timeout time.Duration) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithContext(ctx context.Context) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithHTTPClient(client *http.Client) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithOrganizationID(organizationID string) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithPaginationPageSize(paginationPageSize *int64) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithStates adds the states to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) WithStates(states []string) *InvitationsServiceListOrganizationInvitationsParams {
	o.SetStates(states)
	return o
}

// SetStates adds the states to the invitations service list organization invitations params
func (o *InvitationsServiceListOrganizationInvitationsParams) SetStates(states []string) {
	o.States = states
}

// WriteToRequest writes these params to a swagger request
func (o *InvitationsServiceListOrganizationInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.States != nil {

		// binding items for states
		joinedStates := o.bindParamStates(reg)

		// query array param states
		if err := r.SetQueryParam("states", joinedStates...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamInvitationsServiceListOrganizationInvitations binds the parameter states
func (o *InvitationsServiceListOrganizationInvitationsParams) bindParamStates(formats strfmt.Registry) []string {
	statesIR := o.States

	var statesIC []string
	for _, statesIIR := range statesIR { // explode []string

		statesIIV := statesIIR // string as string
		statesIC = append(statesIC, statesIIV)
	}

	// items.CollectionFormat: "multi"
	statesIS := swag.JoinByFormat(statesIC, "multi")

	return statesIS
}
