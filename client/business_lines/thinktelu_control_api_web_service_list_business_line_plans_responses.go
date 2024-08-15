// Code generated by go-swagger; DO NOT EDIT.

package business_lines

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

// ThinkteluControlAPIWebServiceListBusinessLinePlansReader is a Reader for the ThinkteluControlAPIWebServiceListBusinessLinePlans structure.
type ThinkteluControlAPIWebServiceListBusinessLinePlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceListBusinessLinePlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceListBusinessLinePlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceListBusinessLinePlansOK creates a ThinktelUControlAPIWebServiceListBusinessLinePlansOK with default headers values
func NewThinktelUControlAPIWebServiceListBusinessLinePlansOK() *ThinktelUControlAPIWebServiceListBusinessLinePlansOK {
	return &ThinktelUControlAPIWebServiceListBusinessLinePlansOK{}
}

/*
ThinktelUControlAPIWebServiceListBusinessLinePlansOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceListBusinessLinePlansOK struct {
	Payload []*models.ThinkteluControlAPIDataContractsBusinessLinePlan
}

// IsSuccess returns true when this thinktel u control Api web service list business line plans o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service list business line plans o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service list business line plans o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service list business line plans o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service list business line plans o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service list business line plans o k response
func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans][%d] thinktelUControlApiWebServiceListBusinessLinePlansOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans][%d] thinktelUControlApiWebServiceListBusinessLinePlansOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) GetPayload() []*models.ThinkteluControlAPIDataContractsBusinessLinePlan {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceListBusinessLinePlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceListBusinessLinePlansDefault creates a ThinkteluControlAPIWebServiceListBusinessLinePlansDefault with default headers values
func NewThinkteluControlAPIWebServiceListBusinessLinePlansDefault(code int) *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault {
	return &ThinkteluControlAPIWebServiceListBusinessLinePlansDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceListBusinessLinePlansDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceListBusinessLinePlansDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service list business line plans default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service list business line plans default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service list business line plans default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service list business line plans default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service list business line plans default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service list business line plans default response
func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans][%d] Thinktel.uControl.Api.WebService.ListBusinessLinePlans default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans][%d] Thinktel.uControl.Api.WebService.ListBusinessLinePlans default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceListBusinessLinePlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
