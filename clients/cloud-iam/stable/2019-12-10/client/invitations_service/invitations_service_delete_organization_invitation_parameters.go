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

// NewInvitationsServiceDeleteOrganizationInvitationParams creates a new InvitationsServiceDeleteOrganizationInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvitationsServiceDeleteOrganizationInvitationParams() *InvitationsServiceDeleteOrganizationInvitationParams {
	return &InvitationsServiceDeleteOrganizationInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvitationsServiceDeleteOrganizationInvitationParamsWithTimeout creates a new InvitationsServiceDeleteOrganizationInvitationParams object
// with the ability to set a timeout on a request.
func NewInvitationsServiceDeleteOrganizationInvitationParamsWithTimeout(timeout time.Duration) *InvitationsServiceDeleteOrganizationInvitationParams {
	return &InvitationsServiceDeleteOrganizationInvitationParams{
		timeout: timeout,
	}
}

// NewInvitationsServiceDeleteOrganizationInvitationParamsWithContext creates a new InvitationsServiceDeleteOrganizationInvitationParams object
// with the ability to set a context for a request.
func NewInvitationsServiceDeleteOrganizationInvitationParamsWithContext(ctx context.Context) *InvitationsServiceDeleteOrganizationInvitationParams {
	return &InvitationsServiceDeleteOrganizationInvitationParams{
		Context: ctx,
	}
}

// NewInvitationsServiceDeleteOrganizationInvitationParamsWithHTTPClient creates a new InvitationsServiceDeleteOrganizationInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvitationsServiceDeleteOrganizationInvitationParamsWithHTTPClient(client *http.Client) *InvitationsServiceDeleteOrganizationInvitationParams {
	return &InvitationsServiceDeleteOrganizationInvitationParams{
		HTTPClient: client,
	}
}

/*
InvitationsServiceDeleteOrganizationInvitationParams contains all the parameters to send to the API endpoint

	for the invitations service delete organization invitation operation.

	Typically these are written to a http.Request.
*/
type InvitationsServiceDeleteOrganizationInvitationParams struct {

	/* InvitationID.

	   invitation_id is the id of the invitation to be deleted.
	*/
	InvitationID string

	/* OrganizationID.

	     organization_id is the unique identifier (UUID) of the organization
	for which the invitation is being deleted.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invitations service delete organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithDefaults() *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invitations service delete organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithTimeout(timeout time.Duration) *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithContext(ctx context.Context) *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithHTTPClient(client *http.Client) *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationID adds the invitationID to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithInvitationID(invitationID string) *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetInvitationID(invitationID)
	return o
}

// SetInvitationID adds the invitationId to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetInvitationID(invitationID string) {
	o.InvitationID = invitationID
}

// WithOrganizationID adds the organizationID to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WithOrganizationID(organizationID string) *InvitationsServiceDeleteOrganizationInvitationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the invitations service delete organization invitation params
func (o *InvitationsServiceDeleteOrganizationInvitationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *InvitationsServiceDeleteOrganizationInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation_id
	if err := r.SetPathParam("invitation_id", o.InvitationID); err != nil {
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
