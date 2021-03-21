// Package productwebclient provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package productwebclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAllProducts request
	GetAllProducts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateProduct request  with any body
	CreateProductWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateProduct(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProductById request
	GetProductById(ctx context.Context, productId ProductIdParameter, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllProducts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllProductsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateProductWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateProductRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateProduct(ctx context.Context, body CreateProductJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateProductRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProductById(ctx context.Context, productId ProductIdParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProductByIdRequest(c.Server, productId)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllProductsRequest generates requests for GetAllProducts
func NewGetAllProductsRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/product")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateProductRequest calls the generic CreateProduct builder with application/json body
func NewCreateProductRequest(server string, body CreateProductJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateProductRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateProductRequestWithBody generates requests for CreateProduct with any type of body
func NewCreateProductRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/product")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetProductByIdRequest generates requests for GetProductById
func NewGetProductByIdRequest(server string, productId ProductIdParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "productId", productId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/product/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	req = req.WithContext(ctx)
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAllProducts request
	GetAllProductsWithResponse(ctx context.Context) (*GetAllProductsResponse, error)

	// CreateProduct request  with any body
	CreateProductWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateProductResponse, error)

	CreateProductWithResponse(ctx context.Context, body CreateProductJSONRequestBody) (*CreateProductResponse, error)

	// GetProductById request
	GetProductByIdWithResponse(ctx context.Context, productId ProductIdParameter) (*GetProductByIdResponse, error)
}

type GetAllProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Product
}

// Status returns HTTPResponse.Status
func (r GetAllProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateProductResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ProductCreated
}

// Status returns HTTPResponse.Status
func (r CreateProductResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateProductResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProductByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Product
}

// Status returns HTTPResponse.Status
func (r GetProductByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProductByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllProductsWithResponse request returning *GetAllProductsResponse
func (c *ClientWithResponses) GetAllProductsWithResponse(ctx context.Context) (*GetAllProductsResponse, error) {
	rsp, err := c.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetAllProductsResponse(rsp)
}

// CreateProductWithBodyWithResponse request with arbitrary body returning *CreateProductResponse
func (c *ClientWithResponses) CreateProductWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateProductResponse, error) {
	rsp, err := c.CreateProductWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateProductResponse(rsp)
}

func (c *ClientWithResponses) CreateProductWithResponse(ctx context.Context, body CreateProductJSONRequestBody) (*CreateProductResponse, error) {
	rsp, err := c.CreateProduct(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateProductResponse(rsp)
}

// GetProductByIdWithResponse request returning *GetProductByIdResponse
func (c *ClientWithResponses) GetProductByIdWithResponse(ctx context.Context, productId ProductIdParameter) (*GetProductByIdResponse, error) {
	rsp, err := c.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	return ParseGetProductByIdResponse(rsp)
}

// ParseGetAllProductsResponse parses an HTTP response from a GetAllProductsWithResponse call
func ParseGetAllProductsResponse(rsp *http.Response) (*GetAllProductsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetAllProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Product
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateProductResponse parses an HTTP response from a CreateProductWithResponse call
func ParseCreateProductResponse(rsp *http.Response) (*CreateProductResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateProductResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ProductCreated
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseGetProductByIdResponse parses an HTTP response from a GetProductByIdWithResponse call
func ParseGetProductByIdResponse(rsp *http.Response) (*GetProductByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetProductByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Product
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
