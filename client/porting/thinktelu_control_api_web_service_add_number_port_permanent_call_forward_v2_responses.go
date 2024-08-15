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

// ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Reader is a Reader for the ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2 structure.
type ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK creates a ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK with default headers values
func NewThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK() *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK {
	return &ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK{}
}

/*
ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK struct {
	Payload *models.ThinkteluControlAPIDataContractsTicketResponse
}

// IsSuccess returns true when this thinktel u control Api web service add number port permanent call forward v2 o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add number port permanent call forward v2 o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add number port permanent call forward v2 o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add number port permanent call forward v2 o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add number port permanent call forward v2 o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add number port permanent call forward v2 o k response
func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/PermanentCallForwards][%d] thinktelUControlApiWebServiceAddNumberPortPermanentCallForwardV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/PermanentCallForwards][%d] thinktelUControlApiWebServiceAddNumberPortPermanentCallForwardV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) GetPayload() *models.ThinkteluControlAPIDataContractsTicketResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortPermanentCallForwardV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default creates a ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default with default headers values
func NewThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default(code int) *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default {
	return &ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add number port permanent call forward v2 default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add number port permanent call forward v2 default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add number port permanent call forward v2 default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add number port permanent call forward v2 default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add number port permanent call forward v2 default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add number port permanent call forward v2 default response
func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/PermanentCallForwards][%d] Thinktel.uControl.Api.WebService.AddNumberPortPermanentCallForwardV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/PermanentCallForwards][%d] Thinktel.uControl.Api.WebService.AddNumberPortPermanentCallForwardV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortPermanentCallForwardV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
