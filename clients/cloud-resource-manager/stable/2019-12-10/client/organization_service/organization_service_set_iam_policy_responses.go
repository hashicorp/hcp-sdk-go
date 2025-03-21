// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
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

/*
OrganizationServiceSetIamPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceSetIamPolicyOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationSetIamPolicyResponse
}

// IsSuccess returns true when this organization service set iam policy o k response has a 2xx status code
func (o *OrganizationServiceSetIamPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service set iam policy o k response has a 3xx status code
func (o *OrganizationServiceSetIamPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service set iam policy o k response has a 4xx status code
func (o *OrganizationServiceSetIamPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service set iam policy o k response has a 5xx status code
func (o *OrganizationServiceSetIamPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service set iam policy o k response a status code equal to that given
func (o *OrganizationServiceSetIamPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organization service set iam policy o k response
func (o *OrganizationServiceSetIamPolicyOK) Code() int {
	return 200
}

func (o *OrganizationServiceSetIamPolicyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] organizationServiceSetIamPolicyOK %s", 200, payload)
}

func (o *OrganizationServiceSetIamPolicyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] organizationServiceSetIamPolicyOK %s", 200, payload)
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

/*
OrganizationServiceSetIamPolicyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceSetIamPolicyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this organization service set iam policy default response has a 2xx status code
func (o *OrganizationServiceSetIamPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service set iam policy default response has a 3xx status code
func (o *OrganizationServiceSetIamPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service set iam policy default response has a 4xx status code
func (o *OrganizationServiceSetIamPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service set iam policy default response has a 5xx status code
func (o *OrganizationServiceSetIamPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service set iam policy default response a status code equal to that given
func (o *OrganizationServiceSetIamPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the organization service set iam policy default response
func (o *OrganizationServiceSetIamPolicyDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationServiceSetIamPolicyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] OrganizationService_SetIamPolicy default %s", o._statusCode, payload)
}

func (o *OrganizationServiceSetIamPolicyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/organizations/{id}/iam-policy][%d] OrganizationService_SetIamPolicy default %s", o._statusCode, payload)
}

func (o *OrganizationServiceSetIamPolicyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *OrganizationServiceSetIamPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
OrganizationServiceSetIamPolicyBody see OrganizationService.SetIamPolicy
swagger:model OrganizationServiceSetIamPolicyBody
*/
type OrganizationServiceSetIamPolicyBody struct {

	// Policy is the updated IAM policy for the organization. The policy
	// will be completely replaced and therefore Policy.Etag must be specified
	// in order to prevent concurrent updates.
	Policy *models.HashicorpCloudResourcemanagerPolicy `json:"policy,omitempty"`
}

// Validate validates this organization service set iam policy body
func (o *OrganizationServiceSetIamPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrganizationServiceSetIamPolicyBody) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.Policy) { // not required
		return nil
	}

	if o.Policy != nil {
		if err := o.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organization service set iam policy body based on the context it is used
func (o *OrganizationServiceSetIamPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrganizationServiceSetIamPolicyBody) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if o.Policy != nil {

		if swag.IsZero(o.Policy) { // not required
			return nil
		}

		if err := o.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrganizationServiceSetIamPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrganizationServiceSetIamPolicyBody) UnmarshalBinary(b []byte) error {
	var res OrganizationServiceSetIamPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
