package novapost_test

import (
	"os"
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

const (
	KyivSettlementRef = "0db2df4b-4b3a-11e4-ab6d-005056801329"
	KyivCityRef       = "8d5a980d-391c-11dd-90d9-001a92567626"
)

// newTestClient creates a new client for testing
func newTestClient() *Client {
	return NewClient(os.Getenv("NOVA_POST_API_KEY"))
}

func TestClient_SearchSettlements(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	settlements, err := c.SearchSettlements(SettlementRequest{
		CityName: "Київ",
		Limit:    150,
		Page:     1,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlements.Success {
		t.Log("No settlements found", settlements)
	}
}

func TestClient_SearchSettlementStreets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	streets, err := c.SearchSettlementStreets(SettlementStreetRequest{
		StreetName:    "Хрещатик",
		SettlementRef: KyivSettlementRef,
		Page:          1,
	})
	if err != nil {
		t.Error(err)
	}

	if !streets.Success {
		t.Log("No streets found", streets)
	}
}

func TestClient_GetSettlements(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	settlements, err := c.GetSettlements(GetSettlementsRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlements.Success {
		t.Log("No settlements found", settlements)
	}
}

func TestClient_GetCities(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	cities, err := c.GetCities(CityRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !cities.Success {
		t.Log("No cities found", cities)
	}
}

func TestClient_GetAreas(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	areas, err := c.GetAreas()
	if err != nil {
		t.Error(err)
	}

	if !areas.Success {
		t.Log("No areas found", areas)
	}
}

func TestClient_GetWarehouses(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	warehouses, err := c.GetWarehouses(WarehouseRequest{
		CityName: "Київ",
		Limit:    150,
		Page:     1,
	})
	if err != nil {
		t.Error(err)
	}

	if !warehouses.Success {
		t.Log("No warehouses found", warehouses)
	}
}

func TestClient_GetWarehouseTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	warehouseTypes, err := c.GetWarehouseTypes()
	if err != nil {
		t.Error(err)
	}

	if !warehouseTypes.Success {
		t.Log("No warehouse types found", warehouseTypes)
	}
}

func TestClient_GetStreet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	street, err := c.GetStreet(StreetRequest{
		CityRef: KyivCityRef,
		Limit:   150,
		Page:    1,
	})
	if err != nil {
		t.Error(err)
	}

	if !street.Success {
		t.Log("No streets found", street)
	}
}

func TestClient_GetSettlementCountryRegion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	settlementCountryRegion, err := c.GetSettlementCountryRegion(AreaRef{
		AreaRef: "7150813e-9b87-11de-822f-000c2965ae0e",
	})
	if err != nil {
		t.Error(err)
	}

	if !settlementCountryRegion.Success {
		t.Log("No settlement country regions found", settlementCountryRegion)
	}
}

func TestClient_GetSettlementAreas(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	settlementAreas, err := c.GetSettlementAreas(Ref{
		Ref: KyivSettlementRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlementAreas.Success {
		t.Log("No settlement areas found", settlementAreas)
	}
}
