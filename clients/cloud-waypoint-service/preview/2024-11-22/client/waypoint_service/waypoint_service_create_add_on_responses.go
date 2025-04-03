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

// WaypointServiceCreateAddOnReader is a Reader for the WaypointServiceCreateAddOn structure.
type WaypointServiceCreateAddOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceCreateAddOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceCreateAddOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceCreateAddOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceCreateAddOnOK creates a WaypointServiceCreateAddOnOK with default headers values
func NewWaypointServiceCreateAddOnOK() *WaypointServiceCreateAddOnOK {
	return &WaypointServiceCreateAddOnOK{}
}

/*
WaypointServiceCreateAddOnOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceCreateAddOnOK struct {
	Payload *models.HashicorpCloudWaypointV20241122CreateAddOnResponse
}

// IsSuccess returns true when this waypoint service create add on o k response has a 2xx status code
func (o *WaypointServiceCreateAddOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service create add on o k response has a 3xx status code
func (o *WaypointServiceCreateAddOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service create add on o k response has a 4xx status code
func (o *WaypointServiceCreateAddOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service create add on o k response has a 5xx status code
func (o *WaypointServiceCreateAddOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service create add on o k response a status code equal to that given
func (o *WaypointServiceCreateAddOnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service create add on o k response
func (o *WaypointServiceCreateAddOnOK) Code() int {
	return 200
}

func (o *WaypointServiceCreateAddOnOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on][%d] waypointServiceCreateAddOnOK %s", 200, payload)
}

func (o *WaypointServiceCreateAddOnOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on][%d] waypointServiceCreateAddOnOK %s", 200, payload)
}

func (o *WaypointServiceCreateAddOnOK) GetPayload() *models.HashicorpCloudWaypointV20241122CreateAddOnResponse {
	return o.Payload
}

func (o *WaypointServiceCreateAddOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122CreateAddOnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceCreateAddOnDefault creates a WaypointServiceCreateAddOnDefault with default headers values
func NewWaypointServiceCreateAddOnDefault(code int) *WaypointServiceCreateAddOnDefault {
	return &WaypointServiceCreateAddOnDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceCreateAddOnDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceCreateAddOnDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service create add on default response has a 2xx status code
func (o *WaypointServiceCreateAddOnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service create add on default response has a 3xx status code
func (o *WaypointServiceCreateAddOnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service create add on default response has a 4xx status code
func (o *WaypointServiceCreateAddOnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service create add on default response has a 5xx status code
func (o *WaypointServiceCreateAddOnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service create add on default response a status code equal to that given
func (o *WaypointServiceCreateAddOnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service create add on default response
func (o *WaypointServiceCreateAddOnDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceCreateAddOnDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on][%d] WaypointService_CreateAddOn default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateAddOnDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/add-on][%d] WaypointService_CreateAddOn default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateAddOnDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceCreateAddOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
