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
)

// WaypointDestroyAddOnReader is a Reader for the WaypointDestroyAddOn structure.
type WaypointDestroyAddOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointDestroyAddOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointDestroyAddOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointDestroyAddOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointDestroyAddOnOK creates a WaypointDestroyAddOnOK with default headers values
func NewWaypointDestroyAddOnOK() *WaypointDestroyAddOnOK {
	return &WaypointDestroyAddOnOK{}
}

/*
WaypointDestroyAddOnOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointDestroyAddOnOK struct {
	Payload interface{}
}

// IsSuccess returns true when this waypoint destroy add on o k response has a 2xx status code
func (o *WaypointDestroyAddOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint destroy add on o k response has a 3xx status code
func (o *WaypointDestroyAddOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint destroy add on o k response has a 4xx status code
func (o *WaypointDestroyAddOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint destroy add on o k response has a 5xx status code
func (o *WaypointDestroyAddOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint destroy add on o k response a status code equal to that given
func (o *WaypointDestroyAddOnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint destroy add on o k response
func (o *WaypointDestroyAddOnOK) Code() int {
	return 200
}

func (o *WaypointDestroyAddOnOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/add-on/{add_on.id}][%d] waypointDestroyAddOnOK  %+v", 200, o.Payload)
}

func (o *WaypointDestroyAddOnOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/add-on/{add_on.id}][%d] waypointDestroyAddOnOK  %+v", 200, o.Payload)
}

func (o *WaypointDestroyAddOnOK) GetPayload() interface{} {
	return o.Payload
}

func (o *WaypointDestroyAddOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointDestroyAddOnDefault creates a WaypointDestroyAddOnDefault with default headers values
func NewWaypointDestroyAddOnDefault(code int) *WaypointDestroyAddOnDefault {
	return &WaypointDestroyAddOnDefault{
		_statusCode: code,
	}
}

/*
WaypointDestroyAddOnDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointDestroyAddOnDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint destroy add on default response has a 2xx status code
func (o *WaypointDestroyAddOnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint destroy add on default response has a 3xx status code
func (o *WaypointDestroyAddOnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint destroy add on default response has a 4xx status code
func (o *WaypointDestroyAddOnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint destroy add on default response has a 5xx status code
func (o *WaypointDestroyAddOnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint destroy add on default response a status code equal to that given
func (o *WaypointDestroyAddOnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint destroy add on default response
func (o *WaypointDestroyAddOnDefault) Code() int {
	return o._statusCode
}

func (o *WaypointDestroyAddOnDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/add-on/{add_on.id}][%d] Waypoint_DestroyAddOn default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDestroyAddOnDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/add-on/{add_on.id}][%d] Waypoint_DestroyAddOn default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDestroyAddOnDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointDestroyAddOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}