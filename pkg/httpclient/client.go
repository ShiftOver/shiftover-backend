// Package httpclient provides a retryable HTTP client for making requests to server services
package httpclient

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

const (
	// QueryParam represents a http query param
	QueryParam = "queries"

	// HeaderParam represents http headers
	HeaderParam = "headers"

	// ContentTypeJSON represents http content type of JSON
	ContentTypeJSON = "application/json"
)

// Client is the interface for the HTTP client representing the functions of the client
type Client interface {
	Do(method, targetURL string, body []byte, options SendOptions) (*Response, error)
	Get(targetURL string, options SendOptions) (*Response, error)
	Post(targetURL string, body []byte, options SendOptions) (*Response, error)
	Put(targetURL string, body []byte, options SendOptions) (*Response, error)
	Patch(targetURL string, body []byte, options SendOptions) (*Response, error)
	Delete(targetURL string, body []byte, options SendOptions) (*Response, error)
}

// SendOptions for attached data through a request.
// For example, if header of content-type and custom header "custom" is needed,
// with query params of (http://example.com?query=example&limit=10),
//
//	SendOptions {
//		"queries": map[string]string{
//			"queries": "example",
//			"limit": "10",
//		},
//		"headers": map[string]string{
//			"custom":         	"example",
//			"Content-Type":  "application/json",
//		},
//	}
type SendOptions map[string]map[string]string

// NewSendOptions creates a new SendOptions for the client
func NewSendOptions() SendOptions {
	return SendOptions{}
}

// SetTimeout sets the timeout for the request
func SetTimeout(timeout time.Duration) OptionClient {
	return func(r *retryablehttp.Client) {
		r.HTTPClient.Timeout = timeout
	}
}

// SetMaxRetry sets the maximum number of retries for the request
func SetMaxRetry(maxRetry int) OptionClient {
	return func(r *retryablehttp.Client) {
		r.RetryMax = maxRetry
	}
}

// OptionClient is the type for the options to set for the client
type OptionClient func(*retryablehttp.Client)

type client struct {
	HTTPClient *http.Client
}

// Response represents the response from the server
type Response struct {
	Body       []byte
	Header     http.Header
	StatusCode int
}

// NewClient creates a new HTTP client with the options
func NewClient(options ...OptionClient) Client {
	httpClient := retryablehttp.NewClient()
	for _, option := range options {
		option(httpClient)
	}

	return &client{
		HTTPClient: httpClient.StandardClient(),
	}
}

// SetContentType sets the content type of the request
func (opt SendOptions) SetContentType(t string) SendOptions {
	newOpt := opt
	if newOpt == nil {
		newOpt = make(SendOptions)
	}

	if newOpt[HeaderParam] == nil {
		newOpt[HeaderParam] = make(map[string]string)
	}
	newOpt[HeaderParam]["Content-Type"] = t
	return newOpt
}

// AddHeader adds a header to the request
func (opt SendOptions) AddHeader(key, value string) SendOptions {
	newOpt := opt
	if newOpt == nil {
		newOpt = make(SendOptions)
	}

	if newOpt[HeaderParam] == nil {
		newOpt[HeaderParam] = make(map[string]string)
	}
	newOpt[HeaderParam][http.CanonicalHeaderKey(key)] = value
	return newOpt
}

// SetQueryParam sets the query params for the request
func (opt SendOptions) SetQueryParam(params map[string]string) SendOptions {
	newOpt := opt
	if newOpt == nil {
		newOpt = make(SendOptions)
	}

	newOpt[QueryParam] = params
	return newOpt
}

// Do sends a request to the server with the method, url, body and options
func (c client) Do(method, targetURL string, body []byte, options SendOptions) (*Response, error) {
	log.Infof("[httpclient.Do] sending http request: %s", map[string]interface{}{
		"method": method,
		"url":    targetURL,
		"body":   string(body),
		"opts":   options,
	})

	method = strings.ToUpper(method)
	urlStruct, err := url.Parse(targetURL)
	if err != nil {
		return nil, errors.Wrap(err, "error - [httpclient.Do] failed to parse url")
	}

	queryString := url.Values{}
	if query := options[QueryParam]; len(query) != 0 {
		for key, value := range query {
			queryString.Add(key, value)
		}
	}
	if len(queryString) > 0 {
		urlStruct.RawQuery = queryString.Encode()
	}
	requestAPIUrl := urlStruct.String()

	var b io.Reader
	if body != nil {
		b = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, requestAPIUrl, b)
	if err != nil {
		return nil, errors.Wrap(err, "error - [httpclient.Do] failed to create request")
	}

	if headers := options[HeaderParam]; len(headers) != 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	netResponse, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error - [httpclient.Do] failed to send request")
	}

	content, err := io.ReadAll(netResponse.Body)
	defer func() {
		if err := netResponse.Body.Close(); err != nil {
			log.Errorf("error - [httpclient.Do] failed to close response body: %v", err)
		}
	}()
	if err != nil {
		return nil, errors.Wrap(err, "error - [httpclient.Do] failed to read response body")
	}

	response := &Response{
		Body:       content,
		Header:     netResponse.Header,
		StatusCode: netResponse.StatusCode,
	}

	return response, nil
}

// Get sends a GET request to the server with the url and options
func (c client) Get(targetURL string, options SendOptions) (*Response, error) {
	return c.Do(http.MethodGet, targetURL, nil, options)
}

// Post sends a POST request to the server with the url, body and options
func (c client) Post(targetURL string, body []byte, options SendOptions) (*Response, error) {
	return c.Do(http.MethodPost, targetURL, body, options)
}

// Put sends a PUT request to the server with the url, body and options
func (c client) Put(targetURL string, body []byte, options SendOptions) (*Response, error) {
	return c.Do(http.MethodPut, targetURL, body, options)
}

// Patch sends a PATCH request to the server with the url, body and options
func (c client) Patch(targetURL string, body []byte, options SendOptions) (*Response, error) {
	return c.Do(http.MethodPatch, targetURL, body, options)
}

// Delete sends a DELETE request to the server with the url and options
func (c client) Delete(targetURL string, body []byte, options SendOptions) (*Response, error) {
	return c.Do(http.MethodDelete, targetURL, body, options)
}
