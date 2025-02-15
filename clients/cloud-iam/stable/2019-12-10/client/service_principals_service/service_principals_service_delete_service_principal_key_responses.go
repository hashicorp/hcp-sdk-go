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

// ServicePrincipalsServiceDeleteServicePrincipalKeyReader is a Reader for the ServicePrincipalsServiceDeleteServicePrincipalKey structure.
type ServicePrincipalsServiceDeleteServicePrincipalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceDeleteServicePrincipalKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceDeleteServicePrincipalKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKeyOK creates a ServicePrincipalsServiceDeleteServicePrincipalKeyOK with default headers values
func NewServicePrincipalsServiceDeleteServicePrincipalKeyOK() *ServicePrincipalsServiceDeleteServicePrincipalKeyOK {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyOK{}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKeyOK struct {
	Payload models.HashicorpCloudIamDeleteServicePrincipalKeyResponse
}

// IsSuccess returns true when this service principals service delete service principal key o k response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service delete service principal key o k response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service delete service principal key o k response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service delete service principal key o k response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service delete service principal key o k response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service delete service principal key o k response
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_2}][%d] servicePrincipalsServiceDeleteServicePrincipalKeyOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_2}][%d] servicePrincipalsServiceDeleteServicePrincipalKeyOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) GetPayload() models.HashicorpCloudIamDeleteServicePrincipalKeyResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceDeleteServicePrincipalKeyDefault creates a ServicePrincipalsServiceDeleteServicePrincipalKeyDefault with default headers values
func NewServicePrincipalsServiceDeleteServicePrincipalKeyDefault(code int) *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKeyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service delete service principal key default response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service delete service principal key default response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service delete service principal key default response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service delete service principal key default response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service delete service principal key default response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service delete service principal key default response
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_2}][%d] ServicePrincipalsService_DeleteServicePrincipalKey default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_2}][%d] ServicePrincipalsService_DeleteServicePrincipalKey default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
