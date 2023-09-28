// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIExposedDataRecordBotOperatingSystemV1 api exposed data record bot operating system v1
//
// swagger:model api.ExposedDataRecordBotOperatingSystemV1
type APIExposedDataRecordBotOperatingSystemV1 struct {

	// antivirus
	Antivirus []string `json:"antivirus"`

	// computer name
	ComputerName string `json:"computer_name,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// hardware id
	HardwareID string `json:"hardware_id,omitempty"`

	// installed software
	InstalledSoftware []string `json:"installed_software"`

	// language
	Language string `json:"language,omitempty"`

	// layouts
	Layouts []string `json:"layouts"`

	// os architecture
	OsArchitecture string `json:"os_architecture,omitempty"`

	// os version
	OsVersion string `json:"os_version,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// uac
	Uac string `json:"uac,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this api exposed data record bot operating system v1
func (m *APIExposedDataRecordBotOperatingSystemV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api exposed data record bot operating system v1 based on context it is used
func (m *APIExposedDataRecordBotOperatingSystemV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIExposedDataRecordBotOperatingSystemV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIExposedDataRecordBotOperatingSystemV1) UnmarshalBinary(b []byte) error {
	var res APIExposedDataRecordBotOperatingSystemV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}