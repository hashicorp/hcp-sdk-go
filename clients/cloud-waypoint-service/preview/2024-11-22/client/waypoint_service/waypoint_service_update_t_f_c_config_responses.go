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

// WaypointServiceUpdateTFCConfigReader is a Reader for the WaypointServiceUpdateTFCConfig structure.
type WaypointServiceUpdateTFCConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateTFCConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateTFCConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateTFCConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateTFCConfigOK creates a WaypointServiceUpdateTFCConfigOK with default headers values
func NewWaypointServiceUpdateTFCConfigOK() *WaypointServiceUpdateTFCConfigOK {
	return &WaypointServiceUpdateTFCConfigOK{}
}

/*
WaypointServiceUpdateTFCConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateTFCConfigOK struct {
	Payload *models.HashicorpCloudWaypointV20241122UpdateTFCConfigResponse
}

// IsSuccess returns true when this waypoint service update t f c config o k response has a 2xx status code
func (o *WaypointServiceUpdateTFCConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update t f c config o k response has a 3xx status code
func (o *WaypointServiceUpdateTFCConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update t f c config o k response has a 4xx status code
func (o *WaypointServiceUpdateTFCConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update t f c config o k response has a 5xx status code
func (o *WaypointServiceUpdateTFCConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update t f c config o k response a status code equal to that given
func (o *WaypointServiceUpdateTFCConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update t f c config o k response
func (o *WaypointServiceUpdateTFCConfigOK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateTFCConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] waypointServiceUpdateTFCConfigOK %s", 200, payload)
}

func (o *WaypointServiceUpdateTFCConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] waypointServiceUpdateTFCConfigOK %s", 200, payload)
}

func (o *WaypointServiceUpdateTFCConfigOK) GetPayload() *models.HashicorpCloudWaypointV20241122UpdateTFCConfigResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateTFCConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122UpdateTFCConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateTFCConfigDefault creates a WaypointServiceUpdateTFCConfigDefault with default headers values
func NewWaypointServiceUpdateTFCConfigDefault(code int) *WaypointServiceUpdateTFCConfigDefault {
	return &WaypointServiceUpdateTFCConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateTFCConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateTFCConfigDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update t f c config default response has a 2xx status code
func (o *WaypointServiceUpdateTFCConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update t f c config default response has a 3xx status code
func (o *WaypointServiceUpdateTFCConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update t f c config default response has a 4xx status code
func (o *WaypointServiceUpdateTFCConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update t f c config default response has a 5xx status code
func (o *WaypointServiceUpdateTFCConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update t f c config default response a status code equal to that given
func (o *WaypointServiceUpdateTFCConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update t f c config default response
func (o *WaypointServiceUpdateTFCConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateTFCConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] WaypointService_UpdateTFCConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateTFCConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfcconfig][%d] WaypointService_UpdateTFCConfig default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateTFCConfigDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateTFCConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
