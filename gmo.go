package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ErrorResponse is
type ErrorResponse struct {
	ErrCode string `json:"errCode"`
	ErrInfo string `json:"errInfo"`
}

// OkResponse is
type OkResponse struct {
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
}

// GMOPG is
type GMOPG interface {
	EntryTran(args *EntryTranArgs) (*OkResponse, []ErrorResponse, error)
}

type gMOPG struct {
	baseURL  string
	siteID   string
	sitePass string
	shopID   string
	shopPass string
}

// GmopgConfig is
type GmopgConfig struct {
	baseURL  *string
	siteID   string
	sitePass string
	shopID   string
	shopPass string
}

// EntryTranArgs is
type EntryTranArgs struct {
	OrderID      string
	JobCd        string // Enum対応
	ItemCode     *string
	Amount       int
	Tax          *int
	TdFlag       *string
	TdTenantName *string
}

type entryTranParams struct {
	ShopID   string `json:"shopID"`
	ShopPass string `json:"shopPass"`
	OrderID  string `json:"orderID"`
	JobCd    string `json:"jobCd"`
	Amount   string `json:"amount"`
}

// NewGMOPG is
func NewGMOPG(c *GmopgConfig) GMOPG {
	gmopg := gMOPG{}
	gmopg.siteID = c.siteID
	gmopg.shopID = c.shopID
	gmopg.sitePass = c.sitePass
	gmopg.shopPass = c.shopPass
	fmt.Println(c.baseURL)
	if c.baseURL == nil {
		gmopg.baseURL = "https://pt01.mul-pay.jp"
	} else {
		gmopg.baseURL = *c.baseURL
	}

	return gmopg
}

func (s gMOPG) EntryTran(args *EntryTranArgs) (*OkResponse, []ErrorResponse, error) {
	path := s.baseURL + "/payment/EntryTran.json"
	paramsJSON, _ := json.Marshal(&entryTranParams{
		ShopID:   s.shopID,
		ShopPass: s.shopPass,
		OrderID:  args.OrderID,
		JobCd:    args.JobCd,
		Amount:   strconv.Itoa(args.Amount),
	})

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "GMO PG Client: Unofficial")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil, nil, err
	}

	if res.StatusCode == 200 {
		var ok OkResponse
		err = json.Unmarshal(body, &ok)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}

		return &ok, nil, nil
	}

	var ng []ErrorResponse
	err = json.Unmarshal(body, &ng)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return nil, ng, nil
}
