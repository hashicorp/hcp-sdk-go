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

// WaypointServiceSendStatusLogReader is a Reader for the WaypointServiceSendStatusLog structure.
type WaypointServiceSendStatusLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceSendStatusLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceSendStatusLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceSendStatusLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceSendStatusLogOK creates a WaypointServiceSendStatusLogOK with default headers values
func NewWaypointServiceSendStatusLogOK() *WaypointServiceSendStatusLogOK {
	return &WaypointServiceSendStatusLogOK{}
}

/*
WaypointServiceSendStatusLogOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceSendStatusLogOK struct {
	Payload models.HashicorpCloudWaypointSendStatusLogResponse
}

// IsSuccess returns true when this waypoint service send status log o k response has a 2xx status code
func (o *WaypointServiceSendStatusLogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service send status log o k response has a 3xx status code
func (o *WaypointServiceSendStatusLogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service send status log o k response has a 4xx status code
func (o *WaypointServiceSendStatusLogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service send status log o k response has a 5xx status code
func (o *WaypointServiceSendStatusLogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service send status log o k response a status code equal to that given
func (o *WaypointServiceSendStatusLogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service send status log o k response
func (o *WaypointServiceSendStatusLogOK) Code() int {
	return 200
}

func (o *WaypointServiceSendStatusLogOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2023-08-18/namespace/{namespace.id}/action/{action_config.id}/action-runs/{action_run_seq}/status-log][%d] waypointServiceSendStatusLogOK %s", 200, payload)
}

func (o *WaypointServiceSendStatusLogOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2023-08-18/namespace/{namespace.id}/action/{action_config.id}/action-runs/{action_run_seq}/status-log][%d] waypointServiceSendStatusLogOK %s", 200, payload)
}

func (o *WaypointServiceSendStatusLogOK) GetPayload() models.HashicorpCloudWaypointSendStatusLogResponse {
	return o.Payload
}

func (o *WaypointServiceSendStatusLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceSendStatusLogDefault creates a WaypointServiceSendStatusLogDefault with default headers values
func NewWaypointServiceSendStatusLogDefault(code int) *WaypointServiceSendStatusLogDefault {
	return &WaypointServiceSendStatusLogDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceSendStatusLogDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceSendStatusLogDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service send status log default response has a 2xx status code
func (o *WaypointServiceSendStatusLogDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service send status log default response has a 3xx status code
func (o *WaypointServiceSendStatusLogDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service send status log default response has a 4xx status code
func (o *WaypointServiceSendStatusLogDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service send status log default response has a 5xx status code
func (o *WaypointServiceSendStatusLogDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service send status log default response a status code equal to that given
func (o *WaypointServiceSendStatusLogDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service send status log default response
func (o *WaypointServiceSendStatusLogDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceSendStatusLogDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2023-08-18/namespace/{namespace.id}/action/{action_config.id}/action-runs/{action_run_seq}/status-log][%d] WaypointService_SendStatusLog default %s", o._statusCode, payload)
}

func (o *WaypointServiceSendStatusLogDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2023-08-18/namespace/{namespace.id}/action/{action_config.id}/action-runs/{action_run_seq}/status-log][%d] WaypointService_SendStatusLog default %s", o._statusCode, payload)
}

func (o *WaypointServiceSendStatusLogDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceSendStatusLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
