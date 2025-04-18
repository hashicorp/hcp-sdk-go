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

// UpdateRandomRotatingSecretReader is a Reader for the UpdateRandomRotatingSecret structure.
type UpdateRandomRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRandomRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRandomRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateRandomRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRandomRotatingSecretOK creates a UpdateRandomRotatingSecretOK with default headers values
func NewUpdateRandomRotatingSecretOK() *UpdateRandomRotatingSecretOK {
	return &UpdateRandomRotatingSecretOK{}
}

/*
UpdateRandomRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRandomRotatingSecretOK struct {
	Payload *models.Secrets20231128UpdateRandomRotatingSecretResponse
}

// IsSuccess returns true when this update random rotating secret o k response has a 2xx status code
func (o *UpdateRandomRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update random rotating secret o k response has a 3xx status code
func (o *UpdateRandomRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update random rotating secret o k response has a 4xx status code
func (o *UpdateRandomRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update random rotating secret o k response has a 5xx status code
func (o *UpdateRandomRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update random rotating secret o k response a status code equal to that given
func (o *UpdateRandomRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update random rotating secret o k response
func (o *UpdateRandomRotatingSecretOK) Code() int {
	return 200
}

func (o *UpdateRandomRotatingSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/random/secret/{name}][%d] updateRandomRotatingSecretOK %s", 200, payload)
}

func (o *UpdateRandomRotatingSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/random/secret/{name}][%d] updateRandomRotatingSecretOK %s", 200, payload)
}

func (o *UpdateRandomRotatingSecretOK) GetPayload() *models.Secrets20231128UpdateRandomRotatingSecretResponse {
	return o.Payload
}

func (o *UpdateRandomRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateRandomRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRandomRotatingSecretDefault creates a UpdateRandomRotatingSecretDefault with default headers values
func NewUpdateRandomRotatingSecretDefault(code int) *UpdateRandomRotatingSecretDefault {
	return &UpdateRandomRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateRandomRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateRandomRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update random rotating secret default response has a 2xx status code
func (o *UpdateRandomRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update random rotating secret default response has a 3xx status code
func (o *UpdateRandomRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update random rotating secret default response has a 4xx status code
func (o *UpdateRandomRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update random rotating secret default response has a 5xx status code
func (o *UpdateRandomRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update random rotating secret default response a status code equal to that given
func (o *UpdateRandomRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update random rotating secret default response
func (o *UpdateRandomRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRandomRotatingSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/random/secret/{name}][%d] UpdateRandomRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateRandomRotatingSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/random/secret/{name}][%d] UpdateRandomRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateRandomRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateRandomRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
