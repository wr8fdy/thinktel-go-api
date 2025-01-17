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

	"github.com/wr8fdy/thinktel-go-api/models"
)

// NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams creates a new ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams() *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithContext creates a new ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	return &ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add number port general child product operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams struct {

	/* OrderForm.

	   General child product port-in order.
	*/
	OrderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderGeneralChildProduct

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add number port general child product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WithDefaults() *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add number port general child product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderForm adds the orderForm to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WithOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderGeneralChildProduct) *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams {
	o.SetOrderForm(orderForm)
	return o
}

// SetOrderForm adds the orderForm to the thinktel u control Api web service add number port general child product params
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) SetOrderForm(orderForm *models.ThinkteluControlAPIDataContractsNumberPortOrderGeneralChildProduct) {
	o.OrderForm = orderForm
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddNumberPortGeneralChildProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.OrderForm != nil {
		if err := r.SetBodyParam(o.OrderForm); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
