// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceGetIterationReader is a Reader for the PackerServiceGetIteration structure.
type PackerServiceGetIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetIterationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetIterationOK creates a PackerServiceGetIterationOK with default headers values
func NewPackerServiceGetIterationOK() *PackerServiceGetIterationOK {
	return &PackerServiceGetIterationOK{}
}

/*
PackerServiceGetIterationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceGetIterationOK struct {
	Payload *models.HashicorpCloudPackerGetIterationResponse
}

// IsSuccess returns true when this packer service get iteration o k response has a 2xx status code
func (o *PackerServiceGetIterationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service get iteration o k response has a 3xx status code
func (o *PackerServiceGetIterationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service get iteration o k response has a 4xx status code
func (o *PackerServiceGetIterationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service get iteration o k response has a 5xx status code
func (o *PackerServiceGetIterationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service get iteration o k response a status code equal to that given
func (o *PackerServiceGetIterationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service get iteration o k response
func (o *PackerServiceGetIterationOK) Code() int {
	return 200
}

func (o *PackerServiceGetIterationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iteration][%d] packerServiceGetIterationOK %s", 200, payload)
}

func (o *PackerServiceGetIterationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iteration][%d] packerServiceGetIterationOK %s", 200, payload)
}

func (o *PackerServiceGetIterationOK) GetPayload() *models.HashicorpCloudPackerGetIterationResponse {
	return o.Payload
}

func (o *PackerServiceGetIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetIterationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetIterationDefault creates a PackerServiceGetIterationDefault with default headers values
func NewPackerServiceGetIterationDefault(code int) *PackerServiceGetIterationDefault {
	return &PackerServiceGetIterationDefault{
		_statusCode: code,
	}
}

/*
PackerServiceGetIterationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceGetIterationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service get iteration default response has a 2xx status code
func (o *PackerServiceGetIterationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service get iteration default response has a 3xx status code
func (o *PackerServiceGetIterationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service get iteration default response has a 4xx status code
func (o *PackerServiceGetIterationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service get iteration default response has a 5xx status code
func (o *PackerServiceGetIterationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service get iteration default response a status code equal to that given
func (o *PackerServiceGetIterationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service get iteration default response
func (o *PackerServiceGetIterationDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceGetIterationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iteration][%d] PackerService_GetIteration default %s", o._statusCode, payload)
}

func (o *PackerServiceGetIterationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iteration][%d] PackerService_GetIteration default %s", o._statusCode, payload)
}

func (o *PackerServiceGetIterationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceGetIterationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
