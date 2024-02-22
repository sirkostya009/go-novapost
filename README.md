# go-novapost

A simple library-wrapper for NovaPost API.

[![](https://godoc.org/github.com/sirkostya009/go-novapost?status.svg)](https://godoc.org/github.com/sirkostya009/go-novapost)

## Idea

I was rather surprised by how sparse the availability of NovaPost API libraries for Go is, and so decided to make one of
my own which would be simple to use and try to stick to the original API reference as much as possible.

## Usage

Just go get the libbo.

```bash
go get github.com/sirkostya009/go-novapost
```

### Code:

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

You can also customize your client by passing options to it. Note, however, that none of them are required, with
client defaulting to JSON and no timeout.

```go
c := np.NewClient(os.Getenv("NOVA_POST_API_KEY"),
	WithXML,
	WithJSON, // default
	np.WithTimeout(5 * time.Second),
	WithMarshaler(json.Marshal), // default for JSON
	WithUnmarshaler(xml.Unmarshal), // default for XML
	np.WithURL("https://api.novaposhta.ua/v2.0/json/"), // set automatically by WithXML, WithJSON
)
```

There's also an option for rawdogging requests, in case you wish to do something that is not supported by the library.

```go
type customProps struct {
	Value string `json:"value"`
}

res, err := c.RawRequest("Model", "method", customProps{Value: "foo"}))

for _, m := range res.Data {
	fmt.Println(m["foo"]) // res.Data is a slice of maps
}
```

## TODO:
- [ ] Fix tests
- [ ] Fix XML
- [ ] Add constants for common strings, like "Sender" or "ThirdPerson"

## Contributing

If you wish to fix a bug or perhaps have a better idea of how to implement x, feel free to fork this repo and open a PR.
Make sure existing tests are passing and that you've added tests for your changes, too.
