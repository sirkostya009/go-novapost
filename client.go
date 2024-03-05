package novapost

import (
	"encoding/xml"
	"net/http"
)

const (
	JSONUrl = "https://api.novaposhta.ua/v2.0/json/"
	XMLUrl  = "https://api.novaposhta.ua/v2.0/xml/"
)

type Client struct {
	apiKey      string
	HTTPClient  *http.Client
	Url         string
	Marshaler   func(any) ([]byte, error)
	Unmarshaler func([]byte, any) error
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.Url = url
	}
}

func WithMarshaler(marshaler func(any) ([]byte, error)) Option {
	return func(c *Client) {
		c.Marshaler = marshaler
	}
}

func WithUnmarshaler(unmarshaler func([]byte, any) error) Option {
	return func(c *Client) {
		c.Unmarshaler = unmarshaler
	}
}

func NewClient(apiKey string, options ...Option) *Client {
	c := &Client{apiKey, &http.Client{}, XMLUrl, xml.Marshal, xml.Unmarshal}
	for _, option := range options {
		option(c)
	}
	return c
}
