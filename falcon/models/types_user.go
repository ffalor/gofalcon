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

// TypesUser types user
//
// swagger:model types.User
type TypesUser struct {

	// accepted terms
	// Required: true
	AcceptedTerms *bool `json:"acceptedTerms"`

	// email
	// Required: true
	Email *string `json:"email"`

	// external
	// Required: true
	External *bool `json:"external"`

	// fullname
	// Required: true
	Fullname *string `json:"fullname"`

	// groups
	// Required: true
	Groups []*TypesUserGroup `json:"groups"`

	// job title
	// Required: true
	JobTitle *string `json:"jobTitle"`

	// last login
	// Required: true
	// Format: date-time
	LastLogin *strfmt.DateTime `json:"lastLogin"`

	// role
	// Required: true
	Role *string `json:"role"`

	// status
	// Required: true
	Status *string `json:"status"`

	// user Id
	// Required: true
	UserID *int64 `json:"userId"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this types user
func (m *TypesUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedTerms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesUser) validateAcceptedTerms(formats strfmt.Registry) error {

	if err := validate.Required("acceptedTerms", "body", m.AcceptedTerms); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateExternal(formats strfmt.Registry) error {

	if err := validate.Required("external", "body", m.External); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateFullname(formats strfmt.Registry) error {

	if err := validate.Required("fullname", "body", m.Fullname); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesUser) validateJobTitle(formats strfmt.Registry) error {

	if err := validate.Required("jobTitle", "body", m.JobTitle); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateLastLogin(formats strfmt.Registry) error {

	if err := validate.Required("lastLogin", "body", m.LastLogin); err != nil {
		return err
	}

	if err := validate.FormatOf("lastLogin", "body", "date-time", m.LastLogin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *TypesUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this types user based on the context it is used
func (m *TypesUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesUser) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {

			if swag.IsZero(m.Groups[i]) { // not required
				return nil
			}

			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesUser) UnmarshalBinary(b []byte) error {
	var res TypesUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
