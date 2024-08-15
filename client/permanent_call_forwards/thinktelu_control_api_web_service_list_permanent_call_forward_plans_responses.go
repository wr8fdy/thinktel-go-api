// Code generated by go-swagger; DO NOT EDIT.

package permanent_call_forwards

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

// ThinkteluControlAPIWebServiceListPermanentCallForwardPlansReader is a Reader for the ThinkteluControlAPIWebServiceListPermanentCallForwardPlans structure.
type ThinkteluControlAPIWebServiceListPermanentCallForwardPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK creates a ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK with default headers values
func NewThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK() *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK {
	return &ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK{}
}

/*
ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK struct {
	Payload []*models.ThinkteluControlAPIDataContractsPermanentCallForwardPlan
}

// IsSuccess returns true when this thinktel u control Api web service list permanent call forward plans o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service list permanent call forward plans o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service list permanent call forward plans o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service list permanent call forward plans o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service list permanent call forward plans o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service list permanent call forward plans o k response
func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PermanentCallForwardPlans][%d] thinktelUControlApiWebServiceListPermanentCallForwardPlansOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PermanentCallForwardPlans][%d] thinktelUControlApiWebServiceListPermanentCallForwardPlansOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) GetPayload() []*models.ThinkteluControlAPIDataContractsPermanentCallForwardPlan {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceListPermanentCallForwardPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault creates a ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault with default headers values
func NewThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault(code int) *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault {
	return &ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service list permanent call forward plans default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service list permanent call forward plans default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service list permanent call forward plans default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service list permanent call forward plans default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service list permanent call forward plans default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service list permanent call forward plans default response
func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PermanentCallForwardPlans][%d] Thinktel.uControl.Api.WebService.ListPermanentCallForwardPlans default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PermanentCallForwardPlans][%d] Thinktel.uControl.Api.WebService.ListPermanentCallForwardPlans default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceListPermanentCallForwardPlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
