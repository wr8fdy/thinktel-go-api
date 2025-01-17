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

// ThinkteluControlAPIDataContractsCustInfoRequests thinktel u control Api data contracts cust info requests
//
// swagger:model Thinktel.uControl.Api.DataContracts.CustInfoRequests
type ThinkteluControlAPIDataContractsCustInfoRequests struct {

	// Description of the status of the TNs to be ported. Possible values are: Pending, Click-To-Port, Internal, Closed
	// Max Length: 20
	Assessment string `json:"Assessment,omitempty"`

	// The Number of the ticket in HappyFox associated to the Porting Order.
	// Max Length: 8
	// Min Length: 8
	HappyFoxTicket string `json:"HappyFoxTicket,omitempty"`

	// New Service associated to be created by the Porting Order.
	// Max Length: 60
	// Min Length: 60
	NewService string `json:"NewService,omitempty"`

	// Array of the Numbers to port in the order.
	NumberToPort []*ThinkteluControlAPIDataContractsNumberToPort `json:"NumberToPort"`

	// The parent service TN number of the new service being ordered.
	// Max Length: 10
	// Min Length: 10
	ParentService string `json:"ParentService,omitempty"`

	// The ID of the request in the Requests Array.
	// Required: true
	RequestIdx *int32 `json:"RequestIdx"`
}

// Validate validates this thinktel u control Api data contracts cust info requests
func (m *ThinkteluControlAPIDataContractsCustInfoRequests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHappyFoxTicket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberToPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestIdx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateAssessment(formats strfmt.Registry) error {
	if swag.IsZero(m.Assessment) { // not required
		return nil
	}

	if err := validate.MaxLength("Assessment", "body", m.Assessment, 20); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateHappyFoxTicket(formats strfmt.Registry) error {
	if swag.IsZero(m.HappyFoxTicket) { // not required
		return nil
	}

	if err := validate.MinLength("HappyFoxTicket", "body", m.HappyFoxTicket, 8); err != nil {
		return err
	}

	if err := validate.MaxLength("HappyFoxTicket", "body", m.HappyFoxTicket, 8); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateNewService(formats strfmt.Registry) error {
	if swag.IsZero(m.NewService) { // not required
		return nil
	}

	if err := validate.MinLength("NewService", "body", m.NewService, 60); err != nil {
		return err
	}

	if err := validate.MaxLength("NewService", "body", m.NewService, 60); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateNumberToPort(formats strfmt.Registry) error {
	if swag.IsZero(m.NumberToPort) { // not required
		return nil
	}

	for i := 0; i < len(m.NumberToPort); i++ {
		if swag.IsZero(m.NumberToPort[i]) { // not required
			continue
		}

		if m.NumberToPort[i] != nil {
			if err := m.NumberToPort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NumberToPort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NumberToPort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateParentService(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentService) { // not required
		return nil
	}

	if err := validate.MinLength("ParentService", "body", m.ParentService, 10); err != nil {
		return err
	}

	if err := validate.MaxLength("ParentService", "body", m.ParentService, 10); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) validateRequestIdx(formats strfmt.Registry) error {

	if err := validate.Required("RequestIdx", "body", m.RequestIdx); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts cust info requests based on the context it is used
func (m *ThinkteluControlAPIDataContractsCustInfoRequests) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNumberToPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsCustInfoRequests) contextValidateNumberToPort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NumberToPort); i++ {

		if m.NumberToPort[i] != nil {

			if swag.IsZero(m.NumberToPort[i]) { // not required
				return nil
			}

			if err := m.NumberToPort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NumberToPort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NumberToPort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsCustInfoRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsCustInfoRequests) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsCustInfoRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
