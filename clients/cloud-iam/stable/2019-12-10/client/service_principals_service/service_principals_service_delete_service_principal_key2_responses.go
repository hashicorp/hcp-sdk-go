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

// ServicePrincipalsServiceDeleteServicePrincipalKey2Reader is a Reader for the ServicePrincipalsServiceDeleteServicePrincipalKey2 structure.
type ServicePrincipalsServiceDeleteServicePrincipalKey2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceDeleteServicePrincipalKey2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceDeleteServicePrincipalKey2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKey2OK creates a ServicePrincipalsServiceDeleteServicePrincipalKey2OK with default headers values
func NewServicePrincipalsServiceDeleteServicePrincipalKey2OK() *ServicePrincipalsServiceDeleteServicePrincipalKey2OK {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2OK{}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKey2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKey2OK struct {
	Payload models.HashicorpCloudIamDeleteServicePrincipalKeyResponse
}

// IsSuccess returns true when this service principals service delete service principal key2 o k response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service delete service principal key2 o k response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service delete service principal key2 o k response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service delete service principal key2 o k response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service delete service principal key2 o k response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service delete service principal key2 o k response
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_3}][%d] servicePrincipalsServiceDeleteServicePrincipalKey2OK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_3}][%d] servicePrincipalsServiceDeleteServicePrincipalKey2OK %s", 200, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) GetPayload() models.HashicorpCloudIamDeleteServicePrincipalKeyResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceDeleteServicePrincipalKey2Default creates a ServicePrincipalsServiceDeleteServicePrincipalKey2Default with default headers values
func NewServicePrincipalsServiceDeleteServicePrincipalKey2Default(code int) *ServicePrincipalsServiceDeleteServicePrincipalKey2Default {
	return &ServicePrincipalsServiceDeleteServicePrincipalKey2Default{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKey2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKey2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service delete service principal key2 default response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service delete service principal key2 default response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service delete service principal key2 default response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service delete service principal key2 default response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service delete service principal key2 default response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service delete service principal key2 default response
func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_3}][%d] ServicePrincipalsService_DeleteServicePrincipalKey2 default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /2019-12-10/{resource_name_3}][%d] ServicePrincipalsService_DeleteServicePrincipalKey2 default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteServicePrincipalKey2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
