// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// OrganizationServiceListReader is a Reader for the OrganizationServiceList structure.
type OrganizationServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceListOK creates a OrganizationServiceListOK with default headers values
func NewOrganizationServiceListOK() *OrganizationServiceListOK {
	return &OrganizationServiceListOK{}
}

/*
OrganizationServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceListOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationListResponse
}

// IsSuccess returns true when this organization service list o k response has a 2xx status code
func (o *OrganizationServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service list o k response has a 3xx status code
func (o *OrganizationServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service list o k response has a 4xx status code
func (o *OrganizationServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service list o k response has a 5xx status code
func (o *OrganizationServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service list o k response a status code equal to that given
func (o *OrganizationServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationServiceListOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations][%d] organizationServiceListOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceListOK) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations][%d] organizationServiceListOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceListOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationListResponse {
	return o.Payload
}

func (o *OrganizationServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceListDefault creates a OrganizationServiceListDefault with default headers values
func NewOrganizationServiceListDefault(code int) *OrganizationServiceListDefault {
	return &OrganizationServiceListDefault{
		_statusCode: code,
	}
}

/*
OrganizationServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceListDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service list default response
func (o *OrganizationServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this organization service list default response has a 2xx status code
func (o *OrganizationServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service list default response has a 3xx status code
func (o *OrganizationServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service list default response has a 4xx status code
func (o *OrganizationServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service list default response has a 5xx status code
func (o *OrganizationServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service list default response a status code equal to that given
func (o *OrganizationServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OrganizationServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations][%d] OrganizationService_List default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceListDefault) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations][%d] OrganizationService_List default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceListDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
