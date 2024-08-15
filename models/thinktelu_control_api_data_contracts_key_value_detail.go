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

// ThinkteluControlAPIDataContractsKeyValueDetail Extra detail as a key-value pair.
//
// swagger:model Thinktel.uControl.Api.DataContracts.KeyValueDetail
type ThinkteluControlAPIDataContractsKeyValueDetail struct {

	// Name, label, or subject of extra detail.
	// Required: true
	Key *string `json:"Key"`

	// Value or text of extra detail.
	// Required: true
	Value *string `json:"Value"`
}

// Validate validates this thinktel u control Api data contracts key value detail
func (m *ThinkteluControlAPIDataContractsKeyValueDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsKeyValueDetail) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("Key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsKeyValueDetail) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts key value detail based on context it is used
func (m *ThinkteluControlAPIDataContractsKeyValueDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsKeyValueDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsKeyValueDetail) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsKeyValueDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}