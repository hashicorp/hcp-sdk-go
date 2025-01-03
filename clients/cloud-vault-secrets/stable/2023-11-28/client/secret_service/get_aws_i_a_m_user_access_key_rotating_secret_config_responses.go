// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-11-28/models"
)

// GetAwsIAMUserAccessKeyRotatingSecretConfigReader is a Reader for the GetAwsIAMUserAccessKeyRotatingSecretConfig structure.
type GetAwsIAMUserAccessKeyRotatingSecretConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsIAMUserAccessKeyRotatingSecretConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAwsIAMUserAccessKeyRotatingSecretConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAwsIAMUserAccessKeyRotatingSecretConfigOK creates a GetAwsIAMUserAccessKeyRotatingSecretConfigOK with default headers values
func NewGetAwsIAMUserAccessKeyRotatingSecretConfigOK() *GetAwsIAMUserAccessKeyRotatingSecretConfigOK {
	return &GetAwsIAMUserAccessKeyRotatingSecretConfigOK{}
}

/*
GetAwsIAMUserAccessKeyRotatingSecretConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAwsIAMUserAccessKeyRotatingSecretConfigOK struct {
	Payload *models.Secrets20231128GetAwsIAMUserAccessKeyRotatingSecretConfigResponse
}

// IsSuccess returns true when this get aws i a m user access key rotating secret config o k response has a 2xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get aws i a m user access key rotating secret config o k response has a 3xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws i a m user access key rotating secret config o k response has a 4xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get aws i a m user access key rotating secret config o k response has a 5xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws i a m user access key rotating secret config o k response a status code equal to that given
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get aws i a m user access key rotating secret config o k response
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) Code() int {
	return 200
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] getAwsIAMUserAccessKeyRotatingSecretConfigOK %s", 200, payload)
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] getAwsIAMUserAccessKeyRotatingSecretConfigOK %s", 200, payload)
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) GetPayload() *models.Secrets20231128GetAwsIAMUserAccessKeyRotatingSecretConfigResponse {
	return o.Payload
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetAwsIAMUserAccessKeyRotatingSecretConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsIAMUserAccessKeyRotatingSecretConfigDefault creates a GetAwsIAMUserAccessKeyRotatingSecretConfigDefault with default headers values
func NewGetAwsIAMUserAccessKeyRotatingSecretConfigDefault(code int) *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault {
	return &GetAwsIAMUserAccessKeyRotatingSecretConfigDefault{
		_statusCode: code,
	}
}

/*
GetAwsIAMUserAccessKeyRotatingSecretConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAwsIAMUserAccessKeyRotatingSecretConfigDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get aws i a m user access key rotating secret config default response has a 2xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get aws i a m user access key rotating secret config default response has a 3xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get aws i a m user access key rotating secret config default response has a 4xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get aws i a m user access key rotating secret config default response has a 5xx status code
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get aws i a m user access key rotating secret config default response a status code equal to that given
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get aws i a m user access key rotating secret config default response
func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] GetAwsIAMUserAccessKeyRotatingSecretConfig default %s", o._statusCode, payload)
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/rotating/aws/secret/{name}][%d] GetAwsIAMUserAccessKeyRotatingSecretConfig default %s", o._statusCode, payload)
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetAwsIAMUserAccessKeyRotatingSecretConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
