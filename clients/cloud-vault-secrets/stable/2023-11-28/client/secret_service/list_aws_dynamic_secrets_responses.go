// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// ListAwsDynamicSecretsReader is a Reader for the ListAwsDynamicSecrets structure.
type ListAwsDynamicSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAwsDynamicSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAwsDynamicSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAwsDynamicSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAwsDynamicSecretsOK creates a ListAwsDynamicSecretsOK with default headers values
func NewListAwsDynamicSecretsOK() *ListAwsDynamicSecretsOK {
	return &ListAwsDynamicSecretsOK{}
}

/*
ListAwsDynamicSecretsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListAwsDynamicSecretsOK struct {
	Payload *models.Secrets20231128ListAwsDynamicSecretsResponse
}

// IsSuccess returns true when this list aws dynamic secrets o k response has a 2xx status code
func (o *ListAwsDynamicSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list aws dynamic secrets o k response has a 3xx status code
func (o *ListAwsDynamicSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list aws dynamic secrets o k response has a 4xx status code
func (o *ListAwsDynamicSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list aws dynamic secrets o k response has a 5xx status code
func (o *ListAwsDynamicSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list aws dynamic secrets o k response a status code equal to that given
func (o *ListAwsDynamicSecretsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list aws dynamic secrets o k response
func (o *ListAwsDynamicSecretsOK) Code() int {
	return 200
}

func (o *ListAwsDynamicSecretsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret][%d] listAwsDynamicSecretsOK  %+v", 200, o.Payload)
}

func (o *ListAwsDynamicSecretsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret][%d] listAwsDynamicSecretsOK  %+v", 200, o.Payload)
}

func (o *ListAwsDynamicSecretsOK) GetPayload() *models.Secrets20231128ListAwsDynamicSecretsResponse {
	return o.Payload
}

func (o *ListAwsDynamicSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListAwsDynamicSecretsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAwsDynamicSecretsDefault creates a ListAwsDynamicSecretsDefault with default headers values
func NewListAwsDynamicSecretsDefault(code int) *ListAwsDynamicSecretsDefault {
	return &ListAwsDynamicSecretsDefault{
		_statusCode: code,
	}
}

/*
ListAwsDynamicSecretsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListAwsDynamicSecretsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list aws dynamic secrets default response has a 2xx status code
func (o *ListAwsDynamicSecretsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list aws dynamic secrets default response has a 3xx status code
func (o *ListAwsDynamicSecretsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list aws dynamic secrets default response has a 4xx status code
func (o *ListAwsDynamicSecretsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list aws dynamic secrets default response has a 5xx status code
func (o *ListAwsDynamicSecretsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list aws dynamic secrets default response a status code equal to that given
func (o *ListAwsDynamicSecretsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list aws dynamic secrets default response
func (o *ListAwsDynamicSecretsDefault) Code() int {
	return o._statusCode
}

func (o *ListAwsDynamicSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret][%d] ListAwsDynamicSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAwsDynamicSecretsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/aws/secret][%d] ListAwsDynamicSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAwsDynamicSecretsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListAwsDynamicSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}