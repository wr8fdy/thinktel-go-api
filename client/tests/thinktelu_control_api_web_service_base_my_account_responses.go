// Code generated by go-swagger; DO NOT EDIT.

package tests

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

// ThinkteluControlAPIWebServiceBaseMyAccountReader is a Reader for the ThinkteluControlAPIWebServiceBaseMyAccount structure.
type ThinkteluControlAPIWebServiceBaseMyAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceBaseMyAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceBaseMyAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceBaseMyAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceBaseMyAccountOK creates a ThinktelUControlAPIWebServiceBaseMyAccountOK with default headers values
func NewThinktelUControlAPIWebServiceBaseMyAccountOK() *ThinktelUControlAPIWebServiceBaseMyAccountOK {
	return &ThinktelUControlAPIWebServiceBaseMyAccountOK{}
}

/*
ThinktelUControlAPIWebServiceBaseMyAccountOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceBaseMyAccountOK struct {
	Payload int64
}

// IsSuccess returns true when this thinktel u control Api web service base my account o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service base my account o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service base my account o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service base my account o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service base my account o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service base my account o k response
func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Users/Current/Account][%d] thinktelUControlApiWebServiceBaseMyAccountOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Users/Current/Account][%d] thinktelUControlApiWebServiceBaseMyAccountOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) GetPayload() int64 {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceBaseMyAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceBaseMyAccountDefault creates a ThinkteluControlAPIWebServiceBaseMyAccountDefault with default headers values
func NewThinkteluControlAPIWebServiceBaseMyAccountDefault(code int) *ThinkteluControlAPIWebServiceBaseMyAccountDefault {
	return &ThinkteluControlAPIWebServiceBaseMyAccountDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceBaseMyAccountDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceBaseMyAccountDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service base my account default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service base my account default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service base my account default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service base my account default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service base my account default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service base my account default response
func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Users/Current/Account][%d] Thinktel.uControl.Api.WebServiceBase.MyAccount default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Users/Current/Account][%d] Thinktel.uControl.Api.WebServiceBase.MyAccount default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceBaseMyAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
