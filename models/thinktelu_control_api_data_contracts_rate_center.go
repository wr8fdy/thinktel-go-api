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

// ThinkteluControlAPIDataContractsRateCenter thinktel u control Api data contracts rate center
//
// swagger:model Thinktel.uControl.Api.DataContracts.RateCenter
type ThinkteluControlAPIDataContractsRateCenter struct {

	// available
	Available uint32 `json:"Available,omitempty"`

	// country
	Country string `json:"Country,omitempty"`

	// local
	Local bool `json:"Local,omitempty"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// on net
	OnNet bool `json:"OnNet,omitempty"`
}

// Validate validates this thinktel u control Api data contracts rate center
func (m *ThinkteluControlAPIDataContractsRateCenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsRateCenter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts rate center based on context it is used
func (m *ThinkteluControlAPIDataContractsRateCenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsRateCenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsRateCenter) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsRateCenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
