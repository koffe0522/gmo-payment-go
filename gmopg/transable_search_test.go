package gmopg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchTrade(t *testing.T) {
	type MockResponse struct {
		path, contenttype, body string
	}

	response := &MockResponse{
		path:        "/payment/SearchTrade.json",
		contenttype: "application/json",
		body: `{
            "status":"AUTH",
            "processDate":"20190213200436",
            "jobCd":"AUTH",
            "accessID": "accessid1347e90cdef806b39bd28705",
            "accessPass": "accesspasse14ed03e8cf807a6053813",
            "itemCode":"abc1234",
            "amount":"2000",
            "tax":"200",
            "siteID":"tsite0000000",
            "memberID":"member01",
            "cardNo":"************0000",
            "expire":"2409",
            "method":"1",
            "payTimes":"3",
            "forward":"2FF22F2",
            "TtanID":"1111111111",
            "approve":"123456a",
            "clientField1":"",
            "clientField2":"",
            "clien tField3":""
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

	args := &SearchTradeArgs{
		OrderID: orderID,
	}

	if _, err := gmopg.SearchTrade(args); err != nil {
		t.Fatal(err)
	}
}
