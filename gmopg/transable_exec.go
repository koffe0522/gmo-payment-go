package gmopg

import (
	"encoding/json"
)

// ExecTranArgs is
type ExecTranArgs struct {
	AccessID     string   `json:"accessID"`
	AccessPass   string   `json:"accessPass"`
	OrderID      string   `json:"orderID"`
	Method       Method   `json:"method"`
	PayTimes     *int     `json:"payTimes,omitempty"`
	CardNo       *string  `json:"cardNo,omitempty"`
	Expire       *string  `json:"expired,omitempty"`
	SecurityCode *string  `json:"securityCode,omitempty"`
	Token        *string  `json:"token,omitempty"`
	Pin          *string  `json:"pin,omitempty"`
	SiteID       *string  `json:"siteID,omitempty"`
	SitePass     *string  `json:"sitePass,omitempty"`
	MemberID     *string  `json:"memberID,omitempty"`
	SeqMode      *SeqMode `json:"seqMode,omitempty"`
	CardSeq      *int     `json:"cardSeq,omitempty"`
	CardPass     *int     `json:"cardPass,omitempty"`
	ClientField1 *string  `json:"clientField1,omitempty"`
	ClientField2 *string  `json:"clientField2,omitempty"`
	ClientField3 *string  `json:"clientField3,omitempty"`
}

// ExecTranResult is
type ExecTranResult struct {
	Error        ErrorResults
	Acs          string `json:"acs"`
	OrderID      string `json:"orderID"`
	Forward      string `json:"forward"`
	Method       string `json:"method"`
	PayTimes     string `json:"payTimes"`
	Approve      string `json:"approve"`
	TranID       string `json:"tranID"`
	TranDate     string `json:"tranDate"`
	CheckString  string `json:"checkString"`
	ClientField1 string `json:"clientField1"`
	ClientField2 string `json:"clientField2"`
	ClientField3 string `json:"clientField3"`
}

// ExecTran is
func (g GMOPG) ExecTran(args *ExecTranArgs) (*ExecTranResult, error) {
	paramsJSON, _ := json.Marshal(args)
	resp, err := g.client.Do("/payment/ExecTran.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok ExecTranResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &ExecTranResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
