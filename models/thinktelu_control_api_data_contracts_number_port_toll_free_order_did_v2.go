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

// ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2 Telephone numbers port-in request form for Toll-Free SIP trunk DID service.
//
// swagger:model Thinktel.uControl.Api.DataContracts.NumberPortTollFreeOrderDIDV2
type ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2 struct {

	// DIDs to be added to the specified SIP trunk at end of port-in process.
	// Required: true
	DIDs []*ThinkteluControlAPIDataContractsDIDRequestV2 `json:"DIDs"`

	// Existing pilot number of SIP trunk to add the new DIDs.
	// Required: true
	PilotNumber *int64 `json:"PilotNumber"`

	// port form
	// Required: true
	PortForm *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2 `json:"PortForm"`
}

// Validate validates this thinktel u control Api data contracts number port toll free order DID v2
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotNumber(formats); err != nil {
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

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) validateDIDs(formats strfmt.Registry) error {

	if err := validate.Required("DIDs", "body", m.DIDs); err != nil {
		return err
	}

	for i := 0; i < len(m.DIDs); i++ {
		if swag.IsZero(m.DIDs[i]) { // not required
			continue
		}

		if m.DIDs[i] != nil {
			if err := m.DIDs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DIDs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DIDs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) validatePilotNumber(formats strfmt.Registry) error {

	if err := validate.Required("PilotNumber", "body", m.PilotNumber); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) validatePortForm(formats strfmt.Registry) error {

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

// ContextValidate validate this thinktel u control Api data contracts number port toll free order DID v2 based on the context it is used
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDIDs(ctx, formats); err != nil {
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

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) contextValidateDIDs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DIDs); i++ {

		if m.DIDs[i] != nil {

			if swag.IsZero(m.DIDs[i]) { // not required
				return nil
			}

			if err := m.DIDs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DIDs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DIDs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) contextValidatePortForm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsNumberPortTollFreeOrderDIDV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
