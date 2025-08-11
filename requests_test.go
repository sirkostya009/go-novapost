package novapost_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirkostya009/go-novapost"
)

func TestRawRequest(t *testing.T) {
	const method = "method"
	const model = "model"
	const property = "blah"
	type testT = int
	type testData struct{ String string }

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			var res novapost.Request
			err := json.NewDecoder(r.Body).Decode(&res)
			assertNoError(t, err)

			assertProperty(t, "method", res.CalledMethod, method)
			assertProperty(t, "model", res.ModelName, model)
			assertProperty(t, "property", res.MethodProperties.(map[string]any)["String"].(string), property)

			err = json.NewEncoder(w).Encode(novapost.Response[testT]{
				Success: true,
				Data:    []testT{1},
			})
			assertNoError(t, err)
		}
	}))
	defer serv.Close()

	c := novapost.NewClient("",
		novapost.WithHTTPClient(serv.Client()),
		novapost.WithMarshaler(json.Marshal),
		novapost.WithUnmarshaler(json.Unmarshal),
		novapost.WithURL(serv.URL),
	)

	res, err := novapost.RawRequest[testT](c, model, method, testData{String: property})

	assertNoError(t, err)

	if len(res.Data) != 1 || res.Data[0] != 1 {
		t.Error("unexpected return data")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func assertProperty[T comparable](t *testing.T, name string, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("%s did not match, got %v, want %v", name, got, want)
	}
}
