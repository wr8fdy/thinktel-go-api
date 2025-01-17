// Code generated by go-swagger; DO NOT EDIT.

package surecall_sip_trunks

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

// ThinkteluControlAPIWebServiceGetSurecallReader is a Reader for the ThinkteluControlAPIWebServiceGetSurecall structure.
type ThinkteluControlAPIWebServiceGetSurecallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceGetSurecallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceGetSurecallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceGetSurecallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceGetSurecallOK creates a ThinktelUControlAPIWebServiceGetSurecallOK with default headers values
func NewThinktelUControlAPIWebServiceGetSurecallOK() *ThinktelUControlAPIWebServiceGetSurecallOK {
	return &ThinktelUControlAPIWebServiceGetSurecallOK{}
}

/*
ThinktelUControlAPIWebServiceGetSurecallOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceGetSurecallOK struct {
	Payload []*models.ThinkteluControlAPIDataContractsSurecallNumber
}

// IsSuccess returns true when this thinktel u control Api web service get surecall o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service get surecall o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service get surecall o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service get surecall o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service get surecall o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service get surecall o k response
func (o *ThinktelUControlAPIWebServiceGetSurecallOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceGetSurecallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Surecall/{pilot}][%d] thinktelUControlApiWebServiceGetSurecallOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetSurecallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Surecall/{pilot}][%d] thinktelUControlApiWebServiceGetSurecallOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetSurecallOK) GetPayload() []*models.ThinkteluControlAPIDataContractsSurecallNumber {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceGetSurecallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceGetSurecallDefault creates a ThinkteluControlAPIWebServiceGetSurecallDefault with default headers values
func NewThinkteluControlAPIWebServiceGetSurecallDefault(code int) *ThinkteluControlAPIWebServiceGetSurecallDefault {
	return &ThinkteluControlAPIWebServiceGetSurecallDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceGetSurecallDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceGetSurecallDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service get surecall default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service get surecall default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service get surecall default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service get surecall default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service get surecall default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service get surecall default response
func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Surecall/{pilot}][%d] Thinktel.uControl.Api.WebService.GetSurecall default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Surecall/{pilot}][%d] Thinktel.uControl.Api.WebService.GetSurecall default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceGetSurecallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
