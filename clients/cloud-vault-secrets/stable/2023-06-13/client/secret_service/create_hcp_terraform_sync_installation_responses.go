// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// CreateHcpTerraformSyncInstallationReader is a Reader for the CreateHcpTerraformSyncInstallation structure.
type CreateHcpTerraformSyncInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHcpTerraformSyncInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHcpTerraformSyncInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateHcpTerraformSyncInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHcpTerraformSyncInstallationOK creates a CreateHcpTerraformSyncInstallationOK with default headers values
func NewCreateHcpTerraformSyncInstallationOK() *CreateHcpTerraformSyncInstallationOK {
	return &CreateHcpTerraformSyncInstallationOK{}
}

/*
CreateHcpTerraformSyncInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateHcpTerraformSyncInstallationOK struct {
	Payload *models.Secrets20230613GetSyncInstallationResponse
}

// IsSuccess returns true when this create hcp terraform sync installation o k response has a 2xx status code
func (o *CreateHcpTerraformSyncInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create hcp terraform sync installation o k response has a 3xx status code
func (o *CreateHcpTerraformSyncInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create hcp terraform sync installation o k response has a 4xx status code
func (o *CreateHcpTerraformSyncInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create hcp terraform sync installation o k response has a 5xx status code
func (o *CreateHcpTerraformSyncInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create hcp terraform sync installation o k response a status code equal to that given
func (o *CreateHcpTerraformSyncInstallationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create hcp terraform sync installation o k response
func (o *CreateHcpTerraformSyncInstallationOK) Code() int {
	return 200
}

func (o *CreateHcpTerraformSyncInstallationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/hcp-terraform/installations][%d] createHcpTerraformSyncInstallationOK %s", 200, payload)
}

func (o *CreateHcpTerraformSyncInstallationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/hcp-terraform/installations][%d] createHcpTerraformSyncInstallationOK %s", 200, payload)
}

func (o *CreateHcpTerraformSyncInstallationOK) GetPayload() *models.Secrets20230613GetSyncInstallationResponse {
	return o.Payload
}

func (o *CreateHcpTerraformSyncInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetSyncInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHcpTerraformSyncInstallationDefault creates a CreateHcpTerraformSyncInstallationDefault with default headers values
func NewCreateHcpTerraformSyncInstallationDefault(code int) *CreateHcpTerraformSyncInstallationDefault {
	return &CreateHcpTerraformSyncInstallationDefault{
		_statusCode: code,
	}
}

/*
CreateHcpTerraformSyncInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateHcpTerraformSyncInstallationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create hcp terraform sync installation default response has a 2xx status code
func (o *CreateHcpTerraformSyncInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create hcp terraform sync installation default response has a 3xx status code
func (o *CreateHcpTerraformSyncInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create hcp terraform sync installation default response has a 4xx status code
func (o *CreateHcpTerraformSyncInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create hcp terraform sync installation default response has a 5xx status code
func (o *CreateHcpTerraformSyncInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create hcp terraform sync installation default response a status code equal to that given
func (o *CreateHcpTerraformSyncInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create hcp terraform sync installation default response
func (o *CreateHcpTerraformSyncInstallationDefault) Code() int {
	return o._statusCode
}

func (o *CreateHcpTerraformSyncInstallationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/hcp-terraform/installations][%d] CreateHcpTerraformSyncInstallation default %s", o._statusCode, payload)
}

func (o *CreateHcpTerraformSyncInstallationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/hcp-terraform/installations][%d] CreateHcpTerraformSyncInstallation default %s", o._statusCode, payload)
}

func (o *CreateHcpTerraformSyncInstallationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateHcpTerraformSyncInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateHcpTerraformSyncInstallationBody create hcp terraform sync installation body
swagger:model CreateHcpTerraformSyncInstallationBody
*/
type CreateHcpTerraformSyncInstallationBody struct {

	// location
	Location *CreateHcpTerraformSyncInstallationParamsBodyLocation `json:"location,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this create hcp terraform sync installation body
func (o *CreateHcpTerraformSyncInstallationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateHcpTerraformSyncInstallationBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create hcp terraform sync installation body based on the context it is used
func (o *CreateHcpTerraformSyncInstallationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateHcpTerraformSyncInstallationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {

		if swag.IsZero(o.Location) { // not required
			return nil
		}

		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateHcpTerraformSyncInstallationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateHcpTerraformSyncInstallationBody) UnmarshalBinary(b []byte) error {
	var res CreateHcpTerraformSyncInstallationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateHcpTerraformSyncInstallationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateHcpTerraformSyncInstallationParamsBodyLocation
*/
type CreateHcpTerraformSyncInstallationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create hcp terraform sync installation params body location
func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create hcp terraform sync installation params body location based on the context it is used
func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {

		if swag.IsZero(o.Region) { // not required
			return nil
		}

		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateHcpTerraformSyncInstallationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateHcpTerraformSyncInstallationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
