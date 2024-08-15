// Code generated by go-swagger; DO NOT EDIT.

package sip_bindings

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

// ThinkteluControlAPIWebServiceListSIPBindingsReader is a Reader for the ThinkteluControlAPIWebServiceListSIPBindings structure.
type ThinkteluControlAPIWebServiceListSIPBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceListSIPBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceListSIPBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceListSIPBindingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceListSIPBindingsOK creates a ThinktelUControlAPIWebServiceListSIPBindingsOK with default headers values
func NewThinktelUControlAPIWebServiceListSIPBindingsOK() *ThinktelUControlAPIWebServiceListSIPBindingsOK {
	return &ThinktelUControlAPIWebServiceListSIPBindingsOK{}
}

/*
ThinktelUControlAPIWebServiceListSIPBindingsOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceListSIPBindingsOK struct {
	Payload []*models.ThinkteluControlAPIDataContractsSIPBinding
}

// IsSuccess returns true when this thinktel u control Api web service list Sip bindings o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service list Sip bindings o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service list Sip bindings o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service list Sip bindings o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service list Sip bindings o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service list Sip bindings o k response
func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings][%d] thinktelUControlApiWebServiceListSipBindingsOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings][%d] thinktelUControlApiWebServiceListSipBindingsOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) GetPayload() []*models.ThinkteluControlAPIDataContractsSIPBinding {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceListSIPBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceListSIPBindingsDefault creates a ThinkteluControlAPIWebServiceListSIPBindingsDefault with default headers values
func NewThinkteluControlAPIWebServiceListSIPBindingsDefault(code int) *ThinkteluControlAPIWebServiceListSIPBindingsDefault {
	return &ThinkteluControlAPIWebServiceListSIPBindingsDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceListSIPBindingsDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceListSIPBindingsDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service list Sip bindings default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service list Sip bindings default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service list Sip bindings default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service list Sip bindings default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service list Sip bindings default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service list Sip bindings default response
func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings][%d] Thinktel.uControl.Api.WebService.ListSipBindings default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings][%d] Thinktel.uControl.Api.WebService.ListSipBindings default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceListSIPBindingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}