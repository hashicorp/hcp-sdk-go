// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2021-02-04/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOK creates a CreateOK with default headers values
func NewCreateOK() *CreateOK {
	return &CreateOK{}
}

/*CreateOK handles this case with default header values.

A successful response.
*/
type CreateOK struct {
	Payload *models.HashicorpCloudConsul20210204CreateResponse
}

func (o *CreateOK) Error() string {
	return fmt.Sprintf("[POST /consul/2021-02-04/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] createOK  %+v", 200, o.Payload)
}

func (o *CreateOK) GetPayload() *models.HashicorpCloudConsul20210204CreateResponse {
	return o.Payload
}

func (o *CreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20210204CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefault creates a CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	return &CreateDefault{
		_statusCode: code,
	}
}

/*CreateDefault handles this case with default header values.

An unexpected error response.
*/
type CreateDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create default response
func (o *CreateDefault) Code() int {
	return o._statusCode
}

func (o *CreateDefault) Error() string {
	return fmt.Sprintf("[POST /consul/2021-02-04/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters][%d] Create default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
