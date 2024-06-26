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

// ModelsScanResults models scan results
//
// swagger:model models.ScanResults
type ModelsScanResults struct {

	// applications
	// Required: true
	Applications []*ModelsSnapshotInventoryApplication `json:"applications"`

	// containers
	// Required: true
	Containers []*ModelsContainer `json:"containers"`

	// elf binaries
	// Required: true
	ElfBinaries []*ModelsELFBinary `json:"elf_binaries"`

	// os version
	// Required: true
	OsVersion *string `json:"os_version"`
}

// Validate validates this models scan results
func (m *ModelsScanResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfBinaries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsScanResults) validateApplications(formats strfmt.Registry) error {

	if err := validate.Required("applications", "body", m.Applications); err != nil {
		return err
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsScanResults) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsScanResults) validateElfBinaries(formats strfmt.Registry) error {

	if err := validate.Required("elf_binaries", "body", m.ElfBinaries); err != nil {
		return err
	}

	for i := 0; i < len(m.ElfBinaries); i++ {
		if swag.IsZero(m.ElfBinaries[i]) { // not required
			continue
		}

		if m.ElfBinaries[i] != nil {
			if err := m.ElfBinaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elf_binaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elf_binaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsScanResults) validateOsVersion(formats strfmt.Registry) error {

	if err := validate.Required("os_version", "body", m.OsVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this models scan results based on the context it is used
func (m *ModelsScanResults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfBinaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsScanResults) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Applications); i++ {

		if m.Applications[i] != nil {

			if swag.IsZero(m.Applications[i]) { // not required
				return nil
			}

			if err := m.Applications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsScanResults) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {

			if swag.IsZero(m.Containers[i]) { // not required
				return nil
			}

			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsScanResults) contextValidateElfBinaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ElfBinaries); i++ {

		if m.ElfBinaries[i] != nil {

			if swag.IsZero(m.ElfBinaries[i]) { // not required
				return nil
			}

			if err := m.ElfBinaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elf_binaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elf_binaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsScanResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsScanResults) UnmarshalBinary(b []byte) error {
	var res ModelsScanResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
