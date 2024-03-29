// Code generated by go-swagger; DO NOT EDIT.

package network_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/models"
)

// NewAssociateHVNWithAWSRoute53PrivateHostedZoneParams creates a new AssociateHVNWithAWSRoute53PrivateHostedZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociateHVNWithAWSRoute53PrivateHostedZoneParams() *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	return &AssociateHVNWithAWSRoute53PrivateHostedZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithTimeout creates a new AssociateHVNWithAWSRoute53PrivateHostedZoneParams object
// with the ability to set a timeout on a request.
func NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithTimeout(timeout time.Duration) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	return &AssociateHVNWithAWSRoute53PrivateHostedZoneParams{
		timeout: timeout,
	}
}

// NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithContext creates a new AssociateHVNWithAWSRoute53PrivateHostedZoneParams object
// with the ability to set a context for a request.
func NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithContext(ctx context.Context) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	return &AssociateHVNWithAWSRoute53PrivateHostedZoneParams{
		Context: ctx,
	}
}

// NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithHTTPClient creates a new AssociateHVNWithAWSRoute53PrivateHostedZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociateHVNWithAWSRoute53PrivateHostedZoneParamsWithHTTPClient(client *http.Client) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	return &AssociateHVNWithAWSRoute53PrivateHostedZoneParams{
		HTTPClient: client,
	}
}

/*
AssociateHVNWithAWSRoute53PrivateHostedZoneParams contains all the parameters to send to the API endpoint

	for the associate h v n with a w s route53 private hosted zone operation.

	Typically these are written to a http.Request.
*/
type AssociateHVNWithAWSRoute53PrivateHostedZoneParams struct {

	// Body.
	Body *models.HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest

	/* HvnID.

	   hvn_id is the id of the HVN the private hosted zone should be associated with.
	*/
	HvnID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associate h v n with a w s route53 private hosted zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithDefaults() *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associate h v n with a w s route53 private hosted zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithTimeout(timeout time.Duration) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithContext(ctx context.Context) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithHTTPClient(client *http.Client) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithBody(body *models.HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetBody(body *models.HashicorpCloudNetwork20200907AssociateHVNWithAWSRoute53PrivateHostedZoneRequest) {
	o.Body = body
}

// WithHvnID adds the hvnID to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithHvnID(hvnID string) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetHvnID(hvnID)
	return o
}

// SetHvnID adds the hvnId to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetHvnID(hvnID string) {
	o.HvnID = hvnID
}

// WithLocationOrganizationID adds the locationOrganizationID to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithLocationOrganizationID(locationOrganizationID string) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WithLocationProjectID(locationProjectID string) *AssociateHVNWithAWSRoute53PrivateHostedZoneParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the associate h v n with a w s route53 private hosted zone params
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *AssociateHVNWithAWSRoute53PrivateHostedZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param hvn_id
	if err := r.SetPathParam("hvn_id", o.HvnID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
