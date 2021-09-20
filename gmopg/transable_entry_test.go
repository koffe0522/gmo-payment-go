package gmopg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEntryTran(t *testing.T) {
	type MockResponse struct {
		path, query, contenttype, body string
	}

	response := &MockResponse{
		path:        "/payment/EntryTran.json",
		contenttype: "application/json",
		body:        `{"accessID": "accessid1347e90cdef806b39bd28705","accessPass": "accesspasse14ed03e8cf807a6053813"}`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Send response.
		w.Header().Set("Content-Type", response.contenttype)
		if _, err := io.WriteString(w, response.body); err != nil {
			t.Fatal(err)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	setup(server.URL)

	orderID := "ORDER-vqCQHbDpObQvTDjQ8ZJ8"
	amount := 33000
	args := &EntryTranArgs{
		OrderID: orderID,
		JobCd:   JAuth,
		Amount:  amount,
	}

	if _, err := gmopg.EntryTran(args); err != nil {
		t.Fatal(err)
	}
}
