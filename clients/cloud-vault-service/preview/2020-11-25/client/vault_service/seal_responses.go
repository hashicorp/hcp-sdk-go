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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-11-25/models"
)

// SealReader is a Reader for the Seal structure.
type SealReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SealReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSealOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSealDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSealOK creates a SealOK with default headers values
func NewSealOK() *SealOK {
	return &SealOK{}
}

/* SealOK describes a response with status code 200, with default header values.

A successful response.
*/
type SealOK struct {
	Payload *models.HashicorpCloudVault20201125SealResponse
}

func (o *SealOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/seal][%d] sealOK  %+v", 200, o.Payload)
}
func (o *SealOK) GetPayload() *models.HashicorpCloudVault20201125SealResponse {
	return o.Payload
}

func (o *SealOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125SealResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSealDefault creates a SealDefault with default headers values
func NewSealDefault(code int) *SealDefault {
	return &SealDefault{
		_statusCode: code,
	}
}

/* SealDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SealDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the seal default response
func (o *SealDefault) Code() int {
	return o._statusCode
}

func (o *SealDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/seal][%d] Seal default  %+v", o._statusCode, o.Payload)
}
func (o *SealDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SealDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
