// Code generated by go-swagger; DO NOT EDIT.

package auth_config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// DeleteAuthConnectionFromOrganizationReader is a Reader for the DeleteAuthConnectionFromOrganization structure.
type DeleteAuthConnectionFromOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthConnectionFromOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuthConnectionFromOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAuthConnectionFromOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuthConnectionFromOrganizationOK creates a DeleteAuthConnectionFromOrganizationOK with default headers values
func NewDeleteAuthConnectionFromOrganizationOK() *DeleteAuthConnectionFromOrganizationOK {
	return &DeleteAuthConnectionFromOrganizationOK{}
}

/*
DeleteAuthConnectionFromOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAuthConnectionFromOrganizationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete auth connection from organization o k response has a 2xx status code
func (o *DeleteAuthConnectionFromOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete auth connection from organization o k response has a 3xx status code
func (o *DeleteAuthConnectionFromOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth connection from organization o k response has a 4xx status code
func (o *DeleteAuthConnectionFromOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete auth connection from organization o k response has a 5xx status code
func (o *DeleteAuthConnectionFromOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth connection from organization o k response a status code equal to that given
func (o *DeleteAuthConnectionFromOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAuthConnectionFromOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] deleteAuthConnectionFromOrganizationOK  %+v", 200, o.Payload)
}

func (o *DeleteAuthConnectionFromOrganizationOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] deleteAuthConnectionFromOrganizationOK  %+v", 200, o.Payload)
}

func (o *DeleteAuthConnectionFromOrganizationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAuthConnectionFromOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthConnectionFromOrganizationDefault creates a DeleteAuthConnectionFromOrganizationDefault with default headers values
func NewDeleteAuthConnectionFromOrganizationDefault(code int) *DeleteAuthConnectionFromOrganizationDefault {
	return &DeleteAuthConnectionFromOrganizationDefault{
		_statusCode: code,
	}
}

/*
DeleteAuthConnectionFromOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAuthConnectionFromOrganizationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete auth connection from organization default response
func (o *DeleteAuthConnectionFromOrganizationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete auth connection from organization default response has a 2xx status code
func (o *DeleteAuthConnectionFromOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete auth connection from organization default response has a 3xx status code
func (o *DeleteAuthConnectionFromOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete auth connection from organization default response has a 4xx status code
func (o *DeleteAuthConnectionFromOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete auth connection from organization default response has a 5xx status code
func (o *DeleteAuthConnectionFromOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete auth connection from organization default response a status code equal to that given
func (o *DeleteAuthConnectionFromOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteAuthConnectionFromOrganizationDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] DeleteAuthConnectionFromOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAuthConnectionFromOrganizationDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/auth_connections][%d] DeleteAuthConnectionFromOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAuthConnectionFromOrganizationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteAuthConnectionFromOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
