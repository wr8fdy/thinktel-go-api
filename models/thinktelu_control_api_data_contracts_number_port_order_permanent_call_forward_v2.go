// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2 thinktel u control Api data contracts number port order permanent call forward v2
//
// swagger:model Thinktel.uControl.Api.DataContracts.NumberPortOrderPermanentCallForwardV2
type ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2 struct {

	// forwards
	// Required: true
	Forwards []*ThinkteluControlAPIDataContractsPermanentCallForwardOrderV2 `json:"Forwards"`

	// port form
	// Required: true
	PortForm *ThinkteluControlAPIDataContractsNumberPortOrderFormV2 `json:"PortForm"`
}

// Validate validates this thinktel u control Api data contracts number port order permanent call forward v2
func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForwards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) validateForwards(formats strfmt.Registry) error {

	if err := validate.Required("Forwards", "body", m.Forwards); err != nil {
		return err
	}

	for i := 0; i < len(m.Forwards); i++ {
		if swag.IsZero(m.Forwards[i]) { // not required
			continue
		}

		if m.Forwards[i] != nil {
			if err := m.Forwards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Forwards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Forwards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) validatePortForm(formats strfmt.Registry) error {

	if err := validate.Required("PortForm", "body", m.PortForm); err != nil {
		return err
	}

	if m.PortForm != nil {
		if err := m.PortForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PortForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PortForm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts number port order permanent call forward v2 based on the context it is used
func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateForwards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) contextValidateForwards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Forwards); i++ {

		if m.Forwards[i] != nil {

			if swag.IsZero(m.Forwards[i]) { // not required
				return nil
			}

			if err := m.Forwards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Forwards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Forwards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) contextValidatePortForm(ctx context.Context, formats strfmt.Registry) error {

	if m.PortForm != nil {

		if err := m.PortForm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PortForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PortForm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsNumberPortOrderPermanentCallForwardV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
