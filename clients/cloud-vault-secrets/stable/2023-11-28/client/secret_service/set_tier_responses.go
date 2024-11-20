// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// SetTierReader is a Reader for the SetTier structure.
type SetTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetTierOK creates a SetTierOK with default headers values
func NewSetTierOK() *SetTierOK {
	return &SetTierOK{}
}

/*
SetTierOK describes a response with status code 200, with default header values.

A successful response.
*/
type SetTierOK struct {
	Payload interface{}
}

// IsSuccess returns true when this set tier o k response has a 2xx status code
func (o *SetTierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set tier o k response has a 3xx status code
func (o *SetTierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set tier o k response has a 4xx status code
func (o *SetTierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set tier o k response has a 5xx status code
func (o *SetTierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set tier o k response a status code equal to that given
func (o *SetTierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set tier o k response
func (o *SetTierOK) Code() int {
	return 200
}

func (o *SetTierOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/billing/tier][%d] setTierOK  %+v", 200, o.Payload)
}

func (o *SetTierOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/billing/tier][%d] setTierOK  %+v", 200, o.Payload)
}

func (o *SetTierOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SetTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTierDefault creates a SetTierDefault with default headers values
func NewSetTierDefault(code int) *SetTierDefault {
	return &SetTierDefault{
		_statusCode: code,
	}
}

/*
SetTierDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SetTierDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this set tier default response has a 2xx status code
func (o *SetTierDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set tier default response has a 3xx status code
func (o *SetTierDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set tier default response has a 4xx status code
func (o *SetTierDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set tier default response has a 5xx status code
func (o *SetTierDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set tier default response a status code equal to that given
func (o *SetTierDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set tier default response
func (o *SetTierDefault) Code() int {
	return o._statusCode
}

func (o *SetTierDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/billing/tier][%d] SetTier default  %+v", o._statusCode, o.Payload)
}

func (o *SetTierDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/billing/tier][%d] SetTier default  %+v", o._statusCode, o.Payload)
}

func (o *SetTierDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *SetTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}