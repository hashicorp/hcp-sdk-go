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
)

// NewInvitationsServiceCreateOrganizationInvitationsParams creates a new InvitationsServiceCreateOrganizationInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvitationsServiceCreateOrganizationInvitationsParams() *InvitationsServiceCreateOrganizationInvitationsParams {
	return &InvitationsServiceCreateOrganizationInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvitationsServiceCreateOrganizationInvitationsParamsWithTimeout creates a new InvitationsServiceCreateOrganizationInvitationsParams object
// with the ability to set a timeout on a request.
func NewInvitationsServiceCreateOrganizationInvitationsParamsWithTimeout(timeout time.Duration) *InvitationsServiceCreateOrganizationInvitationsParams {
	return &InvitationsServiceCreateOrganizationInvitationsParams{
		timeout: timeout,
	}
}

// NewInvitationsServiceCreateOrganizationInvitationsParamsWithContext creates a new InvitationsServiceCreateOrganizationInvitationsParams object
// with the ability to set a context for a request.
func NewInvitationsServiceCreateOrganizationInvitationsParamsWithContext(ctx context.Context) *InvitationsServiceCreateOrganizationInvitationsParams {
	return &InvitationsServiceCreateOrganizationInvitationsParams{
		Context: ctx,
	}
}

// NewInvitationsServiceCreateOrganizationInvitationsParamsWithHTTPClient creates a new InvitationsServiceCreateOrganizationInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvitationsServiceCreateOrganizationInvitationsParamsWithHTTPClient(client *http.Client) *InvitationsServiceCreateOrganizationInvitationsParams {
	return &InvitationsServiceCreateOrganizationInvitationsParams{
		HTTPClient: client,
	}
}

/*
InvitationsServiceCreateOrganizationInvitationsParams contains all the parameters to send to the API endpoint

	for the invitations service create organization invitations operation.

	Typically these are written to a http.Request.
*/
type InvitationsServiceCreateOrganizationInvitationsParams struct {

	// Body.
	Body InvitationsServiceCreateOrganizationInvitationsBody

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization
	to which the invitee is being invited.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invitations service create organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithDefaults() *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invitations service create organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithTimeout(timeout time.Duration) *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithContext(ctx context.Context) *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithHTTPClient(client *http.Client) *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithBody(body InvitationsServiceCreateOrganizationInvitationsBody) *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetBody(body InvitationsServiceCreateOrganizationInvitationsBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WithOrganizationID(organizationID string) *InvitationsServiceCreateOrganizationInvitationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the invitations service create organization invitations params
func (o *InvitationsServiceCreateOrganizationInvitationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *InvitationsServiceCreateOrganizationInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
