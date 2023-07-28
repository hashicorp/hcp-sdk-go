// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
)

// NewVerifyDomainOwnershipParams creates a new VerifyDomainOwnershipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyDomainOwnershipParams() *VerifyDomainOwnershipParams {
	return &VerifyDomainOwnershipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyDomainOwnershipParamsWithTimeout creates a new VerifyDomainOwnershipParams object
// with the ability to set a timeout on a request.
func NewVerifyDomainOwnershipParamsWithTimeout(timeout time.Duration) *VerifyDomainOwnershipParams {
	return &VerifyDomainOwnershipParams{
		timeout: timeout,
	}
}

// NewVerifyDomainOwnershipParamsWithContext creates a new VerifyDomainOwnershipParams object
// with the ability to set a context for a request.
func NewVerifyDomainOwnershipParamsWithContext(ctx context.Context) *VerifyDomainOwnershipParams {
	return &VerifyDomainOwnershipParams{
		Context: ctx,
	}
}

// NewVerifyDomainOwnershipParamsWithHTTPClient creates a new VerifyDomainOwnershipParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyDomainOwnershipParamsWithHTTPClient(client *http.Client) *VerifyDomainOwnershipParams {
	return &VerifyDomainOwnershipParams{
		HTTPClient: client,
	}
}

/*
VerifyDomainOwnershipParams contains all the parameters to send to the API endpoint

	for the verify domain ownership operation.

	Typically these are written to a http.Request.
*/
type VerifyDomainOwnershipParams struct {

	// Body.
	Body *models.HashicorpCloudIamVerifyDomainOwnershipRequest

	/* OrganizationID.

	     organization_id is the ID of the organization for which domain ownership
	will be verified.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify domain ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyDomainOwnershipParams) WithDefaults() *VerifyDomainOwnershipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify domain ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyDomainOwnershipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) WithTimeout(timeout time.Duration) *VerifyDomainOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) WithContext(ctx context.Context) *VerifyDomainOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) WithHTTPClient(client *http.Client) *VerifyDomainOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) WithBody(body *models.HashicorpCloudIamVerifyDomainOwnershipRequest) *VerifyDomainOwnershipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) SetBody(body *models.HashicorpCloudIamVerifyDomainOwnershipRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) WithOrganizationID(organizationID string) *VerifyDomainOwnershipParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the verify domain ownership params
func (o *VerifyDomainOwnershipParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyDomainOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
