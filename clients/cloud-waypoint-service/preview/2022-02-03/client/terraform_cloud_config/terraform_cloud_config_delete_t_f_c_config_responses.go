// Code generated by go-swagger; DO NOT EDIT.

package terraform_cloud_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// TerraformCloudConfigDeleteTFCConfigReader is a Reader for the TerraformCloudConfigDeleteTFCConfig structure.
type TerraformCloudConfigDeleteTFCConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerraformCloudConfigDeleteTFCConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTerraformCloudConfigDeleteTFCConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTerraformCloudConfigDeleteTFCConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTerraformCloudConfigDeleteTFCConfigOK creates a TerraformCloudConfigDeleteTFCConfigOK with default headers values
func NewTerraformCloudConfigDeleteTFCConfigOK() *TerraformCloudConfigDeleteTFCConfigOK {
	return &TerraformCloudConfigDeleteTFCConfigOK{}
}

/*
TerraformCloudConfigDeleteTFCConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type TerraformCloudConfigDeleteTFCConfigOK struct {
	Payload interface{}
}

// IsSuccess returns true when this terraform cloud config delete t f c config o k response has a 2xx status code
func (o *TerraformCloudConfigDeleteTFCConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this terraform cloud config delete t f c config o k response has a 3xx status code
func (o *TerraformCloudConfigDeleteTFCConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this terraform cloud config delete t f c config o k response has a 4xx status code
func (o *TerraformCloudConfigDeleteTFCConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this terraform cloud config delete t f c config o k response has a 5xx status code
func (o *TerraformCloudConfigDeleteTFCConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this terraform cloud config delete t f c config o k response a status code equal to that given
func (o *TerraformCloudConfigDeleteTFCConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the terraform cloud config delete t f c config o k response
func (o *TerraformCloudConfigDeleteTFCConfigOK) Code() int {
	return 200
}

func (o *TerraformCloudConfigDeleteTFCConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/tfcconfig][%d] terraformCloudConfigDeleteTFCConfigOK  %+v", 200, o.Payload)
}

func (o *TerraformCloudConfigDeleteTFCConfigOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/tfcconfig][%d] terraformCloudConfigDeleteTFCConfigOK  %+v", 200, o.Payload)
}

func (o *TerraformCloudConfigDeleteTFCConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TerraformCloudConfigDeleteTFCConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerraformCloudConfigDeleteTFCConfigDefault creates a TerraformCloudConfigDeleteTFCConfigDefault with default headers values
func NewTerraformCloudConfigDeleteTFCConfigDefault(code int) *TerraformCloudConfigDeleteTFCConfigDefault {
	return &TerraformCloudConfigDeleteTFCConfigDefault{
		_statusCode: code,
	}
}

/*
TerraformCloudConfigDeleteTFCConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TerraformCloudConfigDeleteTFCConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this terraform cloud config delete t f c config default response has a 2xx status code
func (o *TerraformCloudConfigDeleteTFCConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this terraform cloud config delete t f c config default response has a 3xx status code
func (o *TerraformCloudConfigDeleteTFCConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this terraform cloud config delete t f c config default response has a 4xx status code
func (o *TerraformCloudConfigDeleteTFCConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this terraform cloud config delete t f c config default response has a 5xx status code
func (o *TerraformCloudConfigDeleteTFCConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this terraform cloud config delete t f c config default response a status code equal to that given
func (o *TerraformCloudConfigDeleteTFCConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the terraform cloud config delete t f c config default response
func (o *TerraformCloudConfigDeleteTFCConfigDefault) Code() int {
	return o._statusCode
}

func (o *TerraformCloudConfigDeleteTFCConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/tfcconfig][%d] TerraformCloudConfig_DeleteTFCConfig default  %+v", o._statusCode, o.Payload)
}

func (o *TerraformCloudConfigDeleteTFCConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2022-02-03/namespace/{namespace_id}/tfcconfig][%d] TerraformCloudConfig_DeleteTFCConfig default  %+v", o._statusCode, o.Payload)
}

func (o *TerraformCloudConfigDeleteTFCConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *TerraformCloudConfigDeleteTFCConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}