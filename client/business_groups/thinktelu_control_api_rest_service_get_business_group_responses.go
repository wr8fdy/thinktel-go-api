// Code generated by go-swagger; DO NOT EDIT.

package business_groups

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

// ThinkteluControlAPIRestServiceGetBusinessGroupReader is a Reader for the ThinkteluControlAPIRestServiceGetBusinessGroup structure.
type ThinkteluControlAPIRestServiceGetBusinessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceGetBusinessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceGetBusinessGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceGetBusinessGroupOK creates a ThinktelUControlAPIRestServiceGetBusinessGroupOK with default headers values
func NewThinktelUControlAPIRestServiceGetBusinessGroupOK() *ThinktelUControlAPIRestServiceGetBusinessGroupOK {
	return &ThinktelUControlAPIRestServiceGetBusinessGroupOK{}
}

/*
ThinktelUControlAPIRestServiceGetBusinessGroupOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceGetBusinessGroupOK struct {
	Payload *models.ThinkteluControlAPIDataContractsBusinessGroup
}

// IsSuccess returns true when this thinktel u control Api rest service get business group o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service get business group o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service get business group o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service get business group o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service get business group o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service get business group o k response
func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessGroups/{ID}][%d] thinktelUControlApiRestServiceGetBusinessGroupOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessGroups/{ID}][%d] thinktelUControlApiRestServiceGetBusinessGroupOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) GetPayload() *models.ThinkteluControlAPIDataContractsBusinessGroup {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceGetBusinessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsBusinessGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceGetBusinessGroupDefault creates a ThinkteluControlAPIRestServiceGetBusinessGroupDefault with default headers values
func NewThinkteluControlAPIRestServiceGetBusinessGroupDefault(code int) *ThinkteluControlAPIRestServiceGetBusinessGroupDefault {
	return &ThinkteluControlAPIRestServiceGetBusinessGroupDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceGetBusinessGroupDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceGetBusinessGroupDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service get business group default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service get business group default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service get business group default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service get business group default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service get business group default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service get business group default response
func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessGroups/{ID}][%d] Thinktel.uControl.Api.RestService.GetBusinessGroup default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /BusinessGroups/{ID}][%d] Thinktel.uControl.Api.RestService.GetBusinessGroup default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceGetBusinessGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
