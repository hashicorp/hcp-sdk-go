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

// CreateTwilioIntegrationReader is a Reader for the CreateTwilioIntegration structure.
type CreateTwilioIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTwilioIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTwilioIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTwilioIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTwilioIntegrationOK creates a CreateTwilioIntegrationOK with default headers values
func NewCreateTwilioIntegrationOK() *CreateTwilioIntegrationOK {
	return &CreateTwilioIntegrationOK{}
}

/*
CreateTwilioIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTwilioIntegrationOK struct {
	Payload *models.Secrets20231128CreateTwilioIntegrationResponse
}

// IsSuccess returns true when this create twilio integration o k response has a 2xx status code
func (o *CreateTwilioIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create twilio integration o k response has a 3xx status code
func (o *CreateTwilioIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create twilio integration o k response has a 4xx status code
func (o *CreateTwilioIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create twilio integration o k response has a 5xx status code
func (o *CreateTwilioIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create twilio integration o k response a status code equal to that given
func (o *CreateTwilioIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create twilio integration o k response
func (o *CreateTwilioIntegrationOK) Code() int {
	return 200
}

func (o *CreateTwilioIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/twilio/config][%d] createTwilioIntegrationOK %s", 200, payload)
}

func (o *CreateTwilioIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/twilio/config][%d] createTwilioIntegrationOK %s", 200, payload)
}

func (o *CreateTwilioIntegrationOK) GetPayload() *models.Secrets20231128CreateTwilioIntegrationResponse {
	return o.Payload
}

func (o *CreateTwilioIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateTwilioIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTwilioIntegrationDefault creates a CreateTwilioIntegrationDefault with default headers values
func NewCreateTwilioIntegrationDefault(code int) *CreateTwilioIntegrationDefault {
	return &CreateTwilioIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateTwilioIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTwilioIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create twilio integration default response has a 2xx status code
func (o *CreateTwilioIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create twilio integration default response has a 3xx status code
func (o *CreateTwilioIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create twilio integration default response has a 4xx status code
func (o *CreateTwilioIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create twilio integration default response has a 5xx status code
func (o *CreateTwilioIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create twilio integration default response a status code equal to that given
func (o *CreateTwilioIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create twilio integration default response
func (o *CreateTwilioIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateTwilioIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/twilio/config][%d] CreateTwilioIntegration default %s", o._statusCode, payload)
}

func (o *CreateTwilioIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/twilio/config][%d] CreateTwilioIntegration default %s", o._statusCode, payload)
}

func (o *CreateTwilioIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateTwilioIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
