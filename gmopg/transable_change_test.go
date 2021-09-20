package gmopg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChangeTran(t *testing.T) {
	type MockResponse struct {
		path, contenttype, body string
	}

	response := &MockResponse{
		path:        "/payment/ChangeTran.json",
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
		if _, err := io.WriteString(w, response.body); err != nil {
			t.Fatal(err)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	setup(server.URL)

	args := &ChangeTranArgs{
		ShopID:     "tshop000000",
		ShopPass:   "5tj78kom",
		AccessID:   "accessid1347e90cdef806b39bd28705",
		AccessPass: "accesspasse14ed03e8cf807a6053813",
		JobCd:      JCheck,
		Amount:     2000,
	}

	if _, err := gmopg.ChangeTran(args); err != nil {
		t.Fatal(err)
	}
}
