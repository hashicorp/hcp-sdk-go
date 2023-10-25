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

// ServicePrincipalsServiceListOrganizationServicePrincipalsReader is a Reader for the ServicePrincipalsServiceListOrganizationServicePrincipals structure.
type ServicePrincipalsServiceListOrganizationServicePrincipalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceListOrganizationServicePrincipalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceListOrganizationServicePrincipalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceListOrganizationServicePrincipalsOK creates a ServicePrincipalsServiceListOrganizationServicePrincipalsOK with default headers values
func NewServicePrincipalsServiceListOrganizationServicePrincipalsOK() *ServicePrincipalsServiceListOrganizationServicePrincipalsOK {
	return &ServicePrincipalsServiceListOrganizationServicePrincipalsOK{}
}

/*
ServicePrincipalsServiceListOrganizationServicePrincipalsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceListOrganizationServicePrincipalsOK struct {
	Payload *models.HashicorpCloudIamListOrganizationServicePrincipalsResponse
}

// IsSuccess returns true when this service principals service list organization service principals o k response has a 2xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service list organization service principals o k response has a 3xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service list organization service principals o k response has a 4xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service list organization service principals o k response has a 5xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service list organization service principals o k response a status code equal to that given
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service list organization service principals o k response
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/service-principals][%d] servicePrincipalsServiceListOrganizationServicePrincipalsOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/service-principals][%d] servicePrincipalsServiceListOrganizationServicePrincipalsOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) GetPayload() *models.HashicorpCloudIamListOrganizationServicePrincipalsResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListOrganizationServicePrincipalsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceListOrganizationServicePrincipalsDefault creates a ServicePrincipalsServiceListOrganizationServicePrincipalsDefault with default headers values
func NewServicePrincipalsServiceListOrganizationServicePrincipalsDefault(code int) *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault {
	return &ServicePrincipalsServiceListOrganizationServicePrincipalsDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceListOrganizationServicePrincipalsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceListOrganizationServicePrincipalsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service list organization service principals default response has a 2xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service list organization service principals default response has a 3xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service list organization service principals default response has a 4xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service list organization service principals default response has a 5xx status code
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service list organization service principals default response a status code equal to that given
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service list organization service principals default response
func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/service-principals][%d] ServicePrincipalsService_ListOrganizationServicePrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/service-principals][%d] ServicePrincipalsService_ListOrganizationServicePrincipals default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceListOrganizationServicePrincipalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}