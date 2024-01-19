// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceDeleteBucketReader is a Reader for the PackerServiceDeleteBucket structure.
type PackerServiceDeleteBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceDeleteBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceDeleteBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceDeleteBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceDeleteBucketOK creates a PackerServiceDeleteBucketOK with default headers values
func NewPackerServiceDeleteBucketOK() *PackerServiceDeleteBucketOK {
	return &PackerServiceDeleteBucketOK{}
}

/*
PackerServiceDeleteBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceDeleteBucketOK struct {
	Payload models.HashicorpCloudPackerDeleteBucketResponse
}

// IsSuccess returns true when this packer service delete bucket o k response has a 2xx status code
func (o *PackerServiceDeleteBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service delete bucket o k response has a 3xx status code
func (o *PackerServiceDeleteBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service delete bucket o k response has a 4xx status code
func (o *PackerServiceDeleteBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service delete bucket o k response has a 5xx status code
func (o *PackerServiceDeleteBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service delete bucket o k response a status code equal to that given
func (o *PackerServiceDeleteBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packer service delete bucket o k response
func (o *PackerServiceDeleteBucketOK) Code() int {
	return 200
}

func (o *PackerServiceDeleteBucketOK) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] packerServiceDeleteBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceDeleteBucketOK) String() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] packerServiceDeleteBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceDeleteBucketOK) GetPayload() models.HashicorpCloudPackerDeleteBucketResponse {
	return o.Payload
}

func (o *PackerServiceDeleteBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceDeleteBucketDefault creates a PackerServiceDeleteBucketDefault with default headers values
func NewPackerServiceDeleteBucketDefault(code int) *PackerServiceDeleteBucketDefault {
	return &PackerServiceDeleteBucketDefault{
		_statusCode: code,
	}
}

/*
PackerServiceDeleteBucketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceDeleteBucketDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this packer service delete bucket default response has a 2xx status code
func (o *PackerServiceDeleteBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service delete bucket default response has a 3xx status code
func (o *PackerServiceDeleteBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service delete bucket default response has a 4xx status code
func (o *PackerServiceDeleteBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service delete bucket default response has a 5xx status code
func (o *PackerServiceDeleteBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service delete bucket default response a status code equal to that given
func (o *PackerServiceDeleteBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the packer service delete bucket default response
func (o *PackerServiceDeleteBucketDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceDeleteBucketDefault) Error() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] PackerService_DeleteBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceDeleteBucketDefault) String() string {
	return fmt.Sprintf("[DELETE /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] PackerService_DeleteBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceDeleteBucketDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceDeleteBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
