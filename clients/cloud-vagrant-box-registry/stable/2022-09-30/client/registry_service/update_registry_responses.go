// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/stable/2022-09-30/models"
)

// UpdateRegistryReader is a Reader for the UpdateRegistry structure.
type UpdateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRegistryOK creates a UpdateRegistryOK with default headers values
func NewUpdateRegistryOK() *UpdateRegistryOK {
	return &UpdateRegistryOK{}
}

/*
UpdateRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRegistryOK struct {
	Payload *models.HashicorpCloudVagrant20220930UpdateRegistryResponse
}

// IsSuccess returns true when this update registry o k response has a 2xx status code
func (o *UpdateRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update registry o k response has a 3xx status code
func (o *UpdateRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update registry o k response has a 4xx status code
func (o *UpdateRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update registry o k response has a 5xx status code
func (o *UpdateRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update registry o k response a status code equal to that given
func (o *UpdateRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update registry o k response
func (o *UpdateRegistryOK) Code() int {
	return 200
}

func (o *UpdateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}][%d] updateRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateRegistryOK) String() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}][%d] updateRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateRegistryOK) GetPayload() *models.HashicorpCloudVagrant20220930UpdateRegistryResponse {
	return o.Payload
}

func (o *UpdateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930UpdateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegistryDefault creates a UpdateRegistryDefault with default headers values
func NewUpdateRegistryDefault(code int) *UpdateRegistryDefault {
	return &UpdateRegistryDefault{
		_statusCode: code,
	}
}

/*
UpdateRegistryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateRegistryDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this update registry default response has a 2xx status code
func (o *UpdateRegistryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update registry default response has a 3xx status code
func (o *UpdateRegistryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update registry default response has a 4xx status code
func (o *UpdateRegistryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update registry default response has a 5xx status code
func (o *UpdateRegistryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update registry default response a status code equal to that given
func (o *UpdateRegistryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update registry default response
func (o *UpdateRegistryDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRegistryDefault) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}][%d] UpdateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRegistryDefault) String() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}][%d] UpdateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRegistryDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *UpdateRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
