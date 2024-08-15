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

// ThinkteluControlAPIRestServiceGetSIPBindingReader is a Reader for the ThinkteluControlAPIRestServiceGetSIPBinding structure.
type ThinkteluControlAPIRestServiceGetSIPBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceGetSIPBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceGetSIPBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceGetSIPBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceGetSIPBindingOK creates a ThinktelUControlAPIRestServiceGetSIPBindingOK with default headers values
func NewThinktelUControlAPIRestServiceGetSIPBindingOK() *ThinktelUControlAPIRestServiceGetSIPBindingOK {
	return &ThinktelUControlAPIRestServiceGetSIPBindingOK{}
}

/*
ThinktelUControlAPIRestServiceGetSIPBindingOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceGetSIPBindingOK struct {
	Payload *models.ThinkteluControlAPIDataContractsSIPBinding
}

// IsSuccess returns true when this thinktel u control Api rest service get Sip binding o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service get Sip binding o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service get Sip binding o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service get Sip binding o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service get Sip binding o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service get Sip binding o k response
func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings/{ID}][%d] thinktelUControlApiRestServiceGetSipBindingOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings/{ID}][%d] thinktelUControlApiRestServiceGetSipBindingOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) GetPayload() *models.ThinkteluControlAPIDataContractsSIPBinding {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceGetSIPBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsSIPBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceGetSIPBindingDefault creates a ThinkteluControlAPIRestServiceGetSIPBindingDefault with default headers values
func NewThinkteluControlAPIRestServiceGetSIPBindingDefault(code int) *ThinkteluControlAPIRestServiceGetSIPBindingDefault {
	return &ThinkteluControlAPIRestServiceGetSIPBindingDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceGetSIPBindingDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceGetSIPBindingDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service get Sip binding default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service get Sip binding default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service get Sip binding default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service get Sip binding default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service get Sip binding default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service get Sip binding default response
func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings/{ID}][%d] Thinktel.uControl.Api.RestService.GetSipBinding default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /SipBindings/{ID}][%d] Thinktel.uControl.Api.RestService.GetSipBinding default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceGetSIPBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
