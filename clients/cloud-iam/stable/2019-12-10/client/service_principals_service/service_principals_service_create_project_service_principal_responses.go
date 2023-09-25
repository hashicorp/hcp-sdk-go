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

// ServicePrincipalsServiceCreateProjectServicePrincipalReader is a Reader for the ServicePrincipalsServiceCreateProjectServicePrincipal structure.
type ServicePrincipalsServiceCreateProjectServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceCreateProjectServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceCreateProjectServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceCreateProjectServicePrincipalOK creates a ServicePrincipalsServiceCreateProjectServicePrincipalOK with default headers values
func NewServicePrincipalsServiceCreateProjectServicePrincipalOK() *ServicePrincipalsServiceCreateProjectServicePrincipalOK {
	return &ServicePrincipalsServiceCreateProjectServicePrincipalOK{}
}

/*
ServicePrincipalsServiceCreateProjectServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceCreateProjectServicePrincipalOK struct {
	Payload *models.HashicorpCloudIamCreateProjectServicePrincipalResponse
}

// IsSuccess returns true when this service principals service create project service principal o k response has a 2xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service create project service principal o k response has a 3xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service create project service principal o k response has a 4xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service create project service principal o k response has a 5xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service create project service principal o k response a status code equal to that given
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] servicePrincipalsServiceCreateProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] servicePrincipalsServiceCreateProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) GetPayload() *models.HashicorpCloudIamCreateProjectServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateProjectServicePrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceCreateProjectServicePrincipalDefault creates a ServicePrincipalsServiceCreateProjectServicePrincipalDefault with default headers values
func NewServicePrincipalsServiceCreateProjectServicePrincipalDefault(code int) *ServicePrincipalsServiceCreateProjectServicePrincipalDefault {
	return &ServicePrincipalsServiceCreateProjectServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceCreateProjectServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceCreateProjectServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service create project service principal default response
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service create project service principal default response has a 2xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service create project service principal default response has a 3xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service create project service principal default response has a 4xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service create project service principal default response has a 5xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service create project service principal default response a status code equal to that given
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] ServicePrincipalsService_CreateProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals][%d] ServicePrincipalsService_CreateProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
