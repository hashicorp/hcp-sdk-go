// Code generated by go-swagger; DO NOT EDIT.

package integration_subscription_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// UpdateIntegrationSubscriptionReader is a Reader for the UpdateIntegrationSubscription structure.
type UpdateIntegrationSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIntegrationSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIntegrationSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIntegrationSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIntegrationSubscriptionOK creates a UpdateIntegrationSubscriptionOK with default headers values
func NewUpdateIntegrationSubscriptionOK() *UpdateIntegrationSubscriptionOK {
	return &UpdateIntegrationSubscriptionOK{}
}

/*
UpdateIntegrationSubscriptionOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateIntegrationSubscriptionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update integration subscription o k response has a 2xx status code
func (o *UpdateIntegrationSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update integration subscription o k response has a 3xx status code
func (o *UpdateIntegrationSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update integration subscription o k response has a 4xx status code
func (o *UpdateIntegrationSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update integration subscription o k response has a 5xx status code
func (o *UpdateIntegrationSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update integration subscription o k response a status code equal to that given
func (o *UpdateIntegrationSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update integration subscription o k response
func (o *UpdateIntegrationSubscriptionOK) Code() int {
	return 200
}

func (o *UpdateIntegrationSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] updateIntegrationSubscriptionOK  %+v", 200, o.Payload)
}

func (o *UpdateIntegrationSubscriptionOK) String() string {
	return fmt.Sprintf("[PUT /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] updateIntegrationSubscriptionOK  %+v", 200, o.Payload)
}

func (o *UpdateIntegrationSubscriptionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateIntegrationSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationSubscriptionDefault creates a UpdateIntegrationSubscriptionDefault with default headers values
func NewUpdateIntegrationSubscriptionDefault(code int) *UpdateIntegrationSubscriptionDefault {
	return &UpdateIntegrationSubscriptionDefault{
		_statusCode: code,
	}
}

/*
UpdateIntegrationSubscriptionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateIntegrationSubscriptionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this update integration subscription default response has a 2xx status code
func (o *UpdateIntegrationSubscriptionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update integration subscription default response has a 3xx status code
func (o *UpdateIntegrationSubscriptionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update integration subscription default response has a 4xx status code
func (o *UpdateIntegrationSubscriptionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update integration subscription default response has a 5xx status code
func (o *UpdateIntegrationSubscriptionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update integration subscription default response a status code equal to that given
func (o *UpdateIntegrationSubscriptionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update integration subscription default response
func (o *UpdateIntegrationSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIntegrationSubscriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] UpdateIntegrationSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateIntegrationSubscriptionDefault) String() string {
	return fmt.Sprintf("[PUT /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] UpdateIntegrationSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateIntegrationSubscriptionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpdateIntegrationSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateIntegrationSubscriptionBody update integration subscription body
swagger:model UpdateIntegrationSubscriptionBody
*/
type UpdateIntegrationSubscriptionBody struct {

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// details
	Details string `json:"details,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *UpdateIntegrationSubscriptionParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update integration subscription body
func (o *UpdateIntegrationSubscriptionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateIntegrationSubscriptionBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this update integration subscription body based on the context it is used
func (o *UpdateIntegrationSubscriptionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateIntegrationSubscriptionBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateIntegrationSubscriptionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateIntegrationSubscriptionBody) UnmarshalBinary(b []byte) error {
	var res UpdateIntegrationSubscriptionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateIntegrationSubscriptionParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model UpdateIntegrationSubscriptionParamsBodyLocation
*/
type UpdateIntegrationSubscriptionParamsBodyLocation struct {

	// organization_id is the id of the organization.
	OrganizationID string `json:"organization_id,omitempty"`

	// region
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this update integration subscription params body location
func (o *UpdateIntegrationSubscriptionParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateIntegrationSubscriptionParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this update integration subscription params body location based on the context it is used
func (o *UpdateIntegrationSubscriptionParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateIntegrationSubscriptionParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateIntegrationSubscriptionParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateIntegrationSubscriptionParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpdateIntegrationSubscriptionParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}