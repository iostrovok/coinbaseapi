package face

type MarginWindowType string

const (
	MarginWindowTypeUnspecified MarginWindowType = "FCM_MARGIN_WINDOW_TYPE_UNSPECIFIED"
	MarginWindowTypeOvernight   MarginWindowType = "FCM_MARGIN_WINDOW_TYPE_OVERNIGHT"
	MarginWindowTypeWeekend     MarginWindowType = "FCM_MARGIN_WINDOW_TYPE_WEEKEND"
	MarginWindowTypeIntraday    MarginWindowType = "FCM_MARGIN_WINDOW_TYPE_INTRADAY"
	MarginWindowTypeTransition  MarginWindowType = "FCM_MARGIN_WINDOW_TYPE_TRANSITION"
)

type MarginLevel string

const (
	MarginLevelUnspecified MarginLevel = "MARGIN_LEVEL_TYPE_UNSPECIFIED"
	MarginLevelBase        MarginLevel = "MARGIN_LEVEL_TYPE_BASE"
	MarginLevelWarning     MarginLevel = "MARGIN_LEVEL_TYPE_WARNING"
	MarginLevelDanger      MarginLevel = "MARGIN_LEVEL_TYPE_DANGER"
	MarginLevelLiquidation MarginLevel = "MARGIN_LEVEL_TYPE_LIQUIDATION"
)

// FuturesBalanceSummary
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmbalancesummary
type FuturesBalanceSummary struct {
	FuturesBuyingPower struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"futures_buying_power"`
	TotalUsdBalance struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"total_usd_balance"`
	CbiUsdBalance struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"cbi_usd_balance"`
	CfmUsdBalance struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"cfm_usd_balance"`
	TotalOpenOrdersHoldAmount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"total_open_orders_hold_amount"`
	UnrealizedPnl struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"unrealized_pnl"`
	DailyRealizedPnl struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"daily_realized_pnl"`
	InitialMargin struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"initial_margin"`
	AvailableMargin struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"available_margin"`
	LiquidationThreshold struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"liquidation_threshold"`
	LiquidationBufferAmount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"liquidation_buffer_amount"`
	LiquidationBufferPercentage string `json:"liquidation_buffer_percentage"`
	IntradayMarginWindowMeasure struct {
		MarginWindowType   MarginWindowType `json:"margin_window_type"`
		MarginLevel        MarginLevel      `json:"margin_level"`
		InitialMargin      string           `json:"initial_margin"`
		MaintenanceMargin  string           `json:"maintenance_margin"`
		LiquidationBuffer  string           `json:"liquidation_buffer"`
		TotalHold          string           `json:"total_hold"`
		FuturesBuyingPower string           `json:"futures_buying_power"`
	} `json:"intraday_margin_window_measure"`
	OvernightMarginWindowMeasure struct {
		MarginWindowType   MarginWindowType `json:"margin_window_type"`
		MarginLevel        MarginLevel      `json:"margin_level"`
		InitialMargin      string           `json:"initial_margin"`
		MaintenanceMargin  string           `json:"maintenance_margin"`
		LiquidationBuffer  string           `json:"liquidation_buffer"`
		TotalHold          string           `json:"total_hold"`
		FuturesBuyingPower string           `json:"futures_buying_power"`
	} `json:"overnight_margin_window_measure"`
}
