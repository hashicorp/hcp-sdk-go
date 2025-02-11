// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceDeleteProjectServicePrincipalReader is a Reader for the ServicePrincipalsServiceDeleteProjectServicePrincipal structure.
type ServicePrincipalsServiceDeleteProjectServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceDeleteProjectServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceDeleteProjectServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceDeleteProjectServicePrincipalOK creates a ServicePrincipalsServiceDeleteProjectServicePrincipalOK with default headers values
func NewServicePrincipalsServiceDeleteProjectServicePrincipalOK() *ServicePrincipalsServiceDeleteProjectServicePrincipalOK {
	return &ServicePrincipalsServiceDeleteProjectServicePrincipalOK{}
}

/*
ServicePrincipalsServiceDeleteProjectServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceDeleteProjectServicePrincipalOK struct {
	Payload models.HashicorpCloudIamDeleteProjectServicePrincipalResponse
}

// IsSuccess returns true when this service principals service delete project service principal o k response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service delete project service principal o k response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service delete project service principal o k response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service delete project service principal o k response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service delete project service principal o k response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service delete project service principal o k response
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceDeleteProjectServicePrincipalOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceDeleteProjectServicePrincipalOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) GetPayload() models.HashicorpCloudIamDeleteProjectServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceDeleteProjectServicePrincipalDefault creates a ServicePrincipalsServiceDeleteProjectServicePrincipalDefault with default headers values
func NewServicePrincipalsServiceDeleteProjectServicePrincipalDefault(code int) *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault {
	return &ServicePrincipalsServiceDeleteProjectServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceDeleteProjectServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceDeleteProjectServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service delete project service principal default response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service delete project service principal default response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service delete project service principal default response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service delete project service principal default response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service delete project service principal default response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service delete project service principal default response
func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_DeleteProjectServicePrincipal default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_DeleteProjectServicePrincipal default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteProjectServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
