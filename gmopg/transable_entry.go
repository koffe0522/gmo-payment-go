package gmopg

import (
	"encoding/json"
)

// EntryTranResult is
type EntryTranResult struct {
	Error      ErrorResults
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
}

// EntryTranArgs is
type EntryTranArgs struct {
	OrderID      string  `json:"orderID"`
	JobCd        JobCd   `json:"jobCd"`
	ItemCode     *string `json:"itemCode,omitempty"`
	Amount       int     `json:"amount"`
	Tax          *int    `json:"tax,omitempty"`
	TdFlag       *string `json:"tdFlag,omitempty"`
	TdTenantName *string `json:"tdTenantName,omitempty"`
}

// EntryTran is
func (g *GMOPG) EntryTran(args *EntryTranArgs) (*EntryTranResult, error) {
	param := struct {
		EntryTranArgs
		ShopID   string `json:"shopID"`
		ShopPass string `json:"shopPass"`
	}{
		EntryTranArgs: *args,
		ShopID:        g.shopID,
		ShopPass:      g.shopPass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/EntryTran.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok EntryTranResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &EntryTranResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
