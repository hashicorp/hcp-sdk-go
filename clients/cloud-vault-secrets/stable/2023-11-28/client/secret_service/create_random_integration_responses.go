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

// CreateRandomIntegrationReader is a Reader for the CreateRandomIntegration structure.
type CreateRandomIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRandomIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRandomIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateRandomIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRandomIntegrationOK creates a CreateRandomIntegrationOK with default headers values
func NewCreateRandomIntegrationOK() *CreateRandomIntegrationOK {
	return &CreateRandomIntegrationOK{}
}

/*
CreateRandomIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRandomIntegrationOK struct {
	Payload *models.Secrets20231128CreateRandomIntegrationResponse
}

// IsSuccess returns true when this create random integration o k response has a 2xx status code
func (o *CreateRandomIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create random integration o k response has a 3xx status code
func (o *CreateRandomIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create random integration o k response has a 4xx status code
func (o *CreateRandomIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create random integration o k response has a 5xx status code
func (o *CreateRandomIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create random integration o k response a status code equal to that given
func (o *CreateRandomIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create random integration o k response
func (o *CreateRandomIntegrationOK) Code() int {
	return 200
}

func (o *CreateRandomIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config][%d] createRandomIntegrationOK %s", 200, payload)
}

func (o *CreateRandomIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config][%d] createRandomIntegrationOK %s", 200, payload)
}

func (o *CreateRandomIntegrationOK) GetPayload() *models.Secrets20231128CreateRandomIntegrationResponse {
	return o.Payload
}

func (o *CreateRandomIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateRandomIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRandomIntegrationDefault creates a CreateRandomIntegrationDefault with default headers values
func NewCreateRandomIntegrationDefault(code int) *CreateRandomIntegrationDefault {
	return &CreateRandomIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateRandomIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateRandomIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create random integration default response has a 2xx status code
func (o *CreateRandomIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create random integration default response has a 3xx status code
func (o *CreateRandomIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create random integration default response has a 4xx status code
func (o *CreateRandomIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create random integration default response has a 5xx status code
func (o *CreateRandomIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create random integration default response a status code equal to that given
func (o *CreateRandomIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create random integration default response
func (o *CreateRandomIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateRandomIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config][%d] CreateRandomIntegration default %s", o._statusCode, payload)
}

func (o *CreateRandomIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config][%d] CreateRandomIntegration default %s", o._statusCode, payload)
}

func (o *CreateRandomIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateRandomIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
