// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesAzureLoadBalancerReasonTag types azure load balancer reason tag
//
// swagger:model types.AzureLoadBalancerReasonTag
type TypesAzureLoadBalancerReasonTag struct {

	// name
	Name string `json:"name,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// vm name
	VMName string `json:"vmName,omitempty"`
}

// Validate validates this types azure load balancer reason tag
func (m *TypesAzureLoadBalancerReasonTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types azure load balancer reason tag based on context it is used
func (m *TypesAzureLoadBalancerReasonTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesAzureLoadBalancerReasonTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesAzureLoadBalancerReasonTag) UnmarshalBinary(b []byte) error {
	var res TypesAzureLoadBalancerReasonTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}