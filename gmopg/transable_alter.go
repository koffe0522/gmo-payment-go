package gmopg

import (
	"encoding/json"
)

// AlterTranArgs is
type AlterTranArgs struct {
	ShopID     string `json:"shopID"`
	ShopPass   string `json:"shopPass"`
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
	JobCd      JobCd  `json:"jobCd"`
	Amount     int    `json:"amount"`
	Tax        *int   `json:"tax,omitempty"`
	Method     Method `json:"method"`
	PayTimes   *int   `json:"payTimes,omitempty"`
}

// AlterTranResult is
type AlterTranResult struct {
	Error      ErrorResults
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
	Forward    string `json:"forward"`
	Approve    string `json:"approve"`
	TranID     string `json:"tranID"`
	TranDate   string `json:"tranDate"`
}

// AlterTran is
func (g GMOPG) AlterTran(args *AlterTranArgs) (*AlterTranResult, error) {
	paramsJSON, _ := json.Marshal(args)
	resp, err := g.client.Do("/payment/AlterTran.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok AlterTranResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &AlterTranResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
