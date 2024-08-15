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

// ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures thinktel u control Api data contracts subscriber Sip trunk features
//
// swagger:model Thinktel.uControl.Api.DataContracts.SubscriberSipTrunkFeatures
type ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures struct {

	// external911 notification email
	// Required: true
	External911NotificationEmail *string `json:"External911NotificationEmail"`
}

// Validate validates this thinktel u control Api data contracts subscriber Sip trunk features
func (m *ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternal911NotificationEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures) validateExternal911NotificationEmail(formats strfmt.Registry) error {

	if err := validate.Required("External911NotificationEmail", "body", m.External911NotificationEmail); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this thinktel u control Api data contracts subscriber Sip trunk features based on context it is used
func (m *ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsSubscriberSIPTrunkFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
