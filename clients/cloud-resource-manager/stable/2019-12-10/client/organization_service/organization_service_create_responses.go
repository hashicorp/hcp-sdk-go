// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// OrganizationServiceCreateReader is a Reader for the OrganizationServiceCreate structure.
type OrganizationServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceCreateOK creates a OrganizationServiceCreateOK with default headers values
func NewOrganizationServiceCreateOK() *OrganizationServiceCreateOK {
	return &OrganizationServiceCreateOK{}
}

/*
OrganizationServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceCreateOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationCreateResponse
}

// IsSuccess returns true when this organization service create o k response has a 2xx status code
func (o *OrganizationServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service create o k response has a 3xx status code
func (o *OrganizationServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service create o k response has a 4xx status code
func (o *OrganizationServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service create o k response has a 5xx status code
func (o *OrganizationServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service create o k response a status code equal to that given
func (o *OrganizationServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organization service create o k response
func (o *OrganizationServiceCreateOK) Code() int {
	return 200
}

func (o *OrganizationServiceCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations][%d] organizationServiceCreateOK %s", 200, payload)
}

func (o *OrganizationServiceCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations][%d] organizationServiceCreateOK %s", 200, payload)
}

func (o *OrganizationServiceCreateOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationCreateResponse {
	return o.Payload
}

func (o *OrganizationServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceCreateDefault creates a OrganizationServiceCreateDefault with default headers values
func NewOrganizationServiceCreateDefault(code int) *OrganizationServiceCreateDefault {
	return &OrganizationServiceCreateDefault{
		_statusCode: code,
	}
}

/*
OrganizationServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceCreateDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this organization service create default response has a 2xx status code
func (o *OrganizationServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service create default response has a 3xx status code
func (o *OrganizationServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service create default response has a 4xx status code
func (o *OrganizationServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service create default response has a 5xx status code
func (o *OrganizationServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service create default response a status code equal to that given
func (o *OrganizationServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the organization service create default response
func (o *OrganizationServiceCreateDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations][%d] OrganizationService_Create default %s", o._statusCode, payload)
}

func (o *OrganizationServiceCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations][%d] OrganizationService_Create default %s", o._statusCode, payload)
}

func (o *OrganizationServiceCreateDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *OrganizationServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
