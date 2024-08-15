// Code generated by go-swagger; DO NOT EDIT.

package residential_lines

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

// ThinkteluControlAPIRestServiceCancelResidentialLineReader is a Reader for the ThinkteluControlAPIRestServiceCancelResidentialLine structure.
type ThinkteluControlAPIRestServiceCancelResidentialLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIRestServiceCancelResidentialLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIRestServiceCancelResidentialLineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIRestServiceCancelResidentialLineOK creates a ThinktelUControlAPIRestServiceCancelResidentialLineOK with default headers values
func NewThinktelUControlAPIRestServiceCancelResidentialLineOK() *ThinktelUControlAPIRestServiceCancelResidentialLineOK {
	return &ThinktelUControlAPIRestServiceCancelResidentialLineOK{}
}

/*
ThinktelUControlAPIRestServiceCancelResidentialLineOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIRestServiceCancelResidentialLineOK struct {
	Payload *models.ThinkteluControlAPIDataContractsNumberResponse
}

// IsSuccess returns true when this thinktel u control Api rest service cancel residential line o k response has a 2xx status code
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api rest service cancel residential line o k response has a 3xx status code
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api rest service cancel residential line o k response has a 4xx status code
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api rest service cancel residential line o k response has a 5xx status code
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api rest service cancel residential line o k response a status code equal to that given
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api rest service cancel residential line o k response
func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ResidentialLines/{number}][%d] thinktelUControlApiRestServiceCancelResidentialLineOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ResidentialLines/{number}][%d] thinktelUControlApiRestServiceCancelResidentialLineOK %s", 200, payload)
}

func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) GetPayload() *models.ThinkteluControlAPIDataContractsNumberResponse {
	return o.Payload
}

func (o *ThinktelUControlAPIRestServiceCancelResidentialLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsNumberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIRestServiceCancelResidentialLineDefault creates a ThinkteluControlAPIRestServiceCancelResidentialLineDefault with default headers values
func NewThinkteluControlAPIRestServiceCancelResidentialLineDefault(code int) *ThinkteluControlAPIRestServiceCancelResidentialLineDefault {
	return &ThinkteluControlAPIRestServiceCancelResidentialLineDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIRestServiceCancelResidentialLineDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIRestServiceCancelResidentialLineDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api rest service cancel residential line default response has a 2xx status code
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api rest service cancel residential line default response has a 3xx status code
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api rest service cancel residential line default response has a 4xx status code
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api rest service cancel residential line default response has a 5xx status code
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api rest service cancel residential line default response a status code equal to that given
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api rest service cancel residential line default response
func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ResidentialLines/{number}][%d] Thinktel.uControl.Api.RestService.CancelResidentialLine default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ResidentialLines/{number}][%d] Thinktel.uControl.Api.RestService.CancelResidentialLine default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIRestServiceCancelResidentialLineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
