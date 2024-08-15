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

// ThinkteluControlAPIDataContractsTerseNumber thinktel u control Api data contracts terse number
//
// swagger:model Thinktel.uControl.Api.DataContracts.TerseNumber
type ThinkteluControlAPIDataContractsTerseNumber struct {

	// label
	Label string `json:"Label,omitempty"`

	// number
	// Required: true
	Number *int64 `json:"Number"`
}

// Validate validates this thinktel u control Api data contracts terse number
func (m *ThinkteluControlAPIDataContractsTerseNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsTerseNumber) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("Number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts terse number based on context it is used
func (m *ThinkteluControlAPIDataContractsTerseNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsTerseNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsTerseNumber) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsTerseNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}