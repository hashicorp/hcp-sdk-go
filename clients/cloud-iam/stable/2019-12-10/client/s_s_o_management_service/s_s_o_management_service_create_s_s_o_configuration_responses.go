// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// SSOManagementServiceCreateSSOConfigurationReader is a Reader for the SSOManagementServiceCreateSSOConfiguration structure.
type SSOManagementServiceCreateSSOConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSOManagementServiceCreateSSOConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSOManagementServiceCreateSSOConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSOManagementServiceCreateSSOConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSOManagementServiceCreateSSOConfigurationOK creates a SSOManagementServiceCreateSSOConfigurationOK with default headers values
func NewSSOManagementServiceCreateSSOConfigurationOK() *SSOManagementServiceCreateSSOConfigurationOK {
	return &SSOManagementServiceCreateSSOConfigurationOK{}
}

/*
SSOManagementServiceCreateSSOConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type SSOManagementServiceCreateSSOConfigurationOK struct {
	Payload models.HashicorpCloudIamCreateSSOConfigurationResponse
}

// IsSuccess returns true when this s s o management service create s s o configuration o k response has a 2xx status code
func (o *SSOManagementServiceCreateSSOConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s o management service create s s o configuration o k response has a 3xx status code
func (o *SSOManagementServiceCreateSSOConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s o management service create s s o configuration o k response has a 4xx status code
func (o *SSOManagementServiceCreateSSOConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s o management service create s s o configuration o k response has a 5xx status code
func (o *SSOManagementServiceCreateSSOConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s o management service create s s o configuration o k response a status code equal to that given
func (o *SSOManagementServiceCreateSSOConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s s o management service create s s o configuration o k response
func (o *SSOManagementServiceCreateSSOConfigurationOK) Code() int {
	return 200
}

func (o *SSOManagementServiceCreateSSOConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/sso-configurations][%d] sSOManagementServiceCreateSSOConfigurationOK  %+v", 200, o.Payload)
}

func (o *SSOManagementServiceCreateSSOConfigurationOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/sso-configurations][%d] sSOManagementServiceCreateSSOConfigurationOK  %+v", 200, o.Payload)
}

func (o *SSOManagementServiceCreateSSOConfigurationOK) GetPayload() models.HashicorpCloudIamCreateSSOConfigurationResponse {
	return o.Payload
}

func (o *SSOManagementServiceCreateSSOConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSOManagementServiceCreateSSOConfigurationDefault creates a SSOManagementServiceCreateSSOConfigurationDefault with default headers values
func NewSSOManagementServiceCreateSSOConfigurationDefault(code int) *SSOManagementServiceCreateSSOConfigurationDefault {
	return &SSOManagementServiceCreateSSOConfigurationDefault{
		_statusCode: code,
	}
}

/*
SSOManagementServiceCreateSSOConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SSOManagementServiceCreateSSOConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this s s o management service create s s o configuration default response has a 2xx status code
func (o *SSOManagementServiceCreateSSOConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s s o management service create s s o configuration default response has a 3xx status code
func (o *SSOManagementServiceCreateSSOConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s s o management service create s s o configuration default response has a 4xx status code
func (o *SSOManagementServiceCreateSSOConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s s o management service create s s o configuration default response has a 5xx status code
func (o *SSOManagementServiceCreateSSOConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s s o management service create s s o configuration default response a status code equal to that given
func (o *SSOManagementServiceCreateSSOConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s s o management service create s s o configuration default response
func (o *SSOManagementServiceCreateSSOConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *SSOManagementServiceCreateSSOConfigurationDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/sso-configurations][%d] SSOManagementService_CreateSSOConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *SSOManagementServiceCreateSSOConfigurationDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/sso-configurations][%d] SSOManagementService_CreateSSOConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *SSOManagementServiceCreateSSOConfigurationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SSOManagementServiceCreateSSOConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SSOManagementServiceCreateSSOConfigurationBody s s o management service create s s o configuration body
swagger:model SSOManagementServiceCreateSSOConfigurationBody
*/
type SSOManagementServiceCreateSSOConfigurationBody struct {

	// SsoConfig has the SSO configuration for this organization.
	Config *models.HashicorpCloudIamSSOConfig `json:"config,omitempty"`
}

// Validate validates this s s o management service create s s o configuration body
func (o *SSOManagementServiceCreateSSOConfigurationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SSOManagementServiceCreateSSOConfigurationBody) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.Config) { // not required
		return nil
	}

	if o.Config != nil {
		if err := o.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s s o management service create s s o configuration body based on the context it is used
func (o *SSOManagementServiceCreateSSOConfigurationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SSOManagementServiceCreateSSOConfigurationBody) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.Config != nil {

		if swag.IsZero(o.Config) { // not required
			return nil
		}

		if err := o.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SSOManagementServiceCreateSSOConfigurationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSOManagementServiceCreateSSOConfigurationBody) UnmarshalBinary(b []byte) error {
	var res SSOManagementServiceCreateSSOConfigurationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}