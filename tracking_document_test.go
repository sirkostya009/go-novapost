package novapost_test

import (
	. "github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetStatusDocuments(t *testing.T) {
	documents, err := c.GetStatusDocuments(Documents{
		Documents: []StatusDocumentRequest{
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
