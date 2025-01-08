// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceRegenerateTFCRunTaskHmacKeyReader is a Reader for the PackerServiceRegenerateTFCRunTaskHmacKey structure.
type PackerServiceRegenerateTFCRunTaskHmacKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceRegenerateTFCRunTaskHmacKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyOK creates a PackerServiceRegenerateTFCRunTaskHmacKeyOK with default headers values
func NewPackerServiceRegenerateTFCRunTaskHmacKeyOK() *PackerServiceRegenerateTFCRunTaskHmacKeyOK {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyOK{}
}

/*
PackerServiceRegenerateTFCRunTaskHmacKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyOK struct {
	Payload *models.HashicorpCloudPacker20230101RegenerateTFCRunTaskHmacKeyResponse
}

// IsSuccess returns true when this packer service regenerate t f c run task hmac key o k response has a 2xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service regenerate t f c run task hmac key o k response has a 3xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service regenerate t f c run task hmac key o k response has a 4xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service regenerate t f c run task hmac key o k response has a 5xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service regenerate t f c run task hmac key o k response a status code equal to that given
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service regenerate t f c run task hmac key o k response
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) Code() int {
	return 200
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] packerServiceRegenerateTFCRunTaskHmacKeyOK %s", 200, payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] packerServiceRegenerateTFCRunTaskHmacKeyOK %s", 200, payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) GetPayload() *models.HashicorpCloudPacker20230101RegenerateTFCRunTaskHmacKeyResponse {
	return o.Payload
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101RegenerateTFCRunTaskHmacKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault creates a PackerServiceRegenerateTFCRunTaskHmacKeyDefault with default headers values
func NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault(code int) *PackerServiceRegenerateTFCRunTaskHmacKeyDefault {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyDefault{
		_statusCode: code,
	}
}

/*
PackerServiceRegenerateTFCRunTaskHmacKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service regenerate t f c run task hmac key default response has a 2xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service regenerate t f c run task hmac key default response has a 3xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service regenerate t f c run task hmac key default response has a 4xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service regenerate t f c run task hmac key default response has a 5xx status code
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service regenerate t f c run task hmac key default response a status code equal to that given
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service regenerate t f c run task hmac key default response
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] PackerService_RegenerateTFCRunTaskHmacKey default %s", o._statusCode, payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] PackerService_RegenerateTFCRunTaskHmacKey default %s", o._statusCode, payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceRegenerateTFCRunTaskHmacKeyBody packer service regenerate t f c run task hmac key body
swagger:model PackerServiceRegenerateTFCRunTaskHmacKeyBody
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyBody struct {

	// location
	Location *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation `json:"location,omitempty"`
}

// Validate validates this packer service regenerate t f c run task hmac key body
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this packer service regenerate t f c run task hmac key body based on the context it is used
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceRegenerateTFCRunTaskHmacKeyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this packer service regenerate t f c run task hmac key params body location
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this packer service regenerate t f c run task hmac key params body location based on the context it is used
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res PackerServiceRegenerateTFCRunTaskHmacKeyParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
