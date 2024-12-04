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

// ListPostgresIntegrationsReader is a Reader for the ListPostgresIntegrations structure.
type ListPostgresIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPostgresIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPostgresIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPostgresIntegrationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPostgresIntegrationsOK creates a ListPostgresIntegrationsOK with default headers values
func NewListPostgresIntegrationsOK() *ListPostgresIntegrationsOK {
	return &ListPostgresIntegrationsOK{}
}

/*
ListPostgresIntegrationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPostgresIntegrationsOK struct {
	Payload *models.Secrets20231128ListPostgresIntegrationsResponse
}

// IsSuccess returns true when this list postgres integrations o k response has a 2xx status code
func (o *ListPostgresIntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list postgres integrations o k response has a 3xx status code
func (o *ListPostgresIntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list postgres integrations o k response has a 4xx status code
func (o *ListPostgresIntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list postgres integrations o k response has a 5xx status code
func (o *ListPostgresIntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list postgres integrations o k response a status code equal to that given
func (o *ListPostgresIntegrationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list postgres integrations o k response
func (o *ListPostgresIntegrationsOK) Code() int {
	return 200
}

func (o *ListPostgresIntegrationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/postgres/config][%d] listPostgresIntegrationsOK %s", 200, payload)
}

func (o *ListPostgresIntegrationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/postgres/config][%d] listPostgresIntegrationsOK %s", 200, payload)
}

func (o *ListPostgresIntegrationsOK) GetPayload() *models.Secrets20231128ListPostgresIntegrationsResponse {
	return o.Payload
}

func (o *ListPostgresIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListPostgresIntegrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostgresIntegrationsDefault creates a ListPostgresIntegrationsDefault with default headers values
func NewListPostgresIntegrationsDefault(code int) *ListPostgresIntegrationsDefault {
	return &ListPostgresIntegrationsDefault{
		_statusCode: code,
	}
}

/*
ListPostgresIntegrationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListPostgresIntegrationsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list postgres integrations default response has a 2xx status code
func (o *ListPostgresIntegrationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list postgres integrations default response has a 3xx status code
func (o *ListPostgresIntegrationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list postgres integrations default response has a 4xx status code
func (o *ListPostgresIntegrationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list postgres integrations default response has a 5xx status code
func (o *ListPostgresIntegrationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list postgres integrations default response a status code equal to that given
func (o *ListPostgresIntegrationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list postgres integrations default response
func (o *ListPostgresIntegrationsDefault) Code() int {
	return o._statusCode
}

func (o *ListPostgresIntegrationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/postgres/config][%d] ListPostgresIntegrations default %s", o._statusCode, payload)
}

func (o *ListPostgresIntegrationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/postgres/config][%d] ListPostgresIntegrations default %s", o._statusCode, payload)
}

func (o *ListPostgresIntegrationsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListPostgresIntegrationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
