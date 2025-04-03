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

// WaypointServiceGetVariableReader is a Reader for the WaypointServiceGetVariable structure.
type WaypointServiceGetVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetVariableOK creates a WaypointServiceGetVariableOK with default headers values
func NewWaypointServiceGetVariableOK() *WaypointServiceGetVariableOK {
	return &WaypointServiceGetVariableOK{}
}

/*
WaypointServiceGetVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetVariableOK struct {
	Payload *models.HashicorpCloudWaypointV20241122GetVariableResponse
}

// IsSuccess returns true when this waypoint service get variable o k response has a 2xx status code
func (o *WaypointServiceGetVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get variable o k response has a 3xx status code
func (o *WaypointServiceGetVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get variable o k response has a 4xx status code
func (o *WaypointServiceGetVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get variable o k response has a 5xx status code
func (o *WaypointServiceGetVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get variable o k response a status code equal to that given
func (o *WaypointServiceGetVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get variable o k response
func (o *WaypointServiceGetVariableOK) Code() int {
	return 200
}

func (o *WaypointServiceGetVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] waypointServiceGetVariableOK %s", 200, payload)
}

func (o *WaypointServiceGetVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] waypointServiceGetVariableOK %s", 200, payload)
}

func (o *WaypointServiceGetVariableOK) GetPayload() *models.HashicorpCloudWaypointV20241122GetVariableResponse {
	return o.Payload
}

func (o *WaypointServiceGetVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122GetVariableResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetVariableDefault creates a WaypointServiceGetVariableDefault with default headers values
func NewWaypointServiceGetVariableDefault(code int) *WaypointServiceGetVariableDefault {
	return &WaypointServiceGetVariableDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetVariableDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetVariableDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get variable default response has a 2xx status code
func (o *WaypointServiceGetVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get variable default response has a 3xx status code
func (o *WaypointServiceGetVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get variable default response has a 4xx status code
func (o *WaypointServiceGetVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get variable default response has a 5xx status code
func (o *WaypointServiceGetVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get variable default response a status code equal to that given
func (o *WaypointServiceGetVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get variable default response
func (o *WaypointServiceGetVariableDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] WaypointService_GetVariable default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] WaypointService_GetVariable default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetVariableDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
