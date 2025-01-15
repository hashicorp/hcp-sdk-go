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

// WaypointServiceGetAgentGroupReader is a Reader for the WaypointServiceGetAgentGroup structure.
type WaypointServiceGetAgentGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetAgentGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetAgentGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetAgentGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetAgentGroupOK creates a WaypointServiceGetAgentGroupOK with default headers values
func NewWaypointServiceGetAgentGroupOK() *WaypointServiceGetAgentGroupOK {
	return &WaypointServiceGetAgentGroupOK{}
}

/*
WaypointServiceGetAgentGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetAgentGroupOK struct {
	Payload *models.HashicorpCloudWaypointGetAgentGroupResponse
}

// IsSuccess returns true when this waypoint service get agent group o k response has a 2xx status code
func (o *WaypointServiceGetAgentGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get agent group o k response has a 3xx status code
func (o *WaypointServiceGetAgentGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get agent group o k response has a 4xx status code
func (o *WaypointServiceGetAgentGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get agent group o k response has a 5xx status code
func (o *WaypointServiceGetAgentGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get agent group o k response a status code equal to that given
func (o *WaypointServiceGetAgentGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get agent group o k response
func (o *WaypointServiceGetAgentGroupOK) Code() int {
	return 200
}

func (o *WaypointServiceGetAgentGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] waypointServiceGetAgentGroupOK %s", 200, payload)
}

func (o *WaypointServiceGetAgentGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] waypointServiceGetAgentGroupOK %s", 200, payload)
}

func (o *WaypointServiceGetAgentGroupOK) GetPayload() *models.HashicorpCloudWaypointGetAgentGroupResponse {
	return o.Payload
}

func (o *WaypointServiceGetAgentGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetAgentGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetAgentGroupDefault creates a WaypointServiceGetAgentGroupDefault with default headers values
func NewWaypointServiceGetAgentGroupDefault(code int) *WaypointServiceGetAgentGroupDefault {
	return &WaypointServiceGetAgentGroupDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetAgentGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetAgentGroupDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get agent group default response has a 2xx status code
func (o *WaypointServiceGetAgentGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get agent group default response has a 3xx status code
func (o *WaypointServiceGetAgentGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get agent group default response has a 4xx status code
func (o *WaypointServiceGetAgentGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get agent group default response has a 5xx status code
func (o *WaypointServiceGetAgentGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get agent group default response a status code equal to that given
func (o *WaypointServiceGetAgentGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get agent group default response
func (o *WaypointServiceGetAgentGroupDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetAgentGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] WaypointService_GetAgentGroup default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetAgentGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/agent/group/{name}][%d] WaypointService_GetAgentGroup default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetAgentGroupDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetAgentGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
