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

// ThinkteluControlAPIRestServiceGetBusinessLinePlanReader is a Reader for the ThinkteluControlAPIRestServiceGetBusinessLinePlan structure.
type ThinkteluControlAPIRestServiceGetBusinessLinePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceGetBusinessLinePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceGetBusinessLinePlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceGetBusinessLinePlanOK creates a ThinktelUControlAPIRestServiceGetBusinessLinePlanOK with default headers values
func NewThinktelUControlAPIRestServiceGetBusinessLinePlanOK() *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK {
	return &ThinktelUControlAPIRestServiceGetBusinessLinePlanOK{}
}

/*
ThinktelUControlAPIRestServiceGetBusinessLinePlanOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceGetBusinessLinePlanOK struct {
	Payload *models.ThinkteluControlAPIDataContractsBusinessLinePlan
}

// IsSuccess returns true when this thinktel u control Api rest service get business line plan o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service get business line plan o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service get business line plan o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service get business line plan o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service get business line plan o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service get business line plan o k response
func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans/{ID}][%d] thinktelUControlApiRestServiceGetBusinessLinePlanOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans/{ID}][%d] thinktelUControlApiRestServiceGetBusinessLinePlanOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) GetPayload() *models.ThinkteluControlAPIDataContractsBusinessLinePlan {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceGetBusinessLinePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsBusinessLinePlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceGetBusinessLinePlanDefault creates a ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault with default headers values
func NewThinkteluControlAPIRestServiceGetBusinessLinePlanDefault(code int) *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault {
	return &ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service get business line plan default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service get business line plan default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service get business line plan default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service get business line plan default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service get business line plan default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service get business line plan default response
func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans/{ID}][%d] Thinktel.uControl.Api.RestService.GetBusinessLinePlan default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessLinePlans/{ID}][%d] Thinktel.uControl.Api.RestService.GetBusinessLinePlan default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceGetBusinessLinePlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
