// Code generated by go-swagger; DO NOT EDIT.

package rate_centers

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

// ThinkteluControlAPIWebServiceListCountryNPAsReader is a Reader for the ThinkteluControlAPIWebServiceListCountryNPAs structure.
type ThinkteluControlAPIWebServiceListCountryNPAsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceListCountryNPAsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceListCountryNPAsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceListCountryNPAsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceListCountryNPAsOK creates a ThinktelUControlAPIWebServiceListCountryNPAsOK with default headers values
func NewThinktelUControlAPIWebServiceListCountryNPAsOK() *ThinktelUControlAPIWebServiceListCountryNPAsOK {
	return &ThinktelUControlAPIWebServiceListCountryNPAsOK{}
}

/*
ThinktelUControlAPIWebServiceListCountryNPAsOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceListCountryNPAsOK struct {
	Payload []int64
}

// IsSuccess returns true when this thinktel u control Api web service list country n p as o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service list country n p as o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service list country n p as o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service list country n p as o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service list country n p as o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service list country n p as o k response
func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /RateCenters/Country/{country_alpha_2}/NPAs][%d] thinktelUControlApiWebServiceListCountryNPAsOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /RateCenters/Country/{country_alpha_2}/NPAs][%d] thinktelUControlApiWebServiceListCountryNPAsOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) GetPayload() []int64 {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceListCountryNPAsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceListCountryNPAsDefault creates a ThinkteluControlAPIWebServiceListCountryNPAsDefault with default headers values
func NewThinkteluControlAPIWebServiceListCountryNPAsDefault(code int) *ThinkteluControlAPIWebServiceListCountryNPAsDefault {
	return &ThinkteluControlAPIWebServiceListCountryNPAsDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceListCountryNPAsDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceListCountryNPAsDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service list country n p as default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service list country n p as default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service list country n p as default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service list country n p as default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service list country n p as default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service list country n p as default response
func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /RateCenters/Country/{country_alpha_2}/NPAs][%d] Thinktel.uControl.Api.WebService.ListCountryNPAs default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /RateCenters/Country/{country_alpha_2}/NPAs][%d] Thinktel.uControl.Api.WebService.ListCountryNPAs default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceListCountryNPAsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
