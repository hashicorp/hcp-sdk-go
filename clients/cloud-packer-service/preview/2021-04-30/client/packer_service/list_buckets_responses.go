// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListBucketsReader is a Reader for the ListBuckets structure.
type ListBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListBucketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBucketsOK creates a ListBucketsOK with default headers values
func NewListBucketsOK() *ListBucketsOK {
	return &ListBucketsOK{}
}

/* ListBucketsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListBucketsOK struct {
	Payload *models.HashicorpCloudPackerListBucketsResponse
}

func (o *ListBucketsOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] listBucketsOK  %+v", 200, o.Payload)
}
func (o *ListBucketsOK) GetPayload() *models.HashicorpCloudPackerListBucketsResponse {
	return o.Payload
}

func (o *ListBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerListBucketsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBucketsDefault creates a ListBucketsDefault with default headers values
func NewListBucketsDefault(code int) *ListBucketsDefault {
	return &ListBucketsDefault{
		_statusCode: code,
	}
}

/* ListBucketsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListBucketsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list buckets default response
func (o *ListBucketsDefault) Code() int {
	return o._statusCode
}

func (o *ListBucketsDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] ListBuckets default  %+v", o._statusCode, o.Payload)
}
func (o *ListBucketsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListBucketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
