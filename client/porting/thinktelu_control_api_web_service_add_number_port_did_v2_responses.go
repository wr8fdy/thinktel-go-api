// Code generated by go-swagger; DO NOT EDIT.

package porting

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

// ThinkteluControlAPIWebServiceAddNumberPortDIDV2Reader is a Reader for the ThinkteluControlAPIWebServiceAddNumberPortDIDV2 structure.
type ThinkteluControlAPIWebServiceAddNumberPortDIDV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddNumberPortDIDV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddNumberPortDIDV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddNumberPortDIDV2OK creates a ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK with default headers values
func NewThinktelUControlAPIWebServiceAddNumberPortDIDV2OK() *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK {
	return &ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK{}
}

/*
ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK struct {
	Payload *models.ThinkteluControlAPIDataContractsTicketResponse
}

// IsSuccess returns true when this thinktel u control Api web service add number port Did v2 o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add number port Did v2 o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add number port Did v2 o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add number port Did v2 o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add number port Did v2 o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add number port Did v2 o k response
func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/DIDs][%d] thinktelUControlApiWebServiceAddNumberPortDidV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/DIDs][%d] thinktelUControlApiWebServiceAddNumberPortDidV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) GetPayload() *models.ThinkteluControlAPIDataContractsTicketResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortDIDV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddNumberPortDIDV2Default creates a ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default with default headers values
func NewThinkteluControlAPIWebServiceAddNumberPortDIDV2Default(code int) *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default {
	return &ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add number port DID v2 default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add number port DID v2 default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add number port DID v2 default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add number port DID v2 default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add number port DID v2 default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add number port DID v2 default response
func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/DIDs][%d] Thinktel.uControl.Api.WebService.AddNumberPortDIDV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/DIDs][%d] Thinktel.uControl.Api.WebService.AddNumberPortDIDV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortDIDV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
