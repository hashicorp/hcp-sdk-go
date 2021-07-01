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

// DeleteBuildReader is a Reader for the DeleteBuild structure.
type DeleteBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBuildOK creates a DeleteBuildOK with default headers values
func NewDeleteBuildOK() *DeleteBuildOK {
	return &DeleteBuildOK{}
}

/* DeleteBuildOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteBuildOK struct {
	Payload models.HashicorpCloudPackerDeleteBuildResponse
}

func (o *DeleteBuildOK) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] deleteBuildOK  %+v", 200, o.Payload)
}
func (o *DeleteBuildOK) GetPayload() models.HashicorpCloudPackerDeleteBuildResponse {
	return o.Payload
}

func (o *DeleteBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBuildDefault creates a DeleteBuildDefault with default headers values
func NewDeleteBuildDefault(code int) *DeleteBuildDefault {
	return &DeleteBuildDefault{
		_statusCode: code,
	}
}

/* DeleteBuildDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteBuildDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete build default response
func (o *DeleteBuildDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildDefault) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] DeleteBuild default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteBuildDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeleteBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
