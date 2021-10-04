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

/*PackerServiceListIterationsOK handles this case with default header values.

A successful response.
*/
type PackerServiceListIterationsOK struct {
	Payload *models.HashicorpCloudPackerListIterationsResponse
}

func (o *PackerServiceListIterationsOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] packerServiceListIterationsOK  %+v", 200, o.Payload)
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

/*PackerServiceListIterationsDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceListIterationsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service list iterations default response
func (o *PackerServiceListIterationsDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceListIterationsDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] PackerService_ListIterations default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceListIterationsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceListIterationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
