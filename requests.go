package novapost

import (
	"bytes"
	"net/http"
)

const (
	AddressModel           = "Address"
	AddressGeneralModel    = "AddressGeneral"
	CounterpartyModel      = "Counterparty"
	ContactPersonModel     = "ContactPerson"
	ScanSheetModel         = "ScanSheet"
	CommonModel            = "Common"
	CommonGeneralModel     = "CommonGeneral"
	AdditionalServiceModel = "AdditionalService"
	InternetDocumentModel  = "InternetDocument"
	TrackingDocumentModel  = "TrackingDocument"
)

func request[Res any](c Client, model, method string, props any) (response Response[Res], err error) {
	b, err := c.marshaler(Request{c.apiKey, model, method, props})
	if err != nil {
		return
	}
	r, err := http.NewRequest(http.MethodGet, c.url, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	res, err := c.http.Do(r)
	if err != nil {
		return
	}
	err = c.decoderConstructor(res.Body).Decode(&response)
	if err != nil {
		return
	}
	err = res.Body.Close()
	return
}

func (c Client) RawRequest(model, method string, props any) (Response[map[string]any], error) {
	return request[map[string]any](c, model, method, props)
}
