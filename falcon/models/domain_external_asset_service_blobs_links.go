// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainExternalAssetServiceBlobsLinks domain external asset service blobs links
//
// swagger:model domain.ExternalAssetServiceBlobsLinks
type DomainExternalAssetServiceBlobsLinks struct {

	// download attributes
	DownloadAttributes string `json:"download_attributes,omitempty"`

	// download banner
	DownloadBanner string `json:"download_banner,omitempty"`

	// preview attributes
	PreviewAttributes string `json:"preview_attributes,omitempty"`

	// preview banner
	PreviewBanner string `json:"preview_banner,omitempty"`
}

// Validate validates this domain external asset service blobs links
func (m *DomainExternalAssetServiceBlobsLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain external asset service blobs links based on context it is used
func (m *DomainExternalAssetServiceBlobsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainExternalAssetServiceBlobsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainExternalAssetServiceBlobsLinks) UnmarshalBinary(b []byte) error {
	var res DomainExternalAssetServiceBlobsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
