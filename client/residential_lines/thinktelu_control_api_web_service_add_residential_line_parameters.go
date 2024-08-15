// Code generated by go-swagger; DO NOT EDIT.

package residential_lines

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

// NewThinkteluControlAPIWebServiceAddResidentialLineParams creates a new ThinkteluControlAPIWebServiceAddResidentialLineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceAddResidentialLineParams() *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddResidentialLineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithTimeout creates a new ThinkteluControlAPIWebServiceAddResidentialLineParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddResidentialLineParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithContext creates a new ThinkteluControlAPIWebServiceAddResidentialLineParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddResidentialLineParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceAddResidentialLineParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceAddResidentialLineParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	return &ThinkteluControlAPIWebServiceAddResidentialLineParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceAddResidentialLineParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service add residential line operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceAddResidentialLineParams struct {

	/* Order.

	   Residential line order.
	*/
	Order *models.ThinkteluControlAPIDataContractsResidentialLineOrder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service add residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WithDefaults() *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service add residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WithOrder(order *models.ThinkteluControlAPIDataContractsResidentialLineOrder) *ThinkteluControlAPIWebServiceAddResidentialLineParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the thinktel u control Api web service add residential line params
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) SetOrder(order *models.ThinkteluControlAPIDataContractsResidentialLineOrder) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceAddResidentialLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Order != nil {
		if err := r.SetBodyParam(o.Order); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}