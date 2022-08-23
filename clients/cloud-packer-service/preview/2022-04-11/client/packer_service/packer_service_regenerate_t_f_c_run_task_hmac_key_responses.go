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

// PackerServiceRegenerateTFCRunTaskHmacKeyReader is a Reader for the PackerServiceRegenerateTFCRunTaskHmacKey structure.
type PackerServiceRegenerateTFCRunTaskHmacKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceRegenerateTFCRunTaskHmacKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyOK creates a PackerServiceRegenerateTFCRunTaskHmacKeyOK with default headers values
func NewPackerServiceRegenerateTFCRunTaskHmacKeyOK() *PackerServiceRegenerateTFCRunTaskHmacKeyOK {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyOK{}
}

/*PackerServiceRegenerateTFCRunTaskHmacKeyOK handles this case with default header values.

A successful response.
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyOK struct {
	Payload *models.HashicorpCloudPacker20220411RegenerateTFCRunTaskHmacKeyResponse
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] packerServiceRegenerateTFCRunTaskHmacKeyOK  %+v", 200, o.Payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) GetPayload() *models.HashicorpCloudPacker20220411RegenerateTFCRunTaskHmacKeyResponse {
	return o.Payload
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPacker20220411RegenerateTFCRunTaskHmacKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault creates a PackerServiceRegenerateTFCRunTaskHmacKeyDefault with default headers values
func NewPackerServiceRegenerateTFCRunTaskHmacKeyDefault(code int) *PackerServiceRegenerateTFCRunTaskHmacKeyDefault {
	return &PackerServiceRegenerateTFCRunTaskHmacKeyDefault{
		_statusCode: code,
	}
}

/*PackerServiceRegenerateTFCRunTaskHmacKeyDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceRegenerateTFCRunTaskHmacKeyDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service regenerate t f c run task hmac key default response
func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2022-04-11/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/hmac][%d] PackerService_RegenerateTFCRunTaskHmacKey default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceRegenerateTFCRunTaskHmacKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}