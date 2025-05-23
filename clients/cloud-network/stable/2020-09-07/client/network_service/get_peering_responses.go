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

// GetPeeringReader is a Reader for the GetPeering structure.
type GetPeeringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeeringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeeringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPeeringDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeeringOK creates a GetPeeringOK with default headers values
func NewGetPeeringOK() *GetPeeringOK {
	return &GetPeeringOK{}
}

/*
GetPeeringOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPeeringOK struct {
	Payload *models.HashicorpCloudNetwork20200907GetPeeringResponse
}

// IsSuccess returns true when this get peering o k response has a 2xx status code
func (o *GetPeeringOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get peering o k response has a 3xx status code
func (o *GetPeeringOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peering o k response has a 4xx status code
func (o *GetPeeringOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get peering o k response has a 5xx status code
func (o *GetPeeringOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get peering o k response a status code equal to that given
func (o *GetPeeringOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get peering o k response
func (o *GetPeeringOK) Code() int {
	return 200
}

func (o *GetPeeringOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] getPeeringOK %s", 200, payload)
}

func (o *GetPeeringOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] getPeeringOK %s", 200, payload)
}

func (o *GetPeeringOK) GetPayload() *models.HashicorpCloudNetwork20200907GetPeeringResponse {
	return o.Payload
}

func (o *GetPeeringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907GetPeeringResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeeringDefault creates a GetPeeringDefault with default headers values
func NewGetPeeringDefault(code int) *GetPeeringDefault {
	return &GetPeeringDefault{
		_statusCode: code,
	}
}

/*
GetPeeringDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPeeringDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this get peering default response has a 2xx status code
func (o *GetPeeringDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get peering default response has a 3xx status code
func (o *GetPeeringDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get peering default response has a 4xx status code
func (o *GetPeeringDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get peering default response has a 5xx status code
func (o *GetPeeringDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get peering default response a status code equal to that given
func (o *GetPeeringDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get peering default response
func (o *GetPeeringDefault) Code() int {
	return o._statusCode
}

func (o *GetPeeringDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] GetPeering default %s", o._statusCode, payload)
}

func (o *GetPeeringDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] GetPeering default %s", o._statusCode, payload)
}

func (o *GetPeeringDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetPeeringDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
