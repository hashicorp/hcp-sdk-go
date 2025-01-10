// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ProtobufNullValue `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
// The JSON representation for `NullValue` is JSON `null`.
//
//   - NULL_VALUE: Null value.
//
// swagger:model protobufNullValue
type ProtobufNullValue string

func NewProtobufNullValue(value ProtobufNullValue) *ProtobufNullValue {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ProtobufNullValue.
func (m ProtobufNullValue) Pointer() *ProtobufNullValue {
	return &m
}

const (

	// ProtobufNullValueNULLVALUE captures enum value "NULL_VALUE"
	ProtobufNullValueNULLVALUE ProtobufNullValue = "NULL_VALUE"
)

// for schema
var protobufNullValueEnum []interface{}

func init() {
	var res []ProtobufNullValue
	if err := json.Unmarshal([]byte(`["NULL_VALUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protobufNullValueEnum = append(protobufNullValueEnum, v)
	}
}

func (m ProtobufNullValue) validateProtobufNullValueEnum(path, location string, value ProtobufNullValue) error {
	if err := validate.EnumCase(path, location, value, protobufNullValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this protobuf null value
func (m ProtobufNullValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtobufNullValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this protobuf null value based on context it is used
func (m ProtobufNullValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
