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

	"github.com/wr8fdy/thinktel-go-api/models"
)

// NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams creates a new ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams() *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	return &ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithTimeout creates a new ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams object
// with the ability to set a timeout on a request.
func NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	return &ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams{
		timeout: timeout,
	}
}

// NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithContext creates a new ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams object
// with the ability to set a context for a request.
func NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	return &ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams{
		Context: ctx,
	}
}

// NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithHTTPClient creates a new ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParamsWithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	return &ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams{
		HTTPClient: client,
	}
}

/*
ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams contains all the parameters to send to the API endpoint

	for the thinktel u control Api rest service update Sip trunk Did v2 deprecated operation.

	Typically these are written to a http.Request.
*/
type ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams struct {

	/* DID.

	   Telephone number.
	*/
	DID string

	/* Line.

	   DID.
	*/
	Line *models.ThinkteluControlAPIDataContractsUpdateDID

	/* Number.

	   SIP trunk pilot directory number.
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thinktel u control Api rest service update Sip trunk Did v2 deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithDefaults() *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thinktel u control Api rest service update Sip trunk Did v2 deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithTimeout(timeout time.Duration) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithContext(ctx context.Context) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithHTTPClient(client *http.Client) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDID adds the did to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithDID(did string) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetDID(did)
	return o
}

// SetDID adds the did to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetDID(did string) {
	o.DID = did
}

// WithLine adds the line to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithLine(line *models.ThinkteluControlAPIDataContractsUpdateDID) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetLine(line)
	return o
}

// SetLine adds the line to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetLine(line *models.ThinkteluControlAPIDataContractsUpdateDID) {
	o.Line = line
}

// WithNumber adds the number to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WithNumber(number string) *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the thinktel u control Api rest service update Sip trunk Did v2 deprecated params
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *ThinkteluControlAPIRestServiceUpdateSIPTrunkDIDV2DeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param did
	if err := r.SetPathParam("did", o.DID); err != nil {
		return err
	}
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