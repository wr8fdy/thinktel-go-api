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

// ThinkteluControlAPIWebServiceListV911sv2DeprecatedReader is a Reader for the ThinkteluControlAPIWebServiceListV911sv2Deprecated structure.
type ThinkteluControlAPIWebServiceListV911sv2DeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceListV911sv2DeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceListV911sv2DeprecatedOK creates a ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK with default headers values
func NewThinktelUControlAPIWebServiceListV911sv2DeprecatedOK() *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK {
	return &ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK{}
}

/*
ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK struct {
	Payload []*models.ThinkteluControlAPIDataContractsV911v2
}

// IsSuccess returns true when this thinktel u control Api web service list v911sv2 deprecated o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service list v911sv2 deprecated o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service list v911sv2 deprecated o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service list v911sv2 deprecated o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service list v911sv2 deprecated o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service list v911sv2 deprecated o k response
func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /V911s/v2][%d] thinktelUControlApiWebServiceListV911sv2DeprecatedOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /V911s/v2][%d] thinktelUControlApiWebServiceListV911sv2DeprecatedOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) GetPayload() []*models.ThinkteluControlAPIDataContractsV911v2 {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceListV911sv2DeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault creates a ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault with default headers values
func NewThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault(code int) *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault {
	return &ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service list v911sv2 deprecated default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service list v911sv2 deprecated default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service list v911sv2 deprecated default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service list v911sv2 deprecated default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service list v911sv2 deprecated default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service list v911sv2 deprecated default response
func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /V911s/v2][%d] Thinktel.uControl.Api.WebService.ListV911sv2_Deprecated default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /V911s/v2][%d] Thinktel.uControl.Api.WebService.ListV911sv2_Deprecated default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceListV911sv2DeprecatedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
