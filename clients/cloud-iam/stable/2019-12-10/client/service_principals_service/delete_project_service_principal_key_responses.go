// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// DeleteProjectServicePrincipalKeyReader is a Reader for the DeleteProjectServicePrincipalKey structure.
type DeleteProjectServicePrincipalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectServicePrincipalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectServicePrincipalKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProjectServicePrincipalKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectServicePrincipalKeyOK creates a DeleteProjectServicePrincipalKeyOK with default headers values
func NewDeleteProjectServicePrincipalKeyOK() *DeleteProjectServicePrincipalKeyOK {
	return &DeleteProjectServicePrincipalKeyOK{}
}

/*
DeleteProjectServicePrincipalKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProjectServicePrincipalKeyOK struct {
	Payload models.HashicorpCloudIamDeleteProjectServicePrincipalKeyResponse
}

// IsSuccess returns true when this delete project service principal key o k response has a 2xx status code
func (o *DeleteProjectServicePrincipalKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project service principal key o k response has a 3xx status code
func (o *DeleteProjectServicePrincipalKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project service principal key o k response has a 4xx status code
func (o *DeleteProjectServicePrincipalKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project service principal key o k response has a 5xx status code
func (o *DeleteProjectServicePrincipalKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project service principal key o k response a status code equal to that given
func (o *DeleteProjectServicePrincipalKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteProjectServicePrincipalKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys/{client_id}][%d] deleteProjectServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectServicePrincipalKeyOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys/{client_id}][%d] deleteProjectServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectServicePrincipalKeyOK) GetPayload() models.HashicorpCloudIamDeleteProjectServicePrincipalKeyResponse {
	return o.Payload
}

func (o *DeleteProjectServicePrincipalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectServicePrincipalKeyDefault creates a DeleteProjectServicePrincipalKeyDefault with default headers values
func NewDeleteProjectServicePrincipalKeyDefault(code int) *DeleteProjectServicePrincipalKeyDefault {
	return &DeleteProjectServicePrincipalKeyDefault{
		_statusCode: code,
	}
}

/*
DeleteProjectServicePrincipalKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProjectServicePrincipalKeyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete project service principal key default response
func (o *DeleteProjectServicePrincipalKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete project service principal key default response has a 2xx status code
func (o *DeleteProjectServicePrincipalKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project service principal key default response has a 3xx status code
func (o *DeleteProjectServicePrincipalKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project service principal key default response has a 4xx status code
func (o *DeleteProjectServicePrincipalKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project service principal key default response has a 5xx status code
func (o *DeleteProjectServicePrincipalKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project service principal key default response a status code equal to that given
func (o *DeleteProjectServicePrincipalKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteProjectServicePrincipalKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys/{client_id}][%d] DeleteProjectServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectServicePrincipalKeyDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys/{client_id}][%d] DeleteProjectServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectServicePrincipalKeyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteProjectServicePrincipalKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
