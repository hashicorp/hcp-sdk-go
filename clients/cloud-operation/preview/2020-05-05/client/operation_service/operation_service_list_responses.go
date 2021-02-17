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

// OperationServiceListReader is a Reader for the OperationServiceList structure.
type OperationServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationServiceListOK creates a OperationServiceListOK with default headers values
func NewOperationServiceListOK() *OperationServiceListOK {
	return &OperationServiceListOK{}
}

/*OperationServiceListOK handles this case with default header values.

A successful response.
*/
type OperationServiceListOK struct {
	Payload *models.HashicorpCloudOperationListResponse
}

func (o *OperationServiceListOK) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations][%d] operationServiceListOK  %+v", 200, o.Payload)
}

func (o *OperationServiceListOK) GetPayload() *models.HashicorpCloudOperationListResponse {
	return o.Payload
}

func (o *OperationServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudOperationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationServiceListDefault creates a OperationServiceListDefault with default headers values
func NewOperationServiceListDefault(code int) *OperationServiceListDefault {
	return &OperationServiceListDefault{
		_statusCode: code,
	}
}

/*OperationServiceListDefault handles this case with default header values.

An unexpected error response.
*/
type OperationServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the operation service list default response
func (o *OperationServiceListDefault) Code() int {
	return o._statusCode
}

func (o *OperationServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations][%d] OperationService_List default  %+v", o._statusCode, o.Payload)
}

func (o *OperationServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OperationServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
