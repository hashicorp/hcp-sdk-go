// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// WaypointServiceListNoCodeModules2Reader is a Reader for the WaypointServiceListNoCodeModules2 structure.
type WaypointServiceListNoCodeModules2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListNoCodeModules2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListNoCodeModules2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListNoCodeModules2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListNoCodeModules2OK creates a WaypointServiceListNoCodeModules2OK with default headers values
func NewWaypointServiceListNoCodeModules2OK() *WaypointServiceListNoCodeModules2OK {
	return &WaypointServiceListNoCodeModules2OK{}
}

/*
WaypointServiceListNoCodeModules2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListNoCodeModules2OK struct {
	Payload *models.HashicorpCloudWaypointListNoCodeModulesResponse
}

// IsSuccess returns true when this waypoint service list no code modules2 o k response has a 2xx status code
func (o *WaypointServiceListNoCodeModules2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list no code modules2 o k response has a 3xx status code
func (o *WaypointServiceListNoCodeModules2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list no code modules2 o k response has a 4xx status code
func (o *WaypointServiceListNoCodeModules2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list no code modules2 o k response has a 5xx status code
func (o *WaypointServiceListNoCodeModules2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list no code modules2 o k response a status code equal to that given
func (o *WaypointServiceListNoCodeModules2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list no code modules2 o k response
func (o *WaypointServiceListNoCodeModules2OK) Code() int {
	return 200
}

func (o *WaypointServiceListNoCodeModules2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfmodules][%d] waypointServiceListNoCodeModules2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceListNoCodeModules2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfmodules][%d] waypointServiceListNoCodeModules2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceListNoCodeModules2OK) GetPayload() *models.HashicorpCloudWaypointListNoCodeModulesResponse {
	return o.Payload
}

func (o *WaypointServiceListNoCodeModules2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListNoCodeModulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListNoCodeModules2Default creates a WaypointServiceListNoCodeModules2Default with default headers values
func NewWaypointServiceListNoCodeModules2Default(code int) *WaypointServiceListNoCodeModules2Default {
	return &WaypointServiceListNoCodeModules2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceListNoCodeModules2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListNoCodeModules2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list no code modules2 default response has a 2xx status code
func (o *WaypointServiceListNoCodeModules2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list no code modules2 default response has a 3xx status code
func (o *WaypointServiceListNoCodeModules2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list no code modules2 default response has a 4xx status code
func (o *WaypointServiceListNoCodeModules2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list no code modules2 default response has a 5xx status code
func (o *WaypointServiceListNoCodeModules2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list no code modules2 default response a status code equal to that given
func (o *WaypointServiceListNoCodeModules2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list no code modules2 default response
func (o *WaypointServiceListNoCodeModules2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListNoCodeModules2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfmodules][%d] WaypointService_ListNoCodeModules2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceListNoCodeModules2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfmodules][%d] WaypointService_ListNoCodeModules2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceListNoCodeModules2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListNoCodeModules2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}