// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateIterationReader is a Reader for the PackerServiceUpdateIteration structure.
type PackerServiceUpdateIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateIterationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateIterationOK creates a PackerServiceUpdateIterationOK with default headers values
func NewPackerServiceUpdateIterationOK() *PackerServiceUpdateIterationOK {
	return &PackerServiceUpdateIterationOK{}
}

/*
PackerServiceUpdateIterationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceUpdateIterationOK struct {
	Payload *models.HashicorpCloudPacker20220411UpdateIterationResponse
}

// IsSuccess returns true when this packer service update iteration o k response has a 2xx status code
func (o *PackerServiceUpdateIterationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service update iteration o k response has a 3xx status code
func (o *PackerServiceUpdateIterationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service update iteration o k response has a 4xx status code
func (o *PackerServiceUpdateIterationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service update iteration o k response has a 5xx status code
func (o *PackerServiceUpdateIterationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service update iteration o k response a status code equal to that given
func (o *PackerServiceUpdateIterationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceUpdateIterationOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] packerServiceUpdateIterationOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateIterationOK) String() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] packerServiceUpdateIterationOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateIterationOK) GetPayload() *models.HashicorpCloudPacker20220411UpdateIterationResponse {
	return o.Payload
}

func (o *PackerServiceUpdateIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20220411UpdateIterationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateIterationDefault creates a PackerServiceUpdateIterationDefault with default headers values
func NewPackerServiceUpdateIterationDefault(code int) *PackerServiceUpdateIterationDefault {
	return &PackerServiceUpdateIterationDefault{
		_statusCode: code,
	}
}

/*
PackerServiceUpdateIterationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceUpdateIterationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service update iteration default response
func (o *PackerServiceUpdateIterationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service update iteration default response has a 2xx status code
func (o *PackerServiceUpdateIterationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service update iteration default response has a 3xx status code
func (o *PackerServiceUpdateIterationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service update iteration default response has a 4xx status code
func (o *PackerServiceUpdateIterationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service update iteration default response has a 5xx status code
func (o *PackerServiceUpdateIterationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service update iteration default response a status code equal to that given
func (o *PackerServiceUpdateIterationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceUpdateIterationDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] PackerService_UpdateIteration default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateIterationDefault) String() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] PackerService_UpdateIteration default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateIterationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceUpdateIterationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
