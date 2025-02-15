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

// GetTwilioRotatingSecretConfigReader is a Reader for the GetTwilioRotatingSecretConfig structure.
type GetTwilioRotatingSecretConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTwilioRotatingSecretConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTwilioRotatingSecretConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTwilioRotatingSecretConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTwilioRotatingSecretConfigOK creates a GetTwilioRotatingSecretConfigOK with default headers values
func NewGetTwilioRotatingSecretConfigOK() *GetTwilioRotatingSecretConfigOK {
	return &GetTwilioRotatingSecretConfigOK{}
}

/*
GetTwilioRotatingSecretConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTwilioRotatingSecretConfigOK struct {
	Payload *models.Secrets20231128GetTwilioRotatingSecretConfigResponse
}

// IsSuccess returns true when this get twilio rotating secret config o k response has a 2xx status code
func (o *GetTwilioRotatingSecretConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get twilio rotating secret config o k response has a 3xx status code
func (o *GetTwilioRotatingSecretConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get twilio rotating secret config o k response has a 4xx status code
func (o *GetTwilioRotatingSecretConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get twilio rotating secret config o k response has a 5xx status code
func (o *GetTwilioRotatingSecretConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get twilio rotating secret config o k response a status code equal to that given
func (o *GetTwilioRotatingSecretConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get twilio rotating secret config o k response
func (o *GetTwilioRotatingSecretConfigOK) Code() int {
	return 200
}

func (o *GetTwilioRotatingSecretConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] getTwilioRotatingSecretConfigOK %s", 200, payload)
}

func (o *GetTwilioRotatingSecretConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] getTwilioRotatingSecretConfigOK %s", 200, payload)
}

func (o *GetTwilioRotatingSecretConfigOK) GetPayload() *models.Secrets20231128GetTwilioRotatingSecretConfigResponse {
	return o.Payload
}

func (o *GetTwilioRotatingSecretConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetTwilioRotatingSecretConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTwilioRotatingSecretConfigDefault creates a GetTwilioRotatingSecretConfigDefault with default headers values
func NewGetTwilioRotatingSecretConfigDefault(code int) *GetTwilioRotatingSecretConfigDefault {
	return &GetTwilioRotatingSecretConfigDefault{
		_statusCode: code,
	}
}

/*
GetTwilioRotatingSecretConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTwilioRotatingSecretConfigDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get twilio rotating secret config default response has a 2xx status code
func (o *GetTwilioRotatingSecretConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get twilio rotating secret config default response has a 3xx status code
func (o *GetTwilioRotatingSecretConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get twilio rotating secret config default response has a 4xx status code
func (o *GetTwilioRotatingSecretConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get twilio rotating secret config default response has a 5xx status code
func (o *GetTwilioRotatingSecretConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get twilio rotating secret config default response a status code equal to that given
func (o *GetTwilioRotatingSecretConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get twilio rotating secret config default response
func (o *GetTwilioRotatingSecretConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetTwilioRotatingSecretConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] GetTwilioRotatingSecretConfig default %s", o._statusCode, payload)
}

func (o *GetTwilioRotatingSecretConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/twilio/secret/{name}][%d] GetTwilioRotatingSecretConfig default %s", o._statusCode, payload)
}

func (o *GetTwilioRotatingSecretConfigDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetTwilioRotatingSecretConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
