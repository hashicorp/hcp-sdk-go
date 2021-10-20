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

// PackerServiceGetChildImagesReader is a Reader for the PackerServiceGetChildImages structure.
type PackerServiceGetChildImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetChildImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetChildImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetChildImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetChildImagesOK creates a PackerServiceGetChildImagesOK with default headers values
func NewPackerServiceGetChildImagesOK() *PackerServiceGetChildImagesOK {
	return &PackerServiceGetChildImagesOK{}
}

/*PackerServiceGetChildImagesOK handles this case with default header values.

A successful response.
*/
type PackerServiceGetChildImagesOK struct {
	Payload *models.HashicorpCloudPackerGetChildImagesResponse
}

func (o *PackerServiceGetChildImagesOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{incremental_version}/children][%d] packerServiceGetChildImagesOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetChildImagesOK) GetPayload() *models.HashicorpCloudPackerGetChildImagesResponse {
	return o.Payload
}

func (o *PackerServiceGetChildImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetChildImagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetChildImagesDefault creates a PackerServiceGetChildImagesDefault with default headers values
func NewPackerServiceGetChildImagesDefault(code int) *PackerServiceGetChildImagesDefault {
	return &PackerServiceGetChildImagesDefault{
		_statusCode: code,
	}
}

/*PackerServiceGetChildImagesDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceGetChildImagesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service get child images default response
func (o *PackerServiceGetChildImagesDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceGetChildImagesDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{incremental_version}/children][%d] PackerService_GetChildImages default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetChildImagesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceGetChildImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}