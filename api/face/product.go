package face

import (
	"time"
)

type ProductVenue string

const (
	ProductVenueUnknownVenueType ProductVenue = "UNKNOWN_VENUE_TYPE"
	ProductVenueCBE              ProductVenue = "CBE"
	ProductVenueFCM              ProductVenue = "FCM"
	ProductVenueINTX             ProductVenue = "INTX"
)

type RiskManagedBy string

const (
	RiskManagedByUnknownRiskManagementType RiskManagedBy = "UNKNOWN_RISK_MANAGEMENT_TYPE"
	RiskManagedByFCM                       RiskManagedBy = "MANAGED_BY_FCM"
	RiskManagedByVenue                     RiskManagedBy = "MANAGED_BY_VENUE"
)

type FcmTradingSessionState string

const (
	FcmTradingSessionStateUndefined       FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_UNDEFINED"
	FcmTradingSessionStatePreOpen         FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_PRE_OPEN"
	FcmTradingSessionStatePreOpenNoCancel FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_PRE_OPEN_NO_CANCEL"
	FcmTradingSessionStateOpen            FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_OPEN"
	FcmTradingSessionStateClose           FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_CLOSE"
)

type ProductType string

const (
	ProductTypeEmpty   ProductType = ""
	ProductTypeUnknown ProductType = "UNKNOWN_PRODUCT_TYPE"
	ProductTypeSpot    ProductType = "SPOT"
	ProductTypeFuture  ProductType = "FUTURE"
)

type FcmTradingSessionDetails struct {
	IsSessionOpen                bool                   `json:"is_session_open"`
	OpenTime                     time.Time              `json:"open_time"`
	CloseTime                    time.Time              `json:"close_time"`
	SessionState                 FcmTradingSessionState `json:"session_state"`
	AfterHoursOrderEntryDisabled bool                   `json:"after_hours_order_entry_disabled"`
}

type PerpetualDetails struct {
	OpenInterest   string    `json:"open_interest"`
	FundingRate    string    `json:"funding_rate"`
	FundingTime    time.Time `json:"funding_time"`
	MaxLeverage    string    `json:"max_leverage"`
	BaseAssetUuid  string    `json:"base_asset_uuid"`
	UnderlyingType string    `json:"underlying_type"`
}

type FutureProductDetails struct {
	Venue                  string             `json:"venue"`
	ContractCode           string             `json:"contract_code"`
	ContractExpiry         time.Time          `json:"contract_expiry"`
	ContractSize           string             `json:"contract_size"`
	ContractRootUnit       string             `json:"contract_root_unit"`
	GroupDescription       string             `json:"group_description"`
	ContractExpiryTimezone string             `json:"contract_expiry_timezone"`
	GroupShortDescription  string             `json:"group_short_description"`
	RiskManagedBy          RiskManagedBy      `json:"risk_managed_by"`
	ContractExpiryType     ContractExpiryType `json:"contract_expiry_type"`
	ContractDisplayName    string             `json:"contract_display_name"`
	TimeToExpiryMs         string             `json:"time_to_expiry_ms"`
	NonCrypto              bool               `json:"non_crypto"`
	ContractExpiryName     string             `json:"contract_expiry_name"`
	PerpetualDetails       *PerpetualDetails  `json:"perpetual_details"`
}

type Product struct {
	ProductId                 string                    `json:"product_id"`
	Price                     string                    `json:"price"`
	PricePercentageChange24H  string                    `json:"price_percentage_change_24h"`
	Volume24H                 string                    `json:"volume_24h"`
	VolumePercentageChange24H string                    `json:"volume_percentage_change_24h"`
	BaseIncrement             string                    `json:"base_increment"`
	QuoteIncrement            string                    `json:"quote_increment"`
	QuoteMinSize              string                    `json:"quote_min_size"`
	QuoteMaxSize              string                    `json:"quote_max_size"`
	BaseMinSize               string                    `json:"base_min_size"`
	BaseMaxSize               string                    `json:"base_max_size"`
	BaseName                  string                    `json:"base_name"`
	QuoteName                 string                    `json:"quote_name"`
	Watched                   bool                      `json:"watched"`
	IsDisabled                bool                      `json:"is_disabled"`
	New                       bool                      `json:"new"`
	Status                    string                    `json:"status"`
	CancelOnly                bool                      `json:"cancel_only"`
	LimitOnly                 bool                      `json:"limit_only"`
	PostOnly                  bool                      `json:"post_only"`
	TradingDisabled           bool                      `json:"trading_disabled"`
	AuctionMode               bool                      `json:"auction_mode"`
	ProductType               ProductType               `json:"product_type"`
	QuoteCurrencyId           string                    `json:"quote_currency_id"`
	BaseCurrencyId            string                    `json:"base_currency_id"`
	FcmTradingSessionDetails  *FcmTradingSessionDetails `json:"fcm_trading_session_details"`
	MidMarketPrice            string                    `json:"mid_market_price"`
	Alias                     string                    `json:"alias"`
	AliasTo                   []string                  `json:"alias_to"`
	BaseDisplaySymbol         string                    `json:"base_display_symbol"`
	QuoteDisplaySymbol        string                    `json:"quote_display_symbol"`
	ViewOnly                  bool                      `json:"view_only"`
	PriceIncrement            string                    `json:"price_increment"`
	DisplayName               string                    `json:"display_name"`
	ProductVenue              ProductVenue              `json:"product_venue"`
	ApproximateQuote24HVolume string                    `json:"approximate_quote_24h_volume"`
	FutureProductDetails      *FutureProductDetails     `json:"future_product_details"`
}

type Products struct {
	Products    []*Product `json:"products"`
	NumProducts int        `json:"num_products"`
}
