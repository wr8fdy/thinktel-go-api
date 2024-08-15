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

// ThinkteluControlAPIDataContractsSIPBinding thinktel u control Api data contracts Sip binding
//
// swagger:model Thinktel.uControl.Api.DataContracts.SipBinding
type ThinkteluControlAPIDataContractsSIPBinding struct {

	// channels
	Channels int32 `json:"Channels,omitempty"`

	//
	//
	//     DomainNameSrvLookup = 1
	//     IpAddressAndPort = 2
	//     DomainNameNaptrLookup = 3
	//     DomainNameAOrAaaaLookup = 4
	//
	ContactAddressScheme int64 `json:"ContactAddressScheme,omitempty"`

	// contact domain name
	ContactDomainName string `json:"ContactDomainName,omitempty"`

	// contact IP address
	ContactIPAddress string `json:"ContactIPAddress,omitempty"`

	// contact port
	ContactPort int32 `json:"ContactPort,omitempty"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// IP match required
	IPMatchRequired bool `json:"IPMatchRequired,omitempty"`

	// name
	// Required: true
	Name *string `json:"Name"`

	//
	//
	//     DomainNameSrvLookup = 1
	//     IpAddressAndPort = 2
	//     DomainNameNaptrLookup = 3
	//     DomainNameAOrAaaaLookup = 4
	//
	ProxyAddressScheme int64 `json:"ProxyAddressScheme,omitempty"`

	// proxy domain name
	ProxyDomainName string `json:"ProxyDomainName,omitempty"`

	// proxy IP address
	ProxyIPAddress string `json:"ProxyIPAddress,omitempty"`

	// proxy port
	ProxyPort int32 `json:"ProxyPort,omitempty"`

	// Sip domain name
	SIPDomainName string `json:"SipDomainName,omitempty"`

	// Sip password
	SIPPassword string `json:"SipPassword,omitempty"`

	// Sip username
	SIPUsername string `json:"SipUsername,omitempty"`

	// trunk type
	TrunkType string `json:"TrunkType,omitempty"`
}

// Validate validates this thinktel u control Api data contracts Sip binding
func (m *ThinkteluControlAPIDataContractsSIPBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPBinding) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPBinding) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts Sip binding based on context it is used
func (m *ThinkteluControlAPIDataContractsSIPBinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPBinding) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsSIPBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
