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

// OrganizationServiceListRolesReader is a Reader for the OrganizationServiceListRoles structure.
type OrganizationServiceListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceListRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceListRolesOK creates a OrganizationServiceListRolesOK with default headers values
func NewOrganizationServiceListRolesOK() *OrganizationServiceListRolesOK {
	return &OrganizationServiceListRolesOK{}
}

/*
OrganizationServiceListRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceListRolesOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationListRolesResponse
}

// IsSuccess returns true when this organization service list roles o k response has a 2xx status code
func (o *OrganizationServiceListRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service list roles o k response has a 3xx status code
func (o *OrganizationServiceListRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service list roles o k response has a 4xx status code
func (o *OrganizationServiceListRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service list roles o k response has a 5xx status code
func (o *OrganizationServiceListRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service list roles o k response a status code equal to that given
func (o *OrganizationServiceListRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationServiceListRolesOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/roles][%d] organizationServiceListRolesOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceListRolesOK) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/roles][%d] organizationServiceListRolesOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceListRolesOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationListRolesResponse {
	return o.Payload
}

func (o *OrganizationServiceListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationListRolesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceListRolesDefault creates a OrganizationServiceListRolesDefault with default headers values
func NewOrganizationServiceListRolesDefault(code int) *OrganizationServiceListRolesDefault {
	return &OrganizationServiceListRolesDefault{
		_statusCode: code,
	}
}

/*
OrganizationServiceListRolesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceListRolesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service list roles default response
func (o *OrganizationServiceListRolesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this organization service list roles default response has a 2xx status code
func (o *OrganizationServiceListRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service list roles default response has a 3xx status code
func (o *OrganizationServiceListRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service list roles default response has a 4xx status code
func (o *OrganizationServiceListRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service list roles default response has a 5xx status code
func (o *OrganizationServiceListRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service list roles default response a status code equal to that given
func (o *OrganizationServiceListRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OrganizationServiceListRolesDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/roles][%d] OrganizationService_ListRoles default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceListRolesDefault) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/roles][%d] OrganizationService_ListRoles default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceListRolesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceListRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
