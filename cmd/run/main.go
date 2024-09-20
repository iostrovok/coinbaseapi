package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/iostrovok/coinbaseapi/api"
	"github.com/iostrovok/coinbaseapi/api/auth"
	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/internal/config"
)

func main() {
	cfg := config.Reload()
	a, err := auth.New(cfg.KeyName, cfg.KeySecret)
	if err != nil {
		panic(err)
	}

	ap, err := api.New(a, cfg.CoinbaseHost)
	if err != nil {
		panic(err)
	}

	ap.SetPrintDebugOn()

	if err := PrintAllAccounts(ap); err != nil {
		panic(err)
	}

	//candles, err := GetProductCandles(ap, "BIT-25APR25-CDE")
	candles, err := GetProductCandles(ap, "BTC-USD")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncandles::\n%+v\n\n", candles)
	fmt.Printf("\ncandles::\n%+v\n\n", candles[0])

	for i := range candles {
		fmt.Printf("%d] %+v\n", i, candles[i])
	}

	//listProducts, err := ap.ListProducts(1, 0, face.FutureProductType, nil, "", "", false, false)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\nLAST::\n%+v\n\n", listProducts)
	//fmt.Printf("\nLAST::\n%+v\n\n", listProducts.Products[0])
	////
	//futuresPositions, err := ap.ListFuturesPositions()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\nfuturesPositions::\n%+v\n\n", futuresPositions)

	//futuresBalanceSummary, err := ap.GetFuturesBalanceSummary()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\nfuturesBalanceSummary::\n%+v\n\n", futuresBalanceSummary)

	//cancelOrders, err := ap.CancelOrders("sdasdasdasd-sadasdasd")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\cancelOrders-1::\n%+v\n\n", cancelOrders)
	//fmt.Printf("\cancelOrders-2::\n%+v\n\n", cancelOrders[0])

	//order, clientOrderId, err := CreateOrder(ap, "BIT-27SEP24-CDE")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("new order clientOrderId:: %s\n", clientOrderId)
	//fmt.Printf("new order::\n%+v\n\n", order)

	//if err := PrintAllOrders(ap); err != nil {
	//	panic(err)
	//}
}

// CreateOrder creates new orders
// returns the order info & and internal order id
// just for example
func CreateOrder(ap *api.API, productId string) (*face.CreateOrderResult, string, error) {
	if config.CFG().DebugMode {
		ap.SetPrintDebugOn()
		defer ap.SetPrintDebugOff()
	}

	// check the product
	_, err := ap.GetProduct(productId, true)
	if err != nil {
		return nil, "", err
	}

	clientOrderId := uuid.New().String()
	createOrderRequest := face.NewCreateOrderRequest(clientOrderId, productId, face.OrderSideBUY).
		SetMarketMarketIoc("", "1").
		SetLMP("1.0", "ISOLATED", "")

	order, err := ap.CreateOrder(createOrderRequest)
	if err != nil {
		return nil, clientOrderId, err
	}

	return order, clientOrderId, nil
}

// PrintAllOrders print all your order to console
func PrintAllOrders(ap *api.API) error {
	if config.CFG().DebugMode {
		ap.SetPrintDebugOn()
		defer ap.SetPrintDebugOff()
	}

	orders, err := ap.ListOrders(face.NewListOrdersRequest())
	if err != nil {
		return err
	}

	fmt.Printf("\nAll orders (total: %d):\n", len(orders))
	for i, order := range orders {
		fmt.Printf("%d] %+v\n", i, order)
	}

	return nil
}

// PrintAllAccounts print all your account to console
func PrintAllAccounts(ap *api.API) error {
	if config.CFG().DebugMode {
		ap.SetPrintDebugOn()
		defer ap.SetPrintDebugOff()
	}

	accounts, err := ap.ListAllAccounts()
	if err != nil {
		return err
	}

	fmt.Printf("\nAll accounts (total: %d):\n", len(accounts))
	for i, account := range accounts {
		fmt.Printf("%d] %+v\n", i, account)
	}

	return nil
}

// GetProductCandles gets rates for a single product by product ID, grouped in buckets.
// - productId, string, required, The trading pair (e.g. 'BTC-USD').
// - start time.Time, required, The timestamp indicating the start of the time interval.
// - end, time.Time, required, The timestamp indicating the end of the time interval.
// - granularity, face.GranularityType, required, The timeframe each candle represents.
// - limit, int32, optional, The number of candle buckets to be returned. By default, returns 350 (max 350).
func GetProductCandles(ap *api.API, productId string) ([]*face.CandleStr, error) {
	if config.CFG().DebugMode {
		ap.SetPrintDebugOn()
		defer ap.SetPrintDebugOff()
	}

	start := time.Now().Add(-time.Hour * 3)
	end := time.Now().Add(-time.Hour * 2)

	// check the product
	return ap.GetProductCandles(productId, face.GranularityTypeOneMinute, start, end, 0)
}
