package novapost_test

import (
	. "github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetDocumentPrice(t *testing.T) {
	prices, err := c.GetDocumentPrice(DocumentPriceRequest{
		CitySender:    "",
		CityRecipient: "",
		Weight:        0,
		ServiceType:   CargoType,
		Cost:          0,
		CargoType:     "",
		SeatsAmount:   0,
		RedeliveryCalculate: RedeliveryCalculate{
			CargoType: CargoType,
			Amount:    0,
		},
		PackCount: 0,
		PackRef:   "",
		Amount:    0,
		CargoDetails: []CargoDetail{
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
	dates, err := c.GetDocumentDeliveryDate(DocumentDeliveryDateRequest{
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
	documents, err := c.GetDocumentList(DocumentListRequest{
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
	report, err := c.GenerateReport(ReportRequest{
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
