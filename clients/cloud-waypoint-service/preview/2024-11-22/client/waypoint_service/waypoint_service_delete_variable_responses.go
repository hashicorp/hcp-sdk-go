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

// WaypointServiceDeleteVariableReader is a Reader for the WaypointServiceDeleteVariable structure.
type WaypointServiceDeleteVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceDeleteVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceDeleteVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceDeleteVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceDeleteVariableOK creates a WaypointServiceDeleteVariableOK with default headers values
func NewWaypointServiceDeleteVariableOK() *WaypointServiceDeleteVariableOK {
	return &WaypointServiceDeleteVariableOK{}
}

/*
WaypointServiceDeleteVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceDeleteVariableOK struct {
	Payload models.HashicorpCloudWaypointV20241122DeleteVariableResponse
}

// IsSuccess returns true when this waypoint service delete variable o k response has a 2xx status code
func (o *WaypointServiceDeleteVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service delete variable o k response has a 3xx status code
func (o *WaypointServiceDeleteVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service delete variable o k response has a 4xx status code
func (o *WaypointServiceDeleteVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service delete variable o k response has a 5xx status code
func (o *WaypointServiceDeleteVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service delete variable o k response a status code equal to that given
func (o *WaypointServiceDeleteVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service delete variable o k response
func (o *WaypointServiceDeleteVariableOK) Code() int {
	return 200
}

func (o *WaypointServiceDeleteVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] waypointServiceDeleteVariableOK %s", 200, payload)
}

func (o *WaypointServiceDeleteVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] waypointServiceDeleteVariableOK %s", 200, payload)
}

func (o *WaypointServiceDeleteVariableOK) GetPayload() models.HashicorpCloudWaypointV20241122DeleteVariableResponse {
	return o.Payload
}

func (o *WaypointServiceDeleteVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceDeleteVariableDefault creates a WaypointServiceDeleteVariableDefault with default headers values
func NewWaypointServiceDeleteVariableDefault(code int) *WaypointServiceDeleteVariableDefault {
	return &WaypointServiceDeleteVariableDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceDeleteVariableDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceDeleteVariableDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service delete variable default response has a 2xx status code
func (o *WaypointServiceDeleteVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service delete variable default response has a 3xx status code
func (o *WaypointServiceDeleteVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service delete variable default response has a 4xx status code
func (o *WaypointServiceDeleteVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service delete variable default response has a 5xx status code
func (o *WaypointServiceDeleteVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service delete variable default response a status code equal to that given
func (o *WaypointServiceDeleteVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service delete variable default response
func (o *WaypointServiceDeleteVariableDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceDeleteVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] WaypointService_DeleteVariable default %s", o._statusCode, payload)
}

func (o *WaypointServiceDeleteVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/variable][%d] WaypointService_DeleteVariable default %s", o._statusCode, payload)
}

func (o *WaypointServiceDeleteVariableDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceDeleteVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
