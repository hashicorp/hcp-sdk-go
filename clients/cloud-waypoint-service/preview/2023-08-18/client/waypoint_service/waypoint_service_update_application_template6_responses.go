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

// WaypointServiceUpdateApplicationTemplate6Reader is a Reader for the WaypointServiceUpdateApplicationTemplate6 structure.
type WaypointServiceUpdateApplicationTemplate6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationTemplate6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationTemplate6OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationTemplate6Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationTemplate6OK creates a WaypointServiceUpdateApplicationTemplate6OK with default headers values
func NewWaypointServiceUpdateApplicationTemplate6OK() *WaypointServiceUpdateApplicationTemplate6OK {
	return &WaypointServiceUpdateApplicationTemplate6OK{}
}

/*
WaypointServiceUpdateApplicationTemplate6OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationTemplate6OK struct {
	Payload *models.HashicorpCloudWaypointUpdateApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service update application template6 o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate6OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application template6 o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate6OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application template6 o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate6OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application template6 o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate6OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application template6 o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate6OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application template6 o k response
func (o *WaypointServiceUpdateApplicationTemplate6OK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationTemplate6OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/by-name/{existing_application_template.name}][%d] waypointServiceUpdateApplicationTemplate6OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate6OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/by-name/{existing_application_template.name}][%d] waypointServiceUpdateApplicationTemplate6OK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate6OK) GetPayload() *models.HashicorpCloudWaypointUpdateApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate6OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUpdateApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationTemplate6Default creates a WaypointServiceUpdateApplicationTemplate6Default with default headers values
func NewWaypointServiceUpdateApplicationTemplate6Default(code int) *WaypointServiceUpdateApplicationTemplate6Default {
	return &WaypointServiceUpdateApplicationTemplate6Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationTemplate6Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationTemplate6Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application template6 default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate6Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application template6 default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate6Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application template6 default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate6Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application template6 default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate6Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application template6 default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate6Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application template6 default response
func (o *WaypointServiceUpdateApplicationTemplate6Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationTemplate6Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/by-name/{existing_application_template.name}][%d] WaypointService_UpdateApplicationTemplate6 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate6Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/by-name/{existing_application_template.name}][%d] WaypointService_UpdateApplicationTemplate6 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplate6Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate6Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
