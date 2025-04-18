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

// WaypointServiceUIGetApplicationBundle2Reader is a Reader for the WaypointServiceUIGetApplicationBundle2 structure.
type WaypointServiceUIGetApplicationBundle2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUIGetApplicationBundle2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUIGetApplicationBundle2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUIGetApplicationBundle2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUIGetApplicationBundle2OK creates a WaypointServiceUIGetApplicationBundle2OK with default headers values
func NewWaypointServiceUIGetApplicationBundle2OK() *WaypointServiceUIGetApplicationBundle2OK {
	return &WaypointServiceUIGetApplicationBundle2OK{}
}

/*
WaypointServiceUIGetApplicationBundle2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUIGetApplicationBundle2OK struct {
	Payload *models.HashicorpCloudWaypointV20241122UIGetApplicationBundleResponse
}

// IsSuccess returns true when this waypoint service Ui get application bundle2 o k response has a 2xx status code
func (o *WaypointServiceUIGetApplicationBundle2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service Ui get application bundle2 o k response has a 3xx status code
func (o *WaypointServiceUIGetApplicationBundle2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service Ui get application bundle2 o k response has a 4xx status code
func (o *WaypointServiceUIGetApplicationBundle2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service Ui get application bundle2 o k response has a 5xx status code
func (o *WaypointServiceUIGetApplicationBundle2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service Ui get application bundle2 o k response a status code equal to that given
func (o *WaypointServiceUIGetApplicationBundle2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service Ui get application bundle2 o k response
func (o *WaypointServiceUIGetApplicationBundle2OK) Code() int {
	return 200
}

func (o *WaypointServiceUIGetApplicationBundle2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/by-name/{application.name}][%d] waypointServiceUiGetApplicationBundle2OK %s", 200, payload)
}

func (o *WaypointServiceUIGetApplicationBundle2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/by-name/{application.name}][%d] waypointServiceUiGetApplicationBundle2OK %s", 200, payload)
}

func (o *WaypointServiceUIGetApplicationBundle2OK) GetPayload() *models.HashicorpCloudWaypointV20241122UIGetApplicationBundleResponse {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationBundle2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122UIGetApplicationBundleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUIGetApplicationBundle2Default creates a WaypointServiceUIGetApplicationBundle2Default with default headers values
func NewWaypointServiceUIGetApplicationBundle2Default(code int) *WaypointServiceUIGetApplicationBundle2Default {
	return &WaypointServiceUIGetApplicationBundle2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUIGetApplicationBundle2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUIGetApplicationBundle2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service UI get application bundle2 default response has a 2xx status code
func (o *WaypointServiceUIGetApplicationBundle2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service UI get application bundle2 default response has a 3xx status code
func (o *WaypointServiceUIGetApplicationBundle2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service UI get application bundle2 default response has a 4xx status code
func (o *WaypointServiceUIGetApplicationBundle2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service UI get application bundle2 default response has a 5xx status code
func (o *WaypointServiceUIGetApplicationBundle2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service UI get application bundle2 default response a status code equal to that given
func (o *WaypointServiceUIGetApplicationBundle2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service UI get application bundle2 default response
func (o *WaypointServiceUIGetApplicationBundle2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUIGetApplicationBundle2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/by-name/{application.name}][%d] WaypointService_UI_GetApplicationBundle2 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIGetApplicationBundle2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/by-name/{application.name}][%d] WaypointService_UI_GetApplicationBundle2 default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIGetApplicationBundle2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationBundle2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
