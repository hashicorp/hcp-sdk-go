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

// CreateAwsSmSyncIntegrationReader is a Reader for the CreateAwsSmSyncIntegration structure.
type CreateAwsSmSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAwsSmSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAwsSmSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAwsSmSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAwsSmSyncIntegrationOK creates a CreateAwsSmSyncIntegrationOK with default headers values
func NewCreateAwsSmSyncIntegrationOK() *CreateAwsSmSyncIntegrationOK {
	return &CreateAwsSmSyncIntegrationOK{}
}

/*
CreateAwsSmSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAwsSmSyncIntegrationOK struct {
	Payload *models.Secrets20230613CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create aws sm sync integration o k response has a 2xx status code
func (o *CreateAwsSmSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create aws sm sync integration o k response has a 3xx status code
func (o *CreateAwsSmSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create aws sm sync integration o k response has a 4xx status code
func (o *CreateAwsSmSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create aws sm sync integration o k response has a 5xx status code
func (o *CreateAwsSmSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create aws sm sync integration o k response a status code equal to that given
func (o *CreateAwsSmSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create aws sm sync integration o k response
func (o *CreateAwsSmSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateAwsSmSyncIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm][%d] createAwsSmSyncIntegrationOK %s", 200, payload)
}

func (o *CreateAwsSmSyncIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm][%d] createAwsSmSyncIntegrationOK %s", 200, payload)
}

func (o *CreateAwsSmSyncIntegrationOK) GetPayload() *models.Secrets20230613CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateAwsSmSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAwsSmSyncIntegrationDefault creates a CreateAwsSmSyncIntegrationDefault with default headers values
func NewCreateAwsSmSyncIntegrationDefault(code int) *CreateAwsSmSyncIntegrationDefault {
	return &CreateAwsSmSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateAwsSmSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAwsSmSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create aws sm sync integration default response has a 2xx status code
func (o *CreateAwsSmSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create aws sm sync integration default response has a 3xx status code
func (o *CreateAwsSmSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create aws sm sync integration default response has a 4xx status code
func (o *CreateAwsSmSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create aws sm sync integration default response has a 5xx status code
func (o *CreateAwsSmSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create aws sm sync integration default response a status code equal to that given
func (o *CreateAwsSmSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create aws sm sync integration default response
func (o *CreateAwsSmSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateAwsSmSyncIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm][%d] CreateAwsSmSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateAwsSmSyncIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm][%d] CreateAwsSmSyncIntegration default %s", o._statusCode, payload)
}

func (o *CreateAwsSmSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateAwsSmSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAwsSmSyncIntegrationBody create aws sm sync integration body
swagger:model CreateAwsSmSyncIntegrationBody
*/
type CreateAwsSmSyncIntegrationBody struct {

	// aws sm connection details
	AwsSmConnectionDetails *models.Secrets20230613AwsSmConnectionDetailsRequest `json:"aws_sm_connection_details,omitempty"`

	// location
	Location *CreateAwsSmSyncIntegrationParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create aws sm sync integration body
func (o *CreateAwsSmSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAwsSmConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAwsSmSyncIntegrationBody) validateAwsSmConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.AwsSmConnectionDetails) { // not required
		return nil
	}

	if o.AwsSmConnectionDetails != nil {
		if err := o.AwsSmConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *CreateAwsSmSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this create aws sm sync integration body based on the context it is used
func (o *CreateAwsSmSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAwsSmConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAwsSmSyncIntegrationBody) contextValidateAwsSmConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.AwsSmConnectionDetails != nil {

		if swag.IsZero(o.AwsSmConnectionDetails) { // not required
			return nil
		}

		if err := o.AwsSmConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *CreateAwsSmSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *CreateAwsSmSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAwsSmSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateAwsSmSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAwsSmSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateAwsSmSyncIntegrationParamsBodyLocation
*/
type CreateAwsSmSyncIntegrationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create aws sm sync integration params body location
func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this create aws sm sync integration params body location based on the context it is used
func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAwsSmSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateAwsSmSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
