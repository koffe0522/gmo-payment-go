package main

import (
	"fmt"
)

func main() {
	shop := NewGMOPG(&GmopgConfig{
		siteID:   "",
		shopID:   "",
		sitePass: "",
		shopPass: "",
	})

	entryArg := &EntryTranArgs{
		OrderID: "order-id",
		JobCd:   "AUTH",
		Amount:  100,
	}

	ok, ng, err := shop.EntryTran(entryArg)

	if err != nil {
		fmt.Println("--- error ---")
		fmt.Println(err.Error())
		return
	}

	if ng != nil {
		fmt.Println("--- response: error ---")
		fmt.Println(ng)
		return
	}

	fmt.Println("--- response: ok ---")
	fmt.Println(ok.AccessID)
}
