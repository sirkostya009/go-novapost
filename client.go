package novapost

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

const (
	JSONUrl = "https://api.novaposhta.ua/v2.0/json/"
	XMLUrl  = "https://api.novaposhta.ua/v2.0/xml/"
)

type decoder interface {
	Decode(any) error
}

type Client struct {
	apiKey             string
	http               *http.Client
	url                string
	marshaler          func(any) ([]byte, error)
	decoderConstructor func(at io.Reader) decoder
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
	c.decoderConstructor = func(r io.Reader) decoder { return json.NewDecoder(r) }
}

func WithXML(c *Client) {
	c.url = XMLUrl
	c.marshaler = xml.Marshal
	c.decoderConstructor = func(r io.Reader) decoder { return xml.NewDecoder(r) }
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.url = url
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
