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

	"github.com/hashicorp/cloud-sdk-go/clients/cloud-network/preview/2020-09-07/models"
)

// NewCreateParams creates a new CreateParams object
// with the default values initialized.
func NewCreateParams() *CreateParams {
	var ()
	return &CreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateParamsWithTimeout creates a new CreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateParamsWithTimeout(timeout time.Duration) *CreateParams {
	var ()
	return &CreateParams{

		timeout: timeout,
	}
}

// NewCreateParamsWithContext creates a new CreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateParamsWithContext(ctx context.Context) *CreateParams {
	var ()
	return &CreateParams{

		Context: ctx,
	}
}

// NewCreateParamsWithHTTPClient creates a new CreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateParamsWithHTTPClient(client *http.Client) *CreateParams {
	var ()
	return &CreateParams{
		HTTPClient: client,
	}
}

/*CreateParams contains all the parameters to send to the API endpoint
for the create operation typically these are written to a http.Request
*/
type CreateParams struct {

	/*Body*/
	Body *models.HashicorpCloudNetwork20200907CreateRequest
	/*NetworkLocationOrganizationID
	  organization_id is the id of the organization.

	*/
	NetworkLocationOrganizationID string
	/*NetworkLocationProjectID
	  project_id is the projects id.

	*/
	NetworkLocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create params
func (o *CreateParams) WithTimeout(timeout time.Duration) *CreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create params
func (o *CreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create params
func (o *CreateParams) WithContext(ctx context.Context) *CreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create params
func (o *CreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) WithHTTPClient(client *http.Client) *CreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create params
func (o *CreateParams) WithBody(body *models.HashicorpCloudNetwork20200907CreateRequest) *CreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create params
func (o *CreateParams) SetBody(body *models.HashicorpCloudNetwork20200907CreateRequest) {
	o.Body = body
}

// WithNetworkLocationOrganizationID adds the networkLocationOrganizationID to the create params
func (o *CreateParams) WithNetworkLocationOrganizationID(networkLocationOrganizationID string) *CreateParams {
	o.SetNetworkLocationOrganizationID(networkLocationOrganizationID)
	return o
}

// SetNetworkLocationOrganizationID adds the networkLocationOrganizationId to the create params
func (o *CreateParams) SetNetworkLocationOrganizationID(networkLocationOrganizationID string) {
	o.NetworkLocationOrganizationID = networkLocationOrganizationID
}

// WithNetworkLocationProjectID adds the networkLocationProjectID to the create params
func (o *CreateParams) WithNetworkLocationProjectID(networkLocationProjectID string) *CreateParams {
	o.SetNetworkLocationProjectID(networkLocationProjectID)
	return o
}

// SetNetworkLocationProjectID adds the networkLocationProjectId to the create params
func (o *CreateParams) SetNetworkLocationProjectID(networkLocationProjectID string) {
	o.NetworkLocationProjectID = networkLocationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network.location.organization_id
	if err := r.SetPathParam("network.location.organization_id", o.NetworkLocationOrganizationID); err != nil {
		return err
	}

	// path param network.location.project_id
	if err := r.SetPathParam("network.location.project_id", o.NetworkLocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
