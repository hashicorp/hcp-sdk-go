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

// WaypointServiceUIListActionConfigBundlesReader is a Reader for the WaypointServiceUIListActionConfigBundles structure.
type WaypointServiceUIListActionConfigBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUIListActionConfigBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUIListActionConfigBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUIListActionConfigBundlesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUIListActionConfigBundlesOK creates a WaypointServiceUIListActionConfigBundlesOK with default headers values
func NewWaypointServiceUIListActionConfigBundlesOK() *WaypointServiceUIListActionConfigBundlesOK {
	return &WaypointServiceUIListActionConfigBundlesOK{}
}

/*
WaypointServiceUIListActionConfigBundlesOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUIListActionConfigBundlesOK struct {
	Payload *models.HashicorpCloudWaypointUIListActionConfigResponse
}

// IsSuccess returns true when this waypoint service Ui list action config bundles o k response has a 2xx status code
func (o *WaypointServiceUIListActionConfigBundlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service Ui list action config bundles o k response has a 3xx status code
func (o *WaypointServiceUIListActionConfigBundlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service Ui list action config bundles o k response has a 4xx status code
func (o *WaypointServiceUIListActionConfigBundlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service Ui list action config bundles o k response has a 5xx status code
func (o *WaypointServiceUIListActionConfigBundlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service Ui list action config bundles o k response a status code equal to that given
func (o *WaypointServiceUIListActionConfigBundlesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service Ui list action config bundles o k response
func (o *WaypointServiceUIListActionConfigBundlesOK) Code() int {
	return 200
}

func (o *WaypointServiceUIListActionConfigBundlesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfigs][%d] waypointServiceUiListActionConfigBundlesOK %s", 200, payload)
}

func (o *WaypointServiceUIListActionConfigBundlesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfigs][%d] waypointServiceUiListActionConfigBundlesOK %s", 200, payload)
}

func (o *WaypointServiceUIListActionConfigBundlesOK) GetPayload() *models.HashicorpCloudWaypointUIListActionConfigResponse {
	return o.Payload
}

func (o *WaypointServiceUIListActionConfigBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUIListActionConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUIListActionConfigBundlesDefault creates a WaypointServiceUIListActionConfigBundlesDefault with default headers values
func NewWaypointServiceUIListActionConfigBundlesDefault(code int) *WaypointServiceUIListActionConfigBundlesDefault {
	return &WaypointServiceUIListActionConfigBundlesDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUIListActionConfigBundlesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUIListActionConfigBundlesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service UI list action config bundles default response has a 2xx status code
func (o *WaypointServiceUIListActionConfigBundlesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service UI list action config bundles default response has a 3xx status code
func (o *WaypointServiceUIListActionConfigBundlesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service UI list action config bundles default response has a 4xx status code
func (o *WaypointServiceUIListActionConfigBundlesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service UI list action config bundles default response has a 5xx status code
func (o *WaypointServiceUIListActionConfigBundlesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service UI list action config bundles default response a status code equal to that given
func (o *WaypointServiceUIListActionConfigBundlesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service UI list action config bundles default response
func (o *WaypointServiceUIListActionConfigBundlesDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUIListActionConfigBundlesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfigs][%d] WaypointService_UIListActionConfigBundles default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIListActionConfigBundlesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/actionconfigs][%d] WaypointService_UIListActionConfigBundles default %s", o._statusCode, payload)
}

func (o *WaypointServiceUIListActionConfigBundlesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUIListActionConfigBundlesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
