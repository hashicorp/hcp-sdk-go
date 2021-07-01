// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// UpdateChannelReader is a Reader for the UpdateChannel structure.
type UpdateChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateChannelOK creates a UpdateChannelOK with default headers values
func NewUpdateChannelOK() *UpdateChannelOK {
	return &UpdateChannelOK{}
}

/* UpdateChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateChannelOK struct {
	Payload *models.HashicorpCloudPackerUpdateChannelResponse
}

func (o *UpdateChannelOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] updateChannelOK  %+v", 200, o.Payload)
}
func (o *UpdateChannelOK) GetPayload() *models.HashicorpCloudPackerUpdateChannelResponse {
	return o.Payload
}

func (o *UpdateChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChannelDefault creates a UpdateChannelDefault with default headers values
func NewUpdateChannelDefault(code int) *UpdateChannelDefault {
	return &UpdateChannelDefault{
		_statusCode: code,
	}
}

/* UpdateChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateChannelDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the update channel default response
func (o *UpdateChannelDefault) Code() int {
	return o._statusCode
}

func (o *UpdateChannelDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] UpdateChannel default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateChannelDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdateChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
