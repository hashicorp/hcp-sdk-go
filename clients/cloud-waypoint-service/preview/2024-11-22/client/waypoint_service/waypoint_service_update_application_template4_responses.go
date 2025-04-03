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

// WaypointServiceUpdateApplicationTemplate4Reader is a Reader for the WaypointServiceUpdateApplicationTemplate4 structure.
type WaypointServiceUpdateApplicationTemplate4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationTemplate4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationTemplate4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationTemplate4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationTemplate4OK creates a WaypointServiceUpdateApplicationTemplate4OK with default headers values
func NewWaypointServiceUpdateApplicationTemplate4OK() *WaypointServiceUpdateApplicationTemplate4OK {
	return &WaypointServiceUpdateApplicationTemplate4OK{}
}

/*
WaypointServiceUpdateApplicationTemplate4OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationTemplate4OK struct {
	Payload *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service update application template4 o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate4OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application template4 o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate4OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application template4 o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate4OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application template4 o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate4OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application template4 o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate4OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application template4 o k response
func (o *WaypointServiceUpdateApplicationTemplate4OK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationTemplate4OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/by-name/{existing_application_template.name}][%d] waypointServiceUpdateApplicationTemplate4OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate4OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/by-name/{existing_application_template.name}][%d] waypointServiceUpdateApplicationTemplate4OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate4OK) GetPayload() *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationTemplate4Default creates a WaypointServiceUpdateApplicationTemplate4Default with default headers values
func NewWaypointServiceUpdateApplicationTemplate4Default(code int) *WaypointServiceUpdateApplicationTemplate4Default {
	return &WaypointServiceUpdateApplicationTemplate4Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationTemplate4Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationTemplate4Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application template4 default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate4Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application template4 default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate4Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application template4 default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate4Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application template4 default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate4Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application template4 default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate4Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application template4 default response
func (o *WaypointServiceUpdateApplicationTemplate4Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationTemplate4Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/by-name/{existing_application_template.name}][%d] WaypointService_UpdateApplicationTemplate4 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate4Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/templates/by-name/{existing_application_template.name}][%d] WaypointService_UpdateApplicationTemplate4 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate4Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
