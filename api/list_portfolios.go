package api

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/api/params"
)

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getportfolios

const (
	ListPortfoliosPath    = "/api/v3/brokerage/portfolios"
	callListPortfoliosKey = "ListPortfolios"
)

// ListPortfoliosXRatelimit returns the x-ratelimit headers for the ListPortfolios API call.
func (api *API) ListPortfoliosXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callListPortfoliosKey)
}

// ListPortfolios Get all portfolios of a user.
// This endpoint requires the "view" permission.
//   - portfolio_type string, Only returns portfolios matching this portfolio type.
func (api *API) ListPortfolios(portfolio_type face.PortfolioType) ([]*face.Portfolio, error) {
	p := params.NewParams()

	if portfolio_type != face.PortfolioTypeEmpty {
		p.Add("portfolio_type", portfolio_type)
	}

	u, err := url.JoinPath(api.host.String(), ListOrdersPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	res := &face.ListPortfolios{}
	err = api.GetRequest(u, ListPortfoliosPath, callListPortfoliosKey, p, res)
	if err != nil {
		return nil, errors.Wrap(err, "api.GetRequest")
	}

	return res.GetPortfolios(), nil
}
