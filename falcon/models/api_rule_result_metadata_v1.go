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

// APIRuleResultMetadataV1 api rule result metadata v1
//
// swagger:model api.RuleResultMetadataV1
type APIRuleResultMetadataV1 struct {

	// execution duration
	ExecutionDuration int32 `json:"execution_duration,omitempty"`

	// execution finish
	// Format: date-time
	ExecutionFinish strfmt.DateTime `json:"execution_finish,omitempty"`

	// execution start
	// Format: date-time
	ExecutionStart strfmt.DateTime `json:"execution_start,omitempty"`

	// result count
	ResultCount int32 `json:"result_count,omitempty"`

	// search window end
	// Format: date-time
	SearchWindowEnd strfmt.DateTime `json:"search_window_end,omitempty"`
}

// Validate validates this api rule result metadata v1
func (m *APIRuleResultMetadataV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionFinish(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchWindowEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleResultMetadataV1) validateExecutionFinish(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutionFinish) { // not required
		return nil
	}

	if err := validate.FormatOf("execution_finish", "body", "date-time", m.ExecutionFinish.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleResultMetadataV1) validateExecutionStart(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutionStart) { // not required
		return nil
	}

	if err := validate.FormatOf("execution_start", "body", "date-time", m.ExecutionStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleResultMetadataV1) validateSearchWindowEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchWindowEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("search_window_end", "body", "date-time", m.SearchWindowEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this api rule result metadata v1 based on context it is used
func (m *APIRuleResultMetadataV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleResultMetadataV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleResultMetadataV1) UnmarshalBinary(b []byte) error {
	var res APIRuleResultMetadataV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
