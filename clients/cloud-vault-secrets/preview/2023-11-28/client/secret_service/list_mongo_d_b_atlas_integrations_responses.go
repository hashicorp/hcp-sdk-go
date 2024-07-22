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

// ListMongoDBAtlasIntegrationsReader is a Reader for the ListMongoDBAtlasIntegrations structure.
type ListMongoDBAtlasIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMongoDBAtlasIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMongoDBAtlasIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListMongoDBAtlasIntegrationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMongoDBAtlasIntegrationsOK creates a ListMongoDBAtlasIntegrationsOK with default headers values
func NewListMongoDBAtlasIntegrationsOK() *ListMongoDBAtlasIntegrationsOK {
	return &ListMongoDBAtlasIntegrationsOK{}
}

/*
ListMongoDBAtlasIntegrationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListMongoDBAtlasIntegrationsOK struct {
	Payload *models.Secrets20231128ListMongoDBAtlasIntegrationsResponse
}

// IsSuccess returns true when this list mongo d b atlas integrations o k response has a 2xx status code
func (o *ListMongoDBAtlasIntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list mongo d b atlas integrations o k response has a 3xx status code
func (o *ListMongoDBAtlasIntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mongo d b atlas integrations o k response has a 4xx status code
func (o *ListMongoDBAtlasIntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list mongo d b atlas integrations o k response has a 5xx status code
func (o *ListMongoDBAtlasIntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list mongo d b atlas integrations o k response a status code equal to that given
func (o *ListMongoDBAtlasIntegrationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list mongo d b atlas integrations o k response
func (o *ListMongoDBAtlasIntegrationsOK) Code() int {
	return 200
}

func (o *ListMongoDBAtlasIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/mongodb-atlas/config][%d] listMongoDBAtlasIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListMongoDBAtlasIntegrationsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/mongodb-atlas/config][%d] listMongoDBAtlasIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListMongoDBAtlasIntegrationsOK) GetPayload() *models.Secrets20231128ListMongoDBAtlasIntegrationsResponse {
	return o.Payload
}

func (o *ListMongoDBAtlasIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListMongoDBAtlasIntegrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMongoDBAtlasIntegrationsDefault creates a ListMongoDBAtlasIntegrationsDefault with default headers values
func NewListMongoDBAtlasIntegrationsDefault(code int) *ListMongoDBAtlasIntegrationsDefault {
	return &ListMongoDBAtlasIntegrationsDefault{
		_statusCode: code,
	}
}

/*
ListMongoDBAtlasIntegrationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListMongoDBAtlasIntegrationsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list mongo d b atlas integrations default response has a 2xx status code
func (o *ListMongoDBAtlasIntegrationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list mongo d b atlas integrations default response has a 3xx status code
func (o *ListMongoDBAtlasIntegrationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list mongo d b atlas integrations default response has a 4xx status code
func (o *ListMongoDBAtlasIntegrationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list mongo d b atlas integrations default response has a 5xx status code
func (o *ListMongoDBAtlasIntegrationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list mongo d b atlas integrations default response a status code equal to that given
func (o *ListMongoDBAtlasIntegrationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list mongo d b atlas integrations default response
func (o *ListMongoDBAtlasIntegrationsDefault) Code() int {
	return o._statusCode
}

func (o *ListMongoDBAtlasIntegrationsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/mongodb-atlas/config][%d] ListMongoDBAtlasIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListMongoDBAtlasIntegrationsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/mongodb-atlas/config][%d] ListMongoDBAtlasIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListMongoDBAtlasIntegrationsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListMongoDBAtlasIntegrationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}