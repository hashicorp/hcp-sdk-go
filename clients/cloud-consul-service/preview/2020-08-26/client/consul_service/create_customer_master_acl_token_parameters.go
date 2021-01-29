// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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

	"github.com/hashicorp/cloud-sdk-go/clients/cloud-consul-service/preview/2020-08-26/models"
)

// NewCreateCustomerMasterACLTokenParams creates a new CreateCustomerMasterACLTokenParams object
// with the default values initialized.
func NewCreateCustomerMasterACLTokenParams() *CreateCustomerMasterACLTokenParams {
	var ()
	return &CreateCustomerMasterACLTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomerMasterACLTokenParamsWithTimeout creates a new CreateCustomerMasterACLTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomerMasterACLTokenParamsWithTimeout(timeout time.Duration) *CreateCustomerMasterACLTokenParams {
	var ()
	return &CreateCustomerMasterACLTokenParams{

		timeout: timeout,
	}
}

// NewCreateCustomerMasterACLTokenParamsWithContext creates a new CreateCustomerMasterACLTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomerMasterACLTokenParamsWithContext(ctx context.Context) *CreateCustomerMasterACLTokenParams {
	var ()
	return &CreateCustomerMasterACLTokenParams{

		Context: ctx,
	}
}

// NewCreateCustomerMasterACLTokenParamsWithHTTPClient creates a new CreateCustomerMasterACLTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomerMasterACLTokenParamsWithHTTPClient(client *http.Client) *CreateCustomerMasterACLTokenParams {
	var ()
	return &CreateCustomerMasterACLTokenParams{
		HTTPClient: client,
	}
}

/*CreateCustomerMasterACLTokenParams contains all the parameters to send to the API endpoint
for the create customer master ACL token operation typically these are written to a http.Request
*/
type CreateCustomerMasterACLTokenParams struct {

	/*Body*/
	Body *models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenRequest
	/*ID
	  id is the ID of the cluster to create a token on.

	*/
	ID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithTimeout(timeout time.Duration) *CreateCustomerMasterACLTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithContext(ctx context.Context) *CreateCustomerMasterACLTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithHTTPClient(client *http.Client) *CreateCustomerMasterACLTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithBody(body *models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenRequest) *CreateCustomerMasterACLTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetBody(body *models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenRequest) {
	o.Body = body
}

// WithID adds the id to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithID(id string) *CreateCustomerMasterACLTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithLocationOrganizationID(locationOrganizationID string) *CreateCustomerMasterACLTokenParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) WithLocationProjectID(locationProjectID string) *CreateCustomerMasterACLTokenParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the create customer master ACL token params
func (o *CreateCustomerMasterACLTokenParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomerMasterACLTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
