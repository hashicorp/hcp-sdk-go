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

// PackerServiceListBuildsReader is a Reader for the PackerServiceListBuilds structure.
type PackerServiceListBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceListBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceListBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceListBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceListBuildsOK creates a PackerServiceListBuildsOK with default headers values
func NewPackerServiceListBuildsOK() *PackerServiceListBuildsOK {
	return &PackerServiceListBuildsOK{}
}

/*
PackerServiceListBuildsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceListBuildsOK struct {
	Payload *models.HashicorpCloudPacker20230101ListBuildsResponse
}

// IsSuccess returns true when this packer service list builds o k response has a 2xx status code
func (o *PackerServiceListBuildsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service list builds o k response has a 3xx status code
func (o *PackerServiceListBuildsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service list builds o k response has a 4xx status code
func (o *PackerServiceListBuildsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service list builds o k response has a 5xx status code
func (o *PackerServiceListBuildsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service list builds o k response a status code equal to that given
func (o *PackerServiceListBuildsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service list builds o k response
func (o *PackerServiceListBuildsOK) Code() int {
	return 200
}

func (o *PackerServiceListBuildsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}/builds][%d] packerServiceListBuildsOK %s", 200, payload)
}

func (o *PackerServiceListBuildsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}/builds][%d] packerServiceListBuildsOK %s", 200, payload)
}

func (o *PackerServiceListBuildsOK) GetPayload() *models.HashicorpCloudPacker20230101ListBuildsResponse {
	return o.Payload
}

func (o *PackerServiceListBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20230101ListBuildsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceListBuildsDefault creates a PackerServiceListBuildsDefault with default headers values
func NewPackerServiceListBuildsDefault(code int) *PackerServiceListBuildsDefault {
	return &PackerServiceListBuildsDefault{
		_statusCode: code,
	}
}

/*
PackerServiceListBuildsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceListBuildsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service list builds default response has a 2xx status code
func (o *PackerServiceListBuildsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service list builds default response has a 3xx status code
func (o *PackerServiceListBuildsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service list builds default response has a 4xx status code
func (o *PackerServiceListBuildsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service list builds default response has a 5xx status code
func (o *PackerServiceListBuildsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service list builds default response a status code equal to that given
func (o *PackerServiceListBuildsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service list builds default response
func (o *PackerServiceListBuildsDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceListBuildsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}/builds][%d] PackerService_ListBuilds default %s", o._statusCode, payload)
}

func (o *PackerServiceListBuildsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2023-01-01/organizations/{location.organization_id}/projects/{location.project_id}/buckets/{bucket_name}/versions/{fingerprint}/builds][%d] PackerService_ListBuilds default %s", o._statusCode, payload)
}

func (o *PackerServiceListBuildsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceListBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
