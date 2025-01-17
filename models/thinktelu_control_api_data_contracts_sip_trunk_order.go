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

// ThinkteluControlAPIDataContractsSIPTrunkOrder thinktel u control Api data contracts Sip trunk order
//
// swagger:model Thinktel.uControl.Api.DataContracts.SipTrunkOrder
type ThinkteluControlAPIDataContractsSIPTrunkOrder struct {

	// account
	Account int64 `json:"Account,omitempty"`

	// committed channels
	CommittedChannels int32 `json:"CommittedChannels,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"Enabled"`

	// international call barring enabled
	InternationalCallBarringEnabled bool `json:"InternationalCallBarringEnabled,omitempty"`

	// label
	Label string `json:"Label,omitempty"`

	//
	//
	//     EnglishUS = 0
	//     EnglishUK = 1
	//     EnglishCA = 2
	//     UserDefined1 = 3
	//     UserDefined2 = 4
	//     FrenchCA = 5
	//     SpanishLatinAmerica = 6
	//
	// Required: true
	Locale *int64 `json:"Locale"`

	// number
	// Required: true
	Number *ThinkteluControlAPIDataContractsNumberRequest `json:"Number"`

	// plan ID
	PlanID string `json:"PlanID,omitempty"`

	// premium911 notification enabled
	Premium911NotificationEnabled bool `json:"Premium911NotificationEnabled,omitempty"`

	// premium name enabled
	PremiumNameEnabled bool `json:"PremiumNameEnabled,omitempty"`

	// reseller plan ID
	ResellerPlanID string `json:"ResellerPlanID,omitempty"`

	//
	//
	//     EnglishUS = 0
	//     EnglishUK = 1
	//     EnglishCA = 2
	//     UserDefined1 = 3
	//     UserDefined2 = 4
	//     FrenchCA = 5
	//     SpanishLatinAmerica = 6
	//     None = 9
	//
	// Required: true
	SecondLocale *int64 `json:"SecondLocale"`

	// service suspended
	ServiceSuspended int32 `json:"ServiceSuspended,omitempty"`

	// third party label
	ThirdPartyLabel string `json:"ThirdPartyLabel,omitempty"`
}

// Validate validates this thinktel u control Api data contracts Sip trunk order
func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondLocale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("Locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("Number", "body", m.Number); err != nil {
		return err
	}

	if m.Number != nil {
		if err := m.Number.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Number")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Number")
			}
			return err
		}
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) validateSecondLocale(formats strfmt.Registry) error {

	if err := validate.Required("SecondLocale", "body", m.SecondLocale); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts Sip trunk order based on the context it is used
func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) contextValidateNumber(ctx context.Context, formats strfmt.Registry) error {

	if m.Number != nil {

		if err := m.Number.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Number")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Number")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSIPTrunkOrder) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsSIPTrunkOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
