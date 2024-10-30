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

// DeleteSyncReader is a Reader for the DeleteSync structure.
type DeleteSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSyncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyncOK creates a DeleteSyncOK with default headers values
func NewDeleteSyncOK() *DeleteSyncOK {
	return &DeleteSyncOK{}
}

/*
DeleteSyncOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteSyncOK struct {
	Payload models.Secrets20231128DeleteSyncResponse
}

// IsSuccess returns true when this delete sync o k response has a 2xx status code
func (o *DeleteSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete sync o k response has a 3xx status code
func (o *DeleteSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sync o k response has a 4xx status code
func (o *DeleteSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sync o k response has a 5xx status code
func (o *DeleteSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sync o k response a status code equal to that given
func (o *DeleteSyncOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete sync o k response
func (o *DeleteSyncOK) Code() int {
	return 200
}

func (o *DeleteSyncOK) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/syncs/{name}][%d] deleteSyncOK  %+v", 200, o.Payload)
}

func (o *DeleteSyncOK) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/syncs/{name}][%d] deleteSyncOK  %+v", 200, o.Payload)
}

func (o *DeleteSyncOK) GetPayload() models.Secrets20231128DeleteSyncResponse {
	return o.Payload
}

func (o *DeleteSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSyncDefault creates a DeleteSyncDefault with default headers values
func NewDeleteSyncDefault(code int) *DeleteSyncDefault {
	return &DeleteSyncDefault{
		_statusCode: code,
	}
}

/*
DeleteSyncDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteSyncDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete sync default response has a 2xx status code
func (o *DeleteSyncDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete sync default response has a 3xx status code
func (o *DeleteSyncDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete sync default response has a 4xx status code
func (o *DeleteSyncDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete sync default response has a 5xx status code
func (o *DeleteSyncDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete sync default response a status code equal to that given
func (o *DeleteSyncDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete sync default response
func (o *DeleteSyncDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyncDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/syncs/{name}][%d] DeleteSync default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncDefault) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/syncs/{name}][%d] DeleteSync default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteSyncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
