// Code generated by go-swagger; DO NOT EDIT.

package integration_subscription_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-radar/preview/2023-05-01/models"
)

// CreateIntegrationSubscriptionReader is a Reader for the CreateIntegrationSubscription structure.
type CreateIntegrationSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIntegrationSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIntegrationSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIntegrationSubscriptionOK creates a CreateIntegrationSubscriptionOK with default headers values
func NewCreateIntegrationSubscriptionOK() *CreateIntegrationSubscriptionOK {
	return &CreateIntegrationSubscriptionOK{}
}

/*
CreateIntegrationSubscriptionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateIntegrationSubscriptionOK struct {
	Payload *models.VaultRadar20230501CreateIntegrationSubscriptionResponse
}

// IsSuccess returns true when this create integration subscription o k response has a 2xx status code
func (o *CreateIntegrationSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create integration subscription o k response has a 3xx status code
func (o *CreateIntegrationSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration subscription o k response has a 4xx status code
func (o *CreateIntegrationSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create integration subscription o k response has a 5xx status code
func (o *CreateIntegrationSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration subscription o k response a status code equal to that given
func (o *CreateIntegrationSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create integration subscription o k response
func (o *CreateIntegrationSubscriptionOK) Code() int {
	return 200
}

func (o *CreateIntegrationSubscriptionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] createIntegrationSubscriptionOK %s", 200, payload)
}

func (o *CreateIntegrationSubscriptionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] createIntegrationSubscriptionOK %s", 200, payload)
}

func (o *CreateIntegrationSubscriptionOK) GetPayload() *models.VaultRadar20230501CreateIntegrationSubscriptionResponse {
	return o.Payload
}

func (o *CreateIntegrationSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultRadar20230501CreateIntegrationSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationSubscriptionDefault creates a CreateIntegrationSubscriptionDefault with default headers values
func NewCreateIntegrationSubscriptionDefault(code int) *CreateIntegrationSubscriptionDefault {
	return &CreateIntegrationSubscriptionDefault{
		_statusCode: code,
	}
}

/*
CreateIntegrationSubscriptionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateIntegrationSubscriptionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create integration subscription default response has a 2xx status code
func (o *CreateIntegrationSubscriptionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create integration subscription default response has a 3xx status code
func (o *CreateIntegrationSubscriptionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create integration subscription default response has a 4xx status code
func (o *CreateIntegrationSubscriptionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create integration subscription default response has a 5xx status code
func (o *CreateIntegrationSubscriptionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create integration subscription default response a status code equal to that given
func (o *CreateIntegrationSubscriptionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create integration subscription default response
func (o *CreateIntegrationSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateIntegrationSubscriptionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] CreateIntegrationSubscription default %s", o._statusCode, payload)
}

func (o *CreateIntegrationSubscriptionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /2023-05-01/vault-radar/projects/{location.project_id}/integrations/subscriptions][%d] CreateIntegrationSubscription default %s", o._statusCode, payload)
}

func (o *CreateIntegrationSubscriptionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateIntegrationSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateIntegrationSubscriptionBody create integration subscription body
swagger:model CreateIntegrationSubscriptionBody
*/
type CreateIntegrationSubscriptionBody struct {

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// details
	Details string `json:"details,omitempty"`

	// location
	Location *CreateIntegrationSubscriptionParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create integration subscription body
func (o *CreateIntegrationSubscriptionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationSubscriptionBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this create integration subscription body based on the context it is used
func (o *CreateIntegrationSubscriptionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationSubscriptionBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateIntegrationSubscriptionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationSubscriptionBody) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationSubscriptionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateIntegrationSubscriptionParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateIntegrationSubscriptionParamsBodyLocation
*/
type CreateIntegrationSubscriptionParamsBodyLocation struct {

	// organization_id is the id of the organization.
	OrganizationID string `json:"organization_id,omitempty"`

	// region
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create integration subscription params body location
func (o *CreateIntegrationSubscriptionParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationSubscriptionParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this create integration subscription params body location based on the context it is used
func (o *CreateIntegrationSubscriptionParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationSubscriptionParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateIntegrationSubscriptionParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationSubscriptionParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationSubscriptionParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
