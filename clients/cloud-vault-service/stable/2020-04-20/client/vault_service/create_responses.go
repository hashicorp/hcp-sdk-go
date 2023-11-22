// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-04-20/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOK creates a CreateOK with default headers values
func NewCreateOK() *CreateOK {
	return &CreateOK{}
}

/*
CreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOK struct {
	Payload *models.HashicorpCloudVault20200420CreateResponse
}

// IsSuccess returns true when this create o k response has a 2xx status code
func (o *CreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o k response has a 3xx status code
func (o *CreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o k response has a 4xx status code
func (o *CreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o k response has a 5xx status code
func (o *CreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create o k response a status code equal to that given
func (o *CreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create o k response
func (o *CreateOK) Code() int {
	return 200
}

func (o *CreateOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] createOK  %+v", 200, o.Payload)
}

func (o *CreateOK) String() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] createOK  %+v", 200, o.Payload)
}

func (o *CreateOK) GetPayload() *models.HashicorpCloudVault20200420CreateResponse {
	return o.Payload
}

func (o *CreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefault creates a CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	return &CreateDefault{
		_statusCode: code,
	}
}

/*
CreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this create default response has a 2xx status code
func (o *CreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create default response has a 3xx status code
func (o *CreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create default response has a 4xx status code
func (o *CreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create default response has a 5xx status code
func (o *CreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create default response a status code equal to that given
func (o *CreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create default response
func (o *CreateDefault) Code() int {
	return o._statusCode
}

func (o *CreateDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] Create default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDefault) String() string {
	return fmt.Sprintf("[POST /vault/2020-04-20/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] Create default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
