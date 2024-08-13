package whatsapp_business

import (
	"net/http"
)

type Client struct {
	token      string
	baseUrl    string
	httpClient *http.Client

	// phoneNumberID is required
	// WhatsApp user phone number.
	phoneNumberID string
}

// Option is a function that configures a client
type Option func(*Client)

// NewClient creates a new client with the provided options
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseUrl:    "https://api.default.com",
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithToken sets the token of the client
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// WithBaseUrl sets the base URL of the client
func WithBaseUrl(baseUrl string) Option {
	return func(c *Client) {
		c.baseUrl = baseUrl
	}
}

// WithHttpClient sets the http client of the client
func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
