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

// WaypointServiceCreateTFCConfigReader is a Reader for the WaypointServiceCreateTFCConfig structure.
type WaypointServiceCreateTFCConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceCreateTFCConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceCreateTFCConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceCreateTFCConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceCreateTFCConfigOK creates a WaypointServiceCreateTFCConfigOK with default headers values
func NewWaypointServiceCreateTFCConfigOK() *WaypointServiceCreateTFCConfigOK {
	return &WaypointServiceCreateTFCConfigOK{}
}

/*
WaypointServiceCreateTFCConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceCreateTFCConfigOK struct {
	Payload *models.HashicorpCloudWaypointV20241122CreateTFCConfigResponse
}

// IsSuccess returns true when this waypoint service create t f c config o k response has a 2xx status code
func (o *WaypointServiceCreateTFCConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service create t f c config o k response has a 3xx status code
func (o *WaypointServiceCreateTFCConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service create t f c config o k response has a 4xx status code
func (o *WaypointServiceCreateTFCConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service create t f c config o k response has a 5xx status code
func (o *WaypointServiceCreateTFCConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service create t f c config o k response a status code equal to that given
func (o *WaypointServiceCreateTFCConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service create t f c config o k response
func (o *WaypointServiceCreateTFCConfigOK) Code() int {
	return 200
}

func (o *WaypointServiceCreateTFCConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] waypointServiceCreateTFCConfigOK %s", 200, payload)
}

func (o *WaypointServiceCreateTFCConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] waypointServiceCreateTFCConfigOK %s", 200, payload)
}

func (o *WaypointServiceCreateTFCConfigOK) GetPayload() *models.HashicorpCloudWaypointV20241122CreateTFCConfigResponse {
	return o.Payload
}

func (o *WaypointServiceCreateTFCConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122CreateTFCConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceCreateTFCConfigDefault creates a WaypointServiceCreateTFCConfigDefault with default headers values
func NewWaypointServiceCreateTFCConfigDefault(code int) *WaypointServiceCreateTFCConfigDefault {
	return &WaypointServiceCreateTFCConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceCreateTFCConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceCreateTFCConfigDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service create t f c config default response has a 2xx status code
func (o *WaypointServiceCreateTFCConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service create t f c config default response has a 3xx status code
func (o *WaypointServiceCreateTFCConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service create t f c config default response has a 4xx status code
func (o *WaypointServiceCreateTFCConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service create t f c config default response has a 5xx status code
func (o *WaypointServiceCreateTFCConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service create t f c config default response a status code equal to that given
func (o *WaypointServiceCreateTFCConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service create t f c config default response
func (o *WaypointServiceCreateTFCConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceCreateTFCConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] WaypointService_CreateTFCConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateTFCConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] WaypointService_CreateTFCConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceCreateTFCConfigDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceCreateTFCConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
