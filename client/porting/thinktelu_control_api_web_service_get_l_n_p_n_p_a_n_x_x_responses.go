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

// ThinkteluControlAPIWebServiceGetLNPNPANXXReader is a Reader for the ThinkteluControlAPIWebServiceGetLNPNPANXX structure.
type ThinkteluControlAPIWebServiceGetLNPNPANXXReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceGetLNPNPANXXOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceGetLNPNPANXXDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceGetLNPNPANXXOK creates a ThinktelUControlAPIWebServiceGetLNPNPANXXOK with default headers values
func NewThinktelUControlAPIWebServiceGetLNPNPANXXOK() *ThinktelUControlAPIWebServiceGetLNPNPANXXOK {
	return &ThinktelUControlAPIWebServiceGetLNPNPANXXOK{}
}

/*
ThinktelUControlAPIWebServiceGetLNPNPANXXOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceGetLNPNPANXXOK struct {
	Payload []int64
}

// IsSuccess returns true when this thinktel u control Api web service get l n p n p a n x x o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service get l n p n p a n x x o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service get l n p n p a n x x o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service get l n p n p a n x x o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service get l n p n p a n x x o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service get l n p n p a n x x o k response
func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /NumberPorts/LNP/NPANXX][%d] thinktelUControlApiWebServiceGetLNPNPANXXOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /NumberPorts/LNP/NPANXX][%d] thinktelUControlApiWebServiceGetLNPNPANXXOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) GetPayload() []int64 {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceGetLNPNPANXXOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceGetLNPNPANXXDefault creates a ThinkteluControlAPIWebServiceGetLNPNPANXXDefault with default headers values
func NewThinkteluControlAPIWebServiceGetLNPNPANXXDefault(code int) *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault {
	return &ThinkteluControlAPIWebServiceGetLNPNPANXXDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceGetLNPNPANXXDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceGetLNPNPANXXDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service get l n p n p a n x x default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service get l n p n p a n x x default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service get l n p n p a n x x default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service get l n p n p a n x x default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service get l n p n p a n x x default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service get l n p n p a n x x default response
func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /NumberPorts/LNP/NPANXX][%d] Thinktel.uControl.Api.WebService.GetLNPNPANXX default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /NumberPorts/LNP/NPANXX][%d] Thinktel.uControl.Api.WebService.GetLNPNPANXX default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceGetLNPNPANXXDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
