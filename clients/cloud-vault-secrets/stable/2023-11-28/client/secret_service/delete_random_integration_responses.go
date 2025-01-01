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

// DeleteRandomIntegrationReader is a Reader for the DeleteRandomIntegration structure.
type DeleteRandomIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRandomIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRandomIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRandomIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRandomIntegrationOK creates a DeleteRandomIntegrationOK with default headers values
func NewDeleteRandomIntegrationOK() *DeleteRandomIntegrationOK {
	return &DeleteRandomIntegrationOK{}
}

/*
DeleteRandomIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRandomIntegrationOK struct {
	Payload models.Secrets20231128DeleteRandomIntegrationResponse
}

// IsSuccess returns true when this delete random integration o k response has a 2xx status code
func (o *DeleteRandomIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete random integration o k response has a 3xx status code
func (o *DeleteRandomIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete random integration o k response has a 4xx status code
func (o *DeleteRandomIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete random integration o k response has a 5xx status code
func (o *DeleteRandomIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete random integration o k response a status code equal to that given
func (o *DeleteRandomIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete random integration o k response
func (o *DeleteRandomIntegrationOK) Code() int {
	return 200
}

func (o *DeleteRandomIntegrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config/{name}][%d] deleteRandomIntegrationOK %s", 200, payload)
}

func (o *DeleteRandomIntegrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config/{name}][%d] deleteRandomIntegrationOK %s", 200, payload)
}

func (o *DeleteRandomIntegrationOK) GetPayload() models.Secrets20231128DeleteRandomIntegrationResponse {
	return o.Payload
}

func (o *DeleteRandomIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRandomIntegrationDefault creates a DeleteRandomIntegrationDefault with default headers values
func NewDeleteRandomIntegrationDefault(code int) *DeleteRandomIntegrationDefault {
	return &DeleteRandomIntegrationDefault{
		_statusCode: code,
	}
}

/*
DeleteRandomIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteRandomIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete random integration default response has a 2xx status code
func (o *DeleteRandomIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete random integration default response has a 3xx status code
func (o *DeleteRandomIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete random integration default response has a 4xx status code
func (o *DeleteRandomIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete random integration default response has a 5xx status code
func (o *DeleteRandomIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete random integration default response a status code equal to that given
func (o *DeleteRandomIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete random integration default response
func (o *DeleteRandomIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRandomIntegrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config/{name}][%d] DeleteRandomIntegration default %s", o._statusCode, payload)
}

func (o *DeleteRandomIntegrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/random/config/{name}][%d] DeleteRandomIntegration default %s", o._statusCode, payload)
}

func (o *DeleteRandomIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteRandomIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}