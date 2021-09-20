package gmopg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecTran(t *testing.T) {
	type MockResponse struct {
		path, query, contenttype, body string
	}

	response := &MockResponse{
		path:        "/payment/ExecTran.json",
		contenttype: "application/json",
		body: `{
            "accessID": "accessid1347e90cdef806b39bd28705",
            "accessPass": "accesspasse14ed03e8cf807a6053813",
            "orderID":"ORDER-vqCQHbDpObQvTDjQ8ZJ8",
            "clientField1":"項目１",
            "clientField2":"項目２",
            "clientField3":"項目３",
            "clientFieldFlag":"1",
            "token":"Lg9sRgo5nx6yfefJ51z8bj/1VdNFAaCZYWZ+qLKJyqWwBS7yYvxSiC0zeMVH+O4F"
        }`}
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Check request.
		if g, w := r.URL.Path, response.path; g != w {
			t.Errorf("request got path %s, want %s", g, w)
		}

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
	accessID := "accessid1347e90cdef806b39bd28705"
	accessPass := "accesspasse14ed03e8cf807a6053813"
	cardNo := "3111111111111111"
	expire := "2024"
	securityCode := "123"
	args := &ExecTranArgs{
		AccessID:     accessID,
		AccessPass:   accessPass,
		OrderID:      orderID,
		Method:       Lump,
		CardNo:       &cardNo,
		Expire:       &expire,
		SecurityCode: &securityCode,
	}

	if _, err := gmopg.ExecTran(args); err != nil {
		t.Fatal(err)
	}
}
