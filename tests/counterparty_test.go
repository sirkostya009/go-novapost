package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

const CounterpartyRef = "e718a680-4b3a-11e4-ab6d-005056801329"

func TestGetCounterpartyAddresses(t *testing.T) {
	addresses, err := c.GetCounterpartyAddresses(novapost.CounterpartyAddressRequest{
		Ref:                  CounterpartyRef,
		CounterpartyProperty: "Sender",
	})
	if err != nil {
		t.Error(err)
	}

	if !addresses.Success {
		t.Error("No addresses found", addresses)
	}
}

func TestGetCounterpartyOptions(t *testing.T) {
	options, err := c.GetCounterpartyOptions(novapost.Ref{
		Ref: CounterpartyRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !options.Success {
		t.Error("No options found", options)
	}
}

func TestGetCounterpartyContactPersons(t *testing.T) {
	contactPersons, err := c.GetCounterpartyContactPersons(novapost.CounterpartyContactPersonRequest{
		Ref:  CounterpartyRef,
		Page: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !contactPersons.Success {
		t.Error("No contact persons found", contactPersons)
	}
}

func TestGetCounterparties(t *testing.T) {
	counterparties, err := c.GetCounterparties(novapost.CounterpartiesRequest{
		CounterpartyProperty: "Sender",
		Page:                 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !counterparties.Success {
		t.Error("No counterparties found", counterparties)
	}
}
