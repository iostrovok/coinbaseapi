package api

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_gethistoricalorders

const (
	ListOrdersPath          = "/api/v3/brokerage/orders/historical/batch"
	callListCursorOrdersKey = "ListCursorOrders"
)

// ListCursorOrdersXRatelimit returns the x-ratelimit headers for the ListCursorOrders API call.
func (api *API) ListCursorOrdersXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callListCursorOrdersKey)
}

type ListOrder struct {
	Orders   []*face.Order `json:"orders"`
	Sequence string        `json:"sequence"`
	HasNext  bool          `json:"has_next"`
	Cursor   string        `json:"cursor"`
}

func (lo *ListOrder) GetOrders() []*face.Order {
	var out []*face.Order

	if lo != nil {
		out = lo.Orders
	}

	return out
}

// ListCursorOrders Get a list of orders filtered by optional query parameters (product_id, order_status, etc).
func (api *API) ListCursorOrders(req *face.ListOrdersRequest, cursor string) (*ListOrder, error) {
	// "https://api.coinbase.com/api/v3/brokerage/accounts?limit=10&cursor=20&retail_portfolio_id=30"
	params, err := req.Params()
	if err != nil {
		return nil, errors.Wrap(err, "req.Params")
	}

	if cursor != "" {
		params.Add("cursor", cursor)
	}

	u, err := url.JoinPath(api.host.String(), ListOrdersPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	res := &ListOrder{}
	err = api.GetRequest(u, ListOrdersPath, callListCursorOrdersKey, params, res)
	if err != nil {
		return nil, errors.Wrap(err, "api.GetRequest")
	}

	return res, nil
}

// ListOrders is a wrapper over ListOrders that fetches all orders, but not more than req.Limit
// if req.Limit == 0, then limit = 1_000_000
func (api *API) ListOrders(req *face.ListOrdersRequest) ([]*face.Order, error) {
	var orders []*face.Order
	cursor := ""
	hasNext := true

	limit := req.Limit
	if limit > 250 {
		req.Limit = 250
	} else if limit <= 0 {
		req.Limit = 250
		limit = 1_000_000
	}

	for hasNext {
		nextAccounts, err := api.ListCursorOrders(req, cursor)
		if err != nil {
			return nil, err
		}

		if nextAccounts == nil {
			break
		}

		orders = append(orders, nextAccounts.GetOrders()...)
		cursor = nextAccounts.Cursor
		hasNext = nextAccounts.HasNext

		if len(orders) >= int(limit) {
			break
		}

		req.Limit = limit - int32(len(orders))
		if req.Limit > 250 {
			req.Limit = 250
		}
	}

	// fix changed data for return
	req.Limit = limit

	return orders, nil
}
