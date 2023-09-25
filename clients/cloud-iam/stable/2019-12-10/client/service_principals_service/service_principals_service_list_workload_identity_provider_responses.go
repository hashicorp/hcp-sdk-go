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

// ServicePrincipalsServiceListWorkloadIdentityProviderReader is a Reader for the ServicePrincipalsServiceListWorkloadIdentityProvider structure.
type ServicePrincipalsServiceListWorkloadIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceListWorkloadIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceListWorkloadIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceListWorkloadIdentityProviderOK creates a ServicePrincipalsServiceListWorkloadIdentityProviderOK with default headers values
func NewServicePrincipalsServiceListWorkloadIdentityProviderOK() *ServicePrincipalsServiceListWorkloadIdentityProviderOK {
	return &ServicePrincipalsServiceListWorkloadIdentityProviderOK{}
}

/*
ServicePrincipalsServiceListWorkloadIdentityProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceListWorkloadIdentityProviderOK struct {
	Payload *models.HashicorpCloudIamListWorkloadIdentityProviderResponse
}

// IsSuccess returns true when this service principals service list workload identity provider o k response has a 2xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service list workload identity provider o k response has a 3xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service list workload identity provider o k response has a 4xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service list workload identity provider o k response has a 5xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service list workload identity provider o k response a status code equal to that given
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] servicePrincipalsServiceListWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] servicePrincipalsServiceListWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) GetPayload() *models.HashicorpCloudIamListWorkloadIdentityProviderResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListWorkloadIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceListWorkloadIdentityProviderDefault creates a ServicePrincipalsServiceListWorkloadIdentityProviderDefault with default headers values
func NewServicePrincipalsServiceListWorkloadIdentityProviderDefault(code int) *ServicePrincipalsServiceListWorkloadIdentityProviderDefault {
	return &ServicePrincipalsServiceListWorkloadIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceListWorkloadIdentityProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceListWorkloadIdentityProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service list workload identity provider default response
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service list workload identity provider default response has a 2xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service list workload identity provider default response has a 3xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service list workload identity provider default response has a 4xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service list workload identity provider default response has a 5xx status code
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service list workload identity provider default response a status code equal to that given
func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] ServicePrincipalsService_ListWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) String() string {
	return fmt.Sprintf("[GET /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] ServicePrincipalsService_ListWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceListWorkloadIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
