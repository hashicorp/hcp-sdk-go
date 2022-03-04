// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudPackerUpdateIterationRequest hashicorp cloud packer update iteration request
//
// swagger:model hashicorp.cloud.packer.UpdateIterationRequest
type HashicorpCloudPackerUpdateIterationRequest struct {

	// Human-readable name for the bucket.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// Set to "true" when all builds associated with this iteration have
	// successfully completed and uploaded metadata to the registry. When
	// "complete" is true, this iteration is considered ready to use, and can
	// have channels assigned to it.
	Complete bool `json:"complete,omitempty"`

	// ULID of the iteration.
	IterationID string `json:"iteration_id,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// When set to true, will make a previously revoked iteration valid again.
	Restore bool `json:"restore,omitempty"`

	// Optional field to provide the reason for why this iteration is being revoked.
	RevocationMessage string `json:"revocation_message,omitempty"`

	// revoke_at accepts strings in the [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)
	// format to represent the revocation timestamp. To instantly revoke the iteration, provide the current timestamp.
	// The revoke_at timestamp will always be recorded in UTC (Coordinated Universal Time).
	// This option is equivalent to the 'revoke_in' option and therefore only one of them should be set when updating
	// the iteration.
	// Format: date-time
	RevokeAt strfmt.DateTime `json:"revoke_at,omitempty"`

	// revoke_in accepts a signed sequence of decimal numbers with a unit suffix to represent the duration
	// to the revocation date, such as '30d' or '2h45m'.
	// Valid time units are 's', 'm', 'h', and 'd' as for seconds, minutes, hours, and days.
	// To instantly revoke the iteration, provide the duration of zero seconds ("0s").
	// The revoke_in duration will be used to calculate the iteration revocation timestamp,
	// which will be recorded as UTC (Coordinated Universal Time).
	// This option is equivalent to the 'revoke_at' option and therefore only one of them should be set when updating
	// the iteration.
	RevokeIn string `json:"revoke_in,omitempty"`
}

// Validate validates this hashicorp cloud packer update iteration request
func (m *HashicorpCloudPackerUpdateIterationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevokeAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateIterationRequest) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPackerUpdateIterationRequest) validateRevokeAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RevokeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("revoke_at", "body", "date-time", m.RevokeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateIterationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateIterationRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerUpdateIterationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
