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

// ThinkteluControlAPIDataContractsProvisionedDevice thinktel u control Api data contracts provisioned device
//
// swagger:model Thinktel.uControl.Api.DataContracts.ProvisionedDevice
type ThinkteluControlAPIDataContractsProvisionedDevice struct {

	// account
	// Required: true
	Account *int64 `json:"Account"`

	// external IP address
	ExternalIPAddress string `json:"ExternalIPAddress,omitempty"`

	// IP address
	IPAddress string `json:"IPAddress,omitempty"`

	// last refresh
	// Format: date-time
	LastRefresh strfmt.DateTime `json:"LastRefresh,omitempty"`

	// last update
	// Format: date-time
	LastUpdate strfmt.DateTime `json:"LastUpdate,omitempty"`

	// lines
	Lines []*ThinkteluControlAPIDataContractsProvisionedLine `json:"Lines"`

	// m a c address
	// Required: true
	MACAddress *string `json:"MACAddress"`

	// model
	Model string `json:"Model,omitempty"`

	// serial number
	SerialNumber string `json:"SerialNumber,omitempty"`

	// software version
	SoftwareVersion string `json:"SoftwareVersion,omitempty"`

	// station name
	StationName string `json:"StationName,omitempty"`
}

// Validate validates this thinktel u control Api data contracts provisioned device
func (m *ThinkteluControlAPIDataContractsProvisionedDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRefresh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMACAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("Account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) validateLastRefresh(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRefresh) { // not required
		return nil
	}

	if err := validate.FormatOf("LastRefresh", "body", "date-time", m.LastRefresh.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) validateLastUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) validateLines(formats strfmt.Registry) error {
	if swag.IsZero(m.Lines) { // not required
		return nil
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) validateMACAddress(formats strfmt.Registry) error {

	if err := validate.Required("MACAddress", "body", m.MACAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts provisioned device based on the context it is used
func (m *ThinkteluControlAPIDataContractsProvisionedDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsProvisionedDevice) contextValidateLines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Lines); i++ {

		if m.Lines[i] != nil {

			if swag.IsZero(m.Lines[i]) { // not required
				return nil
			}

			if err := m.Lines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsProvisionedDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsProvisionedDevice) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsProvisionedDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
