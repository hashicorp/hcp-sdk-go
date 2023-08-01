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
)

// NewOrganizationServiceGetRoleParams creates a new OrganizationServiceGetRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationServiceGetRoleParams() *OrganizationServiceGetRoleParams {
	return &OrganizationServiceGetRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationServiceGetRoleParamsWithTimeout creates a new OrganizationServiceGetRoleParams object
// with the ability to set a timeout on a request.
func NewOrganizationServiceGetRoleParamsWithTimeout(timeout time.Duration) *OrganizationServiceGetRoleParams {
	return &OrganizationServiceGetRoleParams{
		timeout: timeout,
	}
}

// NewOrganizationServiceGetRoleParamsWithContext creates a new OrganizationServiceGetRoleParams object
// with the ability to set a context for a request.
func NewOrganizationServiceGetRoleParamsWithContext(ctx context.Context) *OrganizationServiceGetRoleParams {
	return &OrganizationServiceGetRoleParams{
		Context: ctx,
	}
}

// NewOrganizationServiceGetRoleParamsWithHTTPClient creates a new OrganizationServiceGetRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationServiceGetRoleParamsWithHTTPClient(client *http.Client) *OrganizationServiceGetRoleParams {
	return &OrganizationServiceGetRoleParams{
		HTTPClient: client,
	}
}

/*
OrganizationServiceGetRoleParams contains all the parameters to send to the API endpoint

	for the organization service get role operation.

	Typically these are written to a http.Request.
*/
type OrganizationServiceGetRoleParams struct {

	/* ID.

	   ID is the identifier of the organization.
	*/
	ID string

	/* RoleID.

	   RoleId is the identifier of the role to retrieve.
	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization service get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceGetRoleParams) WithDefaults() *OrganizationServiceGetRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization service get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceGetRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization service get role params
func (o *OrganizationServiceGetRoleParams) WithTimeout(timeout time.Duration) *OrganizationServiceGetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization service get role params
func (o *OrganizationServiceGetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization service get role params
func (o *OrganizationServiceGetRoleParams) WithContext(ctx context.Context) *OrganizationServiceGetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization service get role params
func (o *OrganizationServiceGetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization service get role params
func (o *OrganizationServiceGetRoleParams) WithHTTPClient(client *http.Client) *OrganizationServiceGetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization service get role params
func (o *OrganizationServiceGetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organization service get role params
func (o *OrganizationServiceGetRoleParams) WithID(id string) *OrganizationServiceGetRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organization service get role params
func (o *OrganizationServiceGetRoleParams) SetID(id string) {
	o.ID = id
}

// WithRoleID adds the roleID to the organization service get role params
func (o *OrganizationServiceGetRoleParams) WithRoleID(roleID string) *OrganizationServiceGetRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the organization service get role params
func (o *OrganizationServiceGetRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationServiceGetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
