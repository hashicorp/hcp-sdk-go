// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// OrganizationServiceGetReader is a Reader for the OrganizationServiceGet structure.
type OrganizationServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceGetOK creates a OrganizationServiceGetOK with default headers values
func NewOrganizationServiceGetOK() *OrganizationServiceGetOK {
	return &OrganizationServiceGetOK{}
}

/*OrganizationServiceGetOK handles this case with default header values.

A successful response.
*/
type OrganizationServiceGetOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationGetResponse
}

func (o *OrganizationServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}][%d] organizationServiceGetOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceGetOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationGetResponse {
	return o.Payload
}

func (o *OrganizationServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceGetDefault creates a OrganizationServiceGetDefault with default headers values
func NewOrganizationServiceGetDefault(code int) *OrganizationServiceGetDefault {
	return &OrganizationServiceGetDefault{
		_statusCode: code,
	}
}

/*OrganizationServiceGetDefault handles this case with default header values.

An unexpected error response.
*/
type OrganizationServiceGetDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service get default response
func (o *OrganizationServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}][%d] OrganizationService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceGetDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
