package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmpositions

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	ListFuturesPositionsPath = "/api/v3/brokerage/cfm/positions"
)

type ListFuturesPositionsResult struct {
	Positions []*face.FuturePosition `json:"positions"`
}

func (r *ListFuturesPositionsResult) getPositions() []*face.FuturePosition {
	if r == nil {
		return nil
	}

	return r.Positions
}

// ListFuturesPositions Get a list of positions in CFM products
func (api *API) ListFuturesPositions() ([]*face.FuturePosition, error) {
	u, err := url.JoinPath(api.host.String(), ListFuturesPositionsPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	res := &ListFuturesPositionsResult{}
	err = api.GetRequest(u, ListFuturesPositionsPath, nil, res)

	return res.getPositions(), err
}
