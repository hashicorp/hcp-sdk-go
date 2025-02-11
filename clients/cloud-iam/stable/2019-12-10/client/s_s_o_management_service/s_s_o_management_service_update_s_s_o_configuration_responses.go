// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// SSOManagementServiceUpdateSSOConfigurationReader is a Reader for the SSOManagementServiceUpdateSSOConfiguration structure.
type SSOManagementServiceUpdateSSOConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSOManagementServiceUpdateSSOConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSOManagementServiceUpdateSSOConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSOManagementServiceUpdateSSOConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSOManagementServiceUpdateSSOConfigurationOK creates a SSOManagementServiceUpdateSSOConfigurationOK with default headers values
func NewSSOManagementServiceUpdateSSOConfigurationOK() *SSOManagementServiceUpdateSSOConfigurationOK {
	return &SSOManagementServiceUpdateSSOConfigurationOK{}
}

/*
SSOManagementServiceUpdateSSOConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type SSOManagementServiceUpdateSSOConfigurationOK struct {
	Payload models.HashicorpCloudIamUpdateSSOConfigurationResponse
}

// IsSuccess returns true when this s s o management service update s s o configuration o k response has a 2xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s o management service update s s o configuration o k response has a 3xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s o management service update s s o configuration o k response has a 4xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s o management service update s s o configuration o k response has a 5xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s o management service update s s o configuration o k response a status code equal to that given
func (o *SSOManagementServiceUpdateSSOConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s s o management service update s s o configuration o k response
func (o *SSOManagementServiceUpdateSSOConfigurationOK) Code() int {
	return 200
}

func (o *SSOManagementServiceUpdateSSOConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceUpdateSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceUpdateSSOConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] sSOManagementServiceUpdateSSOConfigurationOK %s", 200, payload)
}

func (o *SSOManagementServiceUpdateSSOConfigurationOK) GetPayload() models.HashicorpCloudIamUpdateSSOConfigurationResponse {
	return o.Payload
}

func (o *SSOManagementServiceUpdateSSOConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSOManagementServiceUpdateSSOConfigurationDefault creates a SSOManagementServiceUpdateSSOConfigurationDefault with default headers values
func NewSSOManagementServiceUpdateSSOConfigurationDefault(code int) *SSOManagementServiceUpdateSSOConfigurationDefault {
	return &SSOManagementServiceUpdateSSOConfigurationDefault{
		_statusCode: code,
	}
}

/*
SSOManagementServiceUpdateSSOConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SSOManagementServiceUpdateSSOConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this s s o management service update s s o configuration default response has a 2xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s s o management service update s s o configuration default response has a 3xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s s o management service update s s o configuration default response has a 4xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s s o management service update s s o configuration default response has a 5xx status code
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s s o management service update s s o configuration default response a status code equal to that given
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s s o management service update s s o configuration default response
func (o *SSOManagementServiceUpdateSSOConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *SSOManagementServiceUpdateSSOConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_UpdateSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceUpdateSSOConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] SSOManagementService_UpdateSSOConfiguration default %s", o._statusCode, payload)
}

func (o *SSOManagementServiceUpdateSSOConfigurationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SSOManagementServiceUpdateSSOConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SSOManagementServiceUpdateSSOConfigurationBody s s o management service update s s o configuration body
swagger:model SSOManagementServiceUpdateSSOConfigurationBody
*/
type SSOManagementServiceUpdateSSOConfigurationBody struct {

	// SsoConfig has the SSO configuration for this organization.
	Config *models.HashicorpCloudIamSSOConfig `json:"config,omitempty"`
}

// Validate validates this s s o management service update s s o configuration body
func (o *SSOManagementServiceUpdateSSOConfigurationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SSOManagementServiceUpdateSSOConfigurationBody) validateConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this s s o management service update s s o configuration body based on the context it is used
func (o *SSOManagementServiceUpdateSSOConfigurationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SSOManagementServiceUpdateSSOConfigurationBody) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SSOManagementServiceUpdateSSOConfigurationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSOManagementServiceUpdateSSOConfigurationBody) UnmarshalBinary(b []byte) error {
	var res SSOManagementServiceUpdateSSOConfigurationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
