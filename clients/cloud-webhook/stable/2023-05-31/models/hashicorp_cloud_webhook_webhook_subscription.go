// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWebhookWebhookSubscription The information about a webhook subscription.
// A webhook can either provide the `resource_id` to subscribe to events generated by a specific resource,
// or they can subscribe to the available events for all resources in the webhook's project when the `resource_id` is omitted.
//
// swagger:model hashicorp.cloud.webhook.WebhookSubscription
type HashicorpCloudWebhookWebhookSubscription struct {

	// The list of resource lifecycle events that the webhook is subscribed to.
	// Events are a combination of event source and event action.
	// For example, `hashicorp.packer.version:create`.
	// The wildcard(`*`) subscribes the webhook to all the source's events, for example `hashicorp.packer.version:*`.
	Events []*HashicorpCloudWebhookWebhookSubscriptionEvent `json:"events"`

	// Refers to the resource the webhook is subscribed to.
	// If not set, the webhook subscribes to the emitted events listed in `events`
	// for any resource in the webhook's project.
	ResourceID string `json:"resource_id,omitempty"`
}

// Validate validates this hashicorp cloud webhook webhook subscription
func (m *HashicorpCloudWebhookWebhookSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookWebhookSubscription) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud webhook webhook subscription based on the context it is used
func (m *HashicorpCloudWebhookWebhookSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWebhookWebhookSubscription) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if swag.IsZero(m.Events[i]) { // not required
				return nil
			}

			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookSubscription) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookWebhookSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
