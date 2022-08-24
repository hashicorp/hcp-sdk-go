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

// DeletePathsFilterReader is a Reader for the DeletePathsFilter structure.
type DeletePathsFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePathsFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePathsFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePathsFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePathsFilterOK creates a DeletePathsFilterOK with default headers values
func NewDeletePathsFilterOK() *DeletePathsFilterOK {
	return &DeletePathsFilterOK{}
}

/* DeletePathsFilterOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePathsFilterOK struct {
	Payload *models.HashicorpCloudVault20201125DeletePathsFilterResponse
}

func (o *DeletePathsFilterOK) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/paths-filter/delete][%d] deletePathsFilterOK  %+v", 200, o.Payload)
}
func (o *DeletePathsFilterOK) GetPayload() *models.HashicorpCloudVault20201125DeletePathsFilterResponse {
	return o.Payload
}

func (o *DeletePathsFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125DeletePathsFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePathsFilterDefault creates a DeletePathsFilterDefault with default headers values
func NewDeletePathsFilterDefault(code int) *DeletePathsFilterDefault {
	return &DeletePathsFilterDefault{
		_statusCode: code,
	}
}

/* DeletePathsFilterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePathsFilterDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete paths filter default response
func (o *DeletePathsFilterDefault) Code() int {
	return o._statusCode
}

func (o *DeletePathsFilterDefault) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/paths-filter/delete][%d] DeletePathsFilter default  %+v", o._statusCode, o.Payload)
}
func (o *DeletePathsFilterDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeletePathsFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
