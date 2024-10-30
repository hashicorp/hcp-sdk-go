// Code generated by go-swagger; DO NOT EDIT.

package integration_connection_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// DeleteIntegrationConnectionReader is a Reader for the DeleteIntegrationConnection structure.
type DeleteIntegrationConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIntegrationConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteIntegrationConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIntegrationConnectionOK creates a DeleteIntegrationConnectionOK with default headers values
func NewDeleteIntegrationConnectionOK() *DeleteIntegrationConnectionOK {
	return &DeleteIntegrationConnectionOK{}
}

/*
DeleteIntegrationConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteIntegrationConnectionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete integration connection o k response has a 2xx status code
func (o *DeleteIntegrationConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete integration connection o k response has a 3xx status code
func (o *DeleteIntegrationConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration connection o k response has a 4xx status code
func (o *DeleteIntegrationConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integration connection o k response has a 5xx status code
func (o *DeleteIntegrationConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration connection o k response a status code equal to that given
func (o *DeleteIntegrationConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete integration connection o k response
func (o *DeleteIntegrationConnectionOK) Code() int {
	return 200
}

func (o *DeleteIntegrationConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections/{id}][%d] deleteIntegrationConnectionOK  %+v", 200, o.Payload)
}

func (o *DeleteIntegrationConnectionOK) String() string {
	return fmt.Sprintf("[DELETE /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections/{id}][%d] deleteIntegrationConnectionOK  %+v", 200, o.Payload)
}

func (o *DeleteIntegrationConnectionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteIntegrationConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationConnectionDefault creates a DeleteIntegrationConnectionDefault with default headers values
func NewDeleteIntegrationConnectionDefault(code int) *DeleteIntegrationConnectionDefault {
	return &DeleteIntegrationConnectionDefault{
		_statusCode: code,
	}
}

/*
DeleteIntegrationConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteIntegrationConnectionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this delete integration connection default response has a 2xx status code
func (o *DeleteIntegrationConnectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete integration connection default response has a 3xx status code
func (o *DeleteIntegrationConnectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete integration connection default response has a 4xx status code
func (o *DeleteIntegrationConnectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete integration connection default response has a 5xx status code
func (o *DeleteIntegrationConnectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete integration connection default response a status code equal to that given
func (o *DeleteIntegrationConnectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete integration connection default response
func (o *DeleteIntegrationConnectionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIntegrationConnectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections/{id}][%d] DeleteIntegrationConnection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIntegrationConnectionDefault) String() string {
	return fmt.Sprintf("[DELETE /2023-05-01/vault-radar/projects/{location.project_id}/integrations/connections/{id}][%d] DeleteIntegrationConnection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIntegrationConnectionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DeleteIntegrationConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
