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

// PackerServiceCreateRegistryReader is a Reader for the PackerServiceCreateRegistry structure.
type PackerServiceCreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateRegistryOK creates a PackerServiceCreateRegistryOK with default headers values
func NewPackerServiceCreateRegistryOK() *PackerServiceCreateRegistryOK {
	return &PackerServiceCreateRegistryOK{}
}

/*
PackerServiceCreateRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceCreateRegistryOK struct {
	Payload *models.HashicorpCloudPacker20230101CreateRegistryResponse
}

// IsSuccess returns true when this packer service create registry o k response has a 2xx status code
func (o *PackerServiceCreateRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service create registry o k response has a 3xx status code
func (o *PackerServiceCreateRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service create registry o k response has a 4xx status code
func (o *PackerServiceCreateRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service create registry o k response has a 5xx status code
func (o *PackerServiceCreateRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service create registry o k response a status code equal to that given
func (o *PackerServiceCreateRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service create registry o k response
func (o *PackerServiceCreateRegistryOK) Code() int {
	return 200
}

func (o *PackerServiceCreateRegistryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] packerServiceCreateRegistryOK %s", 200, payload)
}

func (o *PackerServiceCreateRegistryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] packerServiceCreateRegistryOK %s", 200, payload)
}

func (o *PackerServiceCreateRegistryOK) GetPayload() *models.HashicorpCloudPacker20230101CreateRegistryResponse {
	return o.Payload
}

func (o *PackerServiceCreateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101CreateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateRegistryDefault creates a PackerServiceCreateRegistryDefault with default headers values
func NewPackerServiceCreateRegistryDefault(code int) *PackerServiceCreateRegistryDefault {
	return &PackerServiceCreateRegistryDefault{
		_statusCode: code,
	}
}

/*
PackerServiceCreateRegistryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceCreateRegistryDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service create registry default response has a 2xx status code
func (o *PackerServiceCreateRegistryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service create registry default response has a 3xx status code
func (o *PackerServiceCreateRegistryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service create registry default response has a 4xx status code
func (o *PackerServiceCreateRegistryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service create registry default response has a 5xx status code
func (o *PackerServiceCreateRegistryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service create registry default response a status code equal to that given
func (o *PackerServiceCreateRegistryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service create registry default response
func (o *PackerServiceCreateRegistryDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceCreateRegistryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] PackerService_CreateRegistry default %s", o._statusCode, payload)
}

func (o *PackerServiceCreateRegistryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] PackerService_CreateRegistry default %s", o._statusCode, payload)
}

func (o *PackerServiceCreateRegistryDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceCreateRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
