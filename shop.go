package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ShopAPI is
type ShopAPI interface {
	EntryTran(args *EntryTranArgs) (*http.Response, error)
}

// shopAPI is Structure that handles payments such as transaction registration and payment status changes
type shopAPI struct {
	host     string
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

// NewShopAPI is
func NewShopAPI(host string, shopID string, shopPass string) ShopAPI {

	s := shopAPI{}
	s.host = host
	s.shopID = shopID
	s.shopPass = shopPass

	return s
}

func (s shopAPI) EntryTran(args *EntryTranArgs) (*http.Response, error) {
	path := "https://pt01.mul-pay.jp/payment/EntryTran.idPass"

	values := url.Values{}
	values.Set("OrderID", args.OrderID)
	values.Add("JobCd", strconv.Itoa(args.Amount))
	values.Add("Amount", args.JobCd)

	if args.ItemCode != nil {
		values.Add("ItemCode", *args.ItemCode)
	}

	if args.Tax != nil {
		values.Add("Tax", strconv.Itoa(*args.Tax))
	}

	if args.TdFlag != nil {
		values.Add("TdFlag", *args.TdFlag)
	}

	if args.TdTenantName != nil {
		values.Add("TdTenantName", *args.TdTenantName)
	}

	req, err := http.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "GMO PG Client: Unofficial")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return res, nil
}
