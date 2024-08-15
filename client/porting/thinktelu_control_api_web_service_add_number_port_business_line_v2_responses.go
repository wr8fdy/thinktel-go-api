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

// ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Reader is a Reader for the ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2 structure.
type ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK creates a ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK with default headers values
func NewThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK() *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK {
	return &ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK{}
}

/*
ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK struct {
	Payload *models.ThinkteluControlAPIDataContractsTicketResponse
}

// IsSuccess returns true when this thinktel u control Api web service add number port business line v2 o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add number port business line v2 o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add number port business line v2 o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add number port business line v2 o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add number port business line v2 o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add number port business line v2 o k response
func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/BusinessLines][%d] thinktelUControlApiWebServiceAddNumberPortBusinessLineV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/BusinessLines][%d] thinktelUControlApiWebServiceAddNumberPortBusinessLineV2OK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) GetPayload() *models.ThinkteluControlAPIDataContractsTicketResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortBusinessLineV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default creates a ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default with default headers values
func NewThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default(code int) *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default {
	return &ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add number port business line v2 default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add number port business line v2 default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add number port business line v2 default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add number port business line v2 default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add number port business line v2 default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add number port business line v2 default response
func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/BusinessLines][%d] Thinktel.uControl.Api.WebService.AddNumberPortBusinessLineV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v2/NumberPorts/Inbound/BusinessLines][%d] Thinktel.uControl.Api.WebService.AddNumberPortBusinessLineV2 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortBusinessLineV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
