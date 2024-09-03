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

// WaypointServiceGetTFWorkspaceInfo4Reader is a Reader for the WaypointServiceGetTFWorkspaceInfo4 structure.
type WaypointServiceGetTFWorkspaceInfo4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetTFWorkspaceInfo4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetTFWorkspaceInfo4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetTFWorkspaceInfo4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetTFWorkspaceInfo4OK creates a WaypointServiceGetTFWorkspaceInfo4OK with default headers values
func NewWaypointServiceGetTFWorkspaceInfo4OK() *WaypointServiceGetTFWorkspaceInfo4OK {
	return &WaypointServiceGetTFWorkspaceInfo4OK{}
}

/*
WaypointServiceGetTFWorkspaceInfo4OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetTFWorkspaceInfo4OK struct {
	Payload *models.HashicorpCloudWaypointGetTFWorkspaceInfoResponse
}

// IsSuccess returns true when this waypoint service get t f workspace info4 o k response has a 2xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get t f workspace info4 o k response has a 3xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get t f workspace info4 o k response has a 4xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get t f workspace info4 o k response has a 5xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get t f workspace info4 o k response a status code equal to that given
func (o *WaypointServiceGetTFWorkspaceInfo4OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get t f workspace info4 o k response
func (o *WaypointServiceGetTFWorkspaceInfo4OK) Code() int {
	return 200
}

func (o *WaypointServiceGetTFWorkspaceInfo4OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/add-on/{add_on.id}/workspace][%d] waypointServiceGetTFWorkspaceInfo4OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetTFWorkspaceInfo4OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/add-on/{add_on.id}/workspace][%d] waypointServiceGetTFWorkspaceInfo4OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetTFWorkspaceInfo4OK) GetPayload() *models.HashicorpCloudWaypointGetTFWorkspaceInfoResponse {
	return o.Payload
}

func (o *WaypointServiceGetTFWorkspaceInfo4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetTFWorkspaceInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetTFWorkspaceInfo4Default creates a WaypointServiceGetTFWorkspaceInfo4Default with default headers values
func NewWaypointServiceGetTFWorkspaceInfo4Default(code int) *WaypointServiceGetTFWorkspaceInfo4Default {
	return &WaypointServiceGetTFWorkspaceInfo4Default{
		_statusCode: code,
	}
}

/*
WaypointServiceGetTFWorkspaceInfo4Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetTFWorkspaceInfo4Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get t f workspace info4 default response has a 2xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get t f workspace info4 default response has a 3xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get t f workspace info4 default response has a 4xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get t f workspace info4 default response has a 5xx status code
func (o *WaypointServiceGetTFWorkspaceInfo4Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get t f workspace info4 default response a status code equal to that given
func (o *WaypointServiceGetTFWorkspaceInfo4Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get t f workspace info4 default response
func (o *WaypointServiceGetTFWorkspaceInfo4Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetTFWorkspaceInfo4Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/add-on/{add_on.id}/workspace][%d] WaypointService_GetTFWorkspaceInfo4 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetTFWorkspaceInfo4Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/add-on/{add_on.id}/workspace][%d] WaypointService_GetTFWorkspaceInfo4 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetTFWorkspaceInfo4Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetTFWorkspaceInfo4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
