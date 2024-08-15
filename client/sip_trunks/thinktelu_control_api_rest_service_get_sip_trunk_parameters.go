// Code generated by go-swagger; DO NOT EDIT.

package sip_trunks

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

// NewThinkteluControlAPIRestServiceGetSIPTrunkParams creates a new ThinkteluControlAPIRestServiceGetSIPTrunkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceGetSIPTrunkParams() *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithTimeout creates a new ThinkteluControlAPIRestServiceGetSIPTrunkParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithContext creates a new ThinkteluControlAPIRestServiceGetSIPTrunkParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceGetSIPTrunkParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceGetSIPTrunkParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	return &ThinkteluControlAPIRestServiceGetSIPTrunkParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceGetSIPTrunkParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service get Sip trunk operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceGetSIPTrunkParams struct {

	/* Number.

	   SIP trunk pilot directory number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service get Sip trunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WithDefaults() *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service get Sip trunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WithNumber(number string) *ThinkteluControlAPIRestServiceGetSIPTrunkParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service get Sip trunk params
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceGetSIPTrunkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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