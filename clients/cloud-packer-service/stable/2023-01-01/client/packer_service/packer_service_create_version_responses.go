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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceCreateVersionReader is a Reader for the PackerServiceCreateVersion structure.
type PackerServiceCreateVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateVersionOK creates a PackerServiceCreateVersionOK with default headers values
func NewPackerServiceCreateVersionOK() *PackerServiceCreateVersionOK {
	return &PackerServiceCreateVersionOK{}
}

/*
PackerServiceCreateVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceCreateVersionOK struct {
	Payload *models.HashicorpCloudPacker20230101CreateVersionResponse
}

// IsSuccess returns true when this packer service create version o k response has a 2xx status code
func (o *PackerServiceCreateVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service create version o k response has a 3xx status code
func (o *PackerServiceCreateVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service create version o k response has a 4xx status code
func (o *PackerServiceCreateVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service create version o k response has a 5xx status code
func (o *PackerServiceCreateVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service create version o k response a status code equal to that given
func (o *PackerServiceCreateVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service create version o k response
func (o *PackerServiceCreateVersionOK) Code() int {
	return 200
}

func (o *PackerServiceCreateVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] packerServiceCreateVersionOK %s", 200, payload)
}

func (o *PackerServiceCreateVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] packerServiceCreateVersionOK %s", 200, payload)
}

func (o *PackerServiceCreateVersionOK) GetPayload() *models.HashicorpCloudPacker20230101CreateVersionResponse {
	return o.Payload
}

func (o *PackerServiceCreateVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101CreateVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateVersionDefault creates a PackerServiceCreateVersionDefault with default headers values
func NewPackerServiceCreateVersionDefault(code int) *PackerServiceCreateVersionDefault {
	return &PackerServiceCreateVersionDefault{
		_statusCode: code,
	}
}

/*
PackerServiceCreateVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceCreateVersionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service create version default response has a 2xx status code
func (o *PackerServiceCreateVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service create version default response has a 3xx status code
func (o *PackerServiceCreateVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service create version default response has a 4xx status code
func (o *PackerServiceCreateVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service create version default response has a 5xx status code
func (o *PackerServiceCreateVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service create version default response a status code equal to that given
func (o *PackerServiceCreateVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service create version default response
func (o *PackerServiceCreateVersionDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceCreateVersionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] PackerService_CreateVersion default %s", o._statusCode, payload)
}

func (o *PackerServiceCreateVersionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] PackerService_CreateVersion default %s", o._statusCode, payload)
}

func (o *PackerServiceCreateVersionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceCreateVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
