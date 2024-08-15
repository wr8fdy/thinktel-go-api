// Code generated by go-swagger; DO NOT EDIT.

package dids

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

// NewThinkteluControlAPIRestServiceGetDIDParams creates a new ThinkteluControlAPIRestServiceGetDIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetDIDParams() *ThinkteluControlAPIRestServiceGetDIDParams {
	return &ThinkteluControlAPIRestServiceGetDIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetDIDParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetDIDParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetDIDParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetDIDParams {
	return &ThinkteluControlAPIRestServiceGetDIDParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetDIDParamsWithContext creates a new ThinkteluControlAPIRestServiceGetDIDParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetDIDParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetDIDParams {
	return &ThinkteluControlAPIRestServiceGetDIDParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetDIDParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetDIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetDIDParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetDIDParams {
	return &ThinkteluControlAPIRestServiceGetDIDParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetDIDParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get Did operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetDIDParams struct {

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get Did params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WithDefaults() *ThinkteluControlAPIRestServiceGetDIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get Did params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetDIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetDIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetDIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetDIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WithNumber(number string) *ThinkteluControlAPIRestServiceGetDIDParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service get Did params
func (o *ThinkteluControlAPIRestServiceGetDIDParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetDIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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