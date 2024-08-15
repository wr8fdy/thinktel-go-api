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

// NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParams creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParams() *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithTimeout creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithContext creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceGetPortInStatusByTicketParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	return &ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service get port in status by ticket operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams struct {

	/* TicketID.

	   Ticket identifier.
	*/
	TicketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service get port in status by ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WithDefaults() *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service get port in status by ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTicketID adds the ticketID to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WithTicketID(ticketID string) *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams {
	o.SetTicketID(ticketID)
	return o
}

// SetTicketID adds the ticketId to the thinktel u control Api web service get port in status by ticket params
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) SetTicketID(ticketID string) {
	o.TicketID = ticketID
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceGetPortInStatusByTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ticketId
	if err := r.SetPathParam("ticketId", o.TicketID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
