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

// ThinkteluControlAPIDataContractsBusinessLine thinktel u control Api data contracts business line
//
// swagger:model Thinktel.uControl.Api.DataContracts.BusinessLine
type ThinkteluControlAPIDataContractsBusinessLine struct {

	// account
	// Required: true
	Account *int64 `json:"Account"`

	// business group ID
	// Required: true
	BusinessGroupID *string `json:"BusinessGroupID"`

	// call forwarding delay
	CallForwardingDelay int32 `json:"CallForwardingDelay,omitempty"`

	// call forwarding number
	CallForwardingNumber int64 `json:"CallForwardingNumber,omitempty"`

	// caller ID
	CallerID int64 `json:"CallerID,omitempty"`

	// caller name
	CallerName string `json:"CallerName,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"Enabled"`

	// extension
	Extension int32 `json:"Extension,omitempty"`

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
	Number *int64 `json:"Number"`

	// plan ID
	// Required: true
	PlanID *string `json:"PlanID"`

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

	// Sip domain name
	SIPDomainName string `json:"SipDomainName,omitempty"`

	// Sip password
	SIPPassword string `json:"SipPassword,omitempty"`

	// web self care password
	WebSelfCarePassword string `json:"WebSelfCarePassword,omitempty"`
}

// Validate validates this thinktel u control Api data contracts business line
func (m *ThinkteluControlAPIDataContractsBusinessLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusinessGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
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

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("Account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateBusinessGroupID(formats strfmt.Registry) error {

	if err := validate.Required("BusinessGroupID", "body", m.BusinessGroupID); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("Locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("Number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("PlanID", "body", m.PlanID); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsBusinessLine) validateSecondLocale(formats strfmt.Registry) error {

	if err := validate.Required("SecondLocale", "body", m.SecondLocale); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts business line based on context it is used
func (m *ThinkteluControlAPIDataContractsBusinessLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsBusinessLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsBusinessLine) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsBusinessLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
