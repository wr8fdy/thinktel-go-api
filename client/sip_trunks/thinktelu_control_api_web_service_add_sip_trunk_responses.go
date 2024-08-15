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

// ThinkteluControlAPIWebServiceAddSIPTrunkReader is a Reader for the ThinkteluControlAPIWebServiceAddSIPTrunk structure.
type ThinkteluControlAPIWebServiceAddSIPTrunkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddSIPTrunkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddSIPTrunkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddSIPTrunkOK creates a ThinktelUControlAPIWebServiceAddSIPTrunkOK with default headers values
func NewThinktelUControlAPIWebServiceAddSIPTrunkOK() *ThinktelUControlAPIWebServiceAddSIPTrunkOK {
	return &ThinktelUControlAPIWebServiceAddSIPTrunkOK{}
}

/*
ThinktelUControlAPIWebServiceAddSIPTrunkOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddSIPTrunkOK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api web service add Sip trunk o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add Sip trunk o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add Sip trunk o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add Sip trunk o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add Sip trunk o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add Sip trunk o k response
func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SipTrunks][%d] thinktelUControlApiWebServiceAddSipTrunkOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SipTrunks][%d] thinktelUControlApiWebServiceAddSipTrunkOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddSIPTrunkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddSIPTrunkDefault creates a ThinkteluControlAPIWebServiceAddSIPTrunkDefault with default headers values
func NewThinkteluControlAPIWebServiceAddSIPTrunkDefault(code int) *ThinkteluControlAPIWebServiceAddSIPTrunkDefault {
	return &ThinkteluControlAPIWebServiceAddSIPTrunkDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddSIPTrunkDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddSIPTrunkDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add Sip trunk default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add Sip trunk default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add Sip trunk default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add Sip trunk default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add Sip trunk default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add Sip trunk default response
func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SipTrunks][%d] Thinktel.uControl.Api.WebService.AddSipTrunk default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SipTrunks][%d] Thinktel.uControl.Api.WebService.AddSipTrunk default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddSIPTrunkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
