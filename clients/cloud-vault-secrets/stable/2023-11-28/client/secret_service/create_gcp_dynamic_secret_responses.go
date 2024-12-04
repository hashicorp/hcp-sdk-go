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

// CreateGcpDynamicSecretReader is a Reader for the CreateGcpDynamicSecret structure.
type CreateGcpDynamicSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGcpDynamicSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGcpDynamicSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateGcpDynamicSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGcpDynamicSecretOK creates a CreateGcpDynamicSecretOK with default headers values
func NewCreateGcpDynamicSecretOK() *CreateGcpDynamicSecretOK {
	return &CreateGcpDynamicSecretOK{}
}

/*
CreateGcpDynamicSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateGcpDynamicSecretOK struct {
	Payload *models.Secrets20231128CreateGcpDynamicSecretResponse
}

// IsSuccess returns true when this create gcp dynamic secret o k response has a 2xx status code
func (o *CreateGcpDynamicSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create gcp dynamic secret o k response has a 3xx status code
func (o *CreateGcpDynamicSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gcp dynamic secret o k response has a 4xx status code
func (o *CreateGcpDynamicSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gcp dynamic secret o k response has a 5xx status code
func (o *CreateGcpDynamicSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create gcp dynamic secret o k response a status code equal to that given
func (o *CreateGcpDynamicSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create gcp dynamic secret o k response
func (o *CreateGcpDynamicSecretOK) Code() int {
	return 200
}

func (o *CreateGcpDynamicSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] createGcpDynamicSecretOK %s", 200, payload)
}

func (o *CreateGcpDynamicSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] createGcpDynamicSecretOK %s", 200, payload)
}

func (o *CreateGcpDynamicSecretOK) GetPayload() *models.Secrets20231128CreateGcpDynamicSecretResponse {
	return o.Payload
}

func (o *CreateGcpDynamicSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateGcpDynamicSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGcpDynamicSecretDefault creates a CreateGcpDynamicSecretDefault with default headers values
func NewCreateGcpDynamicSecretDefault(code int) *CreateGcpDynamicSecretDefault {
	return &CreateGcpDynamicSecretDefault{
		_statusCode: code,
	}
}

/*
CreateGcpDynamicSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateGcpDynamicSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create gcp dynamic secret default response has a 2xx status code
func (o *CreateGcpDynamicSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create gcp dynamic secret default response has a 3xx status code
func (o *CreateGcpDynamicSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create gcp dynamic secret default response has a 4xx status code
func (o *CreateGcpDynamicSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create gcp dynamic secret default response has a 5xx status code
func (o *CreateGcpDynamicSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create gcp dynamic secret default response a status code equal to that given
func (o *CreateGcpDynamicSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create gcp dynamic secret default response
func (o *CreateGcpDynamicSecretDefault) Code() int {
	return o._statusCode
}

func (o *CreateGcpDynamicSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] CreateGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *CreateGcpDynamicSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] CreateGcpDynamicSecret default %s", o._statusCode, payload)
}

func (o *CreateGcpDynamicSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateGcpDynamicSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
