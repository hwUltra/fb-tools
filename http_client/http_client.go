package http_client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Client is the custom HTTP client
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new HTTP client with the specified timeout
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// DoRequest is a generic method to handle all HTTP methods
func (c *Client) DoRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		switch v := body.(type) {
		case string:
			reqBody = bytes.NewBuffer([]byte(v))
		default:
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reqBody = bytes.NewBuffer(jsonBody)
		}
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.httpClient.Do(req)
}

// Get sends a GET request
func (c *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, url, headers, nil)
}

// Post sends a POST request
func (c *Client) Post(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodPost, url, headers, body)
}

// Put sends a PUT request
func (c *Client) Put(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodPut, url, headers, body)
}

// Delete sends a DELETE request
func (c *Client) Delete(url string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, url, headers, nil)
}

// SSEClient handles Server-Sent Events
type SSEClient struct {
	client *Client
}
