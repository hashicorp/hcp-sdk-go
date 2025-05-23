// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudNetwork20200907ForwardingRule ForwardingRule represents the DNS forwarding rules for a DNS forwarding
//
// swagger:model hashicorp.cloud.network_20200907.ForwardingRule
type HashicorpCloudNetwork20200907ForwardingRule struct {

	// domain_name is the domain name for which DNS Forwarding rule needs to be created
	DomainName string `json:"domain_name,omitempty"`

	// Id is the ID of the dns forwarding rule. It is currently optional, and will default to
	// a UUID if not set.
	ID string `json:"id,omitempty"`

	// InboundEndpointIPs are the IP addresses of the target customer network inbound endpoints
	// to which the DNS requests for the above domain will be forwarded
	InboundEndpointIps []string `json:"inbound_endpoint_ips"`
}

// Validate validates this hashicorp cloud network 20200907 forwarding rule
func (m *HashicorpCloudNetwork20200907ForwardingRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud network 20200907 forwarding rule based on context it is used
func (m *HashicorpCloudNetwork20200907ForwardingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ForwardingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907ForwardingRule) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907ForwardingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
