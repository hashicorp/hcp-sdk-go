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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// WaypointServiceGetActionRunReader is a Reader for the WaypointServiceGetActionRun structure.
type WaypointServiceGetActionRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetActionRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetActionRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetActionRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetActionRunOK creates a WaypointServiceGetActionRunOK with default headers values
func NewWaypointServiceGetActionRunOK() *WaypointServiceGetActionRunOK {
	return &WaypointServiceGetActionRunOK{}
}

/*
WaypointServiceGetActionRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetActionRunOK struct {
	Payload *models.HashicorpCloudWaypointGetActionRunResponse
}

// IsSuccess returns true when this waypoint service get action run o k response has a 2xx status code
func (o *WaypointServiceGetActionRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get action run o k response has a 3xx status code
func (o *WaypointServiceGetActionRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get action run o k response has a 4xx status code
func (o *WaypointServiceGetActionRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get action run o k response has a 5xx status code
func (o *WaypointServiceGetActionRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get action run o k response a status code equal to that given
func (o *WaypointServiceGetActionRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get action run o k response
func (o *WaypointServiceGetActionRunOK) Code() int {
	return 200
}

func (o *WaypointServiceGetActionRunOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/action][%d] waypointServiceGetActionRunOK %s", 200, payload)
}

func (o *WaypointServiceGetActionRunOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/action][%d] waypointServiceGetActionRunOK %s", 200, payload)
}

func (o *WaypointServiceGetActionRunOK) GetPayload() *models.HashicorpCloudWaypointGetActionRunResponse {
	return o.Payload
}

func (o *WaypointServiceGetActionRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetActionRunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetActionRunDefault creates a WaypointServiceGetActionRunDefault with default headers values
func NewWaypointServiceGetActionRunDefault(code int) *WaypointServiceGetActionRunDefault {
	return &WaypointServiceGetActionRunDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetActionRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetActionRunDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get action run default response has a 2xx status code
func (o *WaypointServiceGetActionRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get action run default response has a 3xx status code
func (o *WaypointServiceGetActionRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get action run default response has a 4xx status code
func (o *WaypointServiceGetActionRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get action run default response has a 5xx status code
func (o *WaypointServiceGetActionRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get action run default response a status code equal to that given
func (o *WaypointServiceGetActionRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get action run default response
func (o *WaypointServiceGetActionRunDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetActionRunDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/action][%d] WaypointService_GetActionRun default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetActionRunDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/action][%d] WaypointService_GetActionRun default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetActionRunDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetActionRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
