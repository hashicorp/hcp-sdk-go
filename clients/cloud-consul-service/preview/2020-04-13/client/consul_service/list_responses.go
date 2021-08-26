// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-04-13/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*ListOK handles this case with default header values.

A successful response.
*/
type ListOK struct {
	Payload *models.HashicorpCloudConsul20200413ListResponse
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[GET /consul/2020-04-13/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) GetPayload() *models.HashicorpCloudConsul20200413ListResponse {
	return o.Payload
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200413ListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDefault creates a ListDefault with default headers values
func NewListDefault(code int) *ListDefault {
	return &ListDefault{
		_statusCode: code,
	}
}

/*ListDefault handles this case with default header values.

An unexpected error response.
*/
type ListDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list default response
func (o *ListDefault) Code() int {
	return o._statusCode
}

func (o *ListDefault) Error() string {
	return fmt.Sprintf("[GET /consul/2020-04-13/organizations/{location.organization_id}/projects/{location.project_id}/clusters][%d] List default  %+v", o._statusCode, o.Payload)
}

func (o *ListDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
