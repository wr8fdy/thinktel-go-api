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

// ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2 Toll-Free numbers port-in request form. Contains the common elements for all port-in requests.
//
// swagger:model Thinktel.uControl.Api.DataContracts.NumberPortTollFreeOrderFormV2
type ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2 struct {

	// Additional email addresses that should be included in porting communications.
	AdditionalEmailAddresses []string `json:"AdditionalEmailAddresses"`

	// Comments for the losing carrier.
	CommentsLosingCarrier string `json:"CommentsLosingCarrier,omitempty"`

	// Comments for the porting team.
	CommentsPortTeam string `json:"CommentsPortTeam,omitempty"`

	// Submitter wants to be contacted about their 411 listings (BLIFs).
	ContactMeBLIF bool `json:"ContactMeBLIF,omitempty"`

	// Submitter wants to be contacted about custom scheduling.
	ContactMeScheduling bool `json:"ContactMeScheduling,omitempty"`

	// Current provider (aka losing carrier). Applicable only if CurrentReseller is null.
	CurrentProvider int32 `json:"CurrentProvider,omitempty"`

	// Current provider not mandatory if set to TRUE then allows to bypass Reseller determination logic.
	// Example: false
	CurrentProviderNotMandatory bool `json:"CurrentProviderNotMandatory,omitempty"`

	// Current reseller. Applicable only if CurrentProvider is null.
	CurrentReseller string `json:"CurrentReseller,omitempty"`

	// Customer’s account number as per losing carrier’s records.
	// Required: true
	CustomerAccountNumber *string `json:"CustomerAccountNumber"`

	// Customer’s city as per losing carrier’s records.
	// Required: true
	CustomerCity *string `json:"CustomerCity"`

	// Customer’s country as per losing carrier’s records. Only CA or US accepted at this time.
	// Required: true
	CustomerCountry *string `json:"CustomerCountry"`

	// Customer’s address 'Floor' as per losing carrier’s records.
	CustomerFloor string `json:"CustomerFloor,omitempty"`

	// Customer’s name as per losing carrier’s records.
	// Required: true
	CustomerName *string `json:"CustomerName"`

	// Customer’s address postal code as per losing carrier’s records.
	// Required: true
	CustomerPostalCode *string `json:"CustomerPostalCode"`

	// Customer’s province as per losing carrier’s records.
	// Required: true
	CustomerProvince *string `json:"CustomerProvince"`

	// User provided customer reference number for tracking purposes
	CustomerReferenceNumber string `json:"CustomerReferenceNumber,omitempty"`

	// Customer’s address 'Street Direction' as per losing carrier’s records.
	CustomerStreetDirection string `json:"CustomerStreetDirection,omitempty"`

	// Customer’s address 'Street Name' as per losing carrier’s records.
	// Required: true
	CustomerStreetName *string `json:"CustomerStreetName"`

	// Customer’s address 'Street Number' as per losing carrier’s records.
	// Required: true
	CustomerStreetNumber *string `json:"CustomerStreetNumber"`

	// Customer’s address 'Street Suffix' as per losing carrier’s records.
	CustomerStreetSuffix string `json:"CustomerStreetSuffix,omitempty"`

	// Customer’s address 'Street Type' as per losing carrier’s records.
	CustomerStreetType string `json:"CustomerStreetType,omitempty"`

	// Customer’s address 'Unit' as per losing carrier’s records.
	CustomerUnit string `json:"CustomerUnit,omitempty"`

	// Primary email address for porting communications.
	// Required: true
	EmailAddress *string `json:"EmailAddress"`

	// Extra details. The details provided will not be reviewed before the port request is submitted to the losing carrier. These are for details such as special configuration of the new service once the port has completed.
	ExtraDetails []*ThinkteluControlAPIDataContractsKeyValueDetail `json:"ExtraDetails"`

	// 411 listings (BLIFs) form. Can provide multiple documents if needed. Template can be found here: https://www.thinktel.ca/pdf/?id=Template-for-directory-listing
	FileBlifSpreadsheet []*ThinkteluControlAPIDataContractsBase64EncodedFile `json:"FileBlifSpreadsheet"`

	// Latest customer invoice from losing carrier. Can provide multiple documents if needed.
	FileLastInvoice []*ThinkteluControlAPIDataContractsBase64EncodedFile `json:"FileLastInvoice"`

	// Responsible organization transfer form for toll-free telephone numbers. Can provide multiple documents if needed. Template can be found here: https://www.thinktel.ca/pdf/?id=respOrg-form
	FileRespOrgForm []*ThinkteluControlAPIDataContractsBase64EncodedFile `json:"FileRespOrgForm"`

	// Telephone number(s) to port.
	NumbersToPort []int64 `json:"NumbersToPort"`

	// Requested port date for a timed port. Date must be in YYYY-MM-DD format. Unless RequestedPortTime also has a value, the port will be scheduled to occur sometime during work day hours of the Eastern Prevailing Time (EPT) time zone. If no requested port date is provided, then the order will be implicitly scheduled for the earliest possible port date.
	RequestedPortDate string `json:"RequestedPortDate,omitempty"`

	// Requested port time for a timed port. Applicable only if RequestedPortDate has a value. Time must be in HH:MM 24-hour format and Eastern Prevailing Time(EPT) time zone. Additional charges may apply for specifying a port time.
	RequestedPortTime string `json:"RequestedPortTime,omitempty"`

	// Submitter’s first name.
	SubmitterFirstName string `json:"SubmitterFirstName,omitempty"`

	// Submitter’s last name.
	SubmitterLastName string `json:"SubmitterLastName,omitempty"`

	// Submitter’s phone number.
	SubmitterPhoneNumber int64 `json:"SubmitterPhoneNumber,omitempty"`
}

