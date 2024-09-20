package api

import (
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/api/params"
)

// 'https://api.coinbase.com/api/v3/brokerage/products/BTC/candles?start=12321312&end=534345345&granularity=ONE_MINUTE&limit=234

const (
	// GetProductCandlesPath = "/api/v3/brokerage/products/{product_id}/candles"
	callGetProductCandlesKey = "GetProductCandles"
)

// GetProductCandlesXRatelimit returns the x-ratelimit headers for the GetProductCandles API call.
func (api *API) GetProductCandlesXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetProductCandlesKey)
}

// GET https://api.coinbase.com/api/v3/brokerage/products/{product_id}/candles
//

type GetProductCandlesResult struct {
	Candles []*face.CandleStr `json:"candles"`
}

// GetProductCandles gets rates for a single product by product ID, grouped in buckets.
// - productId, string, required, The trading pair (e.g. 'BTC-USD').
// - start time.Time, required, The timestamp indicating the start of the time interval.
// - end, time.Time, required, The timestamp indicating the end of the time interval.
// - granularity, face.GranularityType, required, The timeframe each candle represents.
// - limit, int32, optional, The number of candle buckets to be returned. By default, returns 350 (max 350).
func (api *API) GetProductCandles(productId string, granularity face.GranularityType,
	start, end time.Time, limit int) ([]*face.CandleStr, error) {
	if productId == "" {
		return nil, errors.New("productId is empty")
	}

	u, err := url.JoinPath(api.host.String(), GetProductPath, productId, "candles")
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	p := params.NewParams().
		Add("granularity", granularity).
		Add("start", start.Unix()).
		Add("end", end.Unix())

	if limit > 0 {
		p.Add("granularity", granularity)
	}

	product := &GetProductCandlesResult{}
	err = api.GetRequest(u, GetProductPath+"/"+productId+"/candles", callGetProductCandlesKey, p, &product)
	if err != nil {
		return nil, err
	}

	return product.Candles, err
}
