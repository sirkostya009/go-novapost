package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestInsertDocuments(t *testing.T) {
	documents, err := c.InsertDocuments(novapost.InsertDocumentsRequest{
		Ref:          "00000000-0000-0000-0000-000000000000",
		DocumentRefs: []string{"00000000-0000-0000-0000-000000000000"},
		Date:         "01.01.2022",
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Error("Error inserting documents", documents)
	}
}

func TestGetScanSheet(t *testing.T) {
	scanSheet, err := c.GetScanSheet(novapost.ScanSheetRequest{
		Ref:             "00000000-0000-0000-0000-000000000000",
		CounterpartyRef: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !scanSheet.Success {
		t.Error("Error getting scan sheet", scanSheet)
	}
}

func TestGetScanSheetList(t *testing.T) {
	scanSheets, err := c.GetScanSheetList()
	if err != nil {
		t.Error(err)
	}

	if !scanSheets.Success {
		t.Error("Error getting scan sheet list", scanSheets)
	}
}

func TestDeleteScanSheet(t *testing.T) {
	scanSheet, err := c.DeleteScanSheet(novapost.ScanSheetRefs{
		ScanSheetRefs: []string{"00000000-0000-0000-0000-000000000000"},
	})
	if err != nil {
		t.Error(err)
	}

	if !scanSheet.Success {
		t.Error("Error deleting scan sheet", scanSheet)
	}
}

func TestRemoveDocuments(t *testing.T) {
	documents, err := c.RemoveDocuments(novapost.RemoveDocumentsRequest{
		DocumentRefs: []string{"00000000-0000-0000-0000-000000000000"},
		Ref:          "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Error("Error removing documents", documents)
	}
}
