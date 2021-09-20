package gmopg

import "github.com/koffe0522/gmo-payment-go/httpc"

var (
	gmopg GMOPG
)

type MockClient struct {
	BaseURL string
}

func setup(baseURL string) {
	gmopg = Init(&Config{
		BaseURL:  &baseURL,
		SiteID:   "siteid",
		ShopID:   "shopid",
		SitePass: "sitepass",
		ShopPass: "shoppass",
	})

	// gmopg.client = &MockClient{
	// 	BaseURL: gmopg.baseURL,
	// }
	gmopg.client = &httpc.Client{
		BaseURL: gmopg.baseURL,
	}
}
