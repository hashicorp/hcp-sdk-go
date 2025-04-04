// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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
	"github.com/go-openapi/validate"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceUpdateWorkloadIdentityProviderReader is a Reader for the ServicePrincipalsServiceUpdateWorkloadIdentityProvider structure.
type ServicePrincipalsServiceUpdateWorkloadIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceUpdateWorkloadIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderOK creates a ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK with default headers values
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderOK() *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK{}
}

/*
ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK struct {
	Payload *models.HashicorpCloudIamUpdateWorkloadIdentityProviderResponse
}

// IsSuccess returns true when this service principals service update workload identity provider o k response has a 2xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service update workload identity provider o k response has a 3xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service update workload identity provider o k response has a 4xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service update workload identity provider o k response has a 5xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service update workload identity provider o k response a status code equal to that given
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service principals service update workload identity provider o k response
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) Code() int {
	return 200
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /2019-12-10/{provider.resource_name}][%d] servicePrincipalsServiceUpdateWorkloadIdentityProviderOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /2019-12-10/{provider.resource_name}][%d] servicePrincipalsServiceUpdateWorkloadIdentityProviderOK %s", 200, payload)
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) GetPayload() *models.HashicorpCloudIamUpdateWorkloadIdentityProviderResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamUpdateWorkloadIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault creates a ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault with default headers values
func NewServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault(code int) *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault {
	return &ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this service principals service update workload identity provider default response has a 2xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service update workload identity provider default response has a 3xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service update workload identity provider default response has a 4xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service update workload identity provider default response has a 5xx status code
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service update workload identity provider default response a status code equal to that given
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the service principals service update workload identity provider default response
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /2019-12-10/{provider.resource_name}][%d] ServicePrincipalsService_UpdateWorkloadIdentityProvider default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /2019-12-10/{provider.resource_name}][%d] ServicePrincipalsService_UpdateWorkloadIdentityProvider default %s", o._statusCode, payload)
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody provider is the workload identity provider to update.
swagger:model ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody
*/
type ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody struct {

	// aws config
	AwsConfig *models.HashicorpCloudIamAWSWorkloadIdentityProviderConfig `json:"aws_config,omitempty"`

	// conditional_access is a go-bexpr string that is evaluated when exchanging
	// tokens. It restricts which upstream identities are allowed to access the
	// service principal.
	ConditionalAccess string `json:"conditional_access,omitempty"`

	// created_at is when the workload identity provider was created.
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// description is an optional description for the workload identity provider.
	Description string `json:"description,omitempty"`

	// oidc config
	OidcConfig *models.HashicorpCloudIamOIDCWorkloadIdentityProviderConfig `json:"oidc_config,omitempty"`

	// resource_id is the resource identifier for this workload identity
	// provider.
	ResourceID string `json:"resource_id,omitempty"`

	// system_managed indicates that the provider is not editable or deletable
	// by the user.
	SystemManaged bool `json:"system_managed,omitempty"`

	// updated_at is when the workload identity provider was last updated.
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this service principals service update workload identity provider body
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAwsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOidcConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) validateAwsConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.AwsConfig) { // not required
		return nil
	}

	if o.AwsConfig != nil {
		if err := o.AwsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider" + "." + "aws_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider" + "." + "aws_config")
			}
			return err
		}
	}

	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("provider"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) validateOidcConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.OidcConfig) { // not required
		return nil
	}

	if o.OidcConfig != nil {
		if err := o.OidcConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider" + "." + "oidc_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider" + "." + "oidc_config")
			}
			return err
		}
	}

	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("provider"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service principals service update workload identity provider body based on the context it is used
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAwsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOidcConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) contextValidateAwsConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.AwsConfig != nil {

		if swag.IsZero(o.AwsConfig) { // not required
			return nil
		}

		if err := o.AwsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider" + "." + "aws_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider" + "." + "aws_config")
			}
			return err
		}
	}

	return nil
}

func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) contextValidateOidcConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.OidcConfig != nil {

		if swag.IsZero(o.OidcConfig) { // not required
			return nil
		}

		if err := o.OidcConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider" + "." + "oidc_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider" + "." + "oidc_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody) UnmarshalBinary(b []byte) error {
	var res ServicePrincipalsServiceUpdateWorkloadIdentityProviderBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
