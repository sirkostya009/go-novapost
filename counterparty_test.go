package novapost_test

import (
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

const CounterpartyRef = "e718a680-4b3a-11e4-ab6d-005056801329"

func TestClient_GetCounterpartyAddresses(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	addresses, err := c.GetCounterpartyAddresses(CounterpartyAddressRequest{
		Ref:                  CounterpartyRef,
		CounterpartyProperty: "Sender",
	})
	if err != nil {
		t.Error(err)
	}

	if !addresses.Success {
		t.Log("No addresses found", addresses)
	}
}

func TestClient_GetCounterpartyOptions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	options, err := c.GetCounterpartyOptions(Ref{
		Ref: CounterpartyRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !options.Success {
		t.Log("No options found", options)
	}
}

func TestClient_GetCounterpartyContactPersons(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	contactPersons, err := c.GetCounterpartyContactPersons(CounterpartyContactPersonRequest{
		Ref:  CounterpartyRef,
		Page: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !contactPersons.Success {
		t.Log("No contact persons found", contactPersons)
	}
}

func TestClient_GetCounterparties(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	counterparties, err := c.GetCounterparties(CounterpartiesRequest{
		CounterpartyProperty: "Sender",
		Page:                 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !counterparties.Success {
		t.Log("No counterparties found", counterparties)
	}
}
