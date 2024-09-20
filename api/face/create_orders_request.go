package face

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_postorder

type StopDirection string

const (
	StopDirectionStopUp   StopDirection = "STOP_DIRECTION_STOP_UP"
	StopDirectionStopDown StopDirection = "STOP_DIRECTION_STOP_DOWN"
)

type CreateOrderRequest struct {
	ClientOrderId      string              `json:"client_order_id"` // required
	ProductId          string              `json:"product_id"`      // required
	Side               OrderSide           `json:"side"`            // required
	OrderConfiguration *OrderConfiguration `json:"order_configuration"`
	Leverage           string              `json:"leverage,omitempty"`
	MarginType         MarginType          `json:"margin_type,omitempty"`
	PreviewId          string              `json:"preview_id,omitempty"`
}

func NewCreateOrderRequest(clientOrderId, productId string, side OrderSide) *CreateOrderRequest {
	return &CreateOrderRequest{
		ClientOrderId:      clientOrderId,
		ProductId:          productId,
		Side:               side,
		OrderConfiguration: &OrderConfiguration{},
	}
}

type MarketMarketIoc struct {
	QuoteSize string `json:"quote_size,omitempty"`
	BaseSize  string `json:"base_size,omitempty"`
}

// SorLimitIoc Buy or sell a specified quantity of an Asset at a specified price.
// The Order will only post to the Order Book if it will immediately Fill;
// any remaining quantity is canceled.
type SorLimitIoc struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
}

// LimitLimitGtc Buy or sell a specified quantity of an Asset at a specified price.
// If posted, the Order will remain on the Order Book until canceled.
type LimitLimitGtc struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
	PostOnly   bool   `json:"post_only"`
}

type LimitLimitGtd struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
	EndTime    string `json:"end_time"`
	PostOnly   bool   `json:"post_only"`
}

type LimitLimitFok struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
}

type StopLimitStopLimitGtc struct {
	BaseSize      string        `json:"base_size"`
	LimitPrice    string        `json:"limit_price"`
	StopPrice     string        `json:"stop_price"`
	StopDirection StopDirection `json:"stop_direction"`
}

type StopLimitStopLimitGtd struct {
	BaseSize      string        `json:"base_size"`
	LimitPrice    string        `json:"limit_price"`
	StopPrice     string        `json:"stop_price"`
	EndTime       string        `json:"end_time"`
	StopDirection StopDirection `json:"stop_direction"`
}

type TriggerBracketGtc struct {
	BaseSize         string `json:"base_size"`
	LimitPrice       string `json:"limit_price"`
	StopTriggerPrice string `json:"stop_trigger_price"`
}

type TriggerBracketGtd struct {
	BaseSize         string `json:"base_size"`
	LimitPrice       string `json:"limit_price"`
	StopTriggerPrice string `json:"stop_trigger_price"`
	EndTime          string `json:"end_time"`
}

type OrderConfiguration struct { // required
	MarketMarketIoc       *MarketMarketIoc       `json:"market_market_ioc,omitempty"`
	SorLimitIoc           *SorLimitIoc           `json:"sor_limit_ioc,omitempty"`
	LimitLimitGtc         *LimitLimitGtc         `json:"limit_limit_gtc,omitempty"`
	LimitLimitGtd         *LimitLimitGtd         `json:"limit_limit_gtd,omitempty"`
	LimitLimitFok         *LimitLimitFok         `json:"limit_limit_fok,omitempty"`
	StopLimitStopLimitGtc *StopLimitStopLimitGtc `json:"stop_limit_stop_limit_gtc,omitempty"`
	StopLimitStopLimitGtd *StopLimitStopLimitGtd `json:"stop_limit_stop_limit_gtd,omitempty"`
	TriggerBracketGtc     *TriggerBracketGtc     `json:"trigger_bracket_gtc,omitempty"`
	TriggerBracketGtd     *TriggerBracketGtd     `json:"trigger_bracket_gtd,omitempty"`
}

