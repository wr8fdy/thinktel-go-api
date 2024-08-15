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

// ThinkteluControlAPIDataContractsSubscription thinktel u control Api data contracts subscription
//
// swagger:model Thinktel.uControl.Api.DataContracts.Subscription
type ThinkteluControlAPIDataContractsSubscription struct {

	// Subscription specific name. Often the billing telephone number, or other relevant identifier for non-voice related subscriptions.
	// Required: true
	Name *string `json:"Name"`

	// product name
	// Required: true
	ProductName *ThinkteluControlAPIDataContractsLocalizedString `json:"ProductName"`

	// Unique identifier for the subscription.
	// Required: true
	SubscriptionID *string `json:"SubscriptionID"`
}

// Validate validates this thinktel u control Api data contracts subscription
func (m *ThinkteluControlAPIDataContractsSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSubscription) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSubscription) validateProductName(formats strfmt.Registry) error {

	if err := validate.Required("ProductName", "body", m.ProductName); err != nil {
		return err
	}

	if m.ProductName != nil {
		if err := m.ProductName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductName")
			}
			return err
		}
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsSubscription) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("SubscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts subscription based on the context it is used
func (m *ThinkteluControlAPIDataContractsSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProductName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsSubscription) contextValidateProductName(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductName != nil {

		if err := m.ProductName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsSubscription) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}