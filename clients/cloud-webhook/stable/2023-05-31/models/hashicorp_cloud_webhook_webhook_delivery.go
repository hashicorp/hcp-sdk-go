// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWebhookWebhookDelivery The information about a webhook delivery, such as the sent request, the received response, and how many times
// the webhook service retried to send the delivery in case of failure.
//
// swagger:model hashicorp.cloud.webhook.WebhookDelivery
type HashicorpCloudWebhookWebhookDelivery struct {

	// The unique identifier for the delivery.
	ID string `json:"id,omitempty"`

	// The request that delivers the event payload.
	Request *HashicorpCloudWebhookWebhookRequest `json:"request,omitempty"`

	// The response received from the external webhook service when a request was sent to deliver the payload.
	Response *HashicorpCloudWebhookWebhookResponse `json:"response,omitempty"`

	// The number of times HCP Packer resent the request before successfully delivering the payload.
	Retries int32 `json:"retries,omitempty"`

	// The ID of the webhook to which the payload was delivered.
	WebhookResourceID string `json:"webhook_resource_id,omitempty"`
}

// Validate validates this hashicorp cloud webhook webhook delivery
func (m *HashicorpCloudWebhookWebhookDelivery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookWebhookDelivery) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWebhookWebhookDelivery) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud webhook webhook delivery based on the context it is used
func (m *HashicorpCloudWebhookWebhookDelivery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookWebhookDelivery) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if swag.IsZero(m.Request) { // not required
			return nil
		}

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWebhookWebhookDelivery) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {

		if swag.IsZero(m.Response) { // not required
			return nil
		}

		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookDelivery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookDelivery) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookWebhookDelivery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}