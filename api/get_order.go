package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_gethistoricalorder

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	GetOrderPath    = "/api/v3/brokerage/orders/historical"
	callGetOrderKey = "GetOrder"
)

// GetOrderXRatelimit returns the x-ratelimit headers for the GetOrder API call.
func (api *API) GetOrderXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetOrderKey)
}

// https://api.coinbase.com/api/v3/brokerage/orders/historical/{order_id}

type GetOrderResult struct {
	Order *face.Order `json:"order"`
}

// GetOrder Get a single order by order ID.
// - order_id string required, The ID of the order.
func (api *API) GetOrder(orderId string) (*face.Order, error) {
	if orderId == "" {
		return nil, errors.New("orderId is empty")
	}

	u, err := url.JoinPath(api.host.String(), GetOrderPath, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	order := &GetOrderResult{}
	err = api.GetRequest(u, GetOrderPath+"/"+orderId, callGetOrderKey, nil, &order)

	return order.Order, err
}
