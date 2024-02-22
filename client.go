package novapost

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"
)

const (
	JSONUrl = "https://api.novaposhta.ua/v2.0/json/"
	XMLUrl  = "https://api.novaposhta.ua/v2.0/xml/"
)

type Client struct {
	apiKey      string
	http        *http.Client
	url         string
	marshaler   func(any) ([]byte, error)
	unmarshaler func([]byte, any) error
}

type Option func(*Client)

func WithTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.http.Timeout = t
	}
}

func WithJSON(c *Client) {
	c.url = JSONUrl
	c.marshaler = json.Marshal
	c.unmarshaler = json.Unmarshal
}

func WithXML(c *Client) {
	c.url = XMLUrl
	c.marshaler = xml.Marshal
	c.unmarshaler = xml.Unmarshal
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

func WithMarshaler(marshaler func(any) ([]byte, error)) Option {
	return func(c *Client) {
		c.marshaler = marshaler
	}
}

func WithUnmarshaler(unmarshaler func([]byte, any) error) Option {
	return func(c *Client) {
		c.unmarshaler = unmarshaler
	}
}

func NewClient(apiKey string, options ...Option) *Client {
	c := &Client{apiKey: apiKey, http: &http.Client{}}
	WithJSON(c)
	for _, option := range options {
		option(c)
	}
	return c
}
