// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wr8fdy/thinktel-go-api/models"
)

// ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedReader is a Reader for the ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecated structure.
type ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK creates a ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK with default headers values
func NewThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK() *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK {
	return &ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK{}
}

/*
ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api rest service update Did premium features deprecated o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service update Did premium features deprecated o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service update Did premium features deprecated o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service update Did premium features deprecated o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service update Did premium features deprecated o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service update Did premium features deprecated o k response
func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunk/{pilot}/Did/{number}/PremiumFeatures][%d] thinktelUControlApiRestServiceUpdateDidPremiumFeaturesDeprecatedOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunk/{pilot}/Did/{number}/PremiumFeatures][%d] thinktelUControlApiRestServiceUpdateDidPremiumFeaturesDeprecatedOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault creates a ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault with default headers values
func NewThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault(code int) *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault {
	return &ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service update DID premium features deprecated default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service update DID premium features deprecated default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service update DID premium features deprecated default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service update DID premium features deprecated default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service update DID premium features deprecated default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service update DID premium features deprecated default response
func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunk/{pilot}/Did/{number}/PremiumFeatures][%d] Thinktel.uControl.Api.RestService.UpdateDIDPremiumFeatures_Deprecated default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunk/{pilot}/Did/{number}/PremiumFeatures][%d] Thinktel.uControl.Api.RestService.UpdateDIDPremiumFeatures_Deprecated default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceUpdateDIDPremiumFeaturesDeprecatedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
