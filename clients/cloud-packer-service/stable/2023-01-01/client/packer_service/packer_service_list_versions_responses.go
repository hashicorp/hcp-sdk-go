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

// PackerServiceListVersionsReader is a Reader for the PackerServiceListVersions structure.
type PackerServiceListVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceListVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceListVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceListVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceListVersionsOK creates a PackerServiceListVersionsOK with default headers values
func NewPackerServiceListVersionsOK() *PackerServiceListVersionsOK {
	return &PackerServiceListVersionsOK{}
}

/*
PackerServiceListVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceListVersionsOK struct {
	Payload *models.HashicorpCloudPacker20230101ListVersionsResponse
}

// IsSuccess returns true when this packer service list versions o k response has a 2xx status code
func (o *PackerServiceListVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service list versions o k response has a 3xx status code
func (o *PackerServiceListVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service list versions o k response has a 4xx status code
func (o *PackerServiceListVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service list versions o k response has a 5xx status code
func (o *PackerServiceListVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service list versions o k response a status code equal to that given
func (o *PackerServiceListVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service list versions o k response
func (o *PackerServiceListVersionsOK) Code() int {
	return 200
}

func (o *PackerServiceListVersionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] packerServiceListVersionsOK %s", 200, payload)
}

func (o *PackerServiceListVersionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] packerServiceListVersionsOK %s", 200, payload)
}

func (o *PackerServiceListVersionsOK) GetPayload() *models.HashicorpCloudPacker20230101ListVersionsResponse {
	return o.Payload
}

func (o *PackerServiceListVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101ListVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceListVersionsDefault creates a PackerServiceListVersionsDefault with default headers values
func NewPackerServiceListVersionsDefault(code int) *PackerServiceListVersionsDefault {
	return &PackerServiceListVersionsDefault{
		_statusCode: code,
	}
}

/*
PackerServiceListVersionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceListVersionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service list versions default response has a 2xx status code
func (o *PackerServiceListVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service list versions default response has a 3xx status code
func (o *PackerServiceListVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service list versions default response has a 4xx status code
func (o *PackerServiceListVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service list versions default response has a 5xx status code
func (o *PackerServiceListVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service list versions default response a status code equal to that given
func (o *PackerServiceListVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service list versions default response
func (o *PackerServiceListVersionsDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceListVersionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] PackerService_ListVersions default %s", o._statusCode, payload)
}

func (o *PackerServiceListVersionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions][%d] PackerService_ListVersions default %s", o._statusCode, payload)
}

func (o *PackerServiceListVersionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceListVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
