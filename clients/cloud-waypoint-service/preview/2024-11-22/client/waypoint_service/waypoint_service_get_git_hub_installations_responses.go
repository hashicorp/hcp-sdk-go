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

// WaypointServiceGetGitHubInstallationsReader is a Reader for the WaypointServiceGetGitHubInstallations structure.
type WaypointServiceGetGitHubInstallationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetGitHubInstallationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetGitHubInstallationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetGitHubInstallationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetGitHubInstallationsOK creates a WaypointServiceGetGitHubInstallationsOK with default headers values
func NewWaypointServiceGetGitHubInstallationsOK() *WaypointServiceGetGitHubInstallationsOK {
	return &WaypointServiceGetGitHubInstallationsOK{}
}

/*
WaypointServiceGetGitHubInstallationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetGitHubInstallationsOK struct {
	Payload *models.HashicorpCloudWaypointV20241122GetGitHubInstallationsResponse
}

// IsSuccess returns true when this waypoint service get git hub installations o k response has a 2xx status code
func (o *WaypointServiceGetGitHubInstallationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get git hub installations o k response has a 3xx status code
func (o *WaypointServiceGetGitHubInstallationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get git hub installations o k response has a 4xx status code
func (o *WaypointServiceGetGitHubInstallationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get git hub installations o k response has a 5xx status code
func (o *WaypointServiceGetGitHubInstallationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get git hub installations o k response a status code equal to that given
func (o *WaypointServiceGetGitHubInstallationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get git hub installations o k response
func (o *WaypointServiceGetGitHubInstallationsOK) Code() int {
	return 200
}

func (o *WaypointServiceGetGitHubInstallationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/github/installations][%d] waypointServiceGetGitHubInstallationsOK %s", 200, payload)
}

func (o *WaypointServiceGetGitHubInstallationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/github/installations][%d] waypointServiceGetGitHubInstallationsOK %s", 200, payload)
}

func (o *WaypointServiceGetGitHubInstallationsOK) GetPayload() *models.HashicorpCloudWaypointV20241122GetGitHubInstallationsResponse {
	return o.Payload
}

func (o *WaypointServiceGetGitHubInstallationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122GetGitHubInstallationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetGitHubInstallationsDefault creates a WaypointServiceGetGitHubInstallationsDefault with default headers values
func NewWaypointServiceGetGitHubInstallationsDefault(code int) *WaypointServiceGetGitHubInstallationsDefault {
	return &WaypointServiceGetGitHubInstallationsDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetGitHubInstallationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetGitHubInstallationsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get git hub installations default response has a 2xx status code
func (o *WaypointServiceGetGitHubInstallationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get git hub installations default response has a 3xx status code
func (o *WaypointServiceGetGitHubInstallationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get git hub installations default response has a 4xx status code
func (o *WaypointServiceGetGitHubInstallationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get git hub installations default response has a 5xx status code
func (o *WaypointServiceGetGitHubInstallationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get git hub installations default response a status code equal to that given
func (o *WaypointServiceGetGitHubInstallationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get git hub installations default response
func (o *WaypointServiceGetGitHubInstallationsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetGitHubInstallationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/github/installations][%d] WaypointService_GetGitHubInstallations default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetGitHubInstallationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/github/installations][%d] WaypointService_GetGitHubInstallations default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetGitHubInstallationsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetGitHubInstallationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
