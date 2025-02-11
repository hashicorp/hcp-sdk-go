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

// WaypointServiceListVariablesReader is a Reader for the WaypointServiceListVariables structure.
type WaypointServiceListVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListVariablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListVariablesOK creates a WaypointServiceListVariablesOK with default headers values
func NewWaypointServiceListVariablesOK() *WaypointServiceListVariablesOK {
	return &WaypointServiceListVariablesOK{}
}

/*
WaypointServiceListVariablesOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListVariablesOK struct {
	Payload *models.HashicorpCloudWaypointListVariablesResponse
}

// IsSuccess returns true when this waypoint service list variables o k response has a 2xx status code
func (o *WaypointServiceListVariablesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list variables o k response has a 3xx status code
func (o *WaypointServiceListVariablesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list variables o k response has a 4xx status code
func (o *WaypointServiceListVariablesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list variables o k response has a 5xx status code
func (o *WaypointServiceListVariablesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list variables o k response a status code equal to that given
func (o *WaypointServiceListVariablesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list variables o k response
func (o *WaypointServiceListVariablesOK) Code() int {
	return 200
}

func (o *WaypointServiceListVariablesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variables][%d] waypointServiceListVariablesOK %s", 200, payload)
}

func (o *WaypointServiceListVariablesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variables][%d] waypointServiceListVariablesOK %s", 200, payload)
}

func (o *WaypointServiceListVariablesOK) GetPayload() *models.HashicorpCloudWaypointListVariablesResponse {
	return o.Payload
}

func (o *WaypointServiceListVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListVariablesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListVariablesDefault creates a WaypointServiceListVariablesDefault with default headers values
func NewWaypointServiceListVariablesDefault(code int) *WaypointServiceListVariablesDefault {
	return &WaypointServiceListVariablesDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceListVariablesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListVariablesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list variables default response has a 2xx status code
func (o *WaypointServiceListVariablesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list variables default response has a 3xx status code
func (o *WaypointServiceListVariablesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list variables default response has a 4xx status code
func (o *WaypointServiceListVariablesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list variables default response has a 5xx status code
func (o *WaypointServiceListVariablesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list variables default response a status code equal to that given
func (o *WaypointServiceListVariablesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list variables default response
func (o *WaypointServiceListVariablesDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListVariablesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variables][%d] WaypointService_ListVariables default %s", o._statusCode, payload)
}

func (o *WaypointServiceListVariablesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variables][%d] WaypointService_ListVariables default %s", o._statusCode, payload)
}

func (o *WaypointServiceListVariablesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListVariablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
