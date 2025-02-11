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

// WaypointServiceGetActionConfigReader is a Reader for the WaypointServiceGetActionConfig structure.
type WaypointServiceGetActionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetActionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetActionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetActionConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetActionConfigOK creates a WaypointServiceGetActionConfigOK with default headers values
func NewWaypointServiceGetActionConfigOK() *WaypointServiceGetActionConfigOK {
	return &WaypointServiceGetActionConfigOK{}
}

/*
WaypointServiceGetActionConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetActionConfigOK struct {
	Payload *models.HashicorpCloudWaypointGetActionConfigResponse
}

// IsSuccess returns true when this waypoint service get action config o k response has a 2xx status code
func (o *WaypointServiceGetActionConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get action config o k response has a 3xx status code
func (o *WaypointServiceGetActionConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get action config o k response has a 4xx status code
func (o *WaypointServiceGetActionConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get action config o k response has a 5xx status code
func (o *WaypointServiceGetActionConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get action config o k response a status code equal to that given
func (o *WaypointServiceGetActionConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get action config o k response
func (o *WaypointServiceGetActionConfigOK) Code() int {
	return 200
}

func (o *WaypointServiceGetActionConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/actionconfig][%d] waypointServiceGetActionConfigOK %s", 200, payload)
}

func (o *WaypointServiceGetActionConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/actionconfig][%d] waypointServiceGetActionConfigOK %s", 200, payload)
}

func (o *WaypointServiceGetActionConfigOK) GetPayload() *models.HashicorpCloudWaypointGetActionConfigResponse {
	return o.Payload
}

func (o *WaypointServiceGetActionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetActionConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetActionConfigDefault creates a WaypointServiceGetActionConfigDefault with default headers values
func NewWaypointServiceGetActionConfigDefault(code int) *WaypointServiceGetActionConfigDefault {
	return &WaypointServiceGetActionConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetActionConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetActionConfigDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get action config default response has a 2xx status code
func (o *WaypointServiceGetActionConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get action config default response has a 3xx status code
func (o *WaypointServiceGetActionConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get action config default response has a 4xx status code
func (o *WaypointServiceGetActionConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get action config default response has a 5xx status code
func (o *WaypointServiceGetActionConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get action config default response a status code equal to that given
func (o *WaypointServiceGetActionConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get action config default response
func (o *WaypointServiceGetActionConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetActionConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/actionconfig][%d] WaypointService_GetActionConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetActionConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/actionconfig][%d] WaypointService_GetActionConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetActionConfigDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetActionConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
