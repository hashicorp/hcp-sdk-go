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
)

// NetworkServiceDeleteTGWAttachmentReader is a Reader for the NetworkServiceDeleteTGWAttachment structure.
type NetworkServiceDeleteTGWAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkServiceDeleteTGWAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkServiceDeleteTGWAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkServiceDeleteTGWAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkServiceDeleteTGWAttachmentOK creates a NetworkServiceDeleteTGWAttachmentOK with default headers values
func NewNetworkServiceDeleteTGWAttachmentOK() *NetworkServiceDeleteTGWAttachmentOK {
	return &NetworkServiceDeleteTGWAttachmentOK{}
}

/*NetworkServiceDeleteTGWAttachmentOK handles this case with default header values.

A successful response.
*/
type NetworkServiceDeleteTGWAttachmentOK struct {
	Payload *models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse
}

func (o *NetworkServiceDeleteTGWAttachmentOK) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments/{id}][%d] networkServiceDeleteTGWAttachmentOK  %+v", 200, o.Payload)
}

func (o *NetworkServiceDeleteTGWAttachmentOK) GetPayload() *models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse {
	return o.Payload
}

func (o *NetworkServiceDeleteTGWAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907DeleteTGWAttachmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkServiceDeleteTGWAttachmentDefault creates a NetworkServiceDeleteTGWAttachmentDefault with default headers values
func NewNetworkServiceDeleteTGWAttachmentDefault(code int) *NetworkServiceDeleteTGWAttachmentDefault {
	return &NetworkServiceDeleteTGWAttachmentDefault{
		_statusCode: code,
	}
}

/*NetworkServiceDeleteTGWAttachmentDefault handles this case with default header values.

An unexpected error response.
*/
type NetworkServiceDeleteTGWAttachmentDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the network service delete t g w attachment default response
func (o *NetworkServiceDeleteTGWAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *NetworkServiceDeleteTGWAttachmentDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/transit-gateway-attachments/{id}][%d] NetworkService_DeleteTGWAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkServiceDeleteTGWAttachmentDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NetworkServiceDeleteTGWAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}