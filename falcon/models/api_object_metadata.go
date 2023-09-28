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

// APIObjectMetadata api object metadata
//
// swagger:model api.ObjectMetadata
type APIObjectMetadata struct {

	// collection name
	// Required: true
	CollectionName *string `json:"collection_name"`

	// last modified time
	// Required: true
	// Format: date-time
	LastModifiedTime *strfmt.DateTime `json:"last_modified_time"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// object key
	// Required: true
	ObjectKey *string `json:"object_key"`

	// schema version
	// Required: true
	SchemaVersion *string `json:"schema_version"`
}

// Validate validates this api object metadata
func (m *APIObjectMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemaVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIObjectMetadata) validateCollectionName(formats strfmt.Registry) error {

	if err := validate.Required("collection_name", "body", m.CollectionName); err != nil {
		return err
	}

	return nil
}

func (m *APIObjectMetadata) validateLastModifiedTime(formats strfmt.Registry) error {

	if err := validate.Required("last_modified_time", "body", m.LastModifiedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified_time", "body", "date-time", m.LastModifiedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIObjectMetadata) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *APIObjectMetadata) validateObjectKey(formats strfmt.Registry) error {

	if err := validate.Required("object_key", "body", m.ObjectKey); err != nil {
		return err
	}

	return nil
}

func (m *APIObjectMetadata) validateSchemaVersion(formats strfmt.Registry) error {

	if err := validate.Required("schema_version", "body", m.SchemaVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this api object metadata based on context it is used
func (m *APIObjectMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIObjectMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIObjectMetadata) UnmarshalBinary(b []byte) error {
	var res APIObjectMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}