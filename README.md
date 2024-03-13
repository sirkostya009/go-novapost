# go-novapost

A simple library-wrapper for NovaPost API.

[![](https://godoc.org/github.com/sirkostya009/go-novapost?status.svg)](https://godoc.org/github.com/sirkostya009/go-novapost)

## Idea

I was rather surprised by how sparse the availability of NovaPost API libraries for Go is, and so decided to make one of
my own which would be simple to use and tries to stick to the original API reference as much as possible.

## Usage

Just go get the libbo.

```bash
go get github.com/sirkostya009/go-novapost
```

### Example:

```go
package main

import (
	"fmt"
	np "github.com/sirkostya009/go-novapost"
	"os"
)

func main() {
	c := np.NewClient(os.Getenv("NOVA_POST_API_KEY"))

	citiesRes, _ := c.GetCities(np.CityRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	for _, city := range citiesRes.Data {
		fmt.Println(city.Description) // prints "Київ" and "Київець"
	}
}
```

You can also customize your client by passing options to it. Note, however, that none of them are required, with client
defaulting to XML and a new `http.Client` instance.

```go
c := np.NewClient(os.Getenv("NOVA_POST_API_KEY"),
	np.WithHTTPClient(&http.Client{}),
	np.WithMarshaler(xml.Marshal),
	np.WithUnmarshaler(xml.Unmarshal),
	np.WithURL(np.XMLUrl),
)
```

You can also initialize the client after construction:

```go
c := np.NewClient(os.Getenv("NOVA_POST_API_KEY"))
c.Url = np.JSONUrl
c.HTTPClient = http.DefaultClient
c.Marshaler = fasterXmlMarshaler
```

As you can see, this library uses XML for request marshalling, due to the nature of data returned by API. You aren't
stripped of the ability to use JSON, though. Do note that the implementation must be a custom one, as most models have
ints and floats whereas JSON response from API for those fields is a string.

Alternatively, you are provided with an option for rawdogging requests, in case you wish to do something that is not
supported by the library.

```go
type customProps struct {
	Value string
}

type responseData struct {
    Foo string
}

res, err := np.RawRequest[responseData](c, "Model", "method", customProps{Value: "foo"}))

for _, m := range res.Data {
	fmt.Println(m.Foo)
}
```

## TODO:
- [x] Fix tests
- [x] Fix XML
- [x] Add constants for common strings, like "Sender" or "ThirdPerson"
- [ ] Add custom marshaler/unmarshaler for JSON

## Contributing

If you wish to fix a bug or perhaps have a better idea of how to implement something, feel free to fork this repo and
open a PR. Make sure existing tests are passing and that you've added tests for your changes, too.
