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

// WaypointServiceListAgentGroupsReader is a Reader for the WaypointServiceListAgentGroups structure.
type WaypointServiceListAgentGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListAgentGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListAgentGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListAgentGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListAgentGroupsOK creates a WaypointServiceListAgentGroupsOK with default headers values
func NewWaypointServiceListAgentGroupsOK() *WaypointServiceListAgentGroupsOK {
	return &WaypointServiceListAgentGroupsOK{}
}

/*
WaypointServiceListAgentGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListAgentGroupsOK struct {
	Payload *models.HashicorpCloudWaypointV20241122ListAgentGroupsResponse
}

// IsSuccess returns true when this waypoint service list agent groups o k response has a 2xx status code
func (o *WaypointServiceListAgentGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list agent groups o k response has a 3xx status code
func (o *WaypointServiceListAgentGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list agent groups o k response has a 4xx status code
func (o *WaypointServiceListAgentGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list agent groups o k response has a 5xx status code
func (o *WaypointServiceListAgentGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list agent groups o k response a status code equal to that given
func (o *WaypointServiceListAgentGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list agent groups o k response
func (o *WaypointServiceListAgentGroupsOK) Code() int {
	return 200
}

func (o *WaypointServiceListAgentGroupsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] waypointServiceListAgentGroupsOK %s", 200, payload)
}

func (o *WaypointServiceListAgentGroupsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] waypointServiceListAgentGroupsOK %s", 200, payload)
}

func (o *WaypointServiceListAgentGroupsOK) GetPayload() *models.HashicorpCloudWaypointV20241122ListAgentGroupsResponse {
	return o.Payload
}

func (o *WaypointServiceListAgentGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122ListAgentGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListAgentGroupsDefault creates a WaypointServiceListAgentGroupsDefault with default headers values
func NewWaypointServiceListAgentGroupsDefault(code int) *WaypointServiceListAgentGroupsDefault {
	return &WaypointServiceListAgentGroupsDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceListAgentGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListAgentGroupsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list agent groups default response has a 2xx status code
func (o *WaypointServiceListAgentGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list agent groups default response has a 3xx status code
func (o *WaypointServiceListAgentGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list agent groups default response has a 4xx status code
func (o *WaypointServiceListAgentGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list agent groups default response has a 5xx status code
func (o *WaypointServiceListAgentGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list agent groups default response a status code equal to that given
func (o *WaypointServiceListAgentGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list agent groups default response
func (o *WaypointServiceListAgentGroupsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListAgentGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] WaypointService_ListAgentGroups default %s", o._statusCode, payload)
}

func (o *WaypointServiceListAgentGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group][%d] WaypointService_ListAgentGroups default %s", o._statusCode, payload)
}

func (o *WaypointServiceListAgentGroupsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListAgentGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
