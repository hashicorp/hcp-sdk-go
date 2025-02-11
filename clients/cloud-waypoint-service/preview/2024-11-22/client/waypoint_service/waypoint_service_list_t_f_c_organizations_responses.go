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

// WaypointServiceListTFCOrganizationsReader is a Reader for the WaypointServiceListTFCOrganizations structure.
type WaypointServiceListTFCOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListTFCOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListTFCOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListTFCOrganizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListTFCOrganizationsOK creates a WaypointServiceListTFCOrganizationsOK with default headers values
func NewWaypointServiceListTFCOrganizationsOK() *WaypointServiceListTFCOrganizationsOK {
	return &WaypointServiceListTFCOrganizationsOK{}
}

/*
WaypointServiceListTFCOrganizationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListTFCOrganizationsOK struct {
	Payload *models.HashicorpCloudWaypointListTFCOrganizationsResponse
}

// IsSuccess returns true when this waypoint service list t f c organizations o k response has a 2xx status code
func (o *WaypointServiceListTFCOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list t f c organizations o k response has a 3xx status code
func (o *WaypointServiceListTFCOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list t f c organizations o k response has a 4xx status code
func (o *WaypointServiceListTFCOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list t f c organizations o k response has a 5xx status code
func (o *WaypointServiceListTFCOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list t f c organizations o k response a status code equal to that given
func (o *WaypointServiceListTFCOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list t f c organizations o k response
func (o *WaypointServiceListTFCOrganizationsOK) Code() int {
	return 200
}

func (o *WaypointServiceListTFCOrganizationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfc-organizations][%d] waypointServiceListTFCOrganizationsOK %s", 200, payload)
}

func (o *WaypointServiceListTFCOrganizationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfc-organizations][%d] waypointServiceListTFCOrganizationsOK %s", 200, payload)
}

func (o *WaypointServiceListTFCOrganizationsOK) GetPayload() *models.HashicorpCloudWaypointListTFCOrganizationsResponse {
	return o.Payload
}

func (o *WaypointServiceListTFCOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListTFCOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListTFCOrganizationsDefault creates a WaypointServiceListTFCOrganizationsDefault with default headers values
func NewWaypointServiceListTFCOrganizationsDefault(code int) *WaypointServiceListTFCOrganizationsDefault {
	return &WaypointServiceListTFCOrganizationsDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceListTFCOrganizationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListTFCOrganizationsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list t f c organizations default response has a 2xx status code
func (o *WaypointServiceListTFCOrganizationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list t f c organizations default response has a 3xx status code
func (o *WaypointServiceListTFCOrganizationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list t f c organizations default response has a 4xx status code
func (o *WaypointServiceListTFCOrganizationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list t f c organizations default response has a 5xx status code
func (o *WaypointServiceListTFCOrganizationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list t f c organizations default response a status code equal to that given
func (o *WaypointServiceListTFCOrganizationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list t f c organizations default response
func (o *WaypointServiceListTFCOrganizationsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListTFCOrganizationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfc-organizations][%d] WaypointService_ListTFCOrganizations default %s", o._statusCode, payload)
}

func (o *WaypointServiceListTFCOrganizationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/tfc-organizations][%d] WaypointService_ListTFCOrganizations default %s", o._statusCode, payload)
}

func (o *WaypointServiceListTFCOrganizationsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListTFCOrganizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
