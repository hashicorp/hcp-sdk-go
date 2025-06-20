// Code generated by go-swagger; DO NOT EDIT.

package billing_account_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/stable/2020-11-05/models"
)

// NewBillingAccountServiceUpdateParams creates a new BillingAccountServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingAccountServiceUpdateParams() *BillingAccountServiceUpdateParams {
	return &BillingAccountServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingAccountServiceUpdateParamsWithTimeout creates a new BillingAccountServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewBillingAccountServiceUpdateParamsWithTimeout(timeout time.Duration) *BillingAccountServiceUpdateParams {
	return &BillingAccountServiceUpdateParams{
		timeout: timeout,
	}
}

// NewBillingAccountServiceUpdateParamsWithContext creates a new BillingAccountServiceUpdateParams object
// with the ability to set a context for a request.
func NewBillingAccountServiceUpdateParamsWithContext(ctx context.Context) *BillingAccountServiceUpdateParams {
	return &BillingAccountServiceUpdateParams{
		Context: ctx,
	}
}

// NewBillingAccountServiceUpdateParamsWithHTTPClient creates a new BillingAccountServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingAccountServiceUpdateParamsWithHTTPClient(client *http.Client) *BillingAccountServiceUpdateParams {
	return &BillingAccountServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
BillingAccountServiceUpdateParams contains all the parameters to send to the API endpoint

	for the billing account service update operation.

	Typically these are written to a http.Request.
*/
type BillingAccountServiceUpdateParams struct {

	// Body.
	Body *models.BillingAccountServiceUpdateBody

	/* ID.

	     id is the user-settable ID that uniquely identifies the Billing Account
	within the organization.
	*/
	ID string

	/* OrganizationID.

	     organization_id is the ID of the organization to which the Billing Account
	belongs.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing account service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingAccountServiceUpdateParams) WithDefaults() *BillingAccountServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing account service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingAccountServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithTimeout(timeout time.Duration) *BillingAccountServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithContext(ctx context.Context) *BillingAccountServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithHTTPClient(client *http.Client) *BillingAccountServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithBody(body *models.BillingAccountServiceUpdateBody) *BillingAccountServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetBody(body *models.BillingAccountServiceUpdateBody) {
	o.Body = body
}

// WithID adds the id to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithID(id string) *BillingAccountServiceUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the billing account service update params
func (o *BillingAccountServiceUpdateParams) WithOrganizationID(organizationID string) *BillingAccountServiceUpdateParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the billing account service update params
func (o *BillingAccountServiceUpdateParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *BillingAccountServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
