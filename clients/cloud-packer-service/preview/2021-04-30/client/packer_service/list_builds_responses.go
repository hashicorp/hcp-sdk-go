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

// ListBuildsReader is a Reader for the ListBuilds structure.
type ListBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBuildsOK creates a ListBuildsOK with default headers values
func NewListBuildsOK() *ListBuildsOK {
	return &ListBuildsOK{}
}

/* ListBuildsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListBuildsOK struct {
	Payload *models.HashicorpCloudPackerListBuildResponse
}

func (o *ListBuildsOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{iteration_id}/builds][%d] listBuildsOK  %+v", 200, o.Payload)
}
func (o *ListBuildsOK) GetPayload() *models.HashicorpCloudPackerListBuildResponse {
	return o.Payload
}

func (o *ListBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerListBuildResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsDefault creates a ListBuildsDefault with default headers values
func NewListBuildsDefault(code int) *ListBuildsDefault {
	return &ListBuildsDefault{
		_statusCode: code,
	}
}

/* ListBuildsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListBuildsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list builds default response
func (o *ListBuildsDefault) Code() int {
	return o._statusCode
}

func (o *ListBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations/{iteration_id}/builds][%d] ListBuilds default  %+v", o._statusCode, o.Payload)
}
func (o *ListBuildsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
