// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// UpdateTwilioRotatingSecretReader is a Reader for the UpdateTwilioRotatingSecret structure.
type UpdateTwilioRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTwilioRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTwilioRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTwilioRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTwilioRotatingSecretOK creates a UpdateTwilioRotatingSecretOK with default headers values
func NewUpdateTwilioRotatingSecretOK() *UpdateTwilioRotatingSecretOK {
	return &UpdateTwilioRotatingSecretOK{}
}

/*
UpdateTwilioRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTwilioRotatingSecretOK struct {
	Payload *models.Secrets20231128UpdateTwilioRotatingSecretResponse
}

// IsSuccess returns true when this update twilio rotating secret o k response has a 2xx status code
func (o *UpdateTwilioRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update twilio rotating secret o k response has a 3xx status code
func (o *UpdateTwilioRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update twilio rotating secret o k response has a 4xx status code
func (o *UpdateTwilioRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update twilio rotating secret o k response has a 5xx status code
func (o *UpdateTwilioRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update twilio rotating secret o k response a status code equal to that given
func (o *UpdateTwilioRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update twilio rotating secret o k response
func (o *UpdateTwilioRotatingSecretOK) Code() int {
	return 200
}

func (o *UpdateTwilioRotatingSecretOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] updateTwilioRotatingSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateTwilioRotatingSecretOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] updateTwilioRotatingSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateTwilioRotatingSecretOK) GetPayload() *models.Secrets20231128UpdateTwilioRotatingSecretResponse {
	return o.Payload
}

func (o *UpdateTwilioRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateTwilioRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTwilioRotatingSecretDefault creates a UpdateTwilioRotatingSecretDefault with default headers values
func NewUpdateTwilioRotatingSecretDefault(code int) *UpdateTwilioRotatingSecretDefault {
	return &UpdateTwilioRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateTwilioRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateTwilioRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update twilio rotating secret default response has a 2xx status code
func (o *UpdateTwilioRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update twilio rotating secret default response has a 3xx status code
func (o *UpdateTwilioRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update twilio rotating secret default response has a 4xx status code
func (o *UpdateTwilioRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update twilio rotating secret default response has a 5xx status code
func (o *UpdateTwilioRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update twilio rotating secret default response a status code equal to that given
func (o *UpdateTwilioRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update twilio rotating secret default response
func (o *UpdateTwilioRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTwilioRotatingSecretDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] UpdateTwilioRotatingSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTwilioRotatingSecretDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] UpdateTwilioRotatingSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTwilioRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateTwilioRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
