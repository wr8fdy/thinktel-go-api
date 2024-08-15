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

// ThinkteluControlAPIDataContractsPermanentCallForwardOrder thinktel u control Api data contracts permanent call forward order
//
// swagger:model Thinktel.uControl.Api.DataContracts.PermanentCallForwardOrder
type ThinkteluControlAPIDataContractsPermanentCallForwardOrder struct {

	// account
	// Required: true
	Account *int64 `json:"Account"`

	// call forwarding number
	// Required: true
	CallForwardingNumber *int64 `json:"CallForwardingNumber"`

	// enabled
	// Required: true
	Enabled *bool `json:"Enabled"`

	// label
	Label string `json:"Label,omitempty"`

	// number
	// Required: true
	Number *ThinkteluControlAPIDataContractsNumberRequest `json:"Number"`

	// plan ID
	// Required: true
	PlanID *string `json:"PlanID"`
}

// Validate validates this thinktel u control Api data contracts permanent call forward order
func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallForwardingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("Account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) validateCallForwardingNumber(formats strfmt.Registry) error {

	if err := validate.Required("CallForwardingNumber", "body", m.CallForwardingNumber); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) validateNumber(formats strfmt.Registry) error {

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

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) validatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("PlanID", "body", m.PlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts permanent call forward order based on the context it is used
func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) contextValidateNumber(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsPermanentCallForwardOrder) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsPermanentCallForwardOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}