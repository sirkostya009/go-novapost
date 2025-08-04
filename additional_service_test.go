package novapost_test

import (
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

func TestClient_CheckPossibilityCreateReturn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	possibility, err := c.CheckPossibilityCreateReturn(Number{
		Number: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Log("Error checking possibility create return", possibility)
	}
}

func TestClient_GetReturnReasons(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	reasons, err := c.GetReturnReasons()
	if err != nil {
		t.Error(err)
	}

	if !reasons.Success {
		t.Log("No return reasons found", reasons)
	}
}

func TestClient_GetReturnReasonsSubtypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	reasons, err := c.GetReturnReasonSubtypes(ReasonRef{
		ReasonRef: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !reasons.Success {
		t.Log("No return reasons subtypes found", reasons)
	}
}

func TestClient_GetReturnOrdersList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	returnOrders, err := c.GetReturnOrdersList(ReturnOrderRequest{
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
		t.Log("No return orders found", returnOrders)
	}
}

func TestClient_DeleteAdditionalService(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	additionalService, err := c.DeleteReturnOrder(Ref{
		Ref: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !additionalService.Success {
		t.Log("Error deleting additional service", additionalService)
	}
}

func TestClient_CheckPossibilityChangeEW(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	possibility, err := c.CheckPossibilityChangeEW(IntDocNumber{
		IntDocNumber: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Log("Error checking possibility change EW", possibility)
	}
}

func TestClient_GetChangeEWOrdersList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	orders, err := c.GetChangeEWOrdersList(ChangeEWRequest{
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
		t.Log("No change EW orders found", orders)
	}
}

func TestClient_CheckPossibilityForRedirecting(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	possibility, err := c.CheckPossibilityForRedirecting(Number{
		Number: "00000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !possibility.Success {
		t.Log("Error checking possibility for redirecting", possibility)
	}
}

func TestClient_GetRedirectionOrdersList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	orders, err := c.GetRedirectionOrdersList(RedirectionRequest{
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
		t.Log("No redirection orders found", orders)
	}
}
