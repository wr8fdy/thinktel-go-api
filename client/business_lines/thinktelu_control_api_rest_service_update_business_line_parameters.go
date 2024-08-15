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

	"github.com/wr8fdy/thinktel-go-api/models"
)

// NewThinkteluControlAPIRestServiceUpdateBusinessLineParams creates a new ThinkteluControlAPIRestServiceUpdateBusinessLineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceUpdateBusinessLineParams() *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	return &ThinkteluControlAPIRestServiceUpdateBusinessLineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithTimeout creates a new ThinkteluControlAPIRestServiceUpdateBusinessLineParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	return &ThinkteluControlAPIRestServiceUpdateBusinessLineParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithContext creates a new ThinkteluControlAPIRestServiceUpdateBusinessLineParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	return &ThinkteluControlAPIRestServiceUpdateBusinessLineParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceUpdateBusinessLineParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceUpdateBusinessLineParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	return &ThinkteluControlAPIRestServiceUpdateBusinessLineParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateBusinessLineParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service update business line operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceUpdateBusinessLineParams struct {

	/* Line.

	   Updated business line.
	*/
	Line *models.ThinkteluControlAPIDataContractsBusinessLine

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service update business line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithDefaults() *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service update business line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLine adds the line to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithLine(line *models.ThinkteluControlAPIDataContractsBusinessLine) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetLine(line)
	return o
}

// SetLine adds the line to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetLine(line *models.ThinkteluControlAPIDataContractsBusinessLine) {
	o.Line = line
}

// WithNumber adds the number to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WithNumber(number string) *ThinkteluControlAPIRestServiceUpdateBusinessLineParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service update business line params
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceUpdateBusinessLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Line != nil {
		if err := r.SetBodyParam(o.Line); err != nil {
			return err
		}
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}