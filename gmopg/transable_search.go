package gmopg

import (
	"encoding/json"
)

// SearchTradeArgs is
type SearchTradeArgs struct {
	ShopID   string `json:"shopID"`
	ShopPass string `json:"shopPass"`
	OrderID  string `json:"orderID"`
}

// SearchTradeResult is
type SearchTradeResult struct {
	Error        ErrorResults
	OrderID      string `json:"orderID"`
	Status       Status `json:"status"`
	ProcessDate  string `json:"processDate"`
	JobCd        JobCd  `json:"jobCd"`
	AccessID     string `json:"accessID"`
	AccessPass   string `json:"accessPass"`
	ItemCode     string `json:"itemCode"`
	Amount       string `json:"amount"`
	Tax          string `json:"tax"`
	SiteID       string `json:"siteID"`
	MemberID     string `json:"memberID"`
	CardNo       string `json:"cardNo"`
	Expire       string `json:"expire"`
	Method       Method `json:"method"`
	PayTimes     string `json:"payTimes"`
	Forward      string `json:"forward"`
	TranID       string `json:"tranID"`
	Approve      string `json:"approve"`
	ClientField1 string `json:"clientField1"`
	ClientField2 string `json:"clientField2"`
	ClientField3 string `json:"clientField3"`
}

// SearchTrade is
func (g GMOPG) SearchTrade(args *SearchTradeArgs) (*SearchTradeResult, error) {
	paramsJSON, _ := json.Marshal(args)
	resp, err := g.client.Do("/payment/SearchTrade.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SearchTradeResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SearchTradeResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
