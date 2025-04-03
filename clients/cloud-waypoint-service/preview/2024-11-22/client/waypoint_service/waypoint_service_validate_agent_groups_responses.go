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

// WaypointServiceValidateAgentGroupsReader is a Reader for the WaypointServiceValidateAgentGroups structure.
type WaypointServiceValidateAgentGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceValidateAgentGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceValidateAgentGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceValidateAgentGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceValidateAgentGroupsOK creates a WaypointServiceValidateAgentGroupsOK with default headers values
func NewWaypointServiceValidateAgentGroupsOK() *WaypointServiceValidateAgentGroupsOK {
	return &WaypointServiceValidateAgentGroupsOK{}
}

/*
WaypointServiceValidateAgentGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceValidateAgentGroupsOK struct {
	Payload *models.HashicorpCloudWaypointV20241122ValidateAgentGroupsResponse
}

// IsSuccess returns true when this waypoint service validate agent groups o k response has a 2xx status code
func (o *WaypointServiceValidateAgentGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service validate agent groups o k response has a 3xx status code
func (o *WaypointServiceValidateAgentGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service validate agent groups o k response has a 4xx status code
func (o *WaypointServiceValidateAgentGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service validate agent groups o k response has a 5xx status code
func (o *WaypointServiceValidateAgentGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service validate agent groups o k response a status code equal to that given
func (o *WaypointServiceValidateAgentGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service validate agent groups o k response
func (o *WaypointServiceValidateAgentGroupsOK) Code() int {
	return 200
}

func (o *WaypointServiceValidateAgentGroupsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group:validate][%d] waypointServiceValidateAgentGroupsOK %s", 200, payload)
}

func (o *WaypointServiceValidateAgentGroupsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group:validate][%d] waypointServiceValidateAgentGroupsOK %s", 200, payload)
}

func (o *WaypointServiceValidateAgentGroupsOK) GetPayload() *models.HashicorpCloudWaypointV20241122ValidateAgentGroupsResponse {
	return o.Payload
}

func (o *WaypointServiceValidateAgentGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122ValidateAgentGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceValidateAgentGroupsDefault creates a WaypointServiceValidateAgentGroupsDefault with default headers values
func NewWaypointServiceValidateAgentGroupsDefault(code int) *WaypointServiceValidateAgentGroupsDefault {
	return &WaypointServiceValidateAgentGroupsDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceValidateAgentGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceValidateAgentGroupsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service validate agent groups default response has a 2xx status code
func (o *WaypointServiceValidateAgentGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service validate agent groups default response has a 3xx status code
func (o *WaypointServiceValidateAgentGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service validate agent groups default response has a 4xx status code
func (o *WaypointServiceValidateAgentGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service validate agent groups default response has a 5xx status code
func (o *WaypointServiceValidateAgentGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service validate agent groups default response a status code equal to that given
func (o *WaypointServiceValidateAgentGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service validate agent groups default response
func (o *WaypointServiceValidateAgentGroupsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceValidateAgentGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group:validate][%d] WaypointService_ValidateAgentGroups default %s", o._statusCode, payload)
}

func (o *WaypointServiceValidateAgentGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/agent/group:validate][%d] WaypointService_ValidateAgentGroups default %s", o._statusCode, payload)
}

func (o *WaypointServiceValidateAgentGroupsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceValidateAgentGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
