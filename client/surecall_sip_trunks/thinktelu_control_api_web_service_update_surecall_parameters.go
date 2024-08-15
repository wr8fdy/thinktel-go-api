// Code generated by go-swagger; DO NOT EDIT.

package surecall_sip_trunks

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

// NewThinkteluControlAPIWebServiceUpdateSurecallParams creates a new ThinkteluControlAPIWebServiceUpdateSurecallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIWebServiceUpdateSurecallParams() *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	return &ThinkteluControlAPIWebServiceUpdateSurecallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithTimeout creates a new ThinkteluControlAPIWebServiceUpdateSurecallParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	return &ThinkteluControlAPIWebServiceUpdateSurecallParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithContext creates a new ThinkteluControlAPIWebServiceUpdateSurecallParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithContext(ctx context.Context) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	return &ThinkteluControlAPIWebServiceUpdateSurecallParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithHTTPClient creates a new ThinkteluControlAPIWebServiceUpdateSurecallParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIWebServiceUpdateSurecallParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	return &ThinkteluControlAPIWebServiceUpdateSurecallParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIWebServiceUpdateSurecallParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api web service update surecall operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIWebServiceUpdateSurecallParams struct {

	/* Numbers.

	   List of updated surecall numbers.
	*/
	Numbers []*models.ThinkteluControlAPIDataContractsSurecallNumber

	/* Pilot.

	   SIP trunk pilot directory number.
	*/
	Pilot string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api web service update surecall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithDefaults() *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api web service update surecall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithContext(ctx context.Context) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumbers adds the numbers to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithNumbers(numbers []*models.ThinkteluControlAPIDataContractsSurecallNumber) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetNumbers(numbers)
	return o
}

// SetNumbers adds the numbers to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetNumbers(numbers []*models.ThinkteluControlAPIDataContractsSurecallNumber) {
	o.Numbers = numbers
}

// WithPilot adds the pilot to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WithPilot(pilot string) *ThinkteluControlAPIWebServiceUpdateSurecallParams {
	o.SetPilot(pilot)
	return o
}

// SetPilot adds the pilot to the thinktel u control Api web service update surecall params
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) SetPilot(pilot string) {
	o.Pilot = pilot
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIWebServiceUpdateSurecallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Numbers != nil {
		if err := r.SetBodyParam(o.Numbers); err != nil {
			return err
		}
	}

	// path param pilot
	if err := r.SetPathParam("pilot", o.Pilot); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
