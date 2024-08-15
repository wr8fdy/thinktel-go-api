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

// ThinkteluControlAPIRestServiceGetSIPTrunkReader is a Reader for the ThinkteluControlAPIRestServiceGetSIPTrunk structure.
type ThinkteluControlAPIRestServiceGetSIPTrunkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceGetSIPTrunkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceGetSIPTrunkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceGetSIPTrunkOK creates a ThinktelUControlAPIRestServiceGetSIPTrunkOK with default headers values
func NewThinktelUControlAPIRestServiceGetSIPTrunkOK() *ThinktelUControlAPIRestServiceGetSIPTrunkOK {
	return &ThinktelUControlAPIRestServiceGetSIPTrunkOK{}
}

/*
ThinktelUControlAPIRestServiceGetSIPTrunkOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceGetSIPTrunkOK struct {
	Payload *models.ThinkteluControlAPIDataContractsSIPTrunk
}

// IsSuccess returns true when this thinktel u control Api rest service get Sip trunk o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service get Sip trunk o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service get Sip trunk o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service get Sip trunk o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service get Sip trunk o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service get Sip trunk o k response
func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipTrunks/{number}][%d] thinktelUControlApiRestServiceGetSipTrunkOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipTrunks/{number}][%d] thinktelUControlApiRestServiceGetSipTrunkOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) GetPayload() *models.ThinkteluControlAPIDataContractsSIPTrunk {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceGetSIPTrunkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsSIPTrunk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkDefault creates a ThinkteluControlAPIRestServiceGetSIPTrunkDefault with default headers values
func NewThinkteluControlAPIRestServiceGetSIPTrunkDefault(code int) *ThinkteluControlAPIRestServiceGetSIPTrunkDefault {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceGetSIPTrunkDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceGetSIPTrunkDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service get Sip trunk default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service get Sip trunk default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service get Sip trunk default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service get Sip trunk default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service get Sip trunk default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service get Sip trunk default response
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipTrunks/{number}][%d] Thinktel.uControl.Api.RestService.GetSipTrunk default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipTrunks/{number}][%d] Thinktel.uControl.Api.RestService.GetSipTrunk default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceGetSIPTrunkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
