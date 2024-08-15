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

// ThinkteluControlAPIRestServiceUpdateV911Reader is a Reader for the ThinkteluControlAPIRestServiceUpdateV911 structure.
type ThinkteluControlAPIRestServiceUpdateV911Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceUpdateV911Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceUpdateV911OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceUpdateV911Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceUpdateV911OK creates a ThinktelUControlAPIRestServiceUpdateV911OK with default headers values
func NewThinktelUControlAPIRestServiceUpdateV911OK() *ThinktelUControlAPIRestServiceUpdateV911OK {
	return &ThinktelUControlAPIRestServiceUpdateV911OK{}
}

/*
ThinktelUControlAPIRestServiceUpdateV911OK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceUpdateV911OK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api rest service update v911 o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service update v911 o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service update v911 o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service update v911 o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service update v911 o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service update v911 o k response
func (o *ThinktelUControlAPIRestServiceUpdateV911OK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceUpdateV911OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /V911s/{number}][%d] thinktelUControlApiRestServiceUpdateV911OK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateV911OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /V911s/{number}][%d] thinktelUControlApiRestServiceUpdateV911OK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceUpdateV911OK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceUpdateV911OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceUpdateV911Default creates a ThinkteluControlAPIRestServiceUpdateV911Default with default headers values
func NewThinkteluControlAPIRestServiceUpdateV911Default(code int) *ThinkteluControlAPIRestServiceUpdateV911Default {
	return &ThinkteluControlAPIRestServiceUpdateV911Default{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateV911Default describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceUpdateV911Default struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service update v911 default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service update v911 default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service update v911 default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service update v911 default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service update v911 default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service update v911 default response
func (o *ThinkteluControlAPIRestServiceUpdateV911Default) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceUpdateV911Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /V911s/{number}][%d] Thinktel.uControl.Api.RestService.UpdateV911 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateV911Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /V911s/{number}][%d] Thinktel.uControl.Api.RestService.UpdateV911 default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceUpdateV911Default) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceUpdateV911Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}