// Validate validates this thinktel u control Api data contracts number port toll free order form v2
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerProvince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerStreetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerStreetNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileBlifSpreadsheet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileLastInvoice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileRespOrgForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("CustomerAccountNumber", "body", m.CustomerAccountNumber); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerCity(formats strfmt.Registry) error {

	if err := validate.Required("CustomerCity", "body", m.CustomerCity); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerCountry(formats strfmt.Registry) error {

	if err := validate.Required("CustomerCountry", "body", m.CustomerCountry); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerName(formats strfmt.Registry) error {

	if err := validate.Required("CustomerName", "body", m.CustomerName); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerPostalCode(formats strfmt.Registry) error {

	if err := validate.Required("CustomerPostalCode", "body", m.CustomerPostalCode); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerProvince(formats strfmt.Registry) error {

	if err := validate.Required("CustomerProvince", "body", m.CustomerProvince); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerStreetName(formats strfmt.Registry) error {

	if err := validate.Required("CustomerStreetName", "body", m.CustomerStreetName); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateCustomerStreetNumber(formats strfmt.Registry) error {

	if err := validate.Required("CustomerStreetNumber", "body", m.CustomerStreetNumber); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("EmailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateExtraDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtraDetails); i++ {
		if swag.IsZero(m.ExtraDetails[i]) { // not required
			continue
		}

		if m.ExtraDetails[i] != nil {
			if err := m.ExtraDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExtraDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExtraDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateFileBlifSpreadsheet(formats strfmt.Registry) error {
	if swag.IsZero(m.FileBlifSpreadsheet) { // not required
		return nil
	}

	for i := 0; i < len(m.FileBlifSpreadsheet); i++ {
		if swag.IsZero(m.FileBlifSpreadsheet[i]) { // not required
			continue
		}

		if m.FileBlifSpreadsheet[i] != nil {
			if err := m.FileBlifSpreadsheet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileBlifSpreadsheet" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileBlifSpreadsheet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateFileLastInvoice(formats strfmt.Registry) error {
	if swag.IsZero(m.FileLastInvoice) { // not required
		return nil
	}

	for i := 0; i < len(m.FileLastInvoice); i++ {
		if swag.IsZero(m.FileLastInvoice[i]) { // not required
			continue
		}

		if m.FileLastInvoice[i] != nil {
			if err := m.FileLastInvoice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileLastInvoice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileLastInvoice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) validateFileRespOrgForm(formats strfmt.Registry) error {
	if swag.IsZero(m.FileRespOrgForm) { // not required
		return nil
	}

	for i := 0; i < len(m.FileRespOrgForm); i++ {
		if swag.IsZero(m.FileRespOrgForm[i]) { // not required
			continue
		}

		if m.FileRespOrgForm[i] != nil {
			if err := m.FileRespOrgForm[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileRespOrgForm" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileRespOrgForm" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this thinktel u control Api data contracts number port toll free order form v2 based on the context it is used
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtraDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileBlifSpreadsheet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileLastInvoice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileRespOrgForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) contextValidateExtraDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtraDetails); i++ {

		if m.ExtraDetails[i] != nil {

			if swag.IsZero(m.ExtraDetails[i]) { // not required
				return nil
			}

			if err := m.ExtraDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExtraDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExtraDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) contextValidateFileBlifSpreadsheet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileBlifSpreadsheet); i++ {

		if m.FileBlifSpreadsheet[i] != nil {

			if swag.IsZero(m.FileBlifSpreadsheet[i]) { // not required
				return nil
			}

			if err := m.FileBlifSpreadsheet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileBlifSpreadsheet" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileBlifSpreadsheet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) contextValidateFileLastInvoice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileLastInvoice); i++ {

		if m.FileLastInvoice[i] != nil {

			if swag.IsZero(m.FileLastInvoice[i]) { // not required
				return nil
			}

			if err := m.FileLastInvoice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileLastInvoice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileLastInvoice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) contextValidateFileRespOrgForm(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileRespOrgForm); i++ {

		if m.FileRespOrgForm[i] != nil {

			if swag.IsZero(m.FileRespOrgForm[i]) { // not required
				return nil
			}

			if err := m.FileRespOrgForm[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FileRespOrgForm" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FileRespOrgForm" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2) UnmarshalBinary(b []byte) error {
	var res ThinkteluControlAPIDataContractsNumberPortTollFreeOrderFormV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
