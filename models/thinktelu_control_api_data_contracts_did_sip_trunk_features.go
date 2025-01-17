// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures thinktel u control Api data contracts DID Sip trunk features
//
// swagger:model Thinktel.uControl.Api.DataContracts.DIDSipTrunkFeatures
type ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures struct {

	// c l ID
	CLID string `json:"CLID,omitempty"`

	// c l ID specified
	CLIDSpecified bool `json:"CLIDSpecified,omitempty"`

	// c n a m
	CNAM string `json:"CNAM,omitempty"`

	// c n a m specified
	CNAMSpecified bool `json:"CNAMSpecified,omitempty"`

	// external911 notification email
	External911NotificationEmail string `json:"External911NotificationEmail,omitempty"`

	// external911 notification email specified
	External911NotificationEmailSpecified bool `json:"External911NotificationEmailSpecified,omitempty"`
}

// Validate validates this thinktel u control Api data contracts DID Sip trunk features
func (m *ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this thinktel u control Api data contracts DID Sip trunk features based on context it is used
func (m *ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsDIDSIPTrunkFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
