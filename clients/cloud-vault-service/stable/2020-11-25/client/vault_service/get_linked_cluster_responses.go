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

// GetLinkedClusterReader is a Reader for the GetLinkedCluster structure.
type GetLinkedClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLinkedClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLinkedClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLinkedClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLinkedClusterOK creates a GetLinkedClusterOK with default headers values
func NewGetLinkedClusterOK() *GetLinkedClusterOK {
	return &GetLinkedClusterOK{}
}

/*GetLinkedClusterOK handles this case with default header values.

A successful response.
*/
type GetLinkedClusterOK struct {
	Payload *models.HashicorpCloudVault20201125GetLinkedClusterResponse
}

func (o *GetLinkedClusterOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/link/clusters/{cluster_id}][%d] getLinkedClusterOK  %+v", 200, o.Payload)
}

func (o *GetLinkedClusterOK) GetPayload() *models.HashicorpCloudVault20201125GetLinkedClusterResponse {
	return o.Payload
}

func (o *GetLinkedClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetLinkedClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLinkedClusterDefault creates a GetLinkedClusterDefault with default headers values
func NewGetLinkedClusterDefault(code int) *GetLinkedClusterDefault {
	return &GetLinkedClusterDefault{
		_statusCode: code,
	}
}

/*GetLinkedClusterDefault handles this case with default header values.

An unexpected error response.
*/
type GetLinkedClusterDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get linked cluster default response
func (o *GetLinkedClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetLinkedClusterDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/link/clusters/{cluster_id}][%d] GetLinkedCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetLinkedClusterDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetLinkedClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
