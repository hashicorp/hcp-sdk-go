// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateIterationReader is a Reader for the CreateIteration structure.
type CreateIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIterationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIterationOK creates a CreateIterationOK with default headers values
func NewCreateIterationOK() *CreateIterationOK {
	return &CreateIterationOK{}
}

/* CreateIterationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateIterationOK struct {
	Payload *models.HashicorpCloudPackerCreateIterationResponse
}

func (o *CreateIterationOK) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] createIterationOK  %+v", 200, o.Payload)
}
func (o *CreateIterationOK) GetPayload() *models.HashicorpCloudPackerCreateIterationResponse {
	return o.Payload
}

func (o *CreateIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerCreateIterationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIterationDefault creates a CreateIterationDefault with default headers values
func NewCreateIterationDefault(code int) *CreateIterationDefault {
	return &CreateIterationDefault{
		_statusCode: code,
	}
}

/* CreateIterationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateIterationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create iteration default response
func (o *CreateIterationDefault) Code() int {
	return o._statusCode
}

func (o *CreateIterationDefault) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] CreateIteration default  %+v", o._statusCode, o.Payload)
}
func (o *CreateIterationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateIterationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
