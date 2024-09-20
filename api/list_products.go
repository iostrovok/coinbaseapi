package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getproducts

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/api/params"
)

const (
	ListProductsPath    = "/api/v3/brokerage/products"
	callListProductsKey = "ListProducts"
)

// ListProductsXRatelimit returns the x-ratelimit headers for the ListProducts API call.
func (api *API) ListProductsXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callListProductsKey)
}

// ListProducts Get a list of the available currency pairs for trading.
// Parameters:
// - limit int32, The number of products to be returned.
// - offset int32, The number of products to skip before returning.
// - productType string, Only returns the orders matching this product type. By default, returns all product types.
// - productIDs string[] The list of trading pairs (e.g. 'BTC-USD').
// - contractExpiryType string, Only returns the orders matching the contract expiry type. Only applicable if product_type is set to FUTURE.
// - expiringContractStatus string, Only returns contracts with this status (default is UNEXPIRED).
// - getTradabilityStatus boolean, Whether or not to populate view_only with the tradability status of the product. This is only enabled for SPOT products.
// - getAllProducts boolean, If true, return all products of all product types (including expired futures contracts).
func (api *API) ListProducts(limit, offset int32, productType face.ProductType, productIDs []string,
	contractExpiryType, expiringContractStatus string, getTradabilityStatus, getAllProducts bool) (*face.Products, error) {
	params := params.NewParams()

	if limit > 0 {
		params.Add("limit", limit)
	}

	if offset > 0 {
		params.Add("offset", offset)
	}

	if productType != face.ProductTypeUnknown {
		params.Add("product_type", productType)
	}

	if len(productIDs) > 0 {
		for _, productID := range productIDs {
			params.Add("product_ids", productID)
		}
	}

	if contractExpiryType != "" {
		params.Add("contract_expiry_type", contractExpiryType)
	}

	if expiringContractStatus != "" {
		params.Add("expiring_contract_status", expiringContractStatus)
	}

	if getTradabilityStatus {
		params.Add("get_tradability_status", "true")
	}

	if getAllProducts {
		params.Add("get_all_products", "true")
	}

	u, err := url.JoinPath(api.host.String(), ListProductsPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	products := &face.Products{}
	err = api.GetRequest(u, ListProductsPath, callListProductsKey, params, products)

	return products, err
}
