// Code generated by go-swagger; DO NOT EDIT.

package organization_service

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

// OrganizationServiceGetRoleReader is a Reader for the OrganizationServiceGetRole structure.
type OrganizationServiceGetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceGetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceGetRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceGetRoleOK creates a OrganizationServiceGetRoleOK with default headers values
func NewOrganizationServiceGetRoleOK() *OrganizationServiceGetRoleOK {
	return &OrganizationServiceGetRoleOK{}
}

/*
OrganizationServiceGetRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceGetRoleOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationGetRoleResponse
}

// IsSuccess returns true when this organization service get role o k response has a 2xx status code
func (o *OrganizationServiceGetRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service get role o k response has a 3xx status code
func (o *OrganizationServiceGetRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service get role o k response has a 4xx status code
func (o *OrganizationServiceGetRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service get role o k response has a 5xx status code
func (o *OrganizationServiceGetRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service get role o k response a status code equal to that given
func (o *OrganizationServiceGetRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationServiceGetRoleOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/{role_id}][%d] organizationServiceGetRoleOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceGetRoleOK) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/{role_id}][%d] organizationServiceGetRoleOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceGetRoleOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationGetRoleResponse {
	return o.Payload
}

func (o *OrganizationServiceGetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationGetRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceGetRoleDefault creates a OrganizationServiceGetRoleDefault with default headers values
func NewOrganizationServiceGetRoleDefault(code int) *OrganizationServiceGetRoleDefault {
	return &OrganizationServiceGetRoleDefault{
		_statusCode: code,
	}
}

/*
OrganizationServiceGetRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceGetRoleDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the organization service get role default response
func (o *OrganizationServiceGetRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this organization service get role default response has a 2xx status code
func (o *OrganizationServiceGetRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service get role default response has a 3xx status code
func (o *OrganizationServiceGetRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service get role default response has a 4xx status code
func (o *OrganizationServiceGetRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service get role default response has a 5xx status code
func (o *OrganizationServiceGetRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service get role default response a status code equal to that given
func (o *OrganizationServiceGetRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OrganizationServiceGetRoleDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/{role_id}][%d] OrganizationService_GetRole default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceGetRoleDefault) String() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/{role_id}][%d] OrganizationService_GetRole default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceGetRoleDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *OrganizationServiceGetRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
