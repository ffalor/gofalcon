// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainParams domain params
//
// swagger:model domain.Params
type DomainParams struct {

	// cookie
	Cookie interface{} `json:"cookie,omitempty"`

	// header
	Header interface{} `json:"header,omitempty"`

	// path
	Path interface{} `json:"path,omitempty"`

	// query
	Query interface{} `json:"query,omitempty"`
}

// Validate validates this domain params
func (m *DomainParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain params based on context it is used
func (m *DomainParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainParams) UnmarshalBinary(b []byte) error {
	var res DomainParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}