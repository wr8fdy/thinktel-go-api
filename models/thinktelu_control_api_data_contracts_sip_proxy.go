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

// ThinkteluControlAPIDataContractsSIPProxy thinktel u control Api data contracts Sip proxy
//
// swagger:model Thinktel.uControl.Api.DataContracts.SipProxy
type ThinkteluControlAPIDataContractsSIPProxy struct {

	// IP address
	// Required: true
	IPAddress *string `json:"IPAddress"`

	// n a t support
	NATSupport bool `json:"NATSupport,omitempty"`

	// port
	// Required: true
	Port *int32 `json:"Port"`

	// Sip domain name
	SIPDomainName string `json:"SipDomainName,omitempty"`

	//
	//
	//     SipTrunk = 0
	//     HostedLine = 1
	//     SipTrunkAndHostedLine = 2
	//     MicrosoftLync = 3
	//
	Type int64 `json:"Type,omitempty"`
}

// Validate validates this thinktel u control Api data contracts Sip proxy
func (m *ThinkteluControlAPIDataContractsSIPProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPProxy) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("IPAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPProxy) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("Port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts Sip proxy based on context it is used
func (m *ThinkteluControlAPIDataContractsSIPProxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPProxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPProxy) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsSIPProxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}