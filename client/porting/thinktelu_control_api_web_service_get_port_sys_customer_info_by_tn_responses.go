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

// ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnReader is a Reader for the ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTn structure.
type ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK creates a ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK with default headers values
func NewThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK() *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK {
	return &ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK{}
}

/*
ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK struct {
	Payload *models.ThinkteluControlAPIDataContractsCustomerInfo
}

// IsSuccess returns true when this thinktel u control Api web service get port sys customer info by tn o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service get port sys customer info by tn o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service get port sys customer info by tn o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service get port sys customer info by tn o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service get port sys customer info by tn o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service get port sys customer info by tn o k response
func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/NumberPorts/Inbound/Dashboard/PortInInfo/OrdersByNumber/{tn}][%d] thinktelUControlApiWebServiceGetPortSysCustomerInfoByTnOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/NumberPorts/Inbound/Dashboard/PortInInfo/OrdersByNumber/{tn}][%d] thinktelUControlApiWebServiceGetPortSysCustomerInfoByTnOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) GetPayload() *models.ThinkteluControlAPIDataContractsCustomerInfo {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceGetPortSysCustomerInfoByTnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsCustomerInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault creates a ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault with default headers values
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault(code int) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service get port sys customer info by tn default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service get port sys customer info by tn default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service get port sys customer info by tn default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service get port sys customer info by tn default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service get port sys customer info by tn default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service get port sys customer info by tn default response
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/NumberPorts/Inbound/Dashboard/PortInInfo/OrdersByNumber/{tn}][%d] Thinktel.uControl.Api.WebService.GetPortSysCustomerInfoByTn default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/NumberPorts/Inbound/Dashboard/PortInInfo/OrdersByNumber/{tn}][%d] Thinktel.uControl.Api.WebService.GetPortSysCustomerInfoByTn default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
