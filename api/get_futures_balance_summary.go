package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmpositions

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	GetFuturesBalanceSummaryPath    = "/api/v3/brokerage/cfm/balance_summary"
	callGetFuturesBalanceSummaryKey = "GetFuturesBalanceSummary"
)

// GetFuturesBalanceSummaryXRatelimit returns the x-ratelimit headers for the GetFuturesBalanceSummary API call.
func (api *API) GetFuturesBalanceSummaryXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetFuturesBalanceSummaryKey)
}

type GetFuturesBalanceSummaryResult struct {
	BalanceSummary *face.FuturesBalanceSummary `json:"balance_summary"`
}

// GetFuturesBalanceSummary Get a summary of balances for CFM trading
func (api *API) GetFuturesBalanceSummary() (*face.FuturesBalanceSummary, error) {
	u, err := url.JoinPath(api.host.String(), GetFuturesBalanceSummaryPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	res := &GetFuturesBalanceSummaryResult{}
	err = api.GetRequest(u, GetFuturesBalanceSummaryPath, callGetFuturesBalanceSummaryKey, nil, &res)

	return res.BalanceSummary, err
}
