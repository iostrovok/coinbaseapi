package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmpositions

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

// Get Futures FuturePosition
//GET
//https://api.coinbase.com//{product_id}
//Get positions for a specific CFM product

// GetFuturesPosition

const (
	GetFuturesPositionPath    = "/api/v3/brokerage/cfm/positions"
	callGetFuturesPositionKey = "GetFuturesPosition"
)

// GetFuturesPositionXRatelimit returns the x-ratelimit headers for the GetFuturesPosition API call.
func (api *API) GetFuturesPositionXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetFuturesPositionKey)
}

type GetFuturesPositionResult struct {
	Position *face.FuturePosition `json:"position"`
}

func (r *GetFuturesPositionResult) getPosition() *face.FuturePosition {
	if r == nil {
		return nil
	}

	return r.Position
}

// GetFuturesPosition Get positions for a specific CFM product
// productId string, required he ticker symbol (e.g. 'BIT-28JUL23-CDE').
func (api *API) GetFuturesPosition(productId string) (*face.FuturePosition, error) {
	if productId == "" {
		return nil, errors.New("productId is empty")
	}

	u, err := url.JoinPath(api.host.String(), GetFuturesPositionPath, productId)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	res := &GetFuturesPositionResult{}
	err = api.GetRequest(u, GetFuturesPositionPath+"/"+productId, callGetFuturesPositionKey, nil, res)

	return res.getPosition(), err
}
