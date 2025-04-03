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

// WaypointServiceUpdateApplicationTemplate3Reader is a Reader for the WaypointServiceUpdateApplicationTemplate3 structure.
type WaypointServiceUpdateApplicationTemplate3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationTemplate3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationTemplate3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationTemplate3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationTemplate3OK creates a WaypointServiceUpdateApplicationTemplate3OK with default headers values
func NewWaypointServiceUpdateApplicationTemplate3OK() *WaypointServiceUpdateApplicationTemplate3OK {
	return &WaypointServiceUpdateApplicationTemplate3OK{}
}

/*
WaypointServiceUpdateApplicationTemplate3OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationTemplate3OK struct {
	Payload *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service update application template3 o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application template3 o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application template3 o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application template3 o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application template3 o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application template3 o k response
func (o *WaypointServiceUpdateApplicationTemplate3OK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationTemplate3OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplate3OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate3OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplate3OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate3OK) GetPayload() *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationTemplate3Default creates a WaypointServiceUpdateApplicationTemplate3Default with default headers values
func NewWaypointServiceUpdateApplicationTemplate3Default(code int) *WaypointServiceUpdateApplicationTemplate3Default {
	return &WaypointServiceUpdateApplicationTemplate3Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationTemplate3Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationTemplate3Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application template3 default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application template3 default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application template3 default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application template3 default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application template3 default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate3Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application template3 default response
func (o *WaypointServiceUpdateApplicationTemplate3Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationTemplate3Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate3 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate3Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate3 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate3Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
