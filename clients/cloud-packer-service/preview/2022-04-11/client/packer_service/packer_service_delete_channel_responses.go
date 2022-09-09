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

// PackerServiceDeleteChannelReader is a Reader for the PackerServiceDeleteChannel structure.
type PackerServiceDeleteChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceDeleteChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceDeleteChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceDeleteChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceDeleteChannelOK creates a PackerServiceDeleteChannelOK with default headers values
func NewPackerServiceDeleteChannelOK() *PackerServiceDeleteChannelOK {
	return &PackerServiceDeleteChannelOK{}
}

/*
PackerServiceDeleteChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceDeleteChannelOK struct {
	Payload models.HashicorpCloudPacker20220411DeleteChannelResponse
}

// IsSuccess returns true when this packer service delete channel o k response has a 2xx status code
func (o *PackerServiceDeleteChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service delete channel o k response has a 3xx status code
func (o *PackerServiceDeleteChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service delete channel o k response has a 4xx status code
func (o *PackerServiceDeleteChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service delete channel o k response has a 5xx status code
func (o *PackerServiceDeleteChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service delete channel o k response a status code equal to that given
func (o *PackerServiceDeleteChannelOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceDeleteChannelOK) Error() string {
	return fmt.Sprintf("[DELETE /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceDeleteChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceDeleteChannelOK) String() string {
	return fmt.Sprintf("[DELETE /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceDeleteChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceDeleteChannelOK) GetPayload() models.HashicorpCloudPacker20220411DeleteChannelResponse {
	return o.Payload
}

func (o *PackerServiceDeleteChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceDeleteChannelDefault creates a PackerServiceDeleteChannelDefault with default headers values
func NewPackerServiceDeleteChannelDefault(code int) *PackerServiceDeleteChannelDefault {
	return &PackerServiceDeleteChannelDefault{
		_statusCode: code,
	}
}

/*
PackerServiceDeleteChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceDeleteChannelDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service delete channel default response
func (o *PackerServiceDeleteChannelDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service delete channel default response has a 2xx status code
func (o *PackerServiceDeleteChannelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service delete channel default response has a 3xx status code
func (o *PackerServiceDeleteChannelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service delete channel default response has a 4xx status code
func (o *PackerServiceDeleteChannelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service delete channel default response has a 5xx status code
func (o *PackerServiceDeleteChannelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service delete channel default response a status code equal to that given
func (o *PackerServiceDeleteChannelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceDeleteChannelDefault) Error() string {
	return fmt.Sprintf("[DELETE /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_DeleteChannel default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceDeleteChannelDefault) String() string {
	return fmt.Sprintf("[DELETE /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_DeleteChannel default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceDeleteChannelDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceDeleteChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
