// Code generated by go-swagger; DO NOT EDIT.

package github_control_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// GithubControlServicePickInstallationReader is a Reader for the GithubControlServicePickInstallation structure.
type GithubControlServicePickInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GithubControlServicePickInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGithubControlServicePickInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGithubControlServicePickInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGithubControlServicePickInstallationOK creates a GithubControlServicePickInstallationOK with default headers values
func NewGithubControlServicePickInstallationOK() *GithubControlServicePickInstallationOK {
	return &GithubControlServicePickInstallationOK{}
}

/*
GithubControlServicePickInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GithubControlServicePickInstallationOK struct {
	Payload *models.HashicorpCloudWaypointPickInstallationResponse
}

// IsSuccess returns true when this github control service pick installation o k response has a 2xx status code
func (o *GithubControlServicePickInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this github control service pick installation o k response has a 3xx status code
func (o *GithubControlServicePickInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this github control service pick installation o k response has a 4xx status code
func (o *GithubControlServicePickInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this github control service pick installation o k response has a 5xx status code
func (o *GithubControlServicePickInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this github control service pick installation o k response a status code equal to that given
func (o *GithubControlServicePickInstallationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the github control service pick installation o k response
func (o *GithubControlServicePickInstallationOK) Code() int {
	return 200
}

func (o *GithubControlServicePickInstallationOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/pick/install][%d] githubControlServicePickInstallationOK  %+v", 200, o.Payload)
}

func (o *GithubControlServicePickInstallationOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/pick/install][%d] githubControlServicePickInstallationOK  %+v", 200, o.Payload)
}

func (o *GithubControlServicePickInstallationOK) GetPayload() *models.HashicorpCloudWaypointPickInstallationResponse {
	return o.Payload
}

func (o *GithubControlServicePickInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointPickInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGithubControlServicePickInstallationDefault creates a GithubControlServicePickInstallationDefault with default headers values
func NewGithubControlServicePickInstallationDefault(code int) *GithubControlServicePickInstallationDefault {
	return &GithubControlServicePickInstallationDefault{
		_statusCode: code,
	}
}

/*
GithubControlServicePickInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GithubControlServicePickInstallationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this github control service pick installation default response has a 2xx status code
func (o *GithubControlServicePickInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this github control service pick installation default response has a 3xx status code
func (o *GithubControlServicePickInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this github control service pick installation default response has a 4xx status code
func (o *GithubControlServicePickInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this github control service pick installation default response has a 5xx status code
func (o *GithubControlServicePickInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this github control service pick installation default response a status code equal to that given
func (o *GithubControlServicePickInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the github control service pick installation default response
func (o *GithubControlServicePickInstallationDefault) Code() int {
	return o._statusCode
}

func (o *GithubControlServicePickInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/pick/install][%d] GithubControlService_PickInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GithubControlServicePickInstallationDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/pick/install][%d] GithubControlService_PickInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GithubControlServicePickInstallationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GithubControlServicePickInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}