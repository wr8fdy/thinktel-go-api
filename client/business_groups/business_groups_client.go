// Code generated by go-swagger; DO NOT EDIT.

package business_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new business groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new business groups API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new business groups API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for business groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationXML sets the Content-Type header to "application/xml".
func WithContentTypeApplicationXML(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/xml"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationXML sets the Accept header to "application/xml".
func WithAcceptApplicationXML(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/xml"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	ThinkteluControlAPIRestServiceCancelBusinessGroup(params *ThinkteluControlAPIRestServiceCancelBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceCancelBusinessGroupOK, error)

	ThinkteluControlAPIRestServiceGetBusinessGroup(params *ThinkteluControlAPIRestServiceGetBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceGetBusinessGroupOK, error)

	ThinkteluControlAPIRestServiceUpdateBusinessGroup(params *ThinkteluControlAPIRestServiceUpdateBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceUpdateBusinessGroupOK, error)

	ThinkteluControlAPIWebServiceAddBusinessGroup(params *ThinkteluControlAPIWebServiceAddBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceAddBusinessGroupOK, error)

	ThinkteluControlAPIWebServiceListBusinessGroups(params *ThinkteluControlAPIWebServiceListBusinessGroupsParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceListBusinessGroupsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ThinkteluControlAPIRestServiceCancelBusinessGroup cancels business group
*/
func (a *Client) ThinkteluControlAPIRestServiceCancelBusinessGroup(params *ThinkteluControlAPIRestServiceCancelBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceCancelBusinessGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIRestServiceCancelBusinessGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.RestService.CancelBusinessGroup",
		Method:             "DELETE",
		PathPattern:        "/BusinessGroups/{ID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIRestServiceCancelBusinessGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIRestServiceCancelBusinessGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIRestServiceCancelBusinessGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ThinkteluControlAPIRestServiceGetBusinessGroup retrieves business group
*/
func (a *Client) ThinkteluControlAPIRestServiceGetBusinessGroup(params *ThinkteluControlAPIRestServiceGetBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceGetBusinessGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIRestServiceGetBusinessGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.RestService.GetBusinessGroup",
		Method:             "GET",
		PathPattern:        "/BusinessGroups/{ID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIRestServiceGetBusinessGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIRestServiceGetBusinessGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIRestServiceGetBusinessGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ThinkteluControlAPIRestServiceUpdateBusinessGroup updates business group
*/
func (a *Client) ThinkteluControlAPIRestServiceUpdateBusinessGroup(params *ThinkteluControlAPIRestServiceUpdateBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIRestServiceUpdateBusinessGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIRestServiceUpdateBusinessGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.RestService.UpdateBusinessGroup",
		Method:             "PUT",
		PathPattern:        "/BusinessGroups/{ID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIRestServiceUpdateBusinessGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIRestServiceUpdateBusinessGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIRestServiceUpdateBusinessGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ThinkteluControlAPIWebServiceAddBusinessGroup orders business group
*/
func (a *Client) ThinkteluControlAPIWebServiceAddBusinessGroup(params *ThinkteluControlAPIWebServiceAddBusinessGroupParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceAddBusinessGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIWebServiceAddBusinessGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.WebService.AddBusinessGroup",
		Method:             "POST",
		PathPattern:        "/BusinessGroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIWebServiceAddBusinessGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIWebServiceAddBusinessGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIWebServiceAddBusinessGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ThinkteluControlAPIWebServiceListBusinessGroups lists business groups
*/
func (a *Client) ThinkteluControlAPIWebServiceListBusinessGroups(params *ThinkteluControlAPIWebServiceListBusinessGroupsParams, opts ...ClientOption) (*ThinktelUControlAPIWebServiceListBusinessGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThinkteluControlAPIWebServiceListBusinessGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Thinktel.uControl.Api.WebService.ListBusinessGroups",
		Method:             "GET",
		PathPattern:        "/BusinessGroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThinkteluControlAPIWebServiceListBusinessGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ThinktelUControlAPIWebServiceListBusinessGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ThinkteluControlAPIWebServiceListBusinessGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
