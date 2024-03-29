// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWebhookWebhookPayload The payload that is delivered as the body of the request made to the webhook's destination URL.
//
// swagger:model hashicorp.cloud.webhook.WebhookPayload
type HashicorpCloudWebhookWebhookPayload struct {

	// The type of action performed by this event, for example `create`.
	EventAction string `json:"event_action,omitempty"`

	// The event description, for example `Created iteration`.
	EventDescription string `json:"event_description,omitempty"`

	// The unique identifier for the event generated by the services that own the resources. The value format is: `<service>.event:<random string>`,
	// for example, `packer.event:t79BRg8WhTmDPBRM`.
	EventID string `json:"event_id,omitempty"`

	// The payload containing information about the resource's lifecycle event. The lifecycle event is defined by the services that own the resources.
	EventPayload interface{} `json:"event_payload,omitempty"`

	// Refers to the source of the event, for example `hashicorp.packer.version`.
	// The event source type may differ from the resource that the webhook is subscribed to when the event is from a
	// descendent resource.
	// For example, webhooks subscribed to `hashicorp.packer.registry` receive events for the descendant's resources,
	// such as a `hashicorp.packer.version`.
	EventSource string `json:"event_source,omitempty"`

	// The version of the payload being sent.
	EventVersion string `json:"event_version,omitempty"`

	// The ID of the resource the event is related to.
	ResourceID string `json:"resource_id,omitempty"`

	// The name of the resource the event is related to.
	ResourceName string `json:"resource_name,omitempty"`
}

// Validate validates this hashicorp cloud webhook webhook payload
func (m *HashicorpCloudWebhookWebhookPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud webhook webhook payload based on context it is used
func (m *HashicorpCloudWebhookWebhookPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWebhookWebhookPayload) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWebhookWebhookPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
