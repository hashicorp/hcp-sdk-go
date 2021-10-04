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

// PackerServiceGetBucketReader is a Reader for the PackerServiceGetBucket structure.
type PackerServiceGetBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetBucketOK creates a PackerServiceGetBucketOK with default headers values
func NewPackerServiceGetBucketOK() *PackerServiceGetBucketOK {
	return &PackerServiceGetBucketOK{}
}

/*PackerServiceGetBucketOK handles this case with default header values.

A successful response.
*/
type PackerServiceGetBucketOK struct {
	Payload *models.HashicorpCloudPackerGetBucketResponse
}

func (o *PackerServiceGetBucketOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] packerServiceGetBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetBucketOK) GetPayload() *models.HashicorpCloudPackerGetBucketResponse {
	return o.Payload
}

func (o *PackerServiceGetBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetBucketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetBucketDefault creates a PackerServiceGetBucketDefault with default headers values
func NewPackerServiceGetBucketDefault(code int) *PackerServiceGetBucketDefault {
	return &PackerServiceGetBucketDefault{
		_statusCode: code,
	}
}

/*PackerServiceGetBucketDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceGetBucketDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service get bucket default response
func (o *PackerServiceGetBucketDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceGetBucketDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] PackerService_GetBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetBucketDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceGetBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
