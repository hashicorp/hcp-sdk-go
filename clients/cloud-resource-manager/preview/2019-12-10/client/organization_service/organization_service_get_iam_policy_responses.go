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

// OrganizationServiceGetIamPolicyReader is a Reader for the OrganizationServiceGetIamPolicy structure.
type OrganizationServiceGetIamPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceGetIamPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceGetIamPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceGetIamPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceGetIamPolicyOK creates a OrganizationServiceGetIamPolicyOK with default headers values
func NewOrganizationServiceGetIamPolicyOK() *OrganizationServiceGetIamPolicyOK {
	return &OrganizationServiceGetIamPolicyOK{}
}

/*OrganizationServiceGetIamPolicyOK handles this case with default header values.

A successful response.
*/
type OrganizationServiceGetIamPolicyOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationGetIamPolicyResponse
}

func (o *OrganizationServiceGetIamPolicyOK) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] organizationServiceGetIamPolicyOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceGetIamPolicyOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationGetIamPolicyResponse {
	return o.Payload
}

func (o *OrganizationServiceGetIamPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationGetIamPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceGetIamPolicyDefault creates a OrganizationServiceGetIamPolicyDefault with default headers values
func NewOrganizationServiceGetIamPolicyDefault(code int) *OrganizationServiceGetIamPolicyDefault {
	return &OrganizationServiceGetIamPolicyDefault{
		_statusCode: code,
	}
}

/*OrganizationServiceGetIamPolicyDefault handles this case with default header values.

An unexpected error response.
*/
type OrganizationServiceGetIamPolicyDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service get iam policy default response
func (o *OrganizationServiceGetIamPolicyDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceGetIamPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] OrganizationService_GetIamPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceGetIamPolicyDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceGetIamPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
