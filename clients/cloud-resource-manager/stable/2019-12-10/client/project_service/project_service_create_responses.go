// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ProjectServiceCreateReader is a Reader for the ProjectServiceCreate structure.
type ProjectServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceCreateOK creates a ProjectServiceCreateOK with default headers values
func NewProjectServiceCreateOK() *ProjectServiceCreateOK {
	return &ProjectServiceCreateOK{}
}

/*
ProjectServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceCreateOK struct {
	Payload *models.HashicorpCloudResourcemanagerProjectCreateResponse
}

// IsSuccess returns true when this project service create o k response has a 2xx status code
func (o *ProjectServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project service create o k response has a 3xx status code
func (o *ProjectServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project service create o k response has a 4xx status code
func (o *ProjectServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project service create o k response has a 5xx status code
func (o *ProjectServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project service create o k response a status code equal to that given
func (o *ProjectServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project service create o k response
func (o *ProjectServiceCreateOK) Code() int {
	return 200
}

func (o *ProjectServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/projects][%d] projectServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/projects][%d] projectServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceCreateOK) GetPayload() *models.HashicorpCloudResourcemanagerProjectCreateResponse {
	return o.Payload
}

func (o *ProjectServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerProjectCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceCreateDefault creates a ProjectServiceCreateDefault with default headers values
func NewProjectServiceCreateDefault(code int) *ProjectServiceCreateDefault {
	return &ProjectServiceCreateDefault{
		_statusCode: code,
	}
}

/*
ProjectServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceCreateDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this project service create default response has a 2xx status code
func (o *ProjectServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this project service create default response has a 3xx status code
func (o *ProjectServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this project service create default response has a 4xx status code
func (o *ProjectServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this project service create default response has a 5xx status code
func (o *ProjectServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this project service create default response a status code equal to that given
func (o *ProjectServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the project service create default response
func (o *ProjectServiceCreateDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/projects][%d] ProjectService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/projects][%d] ProjectService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceCreateDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ProjectServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
