// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateTGWAttachmentReader is a Reader for the CreateTGWAttachment structure.
type CreateTGWAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTGWAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTGWAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTGWAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTGWAttachmentOK creates a CreateTGWAttachmentOK with default headers values
func NewCreateTGWAttachmentOK() *CreateTGWAttachmentOK {
	return &CreateTGWAttachmentOK{}
}

/* CreateTGWAttachmentOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTGWAttachmentOK struct {
	Payload *models.HashicorpCloudNetwork20200907CreateTGWAttachmentResponse
}

func (o *CreateTGWAttachmentOK) Error() string {
	return fmt.Sprintf("[POST /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] createTGWAttachmentOK  %+v", 200, o.Payload)
}
func (o *CreateTGWAttachmentOK) GetPayload() *models.HashicorpCloudNetwork20200907CreateTGWAttachmentResponse {
	return o.Payload
}

func (o *CreateTGWAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907CreateTGWAttachmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTGWAttachmentDefault creates a CreateTGWAttachmentDefault with default headers values
func NewCreateTGWAttachmentDefault(code int) *CreateTGWAttachmentDefault {
	return &CreateTGWAttachmentDefault{
		_statusCode: code,
	}
}

/* CreateTGWAttachmentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTGWAttachmentDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create t g w attachment default response
func (o *CreateTGWAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *CreateTGWAttachmentDefault) Error() string {
	return fmt.Sprintf("[POST /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments][%d] CreateTGWAttachment default  %+v", o._statusCode, o.Payload)
}
func (o *CreateTGWAttachmentDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateTGWAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
