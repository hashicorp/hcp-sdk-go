// Code generated by go-swagger; DO NOT EDIT.

package operation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-operation/preview/2020-05-05/models"
)

// OperationServiceGetReader is a Reader for the OperationServiceGet structure.
type OperationServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationServiceGetOK creates a OperationServiceGetOK with default headers values
func NewOperationServiceGetOK() *OperationServiceGetOK {
	return &OperationServiceGetOK{}
}

/*OperationServiceGetOK handles this case with default header values.

A successful response.
*/
type OperationServiceGetOK struct {
	Payload *models.HashicorpCloudOperationGetResponse
}

func (o *OperationServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}][%d] operationServiceGetOK  %+v", 200, o.Payload)
}

func (o *OperationServiceGetOK) GetPayload() *models.HashicorpCloudOperationGetResponse {
	return o.Payload
}

func (o *OperationServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudOperationGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationServiceGetDefault creates a OperationServiceGetDefault with default headers values
func NewOperationServiceGetDefault(code int) *OperationServiceGetDefault {
	return &OperationServiceGetDefault{
		_statusCode: code,
	}
}

/*OperationServiceGetDefault handles this case with default header values.

An unexpected error response.
*/
type OperationServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the operation service get default response
func (o *OperationServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *OperationServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}][%d] OperationService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *OperationServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OperationServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
