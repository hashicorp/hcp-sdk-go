// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudBoundary20211221Session Session represents single Boundary session.
//
// swagger:model hashicorp.cloud.boundary_20211221.Session
type HashicorpCloudBoundary20211221Session struct {

	// id is the ID of the Session.
	ID string `json:"id,omitempty"`

	// status is the current status of this session.
	Status string `json:"status,omitempty"`

	// type is the type of the Session (e.g. tcp).
	Type string `json:"type,omitempty"`
}

// Validate validates this hashicorp cloud boundary 20211221 session
func (m *HashicorpCloudBoundary20211221Session) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221Session) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudBoundary20211221Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
