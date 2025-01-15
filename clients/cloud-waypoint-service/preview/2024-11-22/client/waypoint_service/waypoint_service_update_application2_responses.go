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

// WaypointServiceUpdateApplication2Reader is a Reader for the WaypointServiceUpdateApplication2 structure.
type WaypointServiceUpdateApplication2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplication2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplication2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplication2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplication2OK creates a WaypointServiceUpdateApplication2OK with default headers values
func NewWaypointServiceUpdateApplication2OK() *WaypointServiceUpdateApplication2OK {
	return &WaypointServiceUpdateApplication2OK{}
}

/*
WaypointServiceUpdateApplication2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplication2OK struct {
	Payload *models.HashicorpCloudWaypointUpdateApplicationResponse
}

// IsSuccess returns true when this waypoint service update application2 o k response has a 2xx status code
func (o *WaypointServiceUpdateApplication2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application2 o k response has a 3xx status code
func (o *WaypointServiceUpdateApplication2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application2 o k response has a 4xx status code
func (o *WaypointServiceUpdateApplication2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application2 o k response has a 5xx status code
func (o *WaypointServiceUpdateApplication2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application2 o k response a status code equal to that given
func (o *WaypointServiceUpdateApplication2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application2 o k response
func (o *WaypointServiceUpdateApplication2OK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplication2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/by-name/{application.name}][%d] waypointServiceUpdateApplication2OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplication2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/by-name/{application.name}][%d] waypointServiceUpdateApplication2OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplication2OK) GetPayload() *models.HashicorpCloudWaypointUpdateApplicationResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplication2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUpdateApplicationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplication2Default creates a WaypointServiceUpdateApplication2Default with default headers values
func NewWaypointServiceUpdateApplication2Default(code int) *WaypointServiceUpdateApplication2Default {
	return &WaypointServiceUpdateApplication2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplication2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplication2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application2 default response has a 2xx status code
func (o *WaypointServiceUpdateApplication2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application2 default response has a 3xx status code
func (o *WaypointServiceUpdateApplication2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application2 default response has a 4xx status code
func (o *WaypointServiceUpdateApplication2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application2 default response has a 5xx status code
func (o *WaypointServiceUpdateApplication2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application2 default response a status code equal to that given
func (o *WaypointServiceUpdateApplication2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application2 default response
func (o *WaypointServiceUpdateApplication2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplication2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/by-name/{application.name}][%d] WaypointService_UpdateApplication2 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplication2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/applications/by-name/{application.name}][%d] WaypointService_UpdateApplication2 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplication2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplication2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
