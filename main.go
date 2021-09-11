package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	shop := NewShopAPI("hoge.mul-pay.jp", "shop_id", "shop_pass")
	strPointer := "tset"
	entryArg := &EntryTranArgs{
		OrderID:  "order_id",
		JobCd:    "AUTH",
		Amount:   100,
		ItemCode: &strPointer,
	}

	resp, err := shop.EntryTran(entryArg)

	if err != nil {
		fmt.Println("--- error ---")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("--- success ---")
	fmt.Println(resp.Status)
	fmt.Println(ioutil.ReadAll(resp.Body))
}
