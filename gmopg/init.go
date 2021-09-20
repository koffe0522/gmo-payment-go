package gmopg

import "github.com/koffe0522/gmo-payment-go/httpc"

// GMOPG is
type GMOPG struct {
	baseURL  string
	siteID   string
	sitePass string
	shopID   string
	shopPass string
	client   httpc.IClinet
}

// Config is
type Config struct {
	BaseURL  *string
	SiteID   string
	SitePass string
	ShopID   string
	ShopPass string
}

// Init is
func Init(c *Config) GMOPG {
	gmopg := GMOPG{}
	gmopg.siteID = c.SiteID
	gmopg.shopID = c.ShopID
	gmopg.sitePass = c.SitePass
	gmopg.shopPass = c.ShopPass
	if c.BaseURL == nil {
		gmopg.baseURL = "https://pt01.mul-pay.jp"
	} else {
		gmopg.baseURL = *c.BaseURL
	}

	gmopg.client = &httpc.Client{
		BaseURL: gmopg.baseURL,
	}

	return gmopg
}
