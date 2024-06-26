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

// ModelsApplicationPackageInfoType models application package info type
//
// swagger:model models.ApplicationPackageInfoType
type ModelsApplicationPackageInfoType struct {

	// libraries
	// Required: true
	Libraries []*ModelsApplicationLibrary `json:"libraries"`

	// type
	// Required: true
	Type *int32 `json:"type"`
}

// Validate validates this models application package info type
func (m *ModelsApplicationPackageInfoType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLibraries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsApplicationPackageInfoType) validateLibraries(formats strfmt.Registry) error {

	if err := validate.Required("libraries", "body", m.Libraries); err != nil {
		return err
	}

	for i := 0; i < len(m.Libraries); i++ {
		if swag.IsZero(m.Libraries[i]) { // not required
			continue
		}

		if m.Libraries[i] != nil {
			if err := m.Libraries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libraries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("libraries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsApplicationPackageInfoType) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this models application package info type based on the context it is used
func (m *ModelsApplicationPackageInfoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLibraries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsApplicationPackageInfoType) contextValidateLibraries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Libraries); i++ {

		if m.Libraries[i] != nil {

			if swag.IsZero(m.Libraries[i]) { // not required
				return nil
			}

			if err := m.Libraries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libraries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("libraries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsApplicationPackageInfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsApplicationPackageInfoType) UnmarshalBinary(b []byte) error {
	var res ModelsApplicationPackageInfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
