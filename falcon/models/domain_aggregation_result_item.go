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

// DomainAggregationResultItem domain aggregation result item
//
// swagger:model domain.AggregationResultItem
type DomainAggregationResultItem struct {

	// count of the documents in the bucket
	// Required: true
	Count *int64 `json:"count"`

	// numerical value of the date or number start of the range
	From float64 `json:"from,omitempty"`

	// string value of the key, usually not populated
	KeyAsString string `json:"key_as_string,omitempty"`

	// label of the bucket, populated all the time, usually value of the aggregated field or range values
	Label DomainAggregationResultItemLabel `json:"label,omitempty"`

	// string value of the from property, usually representing a date or number
	StringFrom string `json:"string_from,omitempty"`

	// string value of the to property, usually representing a date or number
	StringTo string `json:"string_to,omitempty"`

	// sub-aggregations of the bucket
	SubAggregates []*DomainAggregationResult `json:"sub_aggregates"`

	// numerical value of the date or number end of the range
	To float64 `json:"to,omitempty"`

	// numerical value of the bucket
	Value float64 `json:"value,omitempty"`

	// value as a string, usually not populated
	ValueAsString string `json:"value_as_string,omitempty"`
}

// Validate validates this domain aggregation result item
func (m *DomainAggregationResultItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAggregates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAggregationResultItem) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *DomainAggregationResultItem) validateSubAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.SubAggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.SubAggregates); i++ {
		if swag.IsZero(m.SubAggregates[i]) { // not required
			continue
		}

		if m.SubAggregates[i] != nil {
			if err := m.SubAggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain aggregation result item based on the context it is used
func (m *DomainAggregationResultItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAggregationResultItem) contextValidateSubAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubAggregates); i++ {

		if m.SubAggregates[i] != nil {

			if swag.IsZero(m.SubAggregates[i]) { // not required
				return nil
			}

			if err := m.SubAggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAggregationResultItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAggregationResultItem) UnmarshalBinary(b []byte) error {
	var res DomainAggregationResultItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}