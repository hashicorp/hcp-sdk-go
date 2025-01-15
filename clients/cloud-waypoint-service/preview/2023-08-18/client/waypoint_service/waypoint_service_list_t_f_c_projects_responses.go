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

// WaypointServiceListTFCProjectsReader is a Reader for the WaypointServiceListTFCProjects structure.
type WaypointServiceListTFCProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListTFCProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListTFCProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListTFCProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListTFCProjectsOK creates a WaypointServiceListTFCProjectsOK with default headers values
func NewWaypointServiceListTFCProjectsOK() *WaypointServiceListTFCProjectsOK {
	return &WaypointServiceListTFCProjectsOK{}
}

/*
WaypointServiceListTFCProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListTFCProjectsOK struct {
	Payload *models.HashicorpCloudWaypointListTerraformCloudProjectsResponse
}

// IsSuccess returns true when this waypoint service list t f c projects o k response has a 2xx status code
func (o *WaypointServiceListTFCProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list t f c projects o k response has a 3xx status code
func (o *WaypointServiceListTFCProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list t f c projects o k response has a 4xx status code
func (o *WaypointServiceListTFCProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list t f c projects o k response has a 5xx status code
func (o *WaypointServiceListTFCProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list t f c projects o k response a status code equal to that given
func (o *WaypointServiceListTFCProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list t f c projects o k response
func (o *WaypointServiceListTFCProjectsOK) Code() int {
	return 200
}

func (o *WaypointServiceListTFCProjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-projects][%d] waypointServiceListTFCProjectsOK %s", 200, payload)
}

func (o *WaypointServiceListTFCProjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-projects][%d] waypointServiceListTFCProjectsOK %s", 200, payload)
}

func (o *WaypointServiceListTFCProjectsOK) GetPayload() *models.HashicorpCloudWaypointListTerraformCloudProjectsResponse {
	return o.Payload
}

func (o *WaypointServiceListTFCProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListTerraformCloudProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListTFCProjectsDefault creates a WaypointServiceListTFCProjectsDefault with default headers values
func NewWaypointServiceListTFCProjectsDefault(code int) *WaypointServiceListTFCProjectsDefault {
	return &WaypointServiceListTFCProjectsDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceListTFCProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListTFCProjectsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list t f c projects default response has a 2xx status code
func (o *WaypointServiceListTFCProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list t f c projects default response has a 3xx status code
func (o *WaypointServiceListTFCProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list t f c projects default response has a 4xx status code
func (o *WaypointServiceListTFCProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list t f c projects default response has a 5xx status code
func (o *WaypointServiceListTFCProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list t f c projects default response a status code equal to that given
func (o *WaypointServiceListTFCProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list t f c projects default response
func (o *WaypointServiceListTFCProjectsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListTFCProjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-projects][%d] WaypointService_ListTFCProjects default %s", o._statusCode, payload)
}

func (o *WaypointServiceListTFCProjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-projects][%d] WaypointService_ListTFCProjects default %s", o._statusCode, payload)
}

func (o *WaypointServiceListTFCProjectsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListTFCProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
