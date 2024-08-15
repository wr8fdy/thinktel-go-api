// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams creates a new ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams() *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	return &ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	return &ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithContext creates a new ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	return &ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	return &ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get number port c IP status by ticket mask operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams struct {

	/* Ticketmask.

	   Ticket identifier.
	*/
	Ticketmask string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get number port c IP status by ticket mask params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WithDefaults() *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get number port c IP status by ticket mask params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTicketmask adds the ticketmask to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WithTicketmask(ticketmask string) *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams {
	o.SetTicketmask(ticketmask)
	return o
}

// SetTicketmask adds the ticketmask to the thinktel u control Api web service get number port c IP status by ticket mask params
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) SetTicketmask(ticketmask string) {
	o.Ticketmask = ticketmask
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetNumberPortCIPStatusByTicketMaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ticketmask
	if err := r.SetPathParam("ticketmask", o.Ticketmask); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}