// Code generated by go-swagger; DO NOT EDIT.

package provisioned_devices

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

// ThinkteluControlAPIWebServiceGetProvisionedDeviceReader is a Reader for the ThinkteluControlAPIWebServiceGetProvisionedDevice structure.
type ThinkteluControlAPIWebServiceGetProvisionedDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThinktelUControlAPIWebServiceGetProvisionedDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewThinkteluControlAPIWebServiceGetProvisionedDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewThinktelUControlAPIWebServiceGetProvisionedDeviceOK creates a ThinktelUControlAPIWebServiceGetProvisionedDeviceOK with default headers values
func NewThinktelUControlAPIWebServiceGetProvisionedDeviceOK() *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK {
	return &ThinktelUControlAPIWebServiceGetProvisionedDeviceOK{}
}

/*
ThinktelUControlAPIWebServiceGetProvisionedDeviceOK describes a response with status code 200, with default header values.

Success
*/
type ThinktelUControlAPIWebServiceGetProvisionedDeviceOK struct {
	Payload *models.ThinkteluControlAPIDataContractsProvisionedDevice
}

// IsSuccess returns true when this thinktel u control Api web service get provisioned device o k response has a 2xx status code
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thinktel u control Api web service get provisioned device o k response has a 3xx status code
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thinktel u control Api web service get provisioned device o k response has a 4xx status code
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thinktel u control Api web service get provisioned device o k response has a 5xx status code
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thinktel u control Api web service get provisioned device o k response a status code equal to that given
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thinktel u control Api web service get provisioned device o k response
func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) Code() int {
	return 200
}

func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ProvisionedDevices/{MACAddress}][%d] thinktelUControlApiWebServiceGetProvisionedDeviceOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ProvisionedDevices/{MACAddress}][%d] thinktelUControlApiWebServiceGetProvisionedDeviceOK %s", 200, payload)
}

func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) GetPayload() *models.ThinkteluControlAPIDataContractsProvisionedDevice {
	return o.Payload
}

func (o *ThinktelUControlAPIWebServiceGetProvisionedDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsProvisionedDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThinkteluControlAPIWebServiceGetProvisionedDeviceDefault creates a ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault with default headers values
func NewThinkteluControlAPIWebServiceGetProvisionedDeviceDefault(code int) *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault {
	return &ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault{
		_statusCode: code,
	}
}

/*
ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault describes a response with status code -1, with default header values.

Error
*/
type ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault struct {
	_statusCode int

	Payload *models.ThinkteluControlAPIDataContractsException
}

// IsSuccess returns true when this thinktel u control Api web service get provisioned device default response has a 2xx status code
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this thinktel u control Api web service get provisioned device default response has a 3xx status code
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this thinktel u control Api web service get provisioned device default response has a 4xx status code
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this thinktel u control Api web service get provisioned device default response has a 5xx status code
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this thinktel u control Api web service get provisioned device default response a status code equal to that given
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the thinktel u control Api web service get provisioned device default response
func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) Code() int {
	return o._statusCode
}

func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ProvisionedDevices/{MACAddress}][%d] Thinktel.uControl.Api.WebService.GetProvisionedDevice default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ProvisionedDevices/{MACAddress}][%d] Thinktel.uControl.Api.WebService.GetProvisionedDevice default %s", o._statusCode, payload)
}

func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) GetPayload() *models.ThinkteluControlAPIDataContractsException {
	return o.Payload
}

func (o *ThinkteluControlAPIWebServiceGetProvisionedDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThinkteluControlAPIDataContractsException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
