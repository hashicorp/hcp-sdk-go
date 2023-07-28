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

// CreateProjectServicePrincipalReader is a Reader for the CreateProjectServicePrincipal structure.
type CreateProjectServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProjectServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectServicePrincipalOK creates a CreateProjectServicePrincipalOK with default headers values
func NewCreateProjectServicePrincipalOK() *CreateProjectServicePrincipalOK {
	return &CreateProjectServicePrincipalOK{}
}

/*
CreateProjectServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateProjectServicePrincipalOK struct {
	Payload *models.HashicorpCloudIamCreateProjectServicePrincipalResponse
}

// IsSuccess returns true when this create project service principal o k response has a 2xx status code
func (o *CreateProjectServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create project service principal o k response has a 3xx status code
func (o *CreateProjectServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project service principal o k response has a 4xx status code
func (o *CreateProjectServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project service principal o k response has a 5xx status code
func (o *CreateProjectServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create project service principal o k response a status code equal to that given
func (o *CreateProjectServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateProjectServicePrincipalOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] createProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *CreateProjectServicePrincipalOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] createProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *CreateProjectServicePrincipalOK) GetPayload() *models.HashicorpCloudIamCreateProjectServicePrincipalResponse {
	return o.Payload
}

func (o *CreateProjectServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateProjectServicePrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectServicePrincipalDefault creates a CreateProjectServicePrincipalDefault with default headers values
func NewCreateProjectServicePrincipalDefault(code int) *CreateProjectServicePrincipalDefault {
	return &CreateProjectServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
CreateProjectServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateProjectServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the create project service principal default response
func (o *CreateProjectServicePrincipalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create project service principal default response has a 2xx status code
func (o *CreateProjectServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create project service principal default response has a 3xx status code
func (o *CreateProjectServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create project service principal default response has a 4xx status code
func (o *CreateProjectServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create project service principal default response has a 5xx status code
func (o *CreateProjectServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create project service principal default response a status code equal to that given
func (o *CreateProjectServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateProjectServicePrincipalDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] CreateProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectServicePrincipalDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] CreateProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CreateProjectServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
