package novapost_test

import (
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

func TestClient_InsertDocuments(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	documents, err := c.InsertDocuments(InsertDocumentsRequest{
		Ref:          "00000000-0000-0000-0000-000000000000",
		DocumentRefs: []string{"00000000-0000-0000-0000-000000000000"},
		Date:         "01.01.2022",
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Log("Error inserting documents", documents)
	}
}

func TestClient_GetScanSheet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	scanSheet, err := c.GetScanSheet(ScanSheetRequest{
		Ref:             "00000000-0000-0000-0000-000000000000",
		CounterpartyRef: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !scanSheet.Success {
		t.Log("Error getting scan sheet", scanSheet)
	}
}

func TestClient_GetScanSheetList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	scanSheets, err := c.GetScanSheetList()
	if err != nil {
		t.Error(err)
	}

	if !scanSheets.Success {
		t.Log("Error getting scan sheet list", scanSheets)
	}
}

func TestClient_DeleteScanSheet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	scanSheet, err := c.DeleteScanSheet(ScanSheetRefs{
		ScanSheetRefs: []string{"00000000-0000-0000-0000-000000000000"},
	})
	if err != nil {
		t.Error(err)
	}

	if !scanSheet.Success {
		t.Log("Error deleting scan sheet", scanSheet)
	}
}

func TestClient_RemoveDocuments(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	documents, err := c.RemoveDocuments(RemoveDocumentsRequest{
		DocumentRefs: []string{"00000000-0000-0000-0000-000000000000"},
		Ref:          "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Log("Error removing documents", documents)
	}
}
