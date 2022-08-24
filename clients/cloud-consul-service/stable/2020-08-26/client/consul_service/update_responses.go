// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/stable/2020-08-26/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// UpdateReader is a Reader for the Update structure.
type UpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOK creates a UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {
	return &UpdateOK{}
}

/* UpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateOK struct {
	Payload *models.HashicorpCloudConsul20200826UpdateResponse
}

func (o *UpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters/{cluster.id}][%d] updateOK  %+v", 200, o.Payload)
}
func (o *UpdateOK) GetPayload() *models.HashicorpCloudConsul20200826UpdateResponse {
	return o.Payload
}

func (o *UpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826UpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDefault creates a UpdateDefault with default headers values
func NewUpdateDefault(code int) *UpdateDefault {
	return &UpdateDefault{
		_statusCode: code,
	}
}

/* UpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the update default response
func (o *UpdateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /consul/2020-08-26/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters/{cluster.id}][%d] Update default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
