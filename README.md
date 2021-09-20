# GMO Payment API Client for Go: Unofficial

Simple Golang wrapper for [GMO Payment Gateway](https://www.gmo-pg.com/) API.

<a href="https://github.com/koffe0522/gmo-payment-go/blob/master/MIT-LICENSE" title="MIT License"><img src="https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge"></a>

## Installation

```
$ go get github.com/koffe0522/gmo-payment-go
```

## Support API

- credit card payment
  - EntryTran.json
  - ExecTran.json
  - AlterTran.json
  - ChangeTran.json
  - SearchTrade.json

## Usage

### Import

```go
import "github.com/koffe0522/gmo-payment-go/gmopg"
```

```go
package main

import (
	"fmt"
	"gmo-payment-go/gmopg"
)

func main() {
	shop := gmopg.Init(&gmopg.GmopgConfig{
		// defaultURL
		// baseURL: "https://pt01.mul-pay.jp"
		siteID:   "SiteID",
		shopID:   "SitePass",
		sitePass: "ShopID",
		shopPass: "SitePass",
	})

	orderID := "Order-gmo-payment-go"
	amount := 19800

	entryRes, err := shop.EntryTran(&gmopg.EntryTranArgs{
		OrderID: orderID,
		JobCd:   gmopg.JAuth,
		Amount:  amount,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(entryRes.Error) > 0 {
		fmt.Println(entryRes.Error)
		return
	}

	cardNo := "1234123412341234"
	expire := "2024"
	securityCode := "123"
	execRes, err := shop.ExecTran(&gmopg.ExecTranArgs{
		AccessID:     entryRes.AccessID,
		AccessPass:   entryRes.AccessPass,
		OrderID:      orderID,
		Method:       gmopg.Lump,
		CardNo:       &cardNo,
		Expire:       &expire,
		SecurityCode: &securityCode,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(execRes.Error) > 0 {
		fmt.Println(execRes.Error)
		return
	}

	fmt.Println(execRes)
}
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## Contribution

Contributions are always welcome.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[koffe0522](https://github.com/linyows)
