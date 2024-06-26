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

// DomainCPSRating domain c p s rating
//
// swagger:model domain.CPSRating
type DomainCPSRating struct {

	// current rating
	// Required: true
	CurrentRating *DomainCPSRatingHistoryEntry `json:"CurrentRating"`

	// highest rating
	// Required: true
	HighestRating *DomainCPSRatingHistoryEntry `json:"HighestRating"`

	// negative indicators
	// Required: true
	NegativeIndicators []*DomainCPSRatingIndicator `json:"NegativeIndicators"`

	// positive indicators
	// Required: true
	PositiveIndicators []*DomainCPSRatingIndicator `json:"PositiveIndicators"`

	// rating history
	// Required: true
	RatingHistory []*DomainCPSRatingHistoryEntry `json:"RatingHistory"`
}

// Validate validates this domain c p s rating
func (m *DomainCPSRating) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentRating(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHighestRating(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegativeIndicators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositiveIndicators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatingHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCPSRating) validateCurrentRating(formats strfmt.Registry) error {

	if err := validate.Required("CurrentRating", "body", m.CurrentRating); err != nil {
		return err
	}

	if m.CurrentRating != nil {
		if err := m.CurrentRating.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrentRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrentRating")
			}
			return err
		}
	}

	return nil
}

func (m *DomainCPSRating) validateHighestRating(formats strfmt.Registry) error {

	if err := validate.Required("HighestRating", "body", m.HighestRating); err != nil {
		return err
	}

	if m.HighestRating != nil {
		if err := m.HighestRating.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HighestRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HighestRating")
			}
			return err
		}
	}

	return nil
}

func (m *DomainCPSRating) validateNegativeIndicators(formats strfmt.Registry) error {

	if err := validate.Required("NegativeIndicators", "body", m.NegativeIndicators); err != nil {
		return err
	}

	for i := 0; i < len(m.NegativeIndicators); i++ {
		if swag.IsZero(m.NegativeIndicators[i]) { // not required
			continue
		}

		if m.NegativeIndicators[i] != nil {
			if err := m.NegativeIndicators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NegativeIndicators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NegativeIndicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainCPSRating) validatePositiveIndicators(formats strfmt.Registry) error {

	if err := validate.Required("PositiveIndicators", "body", m.PositiveIndicators); err != nil {
		return err
	}

	for i := 0; i < len(m.PositiveIndicators); i++ {
		if swag.IsZero(m.PositiveIndicators[i]) { // not required
			continue
		}

		if m.PositiveIndicators[i] != nil {
			if err := m.PositiveIndicators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PositiveIndicators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PositiveIndicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainCPSRating) validateRatingHistory(formats strfmt.Registry) error {

	if err := validate.Required("RatingHistory", "body", m.RatingHistory); err != nil {
		return err
	}

	for i := 0; i < len(m.RatingHistory); i++ {
		if swag.IsZero(m.RatingHistory[i]) { // not required
			continue
		}

		if m.RatingHistory[i] != nil {
			if err := m.RatingHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RatingHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RatingHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain c p s rating based on the context it is used
func (m *DomainCPSRating) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrentRating(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHighestRating(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNegativeIndicators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePositiveIndicators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatingHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCPSRating) contextValidateCurrentRating(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentRating != nil {

		if err := m.CurrentRating.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrentRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrentRating")
			}
			return err
		}
	}

	return nil
}

func (m *DomainCPSRating) contextValidateHighestRating(ctx context.Context, formats strfmt.Registry) error {

	if m.HighestRating != nil {

		if err := m.HighestRating.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HighestRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HighestRating")
			}
			return err
		}
	}

	return nil
}

func (m *DomainCPSRating) contextValidateNegativeIndicators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NegativeIndicators); i++ {

		if m.NegativeIndicators[i] != nil {

			if swag.IsZero(m.NegativeIndicators[i]) { // not required
				return nil
			}

			if err := m.NegativeIndicators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NegativeIndicators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NegativeIndicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainCPSRating) contextValidatePositiveIndicators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PositiveIndicators); i++ {

		if m.PositiveIndicators[i] != nil {

			if swag.IsZero(m.PositiveIndicators[i]) { // not required
				return nil
			}

			if err := m.PositiveIndicators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PositiveIndicators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PositiveIndicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainCPSRating) contextValidateRatingHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RatingHistory); i++ {

		if m.RatingHistory[i] != nil {

			if swag.IsZero(m.RatingHistory[i]) { // not required
				return nil
			}

			if err := m.RatingHistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RatingHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RatingHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainCPSRating) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCPSRating) UnmarshalBinary(b []byte) error {
	var res DomainCPSRating
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
