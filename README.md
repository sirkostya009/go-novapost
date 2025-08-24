# go-novapost

A simple library-wrapper for Nova Poshta API.

[![](https://godoc.org/github.com/sirkostya009/go-novapost?status.svg)](https://godoc.org/github.com/sirkostya009/go-novapost)
[![Build & Test](https://github.com/sirkostya009/go-novapost/actions/workflows/go.yml/badge.svg)](https://github.com/sirkostya009/go-novapost/actions/workflows/go.yml)

> [!Note]
>
> This is a package for interacting with the awful Ukrainian Nova Poshta API available at https://api.novaposhta.ua/v2.0.
>
> NOT the "international" one https://api.novapost.com/developers/index.html.

## Usage

```bash
go get github.com/sirkostya009/go-novapost
```

```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"

	np "github.com/sirkostya009/go-novapost"
)

func main() {
	c := np.NewClient(os.Getenv("NOVA_POST_API_KEY"),
		// client constructor supports the common ...opts pattern, below are defaults
		np.WithHTTPClient(&http.Client{}),
		np.WithMarshaler(xml.Marshal),
		np.WithUnmarshaler(xml.Unmarshal),
		np.WithURL(np.XMLUrl),
	)
	// you can also change things later, all properties except apiKey are exported
	c.Url = np.JSONUrl
	c.Marshaler = json.Marshal
	c.Unmarshaler = json.Unmarshal
	// tip: use json for smaller network footprint

	citiesRes, _ := c.GetCities(np.CityRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	for _, city := range citiesRes.Data {
		fmt.Println(city.Description) // prints "Київ" and "Київець"
	}

	// a package-level RawRequest func is also exported,
	// that is in fact used by *all* client's methods. (inspect them and you'll see)
	// useful for occasions when something is bugged
	// or wasn't added from the api docs.
	np.RawRequest[np.City](c, np.AddressModel, "getCities", np.CityRequest{})

	// custom types usage example:

	type foo struct {
		Foo int `json:"foo,string,omitempty"`
	}

	type bar struct {
		Bar string
	}

	res, _ := np.RawRequest[bar](c, "Model", "method", foo{Foo: 1000})

	for _, m := range res.Data {
		fmt.Println(m.Bar)
	}
}
```

## TODO:
- [x] Fix tests
- [x] Fix XML
- [x] Add constants for common strings, like "Sender" or "ThirdPerson"
- [ ] ~~Add custom marshaler/unmarshaler for JSON~~

## Contributing

If any new additions to Nova Poshta's API are absent from this package, please don't hesitate to open an issue or a pr.

Same applies to bugs.

## License

MIT.
