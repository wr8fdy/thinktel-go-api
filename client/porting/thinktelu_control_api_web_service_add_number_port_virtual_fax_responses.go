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

// ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxReader is a Reader for the ThinkteluControlAPIWebServiceAddNumberPortVirtualFax structure.
type ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK creates a ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK with default headers values
func NewThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK() *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK {
	return &ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK{}
}

/*
ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK struct {
	Payload *models.ThinkteluControlAPIDataContractsTicketResponse
}

// IsSuccess returns true when this thinktel u control Api web service add number port virtual fax o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add number port virtual fax o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add number port virtual fax o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add number port virtual fax o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add number port virtual fax o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add number port virtual fax o k response
func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/VirtualFaxes][%d] thinktelUControlApiWebServiceAddNumberPortVirtualFaxOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/VirtualFaxes][%d] thinktelUControlApiWebServiceAddNumberPortVirtualFaxOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) GetPayload() *models.ThinkteluControlAPIDataContractsTicketResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortVirtualFaxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault creates a ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault with default headers values
func NewThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault(code int) *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault {
	return &ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add number port virtual fax default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add number port virtual fax default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add number port virtual fax default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add number port virtual fax default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add number port virtual fax default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add number port virtual fax default response
func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/VirtualFaxes][%d] Thinktel.uControl.Api.WebService.AddNumberPortVirtualFax default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/VirtualFaxes][%d] Thinktel.uControl.Api.WebService.AddNumberPortVirtualFax default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortVirtualFaxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
