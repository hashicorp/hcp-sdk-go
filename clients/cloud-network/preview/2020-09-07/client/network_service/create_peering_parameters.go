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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/models"
)

// NewCreatePeeringParams creates a new CreatePeeringParams object
// with the default values initialized.
func NewCreatePeeringParams() *CreatePeeringParams {
	var ()
	return &CreatePeeringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePeeringParamsWithTimeout creates a new CreatePeeringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePeeringParamsWithTimeout(timeout time.Duration) *CreatePeeringParams {
	var ()
	return &CreatePeeringParams{

		timeout: timeout,
	}
}

// NewCreatePeeringParamsWithContext creates a new CreatePeeringParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePeeringParamsWithContext(ctx context.Context) *CreatePeeringParams {
	var ()
	return &CreatePeeringParams{

		Context: ctx,
	}
}

// NewCreatePeeringParamsWithHTTPClient creates a new CreatePeeringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePeeringParamsWithHTTPClient(client *http.Client) *CreatePeeringParams {
	var ()
	return &CreatePeeringParams{
		HTTPClient: client,
	}
}

/*CreatePeeringParams contains all the parameters to send to the API endpoint
for the create peering operation typically these are written to a http.Request
*/
type CreatePeeringParams struct {

	/*Body*/
	Body *models.HashicorpCloudNetwork20200907CreatePeeringRequest
	/*PeeringHvnID
	  id is the identifier for this resource.

	*/
	PeeringHvnID string
	/*PeeringHvnLocationOrganizationID
	  organization_id is the id of the organization.

	*/
	PeeringHvnLocationOrganizationID string
	/*PeeringHvnLocationProjectID
	  project_id is the projects id.

	*/
	PeeringHvnLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create peering params
func (o *CreatePeeringParams) WithTimeout(timeout time.Duration) *CreatePeeringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create peering params
func (o *CreatePeeringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create peering params
func (o *CreatePeeringParams) WithContext(ctx context.Context) *CreatePeeringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create peering params
func (o *CreatePeeringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create peering params
func (o *CreatePeeringParams) WithHTTPClient(client *http.Client) *CreatePeeringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create peering params
func (o *CreatePeeringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create peering params
func (o *CreatePeeringParams) WithBody(body *models.HashicorpCloudNetwork20200907CreatePeeringRequest) *CreatePeeringParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create peering params
func (o *CreatePeeringParams) SetBody(body *models.HashicorpCloudNetwork20200907CreatePeeringRequest) {
	o.Body = body
}

// WithPeeringHvnID adds the peeringHvnID to the create peering params
func (o *CreatePeeringParams) WithPeeringHvnID(peeringHvnID string) *CreatePeeringParams {
	o.SetPeeringHvnID(peeringHvnID)
	return o
}

// SetPeeringHvnID adds the peeringHvnId to the create peering params
func (o *CreatePeeringParams) SetPeeringHvnID(peeringHvnID string) {
	o.PeeringHvnID = peeringHvnID
}

// WithPeeringHvnLocationOrganizationID adds the peeringHvnLocationOrganizationID to the create peering params
func (o *CreatePeeringParams) WithPeeringHvnLocationOrganizationID(peeringHvnLocationOrganizationID string) *CreatePeeringParams {
	o.SetPeeringHvnLocationOrganizationID(peeringHvnLocationOrganizationID)
	return o
}

// SetPeeringHvnLocationOrganizationID adds the peeringHvnLocationOrganizationId to the create peering params
func (o *CreatePeeringParams) SetPeeringHvnLocationOrganizationID(peeringHvnLocationOrganizationID string) {
	o.PeeringHvnLocationOrganizationID = peeringHvnLocationOrganizationID
}

// WithPeeringHvnLocationProjectID adds the peeringHvnLocationProjectID to the create peering params
func (o *CreatePeeringParams) WithPeeringHvnLocationProjectID(peeringHvnLocationProjectID string) *CreatePeeringParams {
	o.SetPeeringHvnLocationProjectID(peeringHvnLocationProjectID)
	return o
}

// SetPeeringHvnLocationProjectID adds the peeringHvnLocationProjectId to the create peering params
func (o *CreatePeeringParams) SetPeeringHvnLocationProjectID(peeringHvnLocationProjectID string) {
	o.PeeringHvnLocationProjectID = peeringHvnLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePeeringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param peering.hvn.id
	if err := r.SetPathParam("peering.hvn.id", o.PeeringHvnID); err != nil {
		return err
	}

	// path param peering.hvn.location.organization_id
	if err := r.SetPathParam("peering.hvn.location.organization_id", o.PeeringHvnLocationOrganizationID); err != nil {
		return err
	}

	// path param peering.hvn.location.project_id
	if err := r.SetPathParam("peering.hvn.location.project_id", o.PeeringHvnLocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
