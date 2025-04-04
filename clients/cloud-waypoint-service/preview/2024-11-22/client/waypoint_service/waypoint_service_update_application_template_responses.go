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

// WaypointServiceUpdateApplicationTemplateReader is a Reader for the WaypointServiceUpdateApplicationTemplate structure.
type WaypointServiceUpdateApplicationTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationTemplateOK creates a WaypointServiceUpdateApplicationTemplateOK with default headers values
func NewWaypointServiceUpdateApplicationTemplateOK() *WaypointServiceUpdateApplicationTemplateOK {
	return &WaypointServiceUpdateApplicationTemplateOK{}
}

/*
WaypointServiceUpdateApplicationTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationTemplateOK struct {
	Payload *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service update application template o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application template o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application template o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application template o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application template o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application template o k response
func (o *WaypointServiceUpdateApplicationTemplateOK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplateOK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplateOK %s", 200, payload)
}

func (o *WaypointServiceUpdateApplicationTemplateOK) GetPayload() *models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointV20241122UpdateApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationTemplateDefault creates a WaypointServiceUpdateApplicationTemplateDefault with default headers values
func NewWaypointServiceUpdateApplicationTemplateDefault(code int) *WaypointServiceUpdateApplicationTemplateDefault {
	return &WaypointServiceUpdateApplicationTemplateDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationTemplateDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application template default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application template default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application template default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application template default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application template default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application template default response
func (o *WaypointServiceUpdateApplicationTemplateDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /waypoint/2024-11-22/organizations/{namespace.location.organization_id}/projects/{namespace.location.project_id}/application-templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate default %s", o._statusCode, payload)
}

func (o *WaypointServiceUpdateApplicationTemplateDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
