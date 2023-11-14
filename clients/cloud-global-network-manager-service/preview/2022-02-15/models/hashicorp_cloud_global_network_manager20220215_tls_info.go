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
	"github.com/go-openapi/validate"
)

// HashicorpCloudGlobalNetworkManager20220215TLSInfo hashicorp cloud global network manager 20220215 TLS info
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.TLSInfo
type HashicorpCloudGlobalNetworkManager20220215TLSInfo struct {

	// cert expiry
	// Format: date-time
	CertExpiry strfmt.DateTime `json:"cert_expiry,omitempty"`

	// cert_issuer is the Subject.CommonName of the issuer of the server
	// certificate.
	CertIssuer string `json:"cert_issuer,omitempty"`

	// cert name
	CertName string `json:"cert_name,omitempty"`

	// cert serial
	CertSerial string `json:"cert_serial,omitempty"`

	// certificate_authorities is the list of trusted certificate authorities
	CertificateAuthorities []*HashicorpCloudGlobalNetworkManager20220215CertificateMetadata `json:"certificate_authorities"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// tls_autorotated is true if HCP manages the cluster's TLS certificates,
	// the active server certificate was issued by the HCP-managed Cluster Signer CA,
	// and the cluster's Consul version supports automatic rotation of certs.
	TLSAutorotated bool `json:"tls_autorotated,omitempty"`

	// verify incoming
	VerifyIncoming bool `json:"verify_incoming,omitempty"`

	// verify outgoing
	VerifyOutgoing bool `json:"verify_outgoing,omitempty"`

	// verify server hostname
	VerifyServerHostname bool `json:"verify_server_hostname,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 TLS info
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateAuthorities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) validateCertExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.CertExpiry) { // not required
		return nil
	}

	if err := validate.FormatOf("cert_expiry", "body", "date-time", m.CertExpiry.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) validateCertificateAuthorities(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateAuthorities) { // not required
		return nil
	}

	for i := 0; i < len(m.CertificateAuthorities); i++ {
		if swag.IsZero(m.CertificateAuthorities[i]) { // not required
			continue
		}

		if m.CertificateAuthorities[i] != nil {
			if err := m.CertificateAuthorities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificate_authorities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificate_authorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 TLS info based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateAuthorities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) contextValidateCertificateAuthorities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CertificateAuthorities); i++ {

		if m.CertificateAuthorities[i] != nil {

			if swag.IsZero(m.CertificateAuthorities[i]) { // not required
				return nil
			}

			if err := m.CertificateAuthorities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificate_authorities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificate_authorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215TLSInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
