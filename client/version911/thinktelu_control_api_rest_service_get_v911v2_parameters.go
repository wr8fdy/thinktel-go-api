// Code generated by go-swagger; DO NOT EDIT.

package version911

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

// NewThinkteluControlAPIRestServiceGetV911v2Params creates a new ThinkteluControlAPIRestServiceGetV911v2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetV911v2Params() *ThinkteluControlAPIRestServiceGetV911v2Params {
	return &ThinkteluControlAPIRestServiceGetV911v2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetV911v2ParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetV911v2Params object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetV911v2ParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetV911v2Params {
	return &ThinkteluControlAPIRestServiceGetV911v2Params{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetV911v2ParamsWithContext creates a new ThinkteluControlAPIRestServiceGetV911v2Params object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetV911v2ParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetV911v2Params {
	return &ThinkteluControlAPIRestServiceGetV911v2Params{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetV911v2ParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetV911v2Params object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetV911v2ParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetV911v2Params {
	return &ThinkteluControlAPIRestServiceGetV911v2Params{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetV911v2Params contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get v911v2 operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetV911v2Params struct {

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get v911v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WithDefaults() *ThinkteluControlAPIRestServiceGetV911v2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get v911v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetV911v2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetV911v2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetV911v2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WithNumber(number string) *ThinkteluControlAPIRestServiceGetV911v2Params {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service get v911v2 params
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetV911v2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}