// Code generated by go-swagger; DO NOT EDIT.

package porting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams() *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithContext creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	return &ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get port sys customer info by tn operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams struct {

	/* Tn.

	   Telephone number.
	*/
	Tn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get port sys customer info by tn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WithDefaults() *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get port sys customer info by tn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTn adds the tn to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WithTn(tn string) *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams {
	o.SetTn(tn)
	return o
}

// SetTn adds the tn to the thinktel u control Api web service get port sys customer info by tn params
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) SetTn(tn string) {
	o.Tn = tn
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetPortSysCustomerInfoByTnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tn
	if err := r.SetPathParam("tn", o.Tn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
