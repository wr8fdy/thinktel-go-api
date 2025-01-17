// Code generated by go-swagger; DO NOT EDIT.

package sip_trunks

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

// ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDReader is a Reader for the ThinkteluControlAPIRestServiceUpdateSIPTrunkDID structure.
type ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK creates a ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK with default headers values
func NewThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK() *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK {
	return &ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK{}
}

/*
ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api rest service update Sip trunk Did o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service update Sip trunk Did o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service update Sip trunk Did o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service update Sip trunk Did o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service update Sip trunk Did o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service update Sip trunk Did o k response
func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunks/{number}/Dids/{did}][%d] thinktelUControlApiRestServiceUpdateSipTrunkDidOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunks/{number}/Dids/{did}][%d] thinktelUControlApiRestServiceUpdateSipTrunkDidOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceUpdateSIPTrunkDIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault creates a ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault with default headers values
func NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault(code int) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault {
	return &ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service update Sip trunk Did default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service update Sip trunk Did default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service update Sip trunk Did default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service update Sip trunk Did default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service update Sip trunk Did default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service update Sip trunk Did default response
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunks/{number}/Dids/{did}][%d] Thinktel.uControl.Api.RestService.UpdateSipTrunkDid default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /SipTrunks/{number}/Dids/{did}][%d] Thinktel.uControl.Api.RestService.UpdateSipTrunkDid default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
