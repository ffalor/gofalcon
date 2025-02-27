// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FigapiAdversary figapi adversary
//
// swagger:model figapi.Adversary
type FigapiAdversary struct {

	// Name of the adversary
	Name string `json:"Name,omitempty"`
}

// Validate validates this figapi adversary
func (m *FigapiAdversary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this figapi adversary based on context it is used
func (m *FigapiAdversary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FigapiAdversary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FigapiAdversary) UnmarshalBinary(b []byte) error {
	var res FigapiAdversary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
