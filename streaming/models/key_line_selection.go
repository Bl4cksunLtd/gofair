// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeyLineSelection key line selection
//
// swagger:model KeyLineSelection
type KeyLineSelection struct {

	// hc
	Hc float64 `json:"hc,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this key line selection
func (m *KeyLineSelection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this key line selection based on context it is used
func (m *KeyLineSelection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyLineSelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyLineSelection) UnmarshalBinary(b []byte) error {
	var res KeyLineSelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
