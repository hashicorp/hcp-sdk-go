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

/*PackerServiceDeleteChannelOK handles this case with default header values.

A successful response.
*/
type PackerServiceDeleteChannelOK struct {
	Payload models.HashicorpCloudPackerDeleteChannelResponse
}

func (o *PackerServiceDeleteChannelOK) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] packerServiceDeleteChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceDeleteChannelOK) GetPayload() models.HashicorpCloudPackerDeleteChannelResponse {
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

/*PackerServiceDeleteChannelDefault handles this case with default header values.

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

func (o *PackerServiceDeleteChannelDefault) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels/{slug}][%d] PackerService_DeleteChannel default  %+v", o._statusCode, o.Payload)
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
