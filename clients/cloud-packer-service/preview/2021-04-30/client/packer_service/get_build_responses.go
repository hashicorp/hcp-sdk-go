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

// GetBuildReader is a Reader for the GetBuild structure.
type GetBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildOK creates a GetBuildOK with default headers values
func NewGetBuildOK() *GetBuildOK {
	return &GetBuildOK{}
}

/* GetBuildOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetBuildOK struct {
	Payload *models.HashicorpCloudPackerGetBuildResponse
}

func (o *GetBuildOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] getBuildOK  %+v", 200, o.Payload)
}
func (o *GetBuildOK) GetPayload() *models.HashicorpCloudPackerGetBuildResponse {
	return o.Payload
}

func (o *GetBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetBuildResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildDefault creates a GetBuildDefault with default headers values
func NewGetBuildDefault(code int) *GetBuildDefault {
	return &GetBuildDefault{
		_statusCode: code,
	}
}

/* GetBuildDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetBuildDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get build default response
func (o *GetBuildDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] GetBuild default  %+v", o._statusCode, o.Payload)
}
func (o *GetBuildDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
