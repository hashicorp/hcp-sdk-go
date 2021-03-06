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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// WaitReader is a Reader for the Wait structure.
type WaitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaitDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaitOK creates a WaitOK with default headers values
func NewWaitOK() *WaitOK {
	return &WaitOK{}
}

/*WaitOK handles this case with default header values.

A successful response.
*/
type WaitOK struct {
	Payload *models.HashicorpCloudOperationWaitResponse
}

func (o *WaitOK) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}/wait][%d] waitOK  %+v", 200, o.Payload)
}

func (o *WaitOK) GetPayload() *models.HashicorpCloudOperationWaitResponse {
	return o.Payload
}

func (o *WaitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudOperationWaitResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaitDefault creates a WaitDefault with default headers values
func NewWaitDefault(code int) *WaitDefault {
	return &WaitDefault{
		_statusCode: code,
	}
}

/*WaitDefault handles this case with default header values.

An unexpected error response.
*/
type WaitDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the wait default response
func (o *WaitDefault) Code() int {
	return o._statusCode
}

func (o *WaitDefault) Error() string {
	return fmt.Sprintf("[GET /operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}/wait][%d] Wait default  %+v", o._statusCode, o.Payload)
}

func (o *WaitDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaitDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
