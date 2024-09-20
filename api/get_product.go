package api

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/api/params"
)

const (
	GetProductPath    = "/api/v3/brokerage/products"
	callGetProductKey = "GetProduct"
)

// GetProductXRatelimit returns the x-ratelimit headers for the GetProduct API call.
func (api *API) GetProductXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetProductKey)
}

// https://api.coinbase.com/api/v3/brokerage/products/{product_id}
//

type GetProductResult struct {
	Account *face.Account `json:"account"`
}

// GetProduct Get information on a single product by product ID.
// - productId string required, The trading pair (e.g. 'BTC-USD').
// - getTradabilityStatus boolean, Whether to populate view_only with the readability status of the product. This is only enabled for SPOT products.
func (api *API) GetProduct(productId string, getTradabilityStatus bool) (*face.Product, error) {
	if productId == "" {
		return nil, errors.New("productId is empty")
	}

	u, err := url.JoinPath(api.host.String(), GetProductPath, productId)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	p := params.NewParams().Add("get_tradability_status", getTradabilityStatus)

	product := &face.Product{}
	err = api.GetRequest(u, GetProductPath+"/"+productId, callGetProductKey, p, &product)

	return product, err
}
