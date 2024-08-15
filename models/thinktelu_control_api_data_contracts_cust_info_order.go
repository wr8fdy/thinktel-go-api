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

// ThinkteluControlAPIDataContractsCustInfoOrder thinktel u control Api data contracts cust info order
//
// swagger:model Thinktel.uControl.Api.DataContracts.CustInfoOrder
type ThinkteluControlAPIDataContractsCustInfoOrder struct {

	// The ID of the Orders in the Orders Array.
	// Required: true
	OrderID *int32 `json:"OrderID"`

	// Array of the requests in the order.
	// Required: true
	Requests []*ThinkteluControlAPIDataContractsCustInfoRequests `json:"Requests"`
}

// Validate validates this thinktel u control Api data contracts cust info order
func (m *ThinkteluControlAPIDataContractsCustInfoOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoOrder) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("OrderID", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoOrder) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("Requests", "body", m.Requests); err != nil {
		return err
	}

	for i := 0; i < len(m.Requests); i++ {
		if swag.IsZero(m.Requests[i]) { // not required
			continue
		}

		if m.Requests[i] != nil {
			if err := m.Requests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts cust info order based on the context it is used
func (m *ThinkteluControlAPIDataContractsCustInfoOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoOrder) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Requests); i++ {

		if m.Requests[i] != nil {

			if swag.IsZero(m.Requests[i]) { // not required
				return nil
			}

			if err := m.Requests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsCustInfoOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsCustInfoOrder) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsCustInfoOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
