package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_closeposition
// https://api.coinbase.com/api/v3/brokerage/orders/close_position

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	ClosePositionPath    = "/api/v3/brokerage/orders/close_position"
	callClosePositionKey = "ClosePosition"
)

// ClosePositionXRatelimit returns the x-ratelimit headers for the ClosePosition API call.
func (api *API) ClosePositionXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callClosePositionKey)
}

type ClosePositionRequest struct {
	ClientOrderId string `json:"client_order_id"`
	ProductId     string `json:"product_id"`
	Size          string `json:"size,omitempty"`
}

// ClosePosition Places an order to close any open positions for a specified product_id.
// parameters:
// - clientOrderId string required, The unique ID provided for the order (used for identification).
// - productId string required, The trading pair (e.g. 'BIT-28JUL23-CDE').
// - size string, The number of contracts that should be closed.
func (api *API) ClosePosition(clientOrderId, productId, size string) (*face.CreateOrderResult, error) {
	if clientOrderId == "" {
		return nil, errors.New("clientOrderId is empty")
	}

	if productId == "" {
		return nil, errors.New("productId is empty")
	}

	u, err := url.JoinPath(api.host.String(), ClosePositionPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	req := &ClosePositionRequest{
		ClientOrderId: clientOrderId,
		ProductId:     productId,
		Size:          size,
	}

	products := &face.CreateOrderResult{}
	err = api.PostRequest(u, ClosePositionPath, callClosePositionKey, req, products)

	return products, err
}
