// Code generated by go-swagger; DO NOT EDIT.

package auth_config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// AuthConfigServiceEditAuthConnectionReader is a Reader for the AuthConfigServiceEditAuthConnection structure.
type AuthConfigServiceEditAuthConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthConfigServiceEditAuthConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthConfigServiceEditAuthConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthConfigServiceEditAuthConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthConfigServiceEditAuthConnectionOK creates a AuthConfigServiceEditAuthConnectionOK with default headers values
func NewAuthConfigServiceEditAuthConnectionOK() *AuthConfigServiceEditAuthConnectionOK {
	return &AuthConfigServiceEditAuthConnectionOK{}
}

/*
AuthConfigServiceEditAuthConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthConfigServiceEditAuthConnectionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this auth config service edit auth connection o k response has a 2xx status code
func (o *AuthConfigServiceEditAuthConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth config service edit auth connection o k response has a 3xx status code
func (o *AuthConfigServiceEditAuthConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth config service edit auth connection o k response has a 4xx status code
func (o *AuthConfigServiceEditAuthConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth config service edit auth connection o k response has a 5xx status code
func (o *AuthConfigServiceEditAuthConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth config service edit auth connection o k response a status code equal to that given
func (o *AuthConfigServiceEditAuthConnectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthConfigServiceEditAuthConnectionOK) Error() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceEditAuthConnectionOK  %+v", 200, o.Payload)
}

func (o *AuthConfigServiceEditAuthConnectionOK) String() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] authConfigServiceEditAuthConnectionOK  %+v", 200, o.Payload)
}

func (o *AuthConfigServiceEditAuthConnectionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthConfigServiceEditAuthConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthConfigServiceEditAuthConnectionDefault creates a AuthConfigServiceEditAuthConnectionDefault with default headers values
func NewAuthConfigServiceEditAuthConnectionDefault(code int) *AuthConfigServiceEditAuthConnectionDefault {
	return &AuthConfigServiceEditAuthConnectionDefault{
		_statusCode: code,
	}
}

/*
AuthConfigServiceEditAuthConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuthConfigServiceEditAuthConnectionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the auth config service edit auth connection default response
func (o *AuthConfigServiceEditAuthConnectionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this auth config service edit auth connection default response has a 2xx status code
func (o *AuthConfigServiceEditAuthConnectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auth config service edit auth connection default response has a 3xx status code
func (o *AuthConfigServiceEditAuthConnectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auth config service edit auth connection default response has a 4xx status code
func (o *AuthConfigServiceEditAuthConnectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auth config service edit auth connection default response has a 5xx status code
func (o *AuthConfigServiceEditAuthConnectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auth config service edit auth connection default response a status code equal to that given
func (o *AuthConfigServiceEditAuthConnectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AuthConfigServiceEditAuthConnectionDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_EditAuthConnection default  %+v", o._statusCode, o.Payload)
}

func (o *AuthConfigServiceEditAuthConnectionDefault) String() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] AuthConfigService_EditAuthConnection default  %+v", o._statusCode, o.Payload)
}

func (o *AuthConfigServiceEditAuthConnectionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *AuthConfigServiceEditAuthConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
