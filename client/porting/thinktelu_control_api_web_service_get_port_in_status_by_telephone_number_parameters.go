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

// NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams() *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithContext creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get port in status by telephone number operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams struct {

	/* Tn.

	   Telephone number.
	*/
	Tn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get port in status by telephone number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WithDefaults() *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get port in status by telephone number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTn adds the tn to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WithTn(tn string) *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams {
	o.SetTn(tn)
	return o
}

// SetTn adds the tn to the thinktel u control Api web service get port in status by telephone number params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) SetTn(tn string) {
	o.Tn = tn
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTelephoneNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
