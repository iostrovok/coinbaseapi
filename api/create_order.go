package api

//
//POST
//https://api.coinbase.com/api/v3/brokerage/orders
//Create an order with a specified product_id (asset-pair), side (buy/sell), etc.
//
//
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_postorder
//
//
//create_order

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	CreateOrderPath    = "/api/v3/brokerage/orders"
	callCreateOrderKey = "CreateOrder"
)

// CreateOrderXRatelimit returns the x-ratelimit headers for the CreateOrder API call.
func (api *API) CreateOrderXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callCreateOrderKey)
}

// CreateOrder Get a list of the available currency pairs for trading.
// Parameters:
// createOrderRequest *face.CreateOrderRequest
func (api *API) CreateOrder(createOrderRequest *face.CreateOrderRequest) (*face.CreateOrderResult, error) {
	u, err := url.JoinPath(api.host.String(), CreateOrderPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	products := &face.CreateOrderResult{}
	err = api.PostRequest(u, CreateOrderPath, callCreateOrderKey, createOrderRequest, products)

	return products, err
}
