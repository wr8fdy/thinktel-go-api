// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

// NewThinkteluControlAPIRestServiceCancelSubscriberParams creates a new ThinkteluControlAPIRestServiceCancelSubscriberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceCancelSubscriberParams() *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	return &ThinkteluControlAPIRestServiceCancelSubscriberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithTimeout creates a new ThinkteluControlAPIRestServiceCancelSubscriberParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	return &ThinkteluControlAPIRestServiceCancelSubscriberParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithContext creates a new ThinkteluControlAPIRestServiceCancelSubscriberParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	return &ThinkteluControlAPIRestServiceCancelSubscriberParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceCancelSubscriberParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceCancelSubscriberParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	return &ThinkteluControlAPIRestServiceCancelSubscriberParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceCancelSubscriberParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service cancel subscriber operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceCancelSubscriberParams struct {

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service cancel subscriber params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WithDefaults() *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service cancel subscriber params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WithNumber(number string) *ThinkteluControlAPIRestServiceCancelSubscriberParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service cancel subscriber params
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceCancelSubscriberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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