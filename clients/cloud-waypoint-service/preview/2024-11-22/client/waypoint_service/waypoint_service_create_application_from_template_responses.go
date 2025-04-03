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

// WaypointServiceCreateApplicationFromTemplateReader is a Reader for the WaypointServiceCreateApplicationFromTemplate structure.
type WaypointServiceCreateApplicationFromTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceCreateApplicationFromTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceCreateApplicationFromTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceCreateApplicationFromTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceCreateApplicationFromTemplateOK creates a WaypointServiceCreateApplicationFromTemplateOK with default headers values
func NewWaypointServiceCreateApplicationFromTemplateOK() *WaypointServiceCreateApplicationFromTemplateOK {
	return &WaypointServiceCreateApplicationFromTemplateOK{}
}

/*
WaypointServiceCreateApplicationFromTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceCreateApplicationFromTemplateOK struct {
	Payload *models.HashicorpCloudWaypointV20241122CreateApplicationFromTemplateResponse
}

// IsSuccess returns true when this waypoint service create application from template o k response has a 2xx status code
func (o *WaypointServiceCreateApplicationFromTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service create application from template o k response has a 3xx status code
func (o *WaypointServiceCreateApplicationFromTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service create application from template o k response has a 4xx status code
func (o *WaypointServiceCreateApplicationFromTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service create application from template o k response has a 5xx status code
func (o *WaypointServiceCreateApplicationFromTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service create application from template o k response a status code equal to that given
func (o *WaypointServiceCreateApplicationFromTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service create application from template o k response
func (o *WaypointServiceCreateApplicationFromTemplateOK) Code() int {
	return 200
}

func (o *WaypointServiceCreateApplicationFromTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/from-template][%d] waypointServiceCreateApplicationFromTemplateOK %s", 200, payload)
}

func (o *WaypointServiceCreateApplicationFromTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/from-template][%d] waypointServiceCreateApplicationFromTemplateOK %s", 200, payload)
}

func (o *WaypointServiceCreateApplicationFromTemplateOK) GetPayload() *models.HashicorpCloudWaypointV20241122CreateApplicationFromTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceCreateApplicationFromTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122CreateApplicationFromTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceCreateApplicationFromTemplateDefault creates a WaypointServiceCreateApplicationFromTemplateDefault with default headers values
func NewWaypointServiceCreateApplicationFromTemplateDefault(code int) *WaypointServiceCreateApplicationFromTemplateDefault {
	return &WaypointServiceCreateApplicationFromTemplateDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceCreateApplicationFromTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceCreateApplicationFromTemplateDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service create application from template default response has a 2xx status code
func (o *WaypointServiceCreateApplicationFromTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service create application from template default response has a 3xx status code
func (o *WaypointServiceCreateApplicationFromTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service create application from template default response has a 4xx status code
func (o *WaypointServiceCreateApplicationFromTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service create application from template default response has a 5xx status code
func (o *WaypointServiceCreateApplicationFromTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service create application from template default response a status code equal to that given
func (o *WaypointServiceCreateApplicationFromTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service create application from template default response
func (o *WaypointServiceCreateApplicationFromTemplateDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceCreateApplicationFromTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/from-template][%d] WaypointService_CreateApplicationFromTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateApplicationFromTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/from-template][%d] WaypointService_CreateApplicationFromTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateApplicationFromTemplateDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceCreateApplicationFromTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
