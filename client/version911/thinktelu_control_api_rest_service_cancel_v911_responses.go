// Code generated by go-swagger; DO NOT EDIT.

package version911

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

// ThinkteluControlAPIRestServiceCancelV911Reader is a Reader for the ThinkteluControlAPIRestServiceCancelV911 structure.
type ThinkteluControlAPIRestServiceCancelV911Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceCancelV911Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceCancelV911OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceCancelV911Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceCancelV911OK creates a ThinktelUControlAPIRestServiceCancelV911OK with default headers values
func NewThinktelUControlAPIRestServiceCancelV911OK() *ThinktelUControlAPIRestServiceCancelV911OK {
	return &ThinktelUControlAPIRestServiceCancelV911OK{}
}

/*
ThinktelUControlAPIRestServiceCancelV911OK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceCancelV911OK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api rest service cancel v911 o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceCancelV911OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service cancel v911 o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceCancelV911OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service cancel v911 o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceCancelV911OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service cancel v911 o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceCancelV911OK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service cancel v911 o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceCancelV911OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service cancel v911 o k response
func (o *ThinktelUControlAPIRestServiceCancelV911OK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceCancelV911OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /V911s/{number}][%d] thinktelUControlApiRestServiceCancelV911OK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceCancelV911OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /V911s/{number}][%d] thinktelUControlApiRestServiceCancelV911OK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceCancelV911OK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceCancelV911OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceCancelV911Default creates a ThinkteluControlAPIRestServiceCancelV911Default with default headers values
func NewThinkteluControlAPIRestServiceCancelV911Default(code int) *ThinkteluControlAPIRestServiceCancelV911Default {
	return &ThinkteluControlAPIRestServiceCancelV911Default{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceCancelV911Default describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceCancelV911Default struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service cancel v911 default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceCancelV911Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service cancel v911 default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceCancelV911Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service cancel v911 default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceCancelV911Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service cancel v911 default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceCancelV911Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service cancel v911 default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceCancelV911Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service cancel v911 default response
func (o *ThinkteluControlAPIRestServiceCancelV911Default) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceCancelV911Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /V911s/{number}][%d] Thinktel.uControl.Api.RestService.CancelV911 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceCancelV911Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /V911s/{number}][%d] Thinktel.uControl.Api.RestService.CancelV911 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceCancelV911Default) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceCancelV911Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
