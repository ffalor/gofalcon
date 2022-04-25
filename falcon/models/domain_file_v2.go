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

// DomainFileV2 domain file v2
//
// swagger:model domain.FileV2
type DomainFileV2 struct {

	// cloud request id
	// Required: true
	CloudRequestID *string `json:"cloud_request_id"`

	// complete
	// Required: true
	Complete *bool `json:"complete"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// deleted at
	// Required: true
	// Format: date-time
	DeletedAt *strfmt.DateTime `json:"deleted_at"`

	// error message
	// Required: true
	ErrorMessage *string `json:"error_message"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// progress
	// Required: true
	Progress *float32 `json:"progress"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`

	// sha256
	// Required: true
	Sha256 *string `json:"sha256"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// stage
	// Required: true
	Stage *string `json:"stage"`

	// status
	// Required: true
	Status *string `json:"status"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this domain file v2
func (m *DomainFileV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainFileV2) validateCloudRequestID(formats strfmt.Registry) error {

	if err := validate.Required("cloud_request_id", "body", m.CloudRequestID); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateComplete(formats strfmt.Registry) error {

	if err := validate.Required("complete", "body", m.Complete); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateDeletedAt(formats strfmt.Registry) error {

	if err := validate.Required("deleted_at", "body", m.DeletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateErrorMessage(formats strfmt.Registry) error {

	if err := validate.Required("error_message", "body", m.ErrorMessage); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateSha256(formats strfmt.Registry) error {

	if err := validate.Required("sha256", "body", m.Sha256); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateStage(formats strfmt.Registry) error {

	if err := validate.Required("stage", "body", m.Stage); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainFileV2) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain file v2 based on context it is used
func (m *DomainFileV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainFileV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainFileV2) UnmarshalBinary(b []byte) error {
	var res DomainFileV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
