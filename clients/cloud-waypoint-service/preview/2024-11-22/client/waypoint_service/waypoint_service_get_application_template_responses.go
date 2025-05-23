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

// WaypointServiceGetApplicationTemplateReader is a Reader for the WaypointServiceGetApplicationTemplate structure.
type WaypointServiceGetApplicationTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetApplicationTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetApplicationTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetApplicationTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetApplicationTemplateOK creates a WaypointServiceGetApplicationTemplateOK with default headers values
func NewWaypointServiceGetApplicationTemplateOK() *WaypointServiceGetApplicationTemplateOK {
	return &WaypointServiceGetApplicationTemplateOK{}
}

/*
WaypointServiceGetApplicationTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetApplicationTemplateOK struct {
	Payload *models.HashicorpCloudWaypointV20241122GetApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service get application template o k response has a 2xx status code
func (o *WaypointServiceGetApplicationTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get application template o k response has a 3xx status code
func (o *WaypointServiceGetApplicationTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get application template o k response has a 4xx status code
func (o *WaypointServiceGetApplicationTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get application template o k response has a 5xx status code
func (o *WaypointServiceGetApplicationTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get application template o k response a status code equal to that given
func (o *WaypointServiceGetApplicationTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get application template o k response
func (o *WaypointServiceGetApplicationTemplateOK) Code() int {
	return 200
}

func (o *WaypointServiceGetApplicationTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{application_template.id}][%d] waypointServiceGetApplicationTemplateOK %s", 200, payload)
}

func (o *WaypointServiceGetApplicationTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{application_template.id}][%d] waypointServiceGetApplicationTemplateOK %s", 200, payload)
}

func (o *WaypointServiceGetApplicationTemplateOK) GetPayload() *models.HashicorpCloudWaypointV20241122GetApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceGetApplicationTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122GetApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetApplicationTemplateDefault creates a WaypointServiceGetApplicationTemplateDefault with default headers values
func NewWaypointServiceGetApplicationTemplateDefault(code int) *WaypointServiceGetApplicationTemplateDefault {
	return &WaypointServiceGetApplicationTemplateDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetApplicationTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetApplicationTemplateDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get application template default response has a 2xx status code
func (o *WaypointServiceGetApplicationTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get application template default response has a 3xx status code
func (o *WaypointServiceGetApplicationTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get application template default response has a 4xx status code
func (o *WaypointServiceGetApplicationTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get application template default response has a 5xx status code
func (o *WaypointServiceGetApplicationTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get application template default response a status code equal to that given
func (o *WaypointServiceGetApplicationTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get application template default response
func (o *WaypointServiceGetApplicationTemplateDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetApplicationTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{application_template.id}][%d] WaypointService_GetApplicationTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetApplicationTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{application_template.id}][%d] WaypointService_GetApplicationTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceGetApplicationTemplateDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetApplicationTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
