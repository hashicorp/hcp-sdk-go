// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetHVNRouteReader is a Reader for the GetHVNRoute structure.
type GetHVNRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHVNRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHVNRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHVNRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHVNRouteOK creates a GetHVNRouteOK with default headers values
func NewGetHVNRouteOK() *GetHVNRouteOK {
	return &GetHVNRouteOK{}
}

/*
GetHVNRouteOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetHVNRouteOK struct {
	Payload *models.HashicorpCloudNetwork20200907GetHVNRouteResponse
}

// IsSuccess returns true when this get h v n route o k response has a 2xx status code
func (o *GetHVNRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get h v n route o k response has a 3xx status code
func (o *GetHVNRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get h v n route o k response has a 4xx status code
func (o *GetHVNRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get h v n route o k response has a 5xx status code
func (o *GetHVNRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get h v n route o k response a status code equal to that given
func (o *GetHVNRouteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get h v n route o k response
func (o *GetHVNRouteOK) Code() int {
	return 200
}

func (o *GetHVNRouteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes/{id}][%d] getHVNRouteOK %s", 200, payload)
}

func (o *GetHVNRouteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes/{id}][%d] getHVNRouteOK %s", 200, payload)
}

func (o *GetHVNRouteOK) GetPayload() *models.HashicorpCloudNetwork20200907GetHVNRouteResponse {
	return o.Payload
}

func (o *GetHVNRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907GetHVNRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHVNRouteDefault creates a GetHVNRouteDefault with default headers values
func NewGetHVNRouteDefault(code int) *GetHVNRouteDefault {
	return &GetHVNRouteDefault{
		_statusCode: code,
	}
}

/*
GetHVNRouteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetHVNRouteDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this get h v n route default response has a 2xx status code
func (o *GetHVNRouteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get h v n route default response has a 3xx status code
func (o *GetHVNRouteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get h v n route default response has a 4xx status code
func (o *GetHVNRouteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get h v n route default response has a 5xx status code
func (o *GetHVNRouteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get h v n route default response a status code equal to that given
func (o *GetHVNRouteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get h v n route default response
func (o *GetHVNRouteDefault) Code() int {
	return o._statusCode
}

func (o *GetHVNRouteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes/{id}][%d] GetHVNRoute default %s", o._statusCode, payload)
}

func (o *GetHVNRouteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes/{id}][%d] GetHVNRoute default %s", o._statusCode, payload)
}

func (o *GetHVNRouteDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetHVNRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
