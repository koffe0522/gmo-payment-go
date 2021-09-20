package gmopg

import (
	"encoding/json"
)

// ChangeTranArgs is
type ChangeTranArgs struct {
	ShopID     string `json:"shopID"`
	ShopPass   string `json:"shopPass"`
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
	JobCd      JobCd  `json:"jobCd"`
	Amount     int    `json:"amount"`
	Tax        *int   `json:"tax,omitempty"`
}

// ChangeTranResult is
type ChangeTranResult struct {
	Error      ErrorResults
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
	Forward    string `json:"forward"`
	Approve    string `json:"approve"`
	TranID     string `json:"tranID"`
	TranDate   string `json:"tranDate"`
}

// ChangeTran is
func (g GMOPG) ChangeTran(args *ChangeTranArgs) (*ChangeTranResult, error) {
	paramsJSON, _ := json.Marshal(args)
	resp, err := g.client.Do("/payment/ChangeTran.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok ChangeTranResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &ChangeTranResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
