package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetStatusDocuments(t *testing.T) {
	documents, err := c.GetStatusDocuments(novapost.Documents{
		Documents: []novapost.StatusDocumentRequest{
			{},
		},
	})
	if err != nil {
		t.Error(err)
	}

	if !documents.Success {
		t.Log("No status documents found", documents)
	}
}
