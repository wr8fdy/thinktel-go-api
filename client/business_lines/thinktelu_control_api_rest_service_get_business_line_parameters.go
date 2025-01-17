// Code generated by go-swagger; DO NOT EDIT.

package business_lines

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

// NewThinkteluControlAPIRestServiceGetBusinessLineParams creates a new ThinkteluControlAPIRestServiceGetBusinessLineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetBusinessLineParams() *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	return &ThinkteluControlAPIRestServiceGetBusinessLineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetBusinessLineParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	return &ThinkteluControlAPIRestServiceGetBusinessLineParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithContext creates a new ThinkteluControlAPIRestServiceGetBusinessLineParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	return &ThinkteluControlAPIRestServiceGetBusinessLineParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetBusinessLineParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetBusinessLineParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	return &ThinkteluControlAPIRestServiceGetBusinessLineParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetBusinessLineParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get business line operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetBusinessLineParams struct {

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get business line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WithDefaults() *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get business line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WithNumber(number string) *ThinkteluControlAPIRestServiceGetBusinessLineParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service get business line params
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetBusinessLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
