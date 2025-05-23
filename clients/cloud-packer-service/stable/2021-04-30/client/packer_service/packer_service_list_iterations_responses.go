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

// PackerServiceListIterationsReader is a Reader for the PackerServiceListIterations structure.
type PackerServiceListIterationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceListIterationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceListIterationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceListIterationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceListIterationsOK creates a PackerServiceListIterationsOK with default headers values
func NewPackerServiceListIterationsOK() *PackerServiceListIterationsOK {
	return &PackerServiceListIterationsOK{}
}

/*
PackerServiceListIterationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceListIterationsOK struct {
	Payload *models.HashicorpCloudPackerListIterationsResponse
}

// IsSuccess returns true when this packer service list iterations o k response has a 2xx status code
func (o *PackerServiceListIterationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service list iterations o k response has a 3xx status code
func (o *PackerServiceListIterationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service list iterations o k response has a 4xx status code
func (o *PackerServiceListIterationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service list iterations o k response has a 5xx status code
func (o *PackerServiceListIterationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service list iterations o k response a status code equal to that given
func (o *PackerServiceListIterationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service list iterations o k response
func (o *PackerServiceListIterationsOK) Code() int {
	return 200
}

func (o *PackerServiceListIterationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] packerServiceListIterationsOK %s", 200, payload)
}

func (o *PackerServiceListIterationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] packerServiceListIterationsOK %s", 200, payload)
}

func (o *PackerServiceListIterationsOK) GetPayload() *models.HashicorpCloudPackerListIterationsResponse {
	return o.Payload
}

func (o *PackerServiceListIterationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerListIterationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceListIterationsDefault creates a PackerServiceListIterationsDefault with default headers values
func NewPackerServiceListIterationsDefault(code int) *PackerServiceListIterationsDefault {
	return &PackerServiceListIterationsDefault{
		_statusCode: code,
	}
}

/*
PackerServiceListIterationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceListIterationsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service list iterations default response has a 2xx status code
func (o *PackerServiceListIterationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service list iterations default response has a 3xx status code
func (o *PackerServiceListIterationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service list iterations default response has a 4xx status code
func (o *PackerServiceListIterationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service list iterations default response has a 5xx status code
func (o *PackerServiceListIterationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service list iterations default response a status code equal to that given
func (o *PackerServiceListIterationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service list iterations default response
func (o *PackerServiceListIterationsDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceListIterationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] PackerService_ListIterations default %s", o._statusCode, payload)
}

func (o *PackerServiceListIterationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] PackerService_ListIterations default %s", o._statusCode, payload)
}

func (o *PackerServiceListIterationsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceListIterationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
