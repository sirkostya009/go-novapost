package novapost_test

import (
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

func TestClient_GetStatusDocuments(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
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
