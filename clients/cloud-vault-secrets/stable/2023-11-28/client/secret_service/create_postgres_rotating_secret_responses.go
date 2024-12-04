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

// CreatePostgresRotatingSecretReader is a Reader for the CreatePostgresRotatingSecret structure.
type CreatePostgresRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePostgresRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePostgresRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePostgresRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePostgresRotatingSecretOK creates a CreatePostgresRotatingSecretOK with default headers values
func NewCreatePostgresRotatingSecretOK() *CreatePostgresRotatingSecretOK {
	return &CreatePostgresRotatingSecretOK{}
}

/*
CreatePostgresRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreatePostgresRotatingSecretOK struct {
	Payload *models.Secrets20231128CreatePostgresRotatingSecretResponse
}

// IsSuccess returns true when this create postgres rotating secret o k response has a 2xx status code
func (o *CreatePostgresRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create postgres rotating secret o k response has a 3xx status code
func (o *CreatePostgresRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create postgres rotating secret o k response has a 4xx status code
func (o *CreatePostgresRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create postgres rotating secret o k response has a 5xx status code
func (o *CreatePostgresRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create postgres rotating secret o k response a status code equal to that given
func (o *CreatePostgresRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create postgres rotating secret o k response
func (o *CreatePostgresRotatingSecretOK) Code() int {
	return 200
}

func (o *CreatePostgresRotatingSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret][%d] createPostgresRotatingSecretOK %s", 200, payload)
}

func (o *CreatePostgresRotatingSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret][%d] createPostgresRotatingSecretOK %s", 200, payload)
}

func (o *CreatePostgresRotatingSecretOK) GetPayload() *models.Secrets20231128CreatePostgresRotatingSecretResponse {
	return o.Payload
}

func (o *CreatePostgresRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreatePostgresRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresRotatingSecretDefault creates a CreatePostgresRotatingSecretDefault with default headers values
func NewCreatePostgresRotatingSecretDefault(code int) *CreatePostgresRotatingSecretDefault {
	return &CreatePostgresRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
CreatePostgresRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreatePostgresRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this create postgres rotating secret default response has a 2xx status code
func (o *CreatePostgresRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create postgres rotating secret default response has a 3xx status code
func (o *CreatePostgresRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create postgres rotating secret default response has a 4xx status code
func (o *CreatePostgresRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create postgres rotating secret default response has a 5xx status code
func (o *CreatePostgresRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create postgres rotating secret default response a status code equal to that given
func (o *CreatePostgresRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create postgres rotating secret default response
func (o *CreatePostgresRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *CreatePostgresRotatingSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret][%d] CreatePostgresRotatingSecret default %s", o._statusCode, payload)
}

func (o *CreatePostgresRotatingSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/postgres/secret][%d] CreatePostgresRotatingSecret default %s", o._statusCode, payload)
}

func (o *CreatePostgresRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreatePostgresRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
