// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// WaypointServiceDeleteActionConfigReader is a Reader for the WaypointServiceDeleteActionConfig structure.
type WaypointServiceDeleteActionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceDeleteActionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceDeleteActionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceDeleteActionConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceDeleteActionConfigOK creates a WaypointServiceDeleteActionConfigOK with default headers values
func NewWaypointServiceDeleteActionConfigOK() *WaypointServiceDeleteActionConfigOK {
	return &WaypointServiceDeleteActionConfigOK{}
}

/*
WaypointServiceDeleteActionConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceDeleteActionConfigOK struct {
	Payload interface{}
}

// IsSuccess returns true when this waypoint service delete action config o k response has a 2xx status code
func (o *WaypointServiceDeleteActionConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service delete action config o k response has a 3xx status code
func (o *WaypointServiceDeleteActionConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service delete action config o k response has a 4xx status code
func (o *WaypointServiceDeleteActionConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service delete action config o k response has a 5xx status code
func (o *WaypointServiceDeleteActionConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service delete action config o k response a status code equal to that given
func (o *WaypointServiceDeleteActionConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service delete action config o k response
func (o *WaypointServiceDeleteActionConfigOK) Code() int {
	return 200
}

func (o *WaypointServiceDeleteActionConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/actionconfig][%d] waypointServiceDeleteActionConfigOK %s", 200, payload)
}

func (o *WaypointServiceDeleteActionConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/actionconfig][%d] waypointServiceDeleteActionConfigOK %s", 200, payload)
}

func (o *WaypointServiceDeleteActionConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *WaypointServiceDeleteActionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceDeleteActionConfigDefault creates a WaypointServiceDeleteActionConfigDefault with default headers values
func NewWaypointServiceDeleteActionConfigDefault(code int) *WaypointServiceDeleteActionConfigDefault {
	return &WaypointServiceDeleteActionConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceDeleteActionConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceDeleteActionConfigDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service delete action config default response has a 2xx status code
func (o *WaypointServiceDeleteActionConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service delete action config default response has a 3xx status code
func (o *WaypointServiceDeleteActionConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service delete action config default response has a 4xx status code
func (o *WaypointServiceDeleteActionConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service delete action config default response has a 5xx status code
func (o *WaypointServiceDeleteActionConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service delete action config default response a status code equal to that given
func (o *WaypointServiceDeleteActionConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service delete action config default response
func (o *WaypointServiceDeleteActionConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceDeleteActionConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/actionconfig][%d] WaypointService_DeleteActionConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceDeleteActionConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/actionconfig][%d] WaypointService_DeleteActionConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceDeleteActionConfigDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceDeleteActionConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
