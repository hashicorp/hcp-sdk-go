// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageOptionDependency Dependency describes a dependency with a Package option.
//
// swagger:model PackageOptionDependency
type PackageOptionDependency struct {

	// item key
	ItemKey string `json:"item_key,omitempty"`

	// option key
	OptionKey string `json:"option_key,omitempty"`
}

// Validate validates this package option dependency
func (m *PackageOptionDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package option dependency based on context it is used
func (m *PackageOptionDependency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageOptionDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageOptionDependency) UnmarshalBinary(b []byte) error {
	var res PackageOptionDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
