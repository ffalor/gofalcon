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

// DomainNotificationDetailsV1 domain notification details v1
//
// swagger:model domain.NotificationDetailsV1
type DomainNotificationDetailsV1 struct {

	// The raw intelligence item author username
	Author string `json:"author,omitempty"`

	// Highlighted content based on the rule that generated the notifications. Highlights are surrounded with a <cs-highlight> tag
	// Required: true
	Content *string `json:"content"`

	// The date when the raw intelligence item was created
	// Required: true
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"created_date"`

	// The raw intelligence item labels. These contain hints around what is actually included in the item (malware, IPs, emails, etc).
	Labels []string `json:"labels"`

	// The raw intelligence item language
	Language string `json:"language,omitempty"`

	// The site where the raw intelligence item was found
	Site string `json:"site,omitempty"`

	// The raw intelligence item title
	Title string `json:"title,omitempty"`

	// The ID of the notifications
	// Required: true
	Type *string `json:"type"`

	// The date when the raw intelligence item was updated
	// Required: true
	// Format: date-time
	UpdatedDate *strfmt.DateTime `json:"updated_date"`

	// The raw intelligence item URL
	URL string `json:"url,omitempty"`
}

// Validate validates this domain notification details v1
func (m *DomainNotificationDetailsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNotificationDetailsV1) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationDetailsV1) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("created_date", "body", m.CreatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("created_date", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationDetailsV1) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationDetailsV1) validateUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("updated_date", "body", m.UpdatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_date", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain notification details v1 based on context it is used
func (m *DomainNotificationDetailsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainNotificationDetailsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainNotificationDetailsV1) UnmarshalBinary(b []byte) error {
	var res DomainNotificationDetailsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}