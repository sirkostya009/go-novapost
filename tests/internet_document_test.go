package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetDocumentPrice(t *testing.T) {
	prices, err := c.GetDocumentPrice(novapost.DocumentPriceRequest{
		CitySender:    "",
		CityRecipient: "",
		Weight:        "",
		ServiceType:   "",
		Cost:          "",
		CargoType:     "",
		SeatsAmount:   "",
		RedeliveryCalculate: novapost.RedeliveryCalculate{
			CargoType: "",
			Amount:    "",
		},
		PackCount: "",
		PackRef:   "",
		Amount:    "",
		CargoDetails: []novapost.CargoDetail{
			{},
		},
		CargoDescription: "",
	})
	if err != nil {
		t.Error(err)
	}

	if !prices.Success {
		t.Error("No document prices found", prices)
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
		t.Error("No document delivery dates found", dates)
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
		t.Error("No documents found", documents)
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
		t.Error("Error generating report", report)
	}
}
