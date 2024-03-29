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

// HashicorpCloudWebhookUpdateWebhookResponse hashicorp cloud webhook update webhook response
//
// swagger:model hashicorp.cloud.webhook.UpdateWebhookResponse
type HashicorpCloudWebhookUpdateWebhookResponse struct {

	// The updated webhook.
	Webhook *HashicorpCloudWebhookWebhook `json:"webhook,omitempty"`
}

// Validate validates this hashicorp cloud webhook update webhook response
func (m *HashicorpCloudWebhookUpdateWebhookResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookUpdateWebhookResponse) validateWebhook(formats strfmt.Registry) error {
	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud webhook update webhook response based on the context it is used
func (m *HashicorpCloudWebhookUpdateWebhookResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebhook(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookUpdateWebhookResponse) contextValidateWebhook(ctx context.Context, formats strfmt.Registry) error {

	if m.Webhook != nil {

		if swag.IsZero(m.Webhook) { // not required
			return nil
		}

		if err := m.Webhook.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookUpdateWebhookResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookUpdateWebhookResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookUpdateWebhookResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
