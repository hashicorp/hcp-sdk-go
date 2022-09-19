// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2020-08-26/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateCustomerMasterACLTokenReader is a Reader for the CreateCustomerMasterACLToken structure.
type CreateCustomerMasterACLTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCustomerMasterACLTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCustomerMasterACLTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCustomerMasterACLTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCustomerMasterACLTokenOK creates a CreateCustomerMasterACLTokenOK with default headers values
func NewCreateCustomerMasterACLTokenOK() *CreateCustomerMasterACLTokenOK {
	return &CreateCustomerMasterACLTokenOK{}
}

/*
CreateCustomerMasterACLTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateCustomerMasterACLTokenOK struct {
	Payload *models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenResponse
}

// IsSuccess returns true when this create customer master Acl token o k response has a 2xx status code
func (o *CreateCustomerMasterACLTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create customer master Acl token o k response has a 3xx status code
func (o *CreateCustomerMasterACLTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create customer master Acl token o k response has a 4xx status code
func (o *CreateCustomerMasterACLTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create customer master Acl token o k response has a 5xx status code
func (o *CreateCustomerMasterACLTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create customer master Acl token o k response a status code equal to that given
func (o *CreateCustomerMasterACLTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateCustomerMasterACLTokenOK) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/master-acl-tokens][%d] createCustomerMasterAclTokenOK  %+v", 200, o.Payload)
}

func (o *CreateCustomerMasterACLTokenOK) String() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/master-acl-tokens][%d] createCustomerMasterAclTokenOK  %+v", 200, o.Payload)
}

func (o *CreateCustomerMasterACLTokenOK) GetPayload() *models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenResponse {
	return o.Payload
}

func (o *CreateCustomerMasterACLTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826CreateCustomerMasterACLTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomerMasterACLTokenDefault creates a CreateCustomerMasterACLTokenDefault with default headers values
func NewCreateCustomerMasterACLTokenDefault(code int) *CreateCustomerMasterACLTokenDefault {
	return &CreateCustomerMasterACLTokenDefault{
		_statusCode: code,
	}
}

/*
CreateCustomerMasterACLTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateCustomerMasterACLTokenDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create customer master ACL token default response
func (o *CreateCustomerMasterACLTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create customer master ACL token default response has a 2xx status code
func (o *CreateCustomerMasterACLTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create customer master ACL token default response has a 3xx status code
func (o *CreateCustomerMasterACLTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create customer master ACL token default response has a 4xx status code
func (o *CreateCustomerMasterACLTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create customer master ACL token default response has a 5xx status code
func (o *CreateCustomerMasterACLTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create customer master ACL token default response a status code equal to that given
func (o *CreateCustomerMasterACLTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateCustomerMasterACLTokenDefault) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/master-acl-tokens][%d] CreateCustomerMasterACLToken default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCustomerMasterACLTokenDefault) String() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/master-acl-tokens][%d] CreateCustomerMasterACLToken default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCustomerMasterACLTokenDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateCustomerMasterACLTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