// SetLMP sets leverage, marginType, previewId
func (cor *CreateOrderRequest) SetLMP(leverage string, marginType MarginType, previewId string) *CreateOrderRequest {
	cor.Leverage = leverage
	cor.MarginType = marginType
	cor.PreviewId = previewId

	return cor
}

func (cor *CreateOrderRequest) SetMarketMarketIoc(quoteSize, baseSize string) *CreateOrderRequest {
	cor.OrderConfiguration.MarketMarketIoc = &MarketMarketIoc{
		QuoteSize: quoteSize,
		BaseSize:  baseSize,
	}

	return cor
}

func (cor *CreateOrderRequest) SetSorLimitIoc(baseSize, limitPrice string) *CreateOrderRequest {
	cor.OrderConfiguration.SorLimitIoc = &SorLimitIoc{
		LimitPrice: limitPrice,
		BaseSize:   baseSize,
	}

	return cor
}

func (cor *CreateOrderRequest) SetLimitLimitGtc(baseSize, limitPrice string, postOnly bool) *CreateOrderRequest {
	cor.OrderConfiguration.LimitLimitGtc = &LimitLimitGtc{
		LimitPrice: limitPrice,
		BaseSize:   baseSize,
		PostOnly:   postOnly,
	}

	return cor
}

func (cor *CreateOrderRequest) SetLimitLimitGtd(baseSize, limitPrice, endTime string, postOnly bool) *CreateOrderRequest {
	cor.OrderConfiguration.LimitLimitGtd = &LimitLimitGtd{
		LimitPrice: limitPrice,
		BaseSize:   baseSize,
		EndTime:    endTime,
		PostOnly:   postOnly,
	}

	return cor
}

func (cor *CreateOrderRequest) SetLimitLimitFok(baseSize, limitPrice string) *CreateOrderRequest {
	cor.OrderConfiguration.LimitLimitFok = &LimitLimitFok{
		LimitPrice: limitPrice,
		BaseSize:   baseSize,
	}

	return cor
}

func (cor *CreateOrderRequest) SetStopLimitStopLimitGtc(baseSize, limitPrice, stopPrice string, stopDirection StopDirection) *CreateOrderRequest {
	cor.OrderConfiguration.StopLimitStopLimitGtc = &StopLimitStopLimitGtc{
		LimitPrice:    limitPrice,
		BaseSize:      baseSize,
		StopPrice:     stopPrice,
		StopDirection: stopDirection,
	}

	return cor
}

func (cor *CreateOrderRequest) SetStopLimitStopLimitGtd(baseSize, limitPrice, stopPrice string, stopDirection StopDirection, endTime string) *CreateOrderRequest {
	cor.OrderConfiguration.StopLimitStopLimitGtd = &StopLimitStopLimitGtd{
		LimitPrice:    limitPrice,
		BaseSize:      baseSize,
		StopPrice:     stopPrice,
		StopDirection: stopDirection,
		EndTime:       endTime,
	}

	return cor
}

func (cor *CreateOrderRequest) SetTriggerBracketGtc(baseSize, limitPrice, stopTriggerPrice string) *CreateOrderRequest {
	cor.OrderConfiguration.TriggerBracketGtc = &TriggerBracketGtc{
		LimitPrice:       limitPrice,
		BaseSize:         baseSize,
		StopTriggerPrice: stopTriggerPrice,
	}

	return cor
}

func (cor *CreateOrderRequest) SetTriggerBracketGtd(baseSize, limitPrice, stopTriggerPrice, endTime string) *CreateOrderRequest {
	cor.OrderConfiguration.TriggerBracketGtd = &TriggerBracketGtd{
		LimitPrice:       limitPrice,
		BaseSize:         baseSize,
		StopTriggerPrice: stopTriggerPrice,
		EndTime:          endTime,
	}

	return cor
}
