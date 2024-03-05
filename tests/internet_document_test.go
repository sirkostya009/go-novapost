package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetDocumentPrice(t *testing.T) {
	prices, err := c.GetDocumentPrice(novapost.DocumentPriceRequest{
		CitySender:    "",
		CityRecipient: "",
		Weight:        0,
		ServiceType:   novapost.CargoType,
		Cost:          0,
		CargoType:     "",
		SeatsAmount:   0,
		RedeliveryCalculate: novapost.RedeliveryCalculate{
			CargoType: novapost.CargoType,
			Amount:    0,
		},
		PackCount: 0,
		PackRef:   "",
		Amount:    0,
		CargoDetails: []novapost.CargoDetail{
			{},
		},
		CargoDescription: "",
	})
	if err != nil {
		t.Error(err)
	}

	if !prices.Success {
		t.Log("No document prices found", prices)
	}
}

func TestGetDocumentDeliveryDate(t *testing.T) {
	dates, err := c.GetDocumentDeliveryDate(novapost.DocumentDeliveryDateRequest{
		ServiceType:   "WarehouseWarehouse",
		CitySender:    KyivCityRef,
		CityRecipient: KyivCityRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !dates.Success {
		t.Log("No document delivery dates found", dates)
	}
}

func TestGetDocumentList(t *testing.T) {
	documents, err := c.GetDocumentList(novapost.DocumentListRequest{
		DateTime: "01.01.2026",
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Log("No documents found", documents)
	}
}

func TestGenerateReport(t *testing.T) {
	report, err := c.GenerateReport(novapost.ReportRequest{
		DocumentRefs: []string{"00000000-0000-0000-0000-000000000000"},
		DateTime:     "01.01.2026",
		Type:         "csv",
	})
	if err != nil {
		t.Error(err)
	}

	if !report.Success {
		t.Log("Error generating report", report)
	}
}
