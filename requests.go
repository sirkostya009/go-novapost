package novapost

import (
	"bytes"
	"io"
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

func RawRequest[Res any](c *Client, model, method string, props any) (*Response[Res], error) {
	b, err := c.Marshaler(Request{c.apiKey, model, method, props})
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Post(c.Url, "", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resp := &Response[Res]{}
	err = c.Unmarshaler(b, resp)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()
	return resp, nil
}
