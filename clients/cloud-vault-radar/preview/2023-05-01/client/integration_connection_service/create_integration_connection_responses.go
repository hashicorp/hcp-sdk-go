// Code generated by go-swagger; DO NOT EDIT.

package integration_connection_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// CreateIntegrationConnectionReader is a Reader for the CreateIntegrationConnection structure.
type CreateIntegrationConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIntegrationConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIntegrationConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIntegrationConnectionOK creates a CreateIntegrationConnectionOK with default headers values
func NewCreateIntegrationConnectionOK() *CreateIntegrationConnectionOK {
	return &CreateIntegrationConnectionOK{}
}

/*
CreateIntegrationConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateIntegrationConnectionOK struct {
	Payload *models.VaultRadar20230501CreateIntegrationConnectionResponse
}

// IsSuccess returns true when this create integration connection o k response has a 2xx status code
func (o *CreateIntegrationConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create integration connection o k response has a 3xx status code
func (o *CreateIntegrationConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration connection o k response has a 4xx status code
func (o *CreateIntegrationConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create integration connection o k response has a 5xx status code
func (o *CreateIntegrationConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration connection o k response a status code equal to that given
func (o *CreateIntegrationConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create integration connection o k response
func (o *CreateIntegrationConnectionOK) Code() int {
	return 200
}

func (o *CreateIntegrationConnectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections][%d] createIntegrationConnectionOK %s", 200, payload)
}

func (o *CreateIntegrationConnectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections][%d] createIntegrationConnectionOK %s", 200, payload)
}

func (o *CreateIntegrationConnectionOK) GetPayload() *models.VaultRadar20230501CreateIntegrationConnectionResponse {
	return o.Payload
}

func (o *CreateIntegrationConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultRadar20230501CreateIntegrationConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationConnectionDefault creates a CreateIntegrationConnectionDefault with default headers values
func NewCreateIntegrationConnectionDefault(code int) *CreateIntegrationConnectionDefault {
	return &CreateIntegrationConnectionDefault{
		_statusCode: code,
	}
}

/*
CreateIntegrationConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateIntegrationConnectionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create integration connection default response has a 2xx status code
func (o *CreateIntegrationConnectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create integration connection default response has a 3xx status code
func (o *CreateIntegrationConnectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create integration connection default response has a 4xx status code
func (o *CreateIntegrationConnectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create integration connection default response has a 5xx status code
func (o *CreateIntegrationConnectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create integration connection default response a status code equal to that given
func (o *CreateIntegrationConnectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create integration connection default response
func (o *CreateIntegrationConnectionDefault) Code() int {
	return o._statusCode
}

func (o *CreateIntegrationConnectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections][%d] CreateIntegrationConnection default %s", o._statusCode, payload)
}

func (o *CreateIntegrationConnectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections][%d] CreateIntegrationConnection default %s", o._statusCode, payload)
}

func (o *CreateIntegrationConnectionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateIntegrationConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateIntegrationConnectionBody create integration connection body
swagger:model CreateIntegrationConnectionBody
*/
type CreateIntegrationConnectionBody struct {

	// auth key
	AuthKey string `json:"auth_key,omitempty"`

	// details
	Details string `json:"details,omitempty"`

	// integration type
	IntegrationType string `json:"integration_type,omitempty"`

	// is sink
	IsSink bool `json:"is_sink,omitempty"`

	// is source
	IsSource bool `json:"is_source,omitempty"`

	// location
	Location *CreateIntegrationConnectionParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create integration connection body
func (o *CreateIntegrationConnectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationConnectionBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this create integration connection body based on the context it is used
func (o *CreateIntegrationConnectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationConnectionBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateIntegrationConnectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationConnectionBody) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationConnectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateIntegrationConnectionParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateIntegrationConnectionParamsBodyLocation
*/
type CreateIntegrationConnectionParamsBodyLocation struct {

	// organization_id is the id of the organization.
	OrganizationID string `json:"organization_id,omitempty"`

	// region
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create integration connection params body location
func (o *CreateIntegrationConnectionParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationConnectionParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this create integration connection params body location based on the context it is used
func (o *CreateIntegrationConnectionParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationConnectionParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateIntegrationConnectionParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationConnectionParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationConnectionParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
