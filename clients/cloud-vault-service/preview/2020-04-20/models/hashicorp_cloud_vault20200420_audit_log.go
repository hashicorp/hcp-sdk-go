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

// HashicorpCloudVault20200420AuditLog AuditLog represents a request for audit logs to download
//
// swagger:model hashicorp.cloud.vault_20200420.AuditLog
type HashicorpCloudVault20200420AuditLog struct {

	// cluster_id is the cluster id that this download backs
	ClusterID string `json:"cluster_id,omitempty"`

	// download_url is the URL which can be used to retrieve the audit logs.
	DownloadURL string `json:"download_url,omitempty"`

	// expires_at is the timestamp of when the cluster was first created.
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// finished_at is the timestamp of when the download was ready.
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// id is the id of the audit log download
	ID string `json:"id,omitempty"`

	// interval_end
	// Format: date-time
	IntervalEnd strfmt.DateTime `json:"interval_end,omitempty"`

	// interval_start
	// Format: date-time
	IntervalStart strfmt.DateTime `json:"interval_start,omitempty"`

	// location is the location of the cluster.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// state is the current state of the download
	State HashicorpCloudVault20200420AuditLogState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 audit log
func (m *HashicorpCloudVault20200420AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420AuditLog) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420AuditLog) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420AuditLog) validateIntervalEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.IntervalEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("interval_end", "body", "date-time", m.IntervalEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420AuditLog) validateIntervalStart(formats strfmt.Registry) error {

	if swag.IsZero(m.IntervalStart) { // not required
		return nil
	}

	if err := validate.FormatOf("interval_start", "body", "date-time", m.IntervalStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420AuditLog) validateLocation(formats strfmt.Registry) error {

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

func (m *HashicorpCloudVault20200420AuditLog) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420AuditLog) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
