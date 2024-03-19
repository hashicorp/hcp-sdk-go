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

// WaypointServiceUIGetActionConfigBundle2Reader is a Reader for the WaypointServiceUIGetActionConfigBundle2 structure.
type WaypointServiceUIGetActionConfigBundle2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUIGetActionConfigBundle2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUIGetActionConfigBundle2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUIGetActionConfigBundle2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUIGetActionConfigBundle2OK creates a WaypointServiceUIGetActionConfigBundle2OK with default headers values
func NewWaypointServiceUIGetActionConfigBundle2OK() *WaypointServiceUIGetActionConfigBundle2OK {
	return &WaypointServiceUIGetActionConfigBundle2OK{}
}

/*
WaypointServiceUIGetActionConfigBundle2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUIGetActionConfigBundle2OK struct {
	Payload *models.HashicorpCloudWaypointUIGetActionConfigResponse
}

// IsSuccess returns true when this waypoint service Ui get action config bundle2 o k response has a 2xx status code
func (o *WaypointServiceUIGetActionConfigBundle2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service Ui get action config bundle2 o k response has a 3xx status code
func (o *WaypointServiceUIGetActionConfigBundle2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service Ui get action config bundle2 o k response has a 4xx status code
func (o *WaypointServiceUIGetActionConfigBundle2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service Ui get action config bundle2 o k response has a 5xx status code
func (o *WaypointServiceUIGetActionConfigBundle2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service Ui get action config bundle2 o k response a status code equal to that given
func (o *WaypointServiceUIGetActionConfigBundle2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service Ui get action config bundle2 o k response
func (o *WaypointServiceUIGetActionConfigBundle2OK) Code() int {
	return 200
}

func (o *WaypointServiceUIGetActionConfigBundle2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfig/by-name/{action_name}][%d] waypointServiceUiGetActionConfigBundle2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUIGetActionConfigBundle2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfig/by-name/{action_name}][%d] waypointServiceUiGetActionConfigBundle2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUIGetActionConfigBundle2OK) GetPayload() *models.HashicorpCloudWaypointUIGetActionConfigResponse {
	return o.Payload
}

func (o *WaypointServiceUIGetActionConfigBundle2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUIGetActionConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUIGetActionConfigBundle2Default creates a WaypointServiceUIGetActionConfigBundle2Default with default headers values
func NewWaypointServiceUIGetActionConfigBundle2Default(code int) *WaypointServiceUIGetActionConfigBundle2Default {
	return &WaypointServiceUIGetActionConfigBundle2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUIGetActionConfigBundle2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUIGetActionConfigBundle2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service UI get action config bundle2 default response has a 2xx status code
func (o *WaypointServiceUIGetActionConfigBundle2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service UI get action config bundle2 default response has a 3xx status code
func (o *WaypointServiceUIGetActionConfigBundle2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service UI get action config bundle2 default response has a 4xx status code
func (o *WaypointServiceUIGetActionConfigBundle2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service UI get action config bundle2 default response has a 5xx status code
func (o *WaypointServiceUIGetActionConfigBundle2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service UI get action config bundle2 default response a status code equal to that given
func (o *WaypointServiceUIGetActionConfigBundle2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service UI get action config bundle2 default response
func (o *WaypointServiceUIGetActionConfigBundle2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUIGetActionConfigBundle2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfig/by-name/{action_name}][%d] WaypointService_UI_GetActionConfigBundle2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUIGetActionConfigBundle2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfig/by-name/{action_name}][%d] WaypointService_UI_GetActionConfigBundle2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUIGetActionConfigBundle2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUIGetActionConfigBundle2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
