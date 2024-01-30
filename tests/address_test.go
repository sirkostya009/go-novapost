package tests

import (
	"github.com/sirkostya009/go-novapost"
	"os"
	"testing"
)

const (
	KyivSettlementRef = "0db2df4b-4b3a-11e4-ab6d-005056801329"
	KyivCityRef       = "8d5a980d-391c-11dd-90d9-001a92567626"
)

var c = novapost.NewClient(os.Getenv("NOVA_POST_API_KEY"))

func TestSearchSettlements(t *testing.T) {
	settlements, err := c.SearchSettlements(novapost.SettlementRequest{
		CityName: "Київ",
		Limit:    150,
		Page:     1,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlements.Success {
		t.Error("No settlements found", settlements)
	}
}

func TestSearchSettlementStreets(t *testing.T) {
	streets, err := c.SearchSettlementStreets(novapost.SettlementStreetRequest{
		StreetName:    "Хрещатик",
		SettlementRef: KyivSettlementRef,
		Page:          1,
	})
	if err != nil {
		t.Error(err)
	}

	if !streets.Success {
		t.Error("No streets found", streets)
	}
}

func TestGetSettlements(t *testing.T) {
	settlements, err := c.GetSettlements(novapost.GetSettlementsRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlements.Success {
		t.Error("No settlements found", settlements)
	}
}

func TestGetCities(t *testing.T) {
	cities, err := c.GetCities(novapost.CityRequest{
		FindByString: "Київ",
		Limit:        150,
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !cities.Success {
		t.Error("No cities found", cities)
	}
}

func TestGetAreas(t *testing.T) {
	areas, err := c.GetAreas()
	if err != nil {
		t.Error(err)
	}

	if !areas.Success {
		t.Error("No areas found", areas)
	}
}

func TestGetWarehouses(t *testing.T) {
	warehouses, err := c.GetWarehouses(novapost.WarehouseRequest{
		CityName: "Київ",
		Limit:    150,
		Page:     1,
	})
	if err != nil {
		t.Error(err)
	}

	if !warehouses.Success {
		t.Error("No warehouses found", warehouses)
	}
}

func TestGetWarehouseTypes(t *testing.T) {
	warehouseTypes, err := c.GetWarehouseTypes()
	if err != nil {
		t.Error(err)
	}

	if !warehouseTypes.Success {
		t.Error("No warehouse types found", warehouseTypes)
	}
}

func TestGetStreet(t *testing.T) {
	street, err := c.GetStreet(novapost.StreetRequest{
		CityRef: KyivCityRef,
		Limit:   150,
		Page:    1,
	})
	if err != nil {
		t.Error(err)
	}

	if !street.Success {
		t.Error("No streets found", street)
	}
}

func TestGetSettlementCountryRegion(t *testing.T) {
	settlementCountryRegion, err := c.GetSettlementCountryRegion(novapost.AreaRef{
		AreaRef: "7150813e-9b87-11de-822f-000c2965ae0e",
	})
	if err != nil {
		t.Error(err)
	}

	if !settlementCountryRegion.Success {
		t.Error("No settlement country regions found", settlementCountryRegion)
	}
}

func TestGetSettlementAreas(t *testing.T) {
	settlementAreas, err := c.GetSettlementAreas(novapost.Ref{
		Ref: KyivSettlementRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !settlementAreas.Success {
		t.Error("No settlement areas found", settlementAreas)
	}
}
