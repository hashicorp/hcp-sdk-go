// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// CreateTwilioRotatingSecretReader is a Reader for the CreateTwilioRotatingSecret structure.
type CreateTwilioRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTwilioRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTwilioRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTwilioRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTwilioRotatingSecretOK creates a CreateTwilioRotatingSecretOK with default headers values
func NewCreateTwilioRotatingSecretOK() *CreateTwilioRotatingSecretOK {
	return &CreateTwilioRotatingSecretOK{}
}

/*
CreateTwilioRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTwilioRotatingSecretOK struct {
	Payload *models.Secrets20231128CreateTwilioRotatingSecretResponse
}

// IsSuccess returns true when this create twilio rotating secret o k response has a 2xx status code
func (o *CreateTwilioRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create twilio rotating secret o k response has a 3xx status code
func (o *CreateTwilioRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create twilio rotating secret o k response has a 4xx status code
func (o *CreateTwilioRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create twilio rotating secret o k response has a 5xx status code
func (o *CreateTwilioRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create twilio rotating secret o k response a status code equal to that given
func (o *CreateTwilioRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create twilio rotating secret o k response
func (o *CreateTwilioRotatingSecretOK) Code() int {
	return 200
}

func (o *CreateTwilioRotatingSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret][%d] createTwilioRotatingSecretOK %s", 200, payload)
}

func (o *CreateTwilioRotatingSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret][%d] createTwilioRotatingSecretOK %s", 200, payload)
}

func (o *CreateTwilioRotatingSecretOK) GetPayload() *models.Secrets20231128CreateTwilioRotatingSecretResponse {
	return o.Payload
}

func (o *CreateTwilioRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateTwilioRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTwilioRotatingSecretDefault creates a CreateTwilioRotatingSecretDefault with default headers values
func NewCreateTwilioRotatingSecretDefault(code int) *CreateTwilioRotatingSecretDefault {
	return &CreateTwilioRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
CreateTwilioRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTwilioRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create twilio rotating secret default response has a 2xx status code
func (o *CreateTwilioRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create twilio rotating secret default response has a 3xx status code
func (o *CreateTwilioRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create twilio rotating secret default response has a 4xx status code
func (o *CreateTwilioRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create twilio rotating secret default response has a 5xx status code
func (o *CreateTwilioRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create twilio rotating secret default response a status code equal to that given
func (o *CreateTwilioRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create twilio rotating secret default response
func (o *CreateTwilioRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *CreateTwilioRotatingSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret][%d] CreateTwilioRotatingSecret default %s", o._statusCode, payload)
}

func (o *CreateTwilioRotatingSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret][%d] CreateTwilioRotatingSecret default %s", o._statusCode, payload)
}

func (o *CreateTwilioRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateTwilioRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
