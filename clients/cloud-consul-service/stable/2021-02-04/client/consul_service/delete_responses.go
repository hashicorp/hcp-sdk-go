// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOK creates a DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*
DeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOK struct {
	Payload *models.HashicorpCloudConsul20210204DeleteResponse
}

// IsSuccess returns true when this delete o k response has a 2xx status code
func (o *DeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete o k response has a 3xx status code
func (o *DeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete o k response has a 4xx status code
func (o *DeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete o k response has a 5xx status code
func (o *DeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete o k response a status code equal to that given
func (o *DeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) String() string {
	return fmt.Sprintf("[DELETE /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) GetPayload() *models.HashicorpCloudConsul20210204DeleteResponse {
	return o.Payload
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204DeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDefault creates a DeleteDefault with default headers values
func NewDeleteDefault(code int) *DeleteDefault {
	return &DeleteDefault{
		_statusCode: code,
	}
}

/*
DeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete default response
func (o *DeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete default response has a 2xx status code
func (o *DeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete default response has a 3xx status code
func (o *DeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete default response has a 4xx status code
func (o *DeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete default response has a 5xx status code
func (o *DeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete default response a status code equal to that given
func (o *DeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] Delete default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /consul/2021-02-04/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}][%d] Delete default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
