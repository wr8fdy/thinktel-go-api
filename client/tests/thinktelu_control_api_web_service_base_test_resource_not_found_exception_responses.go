// Code generated by go-swagger; DO NOT EDIT.

package tests

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

// ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionReader is a Reader for the ThinkteluControlAPIWebServiceBaseTestResourceNotFoundException structure.
type ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceBaseTestResourceNotFoundExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 404:
		result := NewThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /Test/ResourceNotFoundException] Thinktel.uControl.Api.WebServiceBase.TestResourceNotFoundException", response, response.Code())
	}
}

// NewThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound creates a ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound with default headers values
func NewThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound() *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound {
	return &ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound{}
}

/*
ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound describes a response with status code 404, with default header values.

Always generates exception
*/
type ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound struct {
	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service base test resource not found exception not found response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this thinktel u control Api web service base test resource not found exception not found response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service base test resource not found exception not found response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this thinktel u control Api web service base test resource not found exception not found response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service base test resource not found exception not found response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the thinktel u control Api web service base test resource not found exception not found response
func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) Code() int {
	return 404
}

func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Test/ResourceNotFoundException][%d] thinktelUControlApiWebServiceBaseTestResourceNotFoundExceptionNotFound %s", 404, payload)
}

func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Test/ResourceNotFoundException][%d] thinktelUControlApiWebServiceBaseTestResourceNotFoundExceptionNotFound %s", 404, payload)
}

func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceBaseTestResourceNotFoundExceptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
