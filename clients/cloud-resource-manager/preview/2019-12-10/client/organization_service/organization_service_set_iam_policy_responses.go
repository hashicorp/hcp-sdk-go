// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/cloud-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
)

// OrganizationServiceSetIamPolicyReader is a Reader for the OrganizationServiceSetIamPolicy structure.
type OrganizationServiceSetIamPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceSetIamPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceSetIamPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceSetIamPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceSetIamPolicyOK creates a OrganizationServiceSetIamPolicyOK with default headers values
func NewOrganizationServiceSetIamPolicyOK() *OrganizationServiceSetIamPolicyOK {
	return &OrganizationServiceSetIamPolicyOK{}
}

/*OrganizationServiceSetIamPolicyOK handles this case with default header values.

A successful response.
*/
type OrganizationServiceSetIamPolicyOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse
}

func (o *OrganizationServiceSetIamPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] organizationServiceSetIamPolicyOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceSetIamPolicyOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse {
	return o.Payload
}

func (o *OrganizationServiceSetIamPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceSetIamPolicyDefault creates a OrganizationServiceSetIamPolicyDefault with default headers values
func NewOrganizationServiceSetIamPolicyDefault(code int) *OrganizationServiceSetIamPolicyDefault {
	return &OrganizationServiceSetIamPolicyDefault{
		_statusCode: code,
	}
}

/*OrganizationServiceSetIamPolicyDefault handles this case with default header values.

An unexpected error response.
*/
type OrganizationServiceSetIamPolicyDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the organization service set iam policy default response
func (o *OrganizationServiceSetIamPolicyDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceSetIamPolicyDefault) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] OrganizationService_SetIamPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceSetIamPolicyDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *OrganizationServiceSetIamPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
