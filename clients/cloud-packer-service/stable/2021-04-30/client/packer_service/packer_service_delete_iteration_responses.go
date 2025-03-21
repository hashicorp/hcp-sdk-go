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

// PackerServiceDeleteIterationReader is a Reader for the PackerServiceDeleteIteration structure.
type PackerServiceDeleteIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceDeleteIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceDeleteIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceDeleteIterationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceDeleteIterationOK creates a PackerServiceDeleteIterationOK with default headers values
func NewPackerServiceDeleteIterationOK() *PackerServiceDeleteIterationOK {
	return &PackerServiceDeleteIterationOK{}
}

/*
PackerServiceDeleteIterationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceDeleteIterationOK struct {
	Payload models.HashicorpCloudPackerDeleteIterationResponse
}

// IsSuccess returns true when this packer service delete iteration o k response has a 2xx status code
func (o *PackerServiceDeleteIterationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service delete iteration o k response has a 3xx status code
func (o *PackerServiceDeleteIterationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service delete iteration o k response has a 4xx status code
func (o *PackerServiceDeleteIterationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service delete iteration o k response has a 5xx status code
func (o *PackerServiceDeleteIterationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service delete iteration o k response a status code equal to that given
func (o *PackerServiceDeleteIterationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service delete iteration o k response
func (o *PackerServiceDeleteIterationOK) Code() int {
	return 200
}

func (o *PackerServiceDeleteIterationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] packerServiceDeleteIterationOK %s", 200, payload)
}

func (o *PackerServiceDeleteIterationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] packerServiceDeleteIterationOK %s", 200, payload)
}

func (o *PackerServiceDeleteIterationOK) GetPayload() models.HashicorpCloudPackerDeleteIterationResponse {
	return o.Payload
}

func (o *PackerServiceDeleteIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceDeleteIterationDefault creates a PackerServiceDeleteIterationDefault with default headers values
func NewPackerServiceDeleteIterationDefault(code int) *PackerServiceDeleteIterationDefault {
	return &PackerServiceDeleteIterationDefault{
		_statusCode: code,
	}
}

/*
PackerServiceDeleteIterationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceDeleteIterationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service delete iteration default response has a 2xx status code
func (o *PackerServiceDeleteIterationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service delete iteration default response has a 3xx status code
func (o *PackerServiceDeleteIterationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service delete iteration default response has a 4xx status code
func (o *PackerServiceDeleteIterationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service delete iteration default response has a 5xx status code
func (o *PackerServiceDeleteIterationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service delete iteration default response a status code equal to that given
func (o *PackerServiceDeleteIterationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service delete iteration default response
func (o *PackerServiceDeleteIterationDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceDeleteIterationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] PackerService_DeleteIteration default %s", o._statusCode, payload)
}

func (o *PackerServiceDeleteIterationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/iterations/{iteration_id}][%d] PackerService_DeleteIteration default %s", o._statusCode, payload)
}

func (o *PackerServiceDeleteIterationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceDeleteIterationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
