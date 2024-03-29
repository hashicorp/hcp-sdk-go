// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatementResourceLineItem LineItem represents a line item on an invoice.
//
// swagger:model StatementResourceLineItem
type StatementResourceLineItem struct {

	// amount is cost of this line item in USD.
	Amount string `json:"amount,omitempty"`

	// average_quantity is the average units of this product for the time the line item was active.
	AverageQuantity string `json:"average_quantity,omitempty"`

	// average_unit_display_key is human representation label of the average quantity for this product.
	AverageUnitDisplayKey string `json:"average_unit_display_key,omitempty"`

	// description is the name of the product associated with this line item.
	Description string `json:"description,omitempty"`

	// ended_at indicates when usage for the line item ended.
	// Format: date-time
	EndedAt strfmt.DateTime `json:"ended_at,omitempty"`

	// quantity is how many units of this product this item contains.
	Quantity string `json:"quantity,omitempty"`

	// started_at indicates when usage for the line item started.
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// unit_display_key is the human representation label of the unit that was used in this product.
	UnitDisplayKey string `json:"unit_display_key,omitempty"`

	// unit_price is the price per unit of the product associated.
	// TODO: Use existing enum.
	UnitPrice string `json:"unit_price,omitempty"`
}

// Validate validates this statement resource line item
func (m *StatementResourceLineItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatementResourceLineItem) validateEndedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.EndedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ended_at", "body", "date-time", m.EndedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StatementResourceLineItem) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this statement resource line item based on context it is used
func (m *StatementResourceLineItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatementResourceLineItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatementResourceLineItem) UnmarshalBinary(b []byte) error {
	var res StatementResourceLineItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
