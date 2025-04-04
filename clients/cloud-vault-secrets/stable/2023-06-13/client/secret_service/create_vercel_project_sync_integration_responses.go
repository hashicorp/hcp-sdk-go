// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// CreateVercelProjectSyncIntegrationReader is a Reader for the CreateVercelProjectSyncIntegration structure.
type CreateVercelProjectSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVercelProjectSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVercelProjectSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateVercelProjectSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateVercelProjectSyncIntegrationOK creates a CreateVercelProjectSyncIntegrationOK with default headers values
func NewCreateVercelProjectSyncIntegrationOK() *CreateVercelProjectSyncIntegrationOK {
	return &CreateVercelProjectSyncIntegrationOK{}
}

/*
CreateVercelProjectSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateVercelProjectSyncIntegrationOK struct {
	Payload *models.Secrets20230613CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create vercel project sync integration o k response has a 2xx status code
func (o *CreateVercelProjectSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create vercel project sync integration o k response has a 3xx status code
func (o *CreateVercelProjectSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create vercel project sync integration o k response has a 4xx status code
func (o *CreateVercelProjectSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create vercel project sync integration o k response has a 5xx status code
func (o *CreateVercelProjectSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create vercel project sync integration o k response a status code equal to that given
func (o *CreateVercelProjectSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create vercel project sync integration o k response
func (o *CreateVercelProjectSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateVercelProjectSyncIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] createVercelProjectSyncIntegrationOK %s", 200, payload)
}

func (o *CreateVercelProjectSyncIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] createVercelProjectSyncIntegrationOK %s", 200, payload)
}

func (o *CreateVercelProjectSyncIntegrationOK) GetPayload() *models.Secrets20230613CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateVercelProjectSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVercelProjectSyncIntegrationDefault creates a CreateVercelProjectSyncIntegrationDefault with default headers values
func NewCreateVercelProjectSyncIntegrationDefault(code int) *CreateVercelProjectSyncIntegrationDefault {
	return &CreateVercelProjectSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateVercelProjectSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateVercelProjectSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create vercel project sync integration default response has a 2xx status code
func (o *CreateVercelProjectSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create vercel project sync integration default response has a 3xx status code
func (o *CreateVercelProjectSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create vercel project sync integration default response has a 4xx status code
func (o *CreateVercelProjectSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create vercel project sync integration default response has a 5xx status code
func (o *CreateVercelProjectSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create vercel project sync integration default response a status code equal to that given
func (o *CreateVercelProjectSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create vercel project sync integration default response
func (o *CreateVercelProjectSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateVercelProjectSyncIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] CreateVercelProjectSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateVercelProjectSyncIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] CreateVercelProjectSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateVercelProjectSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateVercelProjectSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateVercelProjectSyncIntegrationBody create vercel project sync integration body
swagger:model CreateVercelProjectSyncIntegrationBody
*/
type CreateVercelProjectSyncIntegrationBody struct {

	// location
	Location *CreateVercelProjectSyncIntegrationParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// vercel project connection details
	VercelProjectConnectionDetails *models.Secrets20230613VercelProjectConnectionDetailsRequest `json:"vercel_project_connection_details,omitempty"`
}

// Validate validates this create vercel project sync integration body
func (o *CreateVercelProjectSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVercelProjectConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVercelProjectSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVercelProjectSyncIntegrationBody) validateVercelProjectConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.VercelProjectConnectionDetails) { // not required
		return nil
	}

	if o.VercelProjectConnectionDetails != nil {
		if err := o.VercelProjectConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create vercel project sync integration body based on the context it is used
func (o *CreateVercelProjectSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVercelProjectConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVercelProjectSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {

		if swag.IsZero(o.Location) { // not required
			return nil
		}

		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVercelProjectSyncIntegrationBody) contextValidateVercelProjectConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.VercelProjectConnectionDetails != nil {

		if swag.IsZero(o.VercelProjectConnectionDetails) { // not required
			return nil
		}

		if err := o.VercelProjectConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateVercelProjectSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVercelProjectSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateVercelProjectSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateVercelProjectSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateVercelProjectSyncIntegrationParamsBodyLocation
*/
type CreateVercelProjectSyncIntegrationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create vercel project sync integration params body location
func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create vercel project sync integration params body location based on the context it is used
func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {

		if swag.IsZero(o.Region) { // not required
			return nil
		}

		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVercelProjectSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateVercelProjectSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
