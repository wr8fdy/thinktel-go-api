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

// NewThinkteluControlAPIRestServiceUpdateResidentialLineParams creates a new ThinkteluControlAPIRestServiceUpdateResidentialLineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceUpdateResidentialLineParams() *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	return &ThinkteluControlAPIRestServiceUpdateResidentialLineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithTimeout creates a new ThinkteluControlAPIRestServiceUpdateResidentialLineParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	return &ThinkteluControlAPIRestServiceUpdateResidentialLineParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithContext creates a new ThinkteluControlAPIRestServiceUpdateResidentialLineParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	return &ThinkteluControlAPIRestServiceUpdateResidentialLineParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceUpdateResidentialLineParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceUpdateResidentialLineParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	return &ThinkteluControlAPIRestServiceUpdateResidentialLineParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateResidentialLineParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service update residential line operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceUpdateResidentialLineParams struct {

	/* Line.

	   Updated residential line.
	*/
	Line *models.ThinkteluControlAPIDataContractsResidentialLine

	/* Number.

	   Telephone number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service update residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithDefaults() *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service update residential line params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLine adds the line to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithLine(line *models.ThinkteluControlAPIDataContractsResidentialLine) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetLine(line)
	return o
}

// SetLine adds the line to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetLine(line *models.ThinkteluControlAPIDataContractsResidentialLine) {
	o.Line = line
}

// WithNumber adds the number to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WithNumber(number string) *ThinkteluControlAPIRestServiceUpdateResidentialLineParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service update residential line params
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceUpdateResidentialLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
