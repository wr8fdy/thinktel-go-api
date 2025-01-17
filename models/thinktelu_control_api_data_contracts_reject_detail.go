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

// ThinkteluControlAPIDataContractsRejectDetail thinktel u control Api data contracts reject detail
//
// swagger:model Thinktel.uControl.Api.DataContracts.RejectDetail
type ThinkteluControlAPIDataContractsRejectDetail struct {

	// The end user telephone number to be disconnected, or the first telephone number in a range to be disconnected.
	// Max Length: 12
	// Min Length: 12
	DISC string `json:"DISC #,omitempty"`

	// The final four digits of the last telephone number in a range to be disconnected.
	// Max Length: 4
	// Min Length: 4
	DISCRANGE string `json:"DISC # RANGE,omitempty"`

	// The abbreviation for the field where the error has occurred.
	// Max Length: 20
	FLDABBR string `json:"FLD ABBR,omitempty"`

	// The specific form where the error has occurred.
	// Max Length: 4
	FORM string `json:"FORM,omitempty"`

	// The end user telephone number to be ported, or the first telephone number in a range to be ported.
	// Max Length: 12
	// Min Length: 12
	PORTED string `json:"PORTED #,omitempty"`

	// The final four digits of the last telephone number in a range to be ported.
	// Max Length: 4
	// Min Length: 4
	PORTEDRANGE string `json:"PORTED # RANGE,omitempty"`

	// Free flowing field which can be used to expand upon and clarify other data.
	// Max Length: 360
	REMARKS string `json:"REMARKS,omitempty"`

	// Code which describes the type of error which has occurred.
	// Max Length: 2
	// Min Length: 2
	RJCTCODE string `json:"RJCT CODE,omitempty"`

	// Free flowing field which can be used to expand upon and clarify reject data.
	// Max Length: 360
	RJCTREMARKS string `json:"RJCT REMARKS,omitempty"`
}

// Validate validates this thinktel u control Api data contracts reject detail
func (m *ThinkteluControlAPIDataContractsRejectDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDISC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDISCRANGE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFLDABBR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFORM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePORTED(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePORTEDRANGE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateREMARKS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRJCTCODE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRJCTREMARKS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateDISC(formats strfmt.Registry) error {
	if swag.IsZero(m.DISC) { // not required
		return nil
	}

	if err := validate.MinLength("DISC #", "body", m.DISC, 12); err != nil {
		return err
	}

	if err := validate.MaxLength("DISC #", "body", m.DISC, 12); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateDISCRANGE(formats strfmt.Registry) error {
	if swag.IsZero(m.DISCRANGE) { // not required
		return nil
	}

	if err := validate.MinLength("DISC # RANGE", "body", m.DISCRANGE, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("DISC # RANGE", "body", m.DISCRANGE, 4); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateFLDABBR(formats strfmt.Registry) error {
	if swag.IsZero(m.FLDABBR) { // not required
		return nil
	}

	if err := validate.MaxLength("FLD ABBR", "body", m.FLDABBR, 20); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateFORM(formats strfmt.Registry) error {
	if swag.IsZero(m.FORM) { // not required
		return nil
	}

	if err := validate.MaxLength("FORM", "body", m.FORM, 4); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validatePORTED(formats strfmt.Registry) error {
	if swag.IsZero(m.PORTED) { // not required
		return nil
	}

	if err := validate.MinLength("PORTED #", "body", m.PORTED, 12); err != nil {
		return err
	}

	if err := validate.MaxLength("PORTED #", "body", m.PORTED, 12); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validatePORTEDRANGE(formats strfmt.Registry) error {
	if swag.IsZero(m.PORTEDRANGE) { // not required
		return nil
	}

	if err := validate.MinLength("PORTED # RANGE", "body", m.PORTEDRANGE, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("PORTED # RANGE", "body", m.PORTEDRANGE, 4); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateREMARKS(formats strfmt.Registry) error {
	if swag.IsZero(m.REMARKS) { // not required
		return nil
	}

	if err := validate.MaxLength("REMARKS", "body", m.REMARKS, 360); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateRJCTCODE(formats strfmt.Registry) error {
	if swag.IsZero(m.RJCTCODE) { // not required
		return nil
	}

	if err := validate.MinLength("RJCT CODE", "body", m.RJCTCODE, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("RJCT CODE", "body", m.RJCTCODE, 2); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsRejectDetail) validateRJCTREMARKS(formats strfmt.Registry) error {
	if swag.IsZero(m.RJCTREMARKS) { // not required
		return nil
	}

	if err := validate.MaxLength("RJCT REMARKS", "body", m.RJCTREMARKS, 360); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts reject detail based on context it is used
func (m *ThinkteluControlAPIDataContractsRejectDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsRejectDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsRejectDetail) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsRejectDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
