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

// WaypointServiceGetAddOnReader is a Reader for the WaypointServiceGetAddOn structure.
type WaypointServiceGetAddOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetAddOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetAddOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetAddOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetAddOnOK creates a WaypointServiceGetAddOnOK with default headers values
func NewWaypointServiceGetAddOnOK() *WaypointServiceGetAddOnOK {
	return &WaypointServiceGetAddOnOK{}
}

/*
WaypointServiceGetAddOnOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetAddOnOK struct {
	Payload *models.HashicorpCloudWaypointGetAddOnResponse
}

// IsSuccess returns true when this waypoint service get add on o k response has a 2xx status code
func (o *WaypointServiceGetAddOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get add on o k response has a 3xx status code
func (o *WaypointServiceGetAddOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get add on o k response has a 4xx status code
func (o *WaypointServiceGetAddOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get add on o k response has a 5xx status code
func (o *WaypointServiceGetAddOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get add on o k response a status code equal to that given
func (o *WaypointServiceGetAddOnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get add on o k response
func (o *WaypointServiceGetAddOnOK) Code() int {
	return 200
}

func (o *WaypointServiceGetAddOnOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on/{add_on.id}][%d] waypointServiceGetAddOnOK %s", 200, payload)
}

func (o *WaypointServiceGetAddOnOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on/{add_on.id}][%d] waypointServiceGetAddOnOK %s", 200, payload)
}

func (o *WaypointServiceGetAddOnOK) GetPayload() *models.HashicorpCloudWaypointGetAddOnResponse {
	return o.Payload
}

func (o *WaypointServiceGetAddOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetAddOnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetAddOnDefault creates a WaypointServiceGetAddOnDefault with default headers values
func NewWaypointServiceGetAddOnDefault(code int) *WaypointServiceGetAddOnDefault {
	return &WaypointServiceGetAddOnDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetAddOnDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetAddOnDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get add on default response has a 2xx status code
func (o *WaypointServiceGetAddOnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get add on default response has a 3xx status code
func (o *WaypointServiceGetAddOnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get add on default response has a 4xx status code
func (o *WaypointServiceGetAddOnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get add on default response has a 5xx status code
func (o *WaypointServiceGetAddOnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get add on default response a status code equal to that given
func (o *WaypointServiceGetAddOnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get add on default response
func (o *WaypointServiceGetAddOnDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetAddOnDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on/{add_on.id}][%d] WaypointService_GetAddOn default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetAddOnDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on/{add_on.id}][%d] WaypointService_GetAddOn default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetAddOnDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetAddOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
