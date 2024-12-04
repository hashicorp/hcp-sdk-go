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

// UpdateAwsIAMUserAccessKeyRotatingSecretReader is a Reader for the UpdateAwsIAMUserAccessKeyRotatingSecret structure.
type UpdateAwsIAMUserAccessKeyRotatingSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAwsIAMUserAccessKeyRotatingSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAwsIAMUserAccessKeyRotatingSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAwsIAMUserAccessKeyRotatingSecretOK creates a UpdateAwsIAMUserAccessKeyRotatingSecretOK with default headers values
func NewUpdateAwsIAMUserAccessKeyRotatingSecretOK() *UpdateAwsIAMUserAccessKeyRotatingSecretOK {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretOK{}
}

/*
UpdateAwsIAMUserAccessKeyRotatingSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAwsIAMUserAccessKeyRotatingSecretOK struct {
	Payload *models.Secrets20231128UpdateAwsIAMUserAccessKeyRotatingSecretResponse
}

// IsSuccess returns true when this update aws i a m user access key rotating secret o k response has a 2xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update aws i a m user access key rotating secret o k response has a 3xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update aws i a m user access key rotating secret o k response has a 4xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update aws i a m user access key rotating secret o k response has a 5xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update aws i a m user access key rotating secret o k response a status code equal to that given
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update aws i a m user access key rotating secret o k response
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) Code() int {
	return 200
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] updateAwsIAMUserAccessKeyRotatingSecretOK %s", 200, payload)
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] updateAwsIAMUserAccessKeyRotatingSecretOK %s", 200, payload)
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) GetPayload() *models.Secrets20231128UpdateAwsIAMUserAccessKeyRotatingSecretResponse {
	return o.Payload
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateAwsIAMUserAccessKeyRotatingSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsIAMUserAccessKeyRotatingSecretDefault creates a UpdateAwsIAMUserAccessKeyRotatingSecretDefault with default headers values
func NewUpdateAwsIAMUserAccessKeyRotatingSecretDefault(code int) *UpdateAwsIAMUserAccessKeyRotatingSecretDefault {
	return &UpdateAwsIAMUserAccessKeyRotatingSecretDefault{
		_statusCode: code,
	}
}

/*
UpdateAwsIAMUserAccessKeyRotatingSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAwsIAMUserAccessKeyRotatingSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update aws i a m user access key rotating secret default response has a 2xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update aws i a m user access key rotating secret default response has a 3xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update aws i a m user access key rotating secret default response has a 4xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update aws i a m user access key rotating secret default response has a 5xx status code
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update aws i a m user access key rotating secret default response a status code equal to that given
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update aws i a m user access key rotating secret default response
func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] UpdateAwsIAMUserAccessKeyRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] UpdateAwsIAMUserAccessKeyRotatingSecret default %s", o._statusCode, payload)
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateAwsIAMUserAccessKeyRotatingSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
