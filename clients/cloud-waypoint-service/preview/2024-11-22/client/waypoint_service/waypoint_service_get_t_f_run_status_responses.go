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

// WaypointServiceGetTFRunStatusReader is a Reader for the WaypointServiceGetTFRunStatus structure.
type WaypointServiceGetTFRunStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetTFRunStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetTFRunStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetTFRunStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetTFRunStatusOK creates a WaypointServiceGetTFRunStatusOK with default headers values
func NewWaypointServiceGetTFRunStatusOK() *WaypointServiceGetTFRunStatusOK {
	return &WaypointServiceGetTFRunStatusOK{}
}

/*
WaypointServiceGetTFRunStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetTFRunStatusOK struct {
	Payload *models.HashicorpCloudWaypointGetTFRunStatusResponse
}

// IsSuccess returns true when this waypoint service get t f run status o k response has a 2xx status code
func (o *WaypointServiceGetTFRunStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get t f run status o k response has a 3xx status code
func (o *WaypointServiceGetTFRunStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get t f run status o k response has a 4xx status code
func (o *WaypointServiceGetTFRunStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get t f run status o k response has a 5xx status code
func (o *WaypointServiceGetTFRunStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get t f run status o k response a status code equal to that given
func (o *WaypointServiceGetTFRunStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get t f run status o k response
func (o *WaypointServiceGetTFRunStatusOK) Code() int {
	return 200
}

func (o *WaypointServiceGetTFRunStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tf-run-status/{workspace_name}][%d] waypointServiceGetTFRunStatusOK %s", 200, payload)
}

func (o *WaypointServiceGetTFRunStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tf-run-status/{workspace_name}][%d] waypointServiceGetTFRunStatusOK %s", 200, payload)
}

func (o *WaypointServiceGetTFRunStatusOK) GetPayload() *models.HashicorpCloudWaypointGetTFRunStatusResponse {
	return o.Payload
}

func (o *WaypointServiceGetTFRunStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetTFRunStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetTFRunStatusDefault creates a WaypointServiceGetTFRunStatusDefault with default headers values
func NewWaypointServiceGetTFRunStatusDefault(code int) *WaypointServiceGetTFRunStatusDefault {
	return &WaypointServiceGetTFRunStatusDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetTFRunStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetTFRunStatusDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get t f run status default response has a 2xx status code
func (o *WaypointServiceGetTFRunStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get t f run status default response has a 3xx status code
func (o *WaypointServiceGetTFRunStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get t f run status default response has a 4xx status code
func (o *WaypointServiceGetTFRunStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get t f run status default response has a 5xx status code
func (o *WaypointServiceGetTFRunStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get t f run status default response a status code equal to that given
func (o *WaypointServiceGetTFRunStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get t f run status default response
func (o *WaypointServiceGetTFRunStatusDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetTFRunStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tf-run-status/{workspace_name}][%d] WaypointService_GetTFRunStatus default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetTFRunStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tf-run-status/{workspace_name}][%d] WaypointService_GetTFRunStatus default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetTFRunStatusDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetTFRunStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
