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

// CreateSnapshotReader is a Reader for the CreateSnapshot structure.
type CreateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSnapshotOK creates a CreateSnapshotOK with default headers values
func NewCreateSnapshotOK() *CreateSnapshotOK {
	return &CreateSnapshotOK{}
}

/*
CreateSnapshotOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateSnapshotOK struct {
	Payload *models.HashicorpCloudConsul20200826CreateSnapshotResponse
}

// IsSuccess returns true when this create snapshot o k response has a 2xx status code
func (o *CreateSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create snapshot o k response has a 3xx status code
func (o *CreateSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create snapshot o k response has a 4xx status code
func (o *CreateSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create snapshot o k response has a 5xx status code
func (o *CreateSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create snapshot o k response a status code equal to that given
func (o *CreateSnapshotOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] createSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateSnapshotOK) String() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] createSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateSnapshotOK) GetPayload() *models.HashicorpCloudConsul20200826CreateSnapshotResponse {
	return o.Payload
}

func (o *CreateSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826CreateSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotDefault creates a CreateSnapshotDefault with default headers values
func NewCreateSnapshotDefault(code int) *CreateSnapshotDefault {
	return &CreateSnapshotDefault{
		_statusCode: code,
	}
}

/*
CreateSnapshotDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateSnapshotDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the create snapshot default response
func (o *CreateSnapshotDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create snapshot default response has a 2xx status code
func (o *CreateSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create snapshot default response has a 3xx status code
func (o *CreateSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create snapshot default response has a 4xx status code
func (o *CreateSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create snapshot default response has a 5xx status code
func (o *CreateSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create snapshot default response a status code equal to that given
func (o *CreateSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] CreateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) String() string {
	return fmt.Sprintf("[POST /consul/2020-08-26/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots][%d] CreateSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CreateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
