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

// UpdateAwsIntegrationReader is a Reader for the UpdateAwsIntegration structure.
type UpdateAwsIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAwsIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAwsIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAwsIntegrationOK creates a UpdateAwsIntegrationOK with default headers values
func NewUpdateAwsIntegrationOK() *UpdateAwsIntegrationOK {
	return &UpdateAwsIntegrationOK{}
}

/*
UpdateAwsIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAwsIntegrationOK struct {
	Payload *models.Secrets20231128UpdateAwsIntegrationResponse
}

// IsSuccess returns true when this update aws integration o k response has a 2xx status code
func (o *UpdateAwsIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update aws integration o k response has a 3xx status code
func (o *UpdateAwsIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update aws integration o k response has a 4xx status code
func (o *UpdateAwsIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update aws integration o k response has a 5xx status code
func (o *UpdateAwsIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update aws integration o k response a status code equal to that given
func (o *UpdateAwsIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update aws integration o k response
func (o *UpdateAwsIntegrationOK) Code() int {
	return 200
}

func (o *UpdateAwsIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/aws/config/{name}][%d] updateAwsIntegrationOK %s", 200, payload)
}

func (o *UpdateAwsIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/aws/config/{name}][%d] updateAwsIntegrationOK %s", 200, payload)
}

func (o *UpdateAwsIntegrationOK) GetPayload() *models.Secrets20231128UpdateAwsIntegrationResponse {
	return o.Payload
}

func (o *UpdateAwsIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateAwsIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsIntegrationDefault creates a UpdateAwsIntegrationDefault with default headers values
func NewUpdateAwsIntegrationDefault(code int) *UpdateAwsIntegrationDefault {
	return &UpdateAwsIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpdateAwsIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAwsIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this update aws integration default response has a 2xx status code
func (o *UpdateAwsIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update aws integration default response has a 3xx status code
func (o *UpdateAwsIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update aws integration default response has a 4xx status code
func (o *UpdateAwsIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update aws integration default response has a 5xx status code
func (o *UpdateAwsIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update aws integration default response a status code equal to that given
func (o *UpdateAwsIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update aws integration default response
func (o *UpdateAwsIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAwsIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/aws/config/{name}][%d] UpdateAwsIntegration default %s", o._statusCode, payload)
}

func (o *UpdateAwsIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/aws/config/{name}][%d] UpdateAwsIntegration default %s", o._statusCode, payload)
}

func (o *UpdateAwsIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UpdateAwsIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
