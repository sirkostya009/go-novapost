package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestCheckPossibilityCreateReturn(t *testing.T) {
	possibility, err := c.CheckPossibilityCreateReturn(novapost.Number{
		Number: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Error("Error checking possibility create return", possibility)
	}
}

func TestGetReturnReasons(t *testing.T) {
	reasons, err := c.GetReturnReasons()
	if err != nil {
		t.Error(err)
	}

	if !reasons.Success {
		t.Error("No return reasons found", reasons)
	}
}

func TestGetReturnReasonsSubtypes(t *testing.T) {
	reasons, err := c.GetReturnReasonSubtypes(novapost.ReasonRef{
		ReasonRef: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !reasons.Success {
		t.Error("No return reasons subtypes found", reasons)
	}
}

func TestGetReturnOrdersList(t *testing.T) {
	returnOrders, err := c.GetReturnOrdersList(novapost.ReturnOrderRequest{
		Number:    "102-00003168",
		Ref:       "00000000-0000-0000-0000-000000000000",
		BeginDate: "12/10/15 10:33",
		EndDate:   "12/10/15 10:33",
		Page:      1,
		Limit:     150,
	})
	if err != nil {
		t.Error(err)
	}

	if !returnOrders.Success {
		t.Error("No return orders found", returnOrders)
	}
}

func TestDeleteAdditionalService(t *testing.T) {
	additionalService, err := c.DeleteReturnOrder(novapost.Ref{
		Ref: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !additionalService.Success {
		t.Error("Error deleting additional service", additionalService)
	}
}

func TestCheckPossibilityChangeEW(t *testing.T) {
	possibility, err := c.CheckPossibilityChangeEW(novapost.IntDocNumber{
		IntDocNumber: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Error("Error checking possibility change EW", possibility)
	}
}

func TestGetChangeEWOrdersList(t *testing.T) {
	orders, err := c.GetChangeEWOrdersList(novapost.ChangeEWRequest{
		Number:    "00000000000000",
		Ref:       "00000000-0000-0000-0000-000000000000",
		BeginDate: "12/10/15 10:33",
		EndDate:   "12/10/15 10:33",
		Page:      1,
		Limit:     150,
	})
	if err != nil {
		t.Error(err)
	}

	if !orders.Success {
		t.Error("No change EW orders found", orders)
	}
}

func TestCheckPossibilityForRedirecting(t *testing.T) {
	possibility, err := c.CheckPossibilityForRedirecting(novapost.Number{
		Number: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Error("Error checking possibility for redirecting", possibility)
	}
}

func TestGetRedirectionOrdersList(t *testing.T) {
	orders, err := c.GetRedirectionOrdersList(novapost.RedirectionRequest{
		Number:    "00000000000000",
		Ref:       "00000000-0000-0000-0000-000000000000",
		BeginDate: "12/10/15 10:33",
		EndDate:   "12/10/15 10:33",
		Page:      1,
		Limit:     150,
	})
	if err != nil {
		t.Error(err)
	}

	if !orders.Success {
		t.Error("No redirection orders found", orders)
	}
}
