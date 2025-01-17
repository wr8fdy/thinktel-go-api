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

// ThinkteluControlAPIDataContractsNumberItem thinktel u control Api data contracts number item
//
// swagger:model Thinktel.uControl.Api.DataContracts.NumberItem
type ThinkteluControlAPIDataContractsNumberItem struct {

	// number
	// Required: true
	Number *int64 `json:"Number"`
}

// Validate validates this thinktel u control Api data contracts number item
func (m *ThinkteluControlAPIDataContractsNumberItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberItem) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("Number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts number item based on context it is used
func (m *ThinkteluControlAPIDataContractsNumberItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberItem) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsNumberItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
