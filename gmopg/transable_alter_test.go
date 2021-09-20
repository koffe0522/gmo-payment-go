package gmopg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAlterTran(t *testing.T) {
	type MockResponse struct {
		path, query, contenttype, body string
	}

	response := &MockResponse{
		path:        "/payment/AlterTran.json",
		contenttype: "application/json",
		body: `{
            "accessID":"accessid1347e90cdef806b39bd28705",
            "accessPass":"accesspasse14ed03e8cf807a6053813",
            "forward":"2b 00000",
            "approve":"1234567",
            "tranID":"1111111111",
            "tranDate":"20200920213000"}`}
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Send response.
		w.Header().Set("Content-Type", response.contenttype)
		io.WriteString(w, response.body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	setup(server.URL)

	args := &AlterTranArgs{
		ShopID:     "tshop000000",
		ShopPass:   "5tj78kom",
		AccessID:   "accessid1347e90cdef806b39bd28705",
		AccessPass: "accesspasse14ed03e8cf807a6053813",
		JobCd:      JCheck,
	}

	if _, err := gmopg.AlterTran(args); err != nil {
		t.Fatal(err)
	}
}
