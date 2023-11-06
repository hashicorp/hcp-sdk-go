// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// WaypointGetConfig2Reader is a Reader for the WaypointGetConfig2 structure.
type WaypointGetConfig2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetConfig2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetConfig2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetConfig2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetConfig2OK creates a WaypointGetConfig2OK with default headers values
func NewWaypointGetConfig2OK() *WaypointGetConfig2OK {
	return &WaypointGetConfig2OK{}
}

/*
WaypointGetConfig2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetConfig2OK struct {
	Payload *models.HashicorpWaypointConfigGetResponse
}

// IsSuccess returns true when this waypoint get config2 o k response has a 2xx status code
func (o *WaypointGetConfig2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get config2 o k response has a 3xx status code
func (o *WaypointGetConfig2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get config2 o k response has a 4xx status code
func (o *WaypointGetConfig2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get config2 o k response has a 5xx status code
func (o *WaypointGetConfig2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get config2 o k response a status code equal to that given
func (o *WaypointGetConfig2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get config2 o k response
func (o *WaypointGetConfig2OK) Code() int {
	return 200
}

func (o *WaypointGetConfig2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{project.project}/config][%d] waypointGetConfig2OK  %+v", 200, o.Payload)
}

func (o *WaypointGetConfig2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{project.project}/config][%d] waypointGetConfig2OK  %+v", 200, o.Payload)
}

func (o *WaypointGetConfig2OK) GetPayload() *models.HashicorpWaypointConfigGetResponse {
	return o.Payload
}

func (o *WaypointGetConfig2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointConfigGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetConfig2Default creates a WaypointGetConfig2Default with default headers values
func NewWaypointGetConfig2Default(code int) *WaypointGetConfig2Default {
	return &WaypointGetConfig2Default{
		_statusCode: code,
	}
}

/*
WaypointGetConfig2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetConfig2Default struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get config2 default response has a 2xx status code
func (o *WaypointGetConfig2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get config2 default response has a 3xx status code
func (o *WaypointGetConfig2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get config2 default response has a 4xx status code
func (o *WaypointGetConfig2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get config2 default response has a 5xx status code
func (o *WaypointGetConfig2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get config2 default response a status code equal to that given
func (o *WaypointGetConfig2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get config2 default response
func (o *WaypointGetConfig2Default) Code() int {
	return o._statusCode
}

func (o *WaypointGetConfig2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{project.project}/config][%d] Waypoint_GetConfig2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetConfig2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{project.project}/config][%d] Waypoint_GetConfig2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetConfig2Default) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetConfig2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}