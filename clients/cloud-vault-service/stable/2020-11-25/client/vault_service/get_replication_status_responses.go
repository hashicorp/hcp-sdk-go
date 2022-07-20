// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// GetReplicationStatusReader is a Reader for the GetReplicationStatus structure.
type GetReplicationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReplicationStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReplicationStatusOK creates a GetReplicationStatusOK with default headers values
func NewGetReplicationStatusOK() *GetReplicationStatusOK {
	return &GetReplicationStatusOK{}
}

/*GetReplicationStatusOK handles this case with default header values.

A successful response.
*/
type GetReplicationStatusOK struct {
	Payload *models.HashicorpCloudVault20201125GetReplicationStatusResponse
}

func (o *GetReplicationStatusOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/replication-status][%d] getReplicationStatusOK  %+v", 200, o.Payload)
}

func (o *GetReplicationStatusOK) GetPayload() *models.HashicorpCloudVault20201125GetReplicationStatusResponse {
	return o.Payload
}

func (o *GetReplicationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetReplicationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationStatusDefault creates a GetReplicationStatusDefault with default headers values
func NewGetReplicationStatusDefault(code int) *GetReplicationStatusDefault {
	return &GetReplicationStatusDefault{
		_statusCode: code,
	}
}

/*GetReplicationStatusDefault handles this case with default header values.

An unexpected error response.
*/
type GetReplicationStatusDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get replication status default response
func (o *GetReplicationStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetReplicationStatusDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/replication-status][%d] GetReplicationStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetReplicationStatusDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetReplicationStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
