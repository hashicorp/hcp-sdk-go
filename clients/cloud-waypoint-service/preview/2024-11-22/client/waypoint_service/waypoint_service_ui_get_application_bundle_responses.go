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

// WaypointServiceUIGetApplicationBundleReader is a Reader for the WaypointServiceUIGetApplicationBundle structure.
type WaypointServiceUIGetApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUIGetApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUIGetApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUIGetApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUIGetApplicationBundleOK creates a WaypointServiceUIGetApplicationBundleOK with default headers values
func NewWaypointServiceUIGetApplicationBundleOK() *WaypointServiceUIGetApplicationBundleOK {
	return &WaypointServiceUIGetApplicationBundleOK{}
}

/*
WaypointServiceUIGetApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUIGetApplicationBundleOK struct {
	Payload *models.HashicorpCloudWaypointUIGetApplicationBundleResponse
}

// IsSuccess returns true when this waypoint service Ui get application bundle o k response has a 2xx status code
func (o *WaypointServiceUIGetApplicationBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service Ui get application bundle o k response has a 3xx status code
func (o *WaypointServiceUIGetApplicationBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service Ui get application bundle o k response has a 4xx status code
func (o *WaypointServiceUIGetApplicationBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service Ui get application bundle o k response has a 5xx status code
func (o *WaypointServiceUIGetApplicationBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service Ui get application bundle o k response a status code equal to that given
func (o *WaypointServiceUIGetApplicationBundleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service Ui get application bundle o k response
func (o *WaypointServiceUIGetApplicationBundleOK) Code() int {
	return 200
}

func (o *WaypointServiceUIGetApplicationBundleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/{application.id}][%d] waypointServiceUiGetApplicationBundleOK %s", 200, payload)
}

func (o *WaypointServiceUIGetApplicationBundleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/{application.id}][%d] waypointServiceUiGetApplicationBundleOK %s", 200, payload)
}

func (o *WaypointServiceUIGetApplicationBundleOK) GetPayload() *models.HashicorpCloudWaypointUIGetApplicationBundleResponse {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUIGetApplicationBundleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUIGetApplicationBundleDefault creates a WaypointServiceUIGetApplicationBundleDefault with default headers values
func NewWaypointServiceUIGetApplicationBundleDefault(code int) *WaypointServiceUIGetApplicationBundleDefault {
	return &WaypointServiceUIGetApplicationBundleDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUIGetApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUIGetApplicationBundleDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service UI get application bundle default response has a 2xx status code
func (o *WaypointServiceUIGetApplicationBundleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service UI get application bundle default response has a 3xx status code
func (o *WaypointServiceUIGetApplicationBundleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service UI get application bundle default response has a 4xx status code
func (o *WaypointServiceUIGetApplicationBundleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service UI get application bundle default response has a 5xx status code
func (o *WaypointServiceUIGetApplicationBundleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service UI get application bundle default response a status code equal to that given
func (o *WaypointServiceUIGetApplicationBundleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service UI get application bundle default response
func (o *WaypointServiceUIGetApplicationBundleDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUIGetApplicationBundleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/{application.id}][%d] WaypointService_UI_GetApplicationBundle default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIGetApplicationBundleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/ui/applications/{application.id}][%d] WaypointService_UI_GetApplicationBundle default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIGetApplicationBundleDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
