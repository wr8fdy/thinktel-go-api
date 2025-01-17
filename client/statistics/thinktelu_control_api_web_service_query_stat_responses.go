// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// ThinkteluControlAPIWebServiceQueryStatReader is a Reader for the ThinkteluControlAPIWebServiceQueryStat structure.
type ThinkteluControlAPIWebServiceQueryStatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceQueryStatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceQueryStatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceQueryStatDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceQueryStatOK creates a ThinktelUControlAPIWebServiceQueryStatOK with default headers values
func NewThinktelUControlAPIWebServiceQueryStatOK() *ThinktelUControlAPIWebServiceQueryStatOK {
	return &ThinktelUControlAPIWebServiceQueryStatOK{}
}

/*
ThinktelUControlAPIWebServiceQueryStatOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceQueryStatOK struct {
	Payload *models.ThinkteluControlAPIDataContractsStatResponse
}

// IsSuccess returns true when this thinktel u control Api web service query stat o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceQueryStatOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service query stat o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceQueryStatOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service query stat o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceQueryStatOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service query stat o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceQueryStatOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service query stat o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceQueryStatOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service query stat o k response
func (o *ThinktelUControlAPIWebServiceQueryStatOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceQueryStatOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Stats][%d] thinktelUControlApiWebServiceQueryStatOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceQueryStatOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Stats][%d] thinktelUControlApiWebServiceQueryStatOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceQueryStatOK) GetPayload() *models.ThinkteluControlAPIDataContractsStatResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceQueryStatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsStatResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceQueryStatDefault creates a ThinkteluControlAPIWebServiceQueryStatDefault with default headers values
func NewThinkteluControlAPIWebServiceQueryStatDefault(code int) *ThinkteluControlAPIWebServiceQueryStatDefault {
	return &ThinkteluControlAPIWebServiceQueryStatDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceQueryStatDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceQueryStatDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service query stat default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service query stat default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service query stat default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service query stat default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service query stat default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service query stat default response
func (o *ThinkteluControlAPIWebServiceQueryStatDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceQueryStatDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Stats][%d] Thinktel.uControl.Api.WebService.QueryStat default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceQueryStatDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Stats][%d] Thinktel.uControl.Api.WebService.QueryStat default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceQueryStatDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceQueryStatDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
