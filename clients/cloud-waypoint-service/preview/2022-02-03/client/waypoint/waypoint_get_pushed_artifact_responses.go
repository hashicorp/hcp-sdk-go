// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// WaypointGetPushedArtifactReader is a Reader for the WaypointGetPushedArtifact structure.
type WaypointGetPushedArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetPushedArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetPushedArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetPushedArtifactDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetPushedArtifactOK creates a WaypointGetPushedArtifactOK with default headers values
func NewWaypointGetPushedArtifactOK() *WaypointGetPushedArtifactOK {
	return &WaypointGetPushedArtifactOK{}
}

/*
WaypointGetPushedArtifactOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetPushedArtifactOK struct {
	Payload *models.HashicorpWaypointPushedArtifact
}

// IsSuccess returns true when this waypoint get pushed artifact o k response has a 2xx status code
func (o *WaypointGetPushedArtifactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get pushed artifact o k response has a 3xx status code
func (o *WaypointGetPushedArtifactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get pushed artifact o k response has a 4xx status code
func (o *WaypointGetPushedArtifactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get pushed artifact o k response has a 5xx status code
func (o *WaypointGetPushedArtifactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get pushed artifact o k response a status code equal to that given
func (o *WaypointGetPushedArtifactOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint get pushed artifact o k response
func (o *WaypointGetPushedArtifactOK) Code() int {
	return 200
}

func (o *WaypointGetPushedArtifactOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/artifact/{ref.id}][%d] waypointGetPushedArtifactOK  %+v", 200, o.Payload)
}

func (o *WaypointGetPushedArtifactOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/artifact/{ref.id}][%d] waypointGetPushedArtifactOK  %+v", 200, o.Payload)
}

func (o *WaypointGetPushedArtifactOK) GetPayload() *models.HashicorpWaypointPushedArtifact {
	return o.Payload
}

func (o *WaypointGetPushedArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointPushedArtifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetPushedArtifactDefault creates a WaypointGetPushedArtifactDefault with default headers values
func NewWaypointGetPushedArtifactDefault(code int) *WaypointGetPushedArtifactDefault {
	return &WaypointGetPushedArtifactDefault{
		_statusCode: code,
	}
}

/*
WaypointGetPushedArtifactDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetPushedArtifactDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint get pushed artifact default response has a 2xx status code
func (o *WaypointGetPushedArtifactDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get pushed artifact default response has a 3xx status code
func (o *WaypointGetPushedArtifactDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get pushed artifact default response has a 4xx status code
func (o *WaypointGetPushedArtifactDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get pushed artifact default response has a 5xx status code
func (o *WaypointGetPushedArtifactDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get pushed artifact default response a status code equal to that given
func (o *WaypointGetPushedArtifactDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint get pushed artifact default response
func (o *WaypointGetPushedArtifactDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetPushedArtifactDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/artifact/{ref.id}][%d] Waypoint_GetPushedArtifact default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetPushedArtifactDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/artifact/{ref.id}][%d] Waypoint_GetPushedArtifact default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetPushedArtifactDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetPushedArtifactDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}