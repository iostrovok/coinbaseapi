package face

import (
	"errors"
	"strconv"
	"time"

	"github.com/iostrovok/coinbaseapi/internal/utils"
)

type GranularityType string

const (
	GranularityTypeUnknown       = "UNKNOWN_GRANULARITY"
	GranularityTypeOneMinute     = "ONE_MINUTE"
	GranularityTypeFiveMinute    = "FIVE_MINUTE"
	GranularityTypeFifteenMinute = "FIFTEEN_MINUTE"
	GranularityTypeThirtyMinute  = "THIRTY_MINUTE"
	GranularityTypeTwoHour       = "TWO_HOUR"
	GranularityTypeOneHour       = "ONE_HOUR"
	GranularityTypeSixHour       = "SIX_HOUR"
	GranularityTypeOneDay        = "ONE_DAY"
)

func (g GranularityType) Duration() time.Duration {
	switch g {
	case GranularityTypeOneMinute:
		return time.Minute
	case GranularityTypeFiveMinute:
		return 5 * time.Minute
	case GranularityTypeFifteenMinute:
		return 15 * time.Minute
	case GranularityTypeThirtyMinute:
		return 30 * time.Minute
	case GranularityTypeTwoHour:
		return 2 * time.Hour
	case GranularityTypeOneHour:
		return time.Hour
	case GranularityTypeSixHour:
		return 6 * time.Hour
	case GranularityTypeOneDay:
		return 24 * time.Hour
	default:
		return 0
	}
}

type Candle struct {
	Start  int64   `json:"start"`  // The UNIX timestamp indicating the start of the time interval, Example: 1639508050
	Low    float32 `json:"low"`    // Lowest price during the bucket interval. Example: 140.21
	High   float32 `json:"high"`   // Highest price during the bucket interval. Example: 140.21
	Open   float32 `json:"open"`   // Opening price (first trade) in the bucket interval. Example: 140.21
	Close  float32 `json:"close"`  // Closing price (last trade) in the bucket interval. Example: 140.21
	Volume float32 `json:"volume"` // Volume of trading activity during the bucket interval. Example: 56437345
}

type CandleStr struct {
	Start  string `json:"start"`  // The UNIX timestamp indicating the start of the time interval, Example: 1639508050
	Low    string `json:"low"`    // Lowest price during the bucket interval. Example: 140.21
	High   string `json:"high"`   // Highest price during the bucket interval. Example: 140.21
	Open   string `json:"open"`   // Opening price (first trade) in the bucket interval. Example: 140.21
	Close  string `json:"close"`  // Closing price (last trade) in the bucket interval. Example: 140.21
	Volume string `json:"volume"` // Volume of trading activity during the bucket interval. Example: 56437345
}

func (c *CandleStr) Candle() (*Candle, error) {
	out := &Candle{}
	if c == nil {
		return nil, errors.New("CandleStr is nil")
	}
	var err error

	out.Start, err = strconv.ParseInt(c.Start, 10, 64)
	if err != nil {
		return nil, err
	}

	out.Low, err = utils.ParseFloat32(c.Low)
	if err != nil {
		return nil, err
	}

	out.High, err = utils.ParseFloat32(c.High)
	if err != nil {
		return nil, err
	}

	out.Open, err = utils.ParseFloat32(c.Open)
	if err != nil {
		return nil, err
	}

	out.Close, err = utils.ParseFloat32(c.Close)
	if err != nil {
		return nil, err
	}

	out.Volume, err = utils.ParseFloat32(c.Volume)
	if err != nil {
		return nil, err
	}

	return out, nil
}
