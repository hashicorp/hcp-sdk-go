// Code generated by go-swagger; DO NOT EDIT.

package statement_service

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

// NewStatementServiceGetStatementParams creates a new StatementServiceGetStatementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatementServiceGetStatementParams() *StatementServiceGetStatementParams {
	return &StatementServiceGetStatementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatementServiceGetStatementParamsWithTimeout creates a new StatementServiceGetStatementParams object
// with the ability to set a timeout on a request.
func NewStatementServiceGetStatementParamsWithTimeout(timeout time.Duration) *StatementServiceGetStatementParams {
	return &StatementServiceGetStatementParams{
		timeout: timeout,
	}
}

// NewStatementServiceGetStatementParamsWithContext creates a new StatementServiceGetStatementParams object
// with the ability to set a context for a request.
func NewStatementServiceGetStatementParamsWithContext(ctx context.Context) *StatementServiceGetStatementParams {
	return &StatementServiceGetStatementParams{
		Context: ctx,
	}
}

// NewStatementServiceGetStatementParamsWithHTTPClient creates a new StatementServiceGetStatementParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatementServiceGetStatementParamsWithHTTPClient(client *http.Client) *StatementServiceGetStatementParams {
	return &StatementServiceGetStatementParams{
		HTTPClient: client,
	}
}

/*
StatementServiceGetStatementParams contains all the parameters to send to the API endpoint

	for the statement service get statement operation.

	Typically these are written to a http.Request.
*/
type StatementServiceGetStatementParams struct {

	/* BillingAccountID.

	     id is the user-settable ID that uniquely identifies the Billing Account
	within the organization.
	*/
	BillingAccountID string

	/* OrganizationID.

	     organization_id is the ID of the organization to which the Billing Account
	belongs.
	*/
	OrganizationID string

	/* StatementID.

	   id is the ID that uniquely identifies the monthly statement
	*/
	StatementID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the statement service get statement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatementServiceGetStatementParams) WithDefaults() *StatementServiceGetStatementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the statement service get statement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatementServiceGetStatementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithTimeout(timeout time.Duration) *StatementServiceGetStatementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithContext(ctx context.Context) *StatementServiceGetStatementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithHTTPClient(client *http.Client) *StatementServiceGetStatementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithBillingAccountID(billingAccountID string) *StatementServiceGetStatementParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithOrganizationID adds the organizationID to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithOrganizationID(organizationID string) *StatementServiceGetStatementParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithStatementID adds the statementID to the statement service get statement params
func (o *StatementServiceGetStatementParams) WithStatementID(statementID string) *StatementServiceGetStatementParams {
	o.SetStatementID(statementID)
	return o
}

// SetStatementID adds the statementId to the statement service get statement params
func (o *StatementServiceGetStatementParams) SetStatementID(statementID string) {
	o.StatementID = statementID
}

// WriteToRequest writes these params to a swagger request
func (o *StatementServiceGetStatementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billing_account_id
	if err := r.SetPathParam("billing_account_id", o.BillingAccountID); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param statement_id
	if err := r.SetPathParam("statement_id", o.StatementID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
