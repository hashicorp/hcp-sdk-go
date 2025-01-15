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

// WaypointServiceUpdateApplicationReader is a Reader for the WaypointServiceUpdateApplication structure.
type WaypointServiceUpdateApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationOK creates a WaypointServiceUpdateApplicationOK with default headers values
func NewWaypointServiceUpdateApplicationOK() *WaypointServiceUpdateApplicationOK {
	return &WaypointServiceUpdateApplicationOK{}
}

/*
WaypointServiceUpdateApplicationOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationOK struct {
	Payload *models.HashicorpCloudWaypointUpdateApplicationResponse
}

// IsSuccess returns true when this waypoint service update application o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application o k response
func (o *WaypointServiceUpdateApplicationOK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] waypointServiceUpdateApplicationOK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] waypointServiceUpdateApplicationOK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationOK) GetPayload() *models.HashicorpCloudWaypointUpdateApplicationResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUpdateApplicationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationDefault creates a WaypointServiceUpdateApplicationDefault with default headers values
func NewWaypointServiceUpdateApplicationDefault(code int) *WaypointServiceUpdateApplicationDefault {
	return &WaypointServiceUpdateApplicationDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application default response
func (o *WaypointServiceUpdateApplicationDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] WaypointService_UpdateApplication default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] WaypointService_UpdateApplication default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
