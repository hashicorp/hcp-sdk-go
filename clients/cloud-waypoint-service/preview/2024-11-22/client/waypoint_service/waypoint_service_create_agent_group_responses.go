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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2024-11-22/models"
)

// WaypointServiceCreateAgentGroupReader is a Reader for the WaypointServiceCreateAgentGroup structure.
type WaypointServiceCreateAgentGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceCreateAgentGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceCreateAgentGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceCreateAgentGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceCreateAgentGroupOK creates a WaypointServiceCreateAgentGroupOK with default headers values
func NewWaypointServiceCreateAgentGroupOK() *WaypointServiceCreateAgentGroupOK {
	return &WaypointServiceCreateAgentGroupOK{}
}

/*
WaypointServiceCreateAgentGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceCreateAgentGroupOK struct {
	Payload models.HashicorpCloudWaypointV20241122CreateAgentGroupResponse
}

// IsSuccess returns true when this waypoint service create agent group o k response has a 2xx status code
func (o *WaypointServiceCreateAgentGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service create agent group o k response has a 3xx status code
func (o *WaypointServiceCreateAgentGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service create agent group o k response has a 4xx status code
func (o *WaypointServiceCreateAgentGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service create agent group o k response has a 5xx status code
func (o *WaypointServiceCreateAgentGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service create agent group o k response a status code equal to that given
func (o *WaypointServiceCreateAgentGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service create agent group o k response
func (o *WaypointServiceCreateAgentGroupOK) Code() int {
	return 200
}

func (o *WaypointServiceCreateAgentGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] waypointServiceCreateAgentGroupOK %s", 200, payload)
}

func (o *WaypointServiceCreateAgentGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] waypointServiceCreateAgentGroupOK %s", 200, payload)
}

func (o *WaypointServiceCreateAgentGroupOK) GetPayload() models.HashicorpCloudWaypointV20241122CreateAgentGroupResponse {
	return o.Payload
}

func (o *WaypointServiceCreateAgentGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceCreateAgentGroupDefault creates a WaypointServiceCreateAgentGroupDefault with default headers values
func NewWaypointServiceCreateAgentGroupDefault(code int) *WaypointServiceCreateAgentGroupDefault {
	return &WaypointServiceCreateAgentGroupDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceCreateAgentGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceCreateAgentGroupDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service create agent group default response has a 2xx status code
func (o *WaypointServiceCreateAgentGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service create agent group default response has a 3xx status code
func (o *WaypointServiceCreateAgentGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service create agent group default response has a 4xx status code
func (o *WaypointServiceCreateAgentGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service create agent group default response has a 5xx status code
func (o *WaypointServiceCreateAgentGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service create agent group default response a status code equal to that given
func (o *WaypointServiceCreateAgentGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service create agent group default response
func (o *WaypointServiceCreateAgentGroupDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceCreateAgentGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] WaypointService_CreateAgentGroup default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateAgentGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] WaypointService_CreateAgentGroup default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateAgentGroupDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceCreateAgentGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
