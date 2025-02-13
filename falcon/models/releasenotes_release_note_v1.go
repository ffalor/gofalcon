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

// ReleasenotesReleaseNoteV1 releasenotes release note v1
//
// swagger:model releasenotes.ReleaseNoteV1
type ReleasenotesReleaseNoteV1 struct {

	// categories
	// Required: true
	Categories []string `json:"categories"`

	// early access
	EarlyAccess string `json:"early_access,omitempty"`

	// general availability
	GeneralAvailability string `json:"general_availability,omitempty"`

	// notes
	// Required: true
	Notes []*ReleasenotesNoteDetail `json:"notes"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this releasenotes release note v1
func (m *ReleasenotesReleaseNoteV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleasenotesReleaseNoteV1) validateCategories(formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	return nil
}

func (m *ReleasenotesReleaseNoteV1) validateNotes(formats strfmt.Registry) error {

	if err := validate.Required("notes", "body", m.Notes); err != nil {
		return err
	}

	for i := 0; i < len(m.Notes); i++ {
		if swag.IsZero(m.Notes[i]) { // not required
			continue
		}

		if m.Notes[i] != nil {
			if err := m.Notes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleasenotesReleaseNoteV1) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleasenotesReleaseNoteV1) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this releasenotes release note v1 based on the context it is used
func (m *ReleasenotesReleaseNoteV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleasenotesReleaseNoteV1) contextValidateNotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notes); i++ {

		if m.Notes[i] != nil {

			if swag.IsZero(m.Notes[i]) { // not required
				return nil
			}

			if err := m.Notes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleasenotesReleaseNoteV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleasenotesReleaseNoteV1) UnmarshalBinary(b []byte) error {
	var res ReleasenotesReleaseNoteV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
