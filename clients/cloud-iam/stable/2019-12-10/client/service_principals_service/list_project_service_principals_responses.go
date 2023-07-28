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

// ListProjectServicePrincipalsReader is a Reader for the ListProjectServicePrincipals structure.
type ListProjectServicePrincipalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectServicePrincipalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectServicePrincipalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectServicePrincipalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectServicePrincipalsOK creates a ListProjectServicePrincipalsOK with default headers values
func NewListProjectServicePrincipalsOK() *ListProjectServicePrincipalsOK {
	return &ListProjectServicePrincipalsOK{}
}

/*
ListProjectServicePrincipalsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectServicePrincipalsOK struct {
	Payload *models.HashicorpCloudIamListProjectServicePrincipalsResponse
}

// IsSuccess returns true when this list project service principals o k response has a 2xx status code
func (o *ListProjectServicePrincipalsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project service principals o k response has a 3xx status code
func (o *ListProjectServicePrincipalsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service principals o k response has a 4xx status code
func (o *ListProjectServicePrincipalsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project service principals o k response has a 5xx status code
func (o *ListProjectServicePrincipalsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service principals o k response a status code equal to that given
func (o *ListProjectServicePrincipalsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectServicePrincipalsOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] listProjectServicePrincipalsOK  %+v", 200, o.Payload)
}

func (o *ListProjectServicePrincipalsOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] listProjectServicePrincipalsOK  %+v", 200, o.Payload)
}

func (o *ListProjectServicePrincipalsOK) GetPayload() *models.HashicorpCloudIamListProjectServicePrincipalsResponse {
	return o.Payload
}

func (o *ListProjectServicePrincipalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListProjectServicePrincipalsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectServicePrincipalsDefault creates a ListProjectServicePrincipalsDefault with default headers values
func NewListProjectServicePrincipalsDefault(code int) *ListProjectServicePrincipalsDefault {
	return &ListProjectServicePrincipalsDefault{
		_statusCode: code,
	}
}

/*
ListProjectServicePrincipalsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectServicePrincipalsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the list project service principals default response
func (o *ListProjectServicePrincipalsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project service principals default response has a 2xx status code
func (o *ListProjectServicePrincipalsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project service principals default response has a 3xx status code
func (o *ListProjectServicePrincipalsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project service principals default response has a 4xx status code
func (o *ListProjectServicePrincipalsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project service principals default response has a 5xx status code
func (o *ListProjectServicePrincipalsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project service principals default response a status code equal to that given
func (o *ListProjectServicePrincipalsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectServicePrincipalsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] ListProjectServicePrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectServicePrincipalsDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] ListProjectServicePrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectServicePrincipalsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListProjectServicePrincipalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
