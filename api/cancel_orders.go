package api

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

// Initiate cancel requests for one or more orders.

const (
	CancelOrdersPath = "/api/v3/brokerage/orders/batch_cancel"
)

type CancelOrders struct {
	Results []*face.CancelOrder `json:"results"`
}

// CancelOrders Initiate cancel requests for one or more orders.
// Parameters:
// order_ids string[] required, The order IDs that cancel requests should be initiated for.
func (api *API) CancelOrders(orderIDs ...string) ([]*face.CancelOrder, error) {
	u, err := url.JoinPath(api.host.String(), CancelOrdersPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	req := map[string][]string{
		"order_ids": orderIDs,
	}

	products := &CancelOrders{}
	err = api.PostRequest(u, CancelOrdersPath, req, products)

	fmt.Printf("CancelOrders:\n%v\n", products)

	return products.Results, err
}
