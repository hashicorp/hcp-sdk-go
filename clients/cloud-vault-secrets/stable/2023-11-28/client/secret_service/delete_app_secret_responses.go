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

// DeleteAppSecretReader is a Reader for the DeleteAppSecret structure.
type DeleteAppSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAppSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAppSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAppSecretOK creates a DeleteAppSecretOK with default headers values
func NewDeleteAppSecretOK() *DeleteAppSecretOK {
	return &DeleteAppSecretOK{}
}

/*
DeleteAppSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAppSecretOK struct {
	Payload models.Secrets20231128DeleteAppSecretResponse
}

// IsSuccess returns true when this delete app secret o k response has a 2xx status code
func (o *DeleteAppSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete app secret o k response has a 3xx status code
func (o *DeleteAppSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app secret o k response has a 4xx status code
func (o *DeleteAppSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete app secret o k response has a 5xx status code
func (o *DeleteAppSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app secret o k response a status code equal to that given
func (o *DeleteAppSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete app secret o k response
func (o *DeleteAppSecretOK) Code() int {
	return 200
}

func (o *DeleteAppSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}][%d] deleteAppSecretOK %s", 200, payload)
}

func (o *DeleteAppSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}][%d] deleteAppSecretOK %s", 200, payload)
}

func (o *DeleteAppSecretOK) GetPayload() models.Secrets20231128DeleteAppSecretResponse {
	return o.Payload
}

func (o *DeleteAppSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppSecretDefault creates a DeleteAppSecretDefault with default headers values
func NewDeleteAppSecretDefault(code int) *DeleteAppSecretDefault {
	return &DeleteAppSecretDefault{
		_statusCode: code,
	}
}

/*
DeleteAppSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAppSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete app secret default response has a 2xx status code
func (o *DeleteAppSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete app secret default response has a 3xx status code
func (o *DeleteAppSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete app secret default response has a 4xx status code
func (o *DeleteAppSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete app secret default response has a 5xx status code
func (o *DeleteAppSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete app secret default response a status code equal to that given
func (o *DeleteAppSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete app secret default response
func (o *DeleteAppSecretDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAppSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}][%d] DeleteAppSecret default %s", o._statusCode, payload)
}

func (o *DeleteAppSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}][%d] DeleteAppSecret default %s", o._statusCode, payload)
}

func (o *DeleteAppSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteAppSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
