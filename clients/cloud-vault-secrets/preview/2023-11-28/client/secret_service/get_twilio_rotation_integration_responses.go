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

// GetTwilioRotationIntegrationReader is a Reader for the GetTwilioRotationIntegration structure.
type GetTwilioRotationIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTwilioRotationIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTwilioRotationIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTwilioRotationIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTwilioRotationIntegrationOK creates a GetTwilioRotationIntegrationOK with default headers values
func NewGetTwilioRotationIntegrationOK() *GetTwilioRotationIntegrationOK {
	return &GetTwilioRotationIntegrationOK{}
}

/*
GetTwilioRotationIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTwilioRotationIntegrationOK struct {
	Payload *models.Secrets20231128GetTwilioRotationIntegrationResponse
}

// IsSuccess returns true when this get twilio rotation integration o k response has a 2xx status code
func (o *GetTwilioRotationIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get twilio rotation integration o k response has a 3xx status code
func (o *GetTwilioRotationIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get twilio rotation integration o k response has a 4xx status code
func (o *GetTwilioRotationIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get twilio rotation integration o k response has a 5xx status code
func (o *GetTwilioRotationIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get twilio rotation integration o k response a status code equal to that given
func (o *GetTwilioRotationIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get twilio rotation integration o k response
func (o *GetTwilioRotationIntegrationOK) Code() int {
	return 200
}

func (o *GetTwilioRotationIntegrationOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/twilio/{integration_name}][%d] getTwilioRotationIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetTwilioRotationIntegrationOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/twilio/{integration_name}][%d] getTwilioRotationIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetTwilioRotationIntegrationOK) GetPayload() *models.Secrets20231128GetTwilioRotationIntegrationResponse {
	return o.Payload
}

func (o *GetTwilioRotationIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetTwilioRotationIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTwilioRotationIntegrationDefault creates a GetTwilioRotationIntegrationDefault with default headers values
func NewGetTwilioRotationIntegrationDefault(code int) *GetTwilioRotationIntegrationDefault {
	return &GetTwilioRotationIntegrationDefault{
		_statusCode: code,
	}
}

/*
GetTwilioRotationIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTwilioRotationIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get twilio rotation integration default response has a 2xx status code
func (o *GetTwilioRotationIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get twilio rotation integration default response has a 3xx status code
func (o *GetTwilioRotationIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get twilio rotation integration default response has a 4xx status code
func (o *GetTwilioRotationIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get twilio rotation integration default response has a 5xx status code
func (o *GetTwilioRotationIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get twilio rotation integration default response a status code equal to that given
func (o *GetTwilioRotationIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get twilio rotation integration default response
func (o *GetTwilioRotationIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *GetTwilioRotationIntegrationDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/twilio/{integration_name}][%d] GetTwilioRotationIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetTwilioRotationIntegrationDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/twilio/{integration_name}][%d] GetTwilioRotationIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetTwilioRotationIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetTwilioRotationIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
