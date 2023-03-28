// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// DeletePeeringConnectionReader is a Reader for the DeletePeeringConnection structure.
type DeletePeeringConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeeringConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePeeringConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePeeringConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePeeringConnectionOK creates a DeletePeeringConnectionOK with default headers values
func NewDeletePeeringConnectionOK() *DeletePeeringConnectionOK {
	return &DeletePeeringConnectionOK{}
}

/*
DeletePeeringConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePeeringConnectionOK struct {
	Payload models.HashicorpCloudGlobalNetworkManager20220215DeletePeeringConnectionResponse
}

// IsSuccess returns true when this delete peering connection o k response has a 2xx status code
func (o *DeletePeeringConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete peering connection o k response has a 3xx status code
func (o *DeletePeeringConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peering connection o k response has a 4xx status code
func (o *DeletePeeringConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete peering connection o k response has a 5xx status code
func (o *DeletePeeringConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peering connection o k response a status code equal to that given
func (o *DeletePeeringConnectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeletePeeringConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/peering_connections/{peering_connection_id}][%d] deletePeeringConnectionOK  %+v", 200, o.Payload)
}

func (o *DeletePeeringConnectionOK) String() string {
	return fmt.Sprintf("[DELETE /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/peering_connections/{peering_connection_id}][%d] deletePeeringConnectionOK  %+v", 200, o.Payload)
}

func (o *DeletePeeringConnectionOK) GetPayload() models.HashicorpCloudGlobalNetworkManager20220215DeletePeeringConnectionResponse {
	return o.Payload
}

func (o *DeletePeeringConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePeeringConnectionDefault creates a DeletePeeringConnectionDefault with default headers values
func NewDeletePeeringConnectionDefault(code int) *DeletePeeringConnectionDefault {
	return &DeletePeeringConnectionDefault{
		_statusCode: code,
	}
}

/*
DeletePeeringConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePeeringConnectionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete peering connection default response
func (o *DeletePeeringConnectionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete peering connection default response has a 2xx status code
func (o *DeletePeeringConnectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete peering connection default response has a 3xx status code
func (o *DeletePeeringConnectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete peering connection default response has a 4xx status code
func (o *DeletePeeringConnectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete peering connection default response has a 5xx status code
func (o *DeletePeeringConnectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete peering connection default response a status code equal to that given
func (o *DeletePeeringConnectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePeeringConnectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/peering_connections/{peering_connection_id}][%d] DeletePeeringConnection default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeeringConnectionDefault) String() string {
	return fmt.Sprintf("[DELETE /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/peering_connections/{peering_connection_id}][%d] DeletePeeringConnection default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeeringConnectionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeletePeeringConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
