// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// ThinkteluControlAPIWebServiceAddNumberPortResidentialLineReader is a Reader for the ThinkteluControlAPIWebServiceAddNumberPortResidentialLine structure.
type ThinkteluControlAPIWebServiceAddNumberPortResidentialLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK creates a ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK with default headers values
func NewThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK() *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK {
	return &ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK{}
}

/*
ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK struct {
	Payload *models.ThinkteluControlAPIDataContractsTicketResponse
}

// IsSuccess returns true when this thinktel u control Api web service add number port residential line o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service add number port residential line o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service add number port residential line o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service add number port residential line o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service add number port residential line o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service add number port residential line o k response
func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/ResidentialLines][%d] thinktelUControlApiWebServiceAddNumberPortResidentialLineOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/ResidentialLines][%d] thinktelUControlApiWebServiceAddNumberPortResidentialLineOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) GetPayload() *models.ThinkteluControlAPIDataContractsTicketResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceAddNumberPortResidentialLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault creates a ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault with default headers values
func NewThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault(code int) *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault {
	return &ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service add number port residential line default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service add number port residential line default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service add number port residential line default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service add number port residential line default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service add number port residential line default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service add number port residential line default response
func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/ResidentialLines][%d] Thinktel.uControl.Api.WebService.AddNumberPortResidentialLine default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /NumberPorts/Inbound/ResidentialLines][%d] Thinktel.uControl.Api.WebService.AddNumberPortResidentialLine default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceAddNumberPortResidentialLineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
