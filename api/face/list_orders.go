package face

/*
	How to find the product_id? See web page with product. For example:
	https://www.coinbase.com/advanced-trade/futures/BIT-27SEP24-CDE
	Here product_id is "BIT-27SEP24-CDE"
*/

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/params"
)

type OrderPlacementSource string

const (
	OrderPlacementSourceUnknownPlacementSource OrderPlacementSource = "UNKNOWN_PLACEMENT_SOURCE"
	OrderPlacementSourceRetailSimple           OrderPlacementSource = "RETAIL_SIMPLE"
	OrderPlacementSourceRetailAdvanced         OrderPlacementSource = "RETAIL_ADVANCED"
)

type ContractExpiryType string

const (
	ContractExpiryTypeExpiring                  ContractExpiryType = "EXPIRING"
	ContractExpiryTypePerpetual                 ContractExpiryType = "PERPETUAL"
	ContractExpiryTypeUnknownContractExpiryType ContractExpiryType = "UNKNOWN_CONTRACT_EXPIRY_TYPE"
)

type SortBy string

const (
	SortByUnknown      SortBy = "UNKNOWN_SORT_BY"
	SortByLastFillTime SortBy = "LAST_FILL_TIME"
	SortByLimitPrice   SortBy = "LIMIT_PRICE"
)

type ListOrdersRequest struct {
	OrderIDs     []string      `json:"order_ids"`
	ProductIds   []string      `json:"product_ids"`
	ProductType  ProductType   `json:"product_type"`
	OrderStatus  []OrderStatus `json:"order_status"` // OPEN, CANCELLED, EXPIRED, FILLED
	TimeInForces []string      `json:"time_in_forces"`
	OrderTypes   []string      `json:"order_types"`

	Limit                int32                `json:"limit"`
	OrderSide            OrderSide            `json:"order_side"`
	StartDate            string               `json:"start_date"`             // RFC3339 Timestamp
	EndDate              string               `json:"end_date"`               // RFC3339 Timestamp
	OrderPlacementSource OrderPlacementSource `json:"order_placement_source"` // RETAIL_SIMPLE, RETAIL_ADVANCED, UNKNOWN_PLACEMENT_SOURCE
	ContractExpiryType   ContractExpiryType   `json:"contract_expiry_type"`   // EXPIRING, PERPETUAL, UNKNOWN_CONTRACT_EXPIRY_TYPE
	AssetFilters         []string             `json:"asset_filters"`
	RetailPortfolioId    string               `json:"retail_portfolio_id"`
	SortBy               SortBy               `json:"sort_by"`
	UserNativeCurrency   string               `json:"user_native_currency"`
}

func NewListOrdersRequest() *ListOrdersRequest {
	return &ListOrdersRequest{}
}

func (l *ListOrdersRequest) Params() (*params.Params, error) {
	p := params.NewParams()

	if len(l.OrderStatus) > 1 {
		if l.OrderStatus[0] == "OPEN" {
			return p, errors.New("order_status OPEN is not supported with other statuses")
		}
	}

	params.AddFilledList(p, "product_ids", l.ProductIds)
	params.AddFilledList(p, "order_status", ListOrderStatus(l.OrderStatus).String())
	params.AddFilledList(p, "time_in_forces", l.TimeInForces)
	params.AddFilledList(p, "order_types", l.OrderTypes)
	params.AddFilledList(p, "asset_filters", l.AssetFilters)

	p.AddFilled("order_ids", l.OrderIDs).
		AddFilled("product_type", l.ProductType).
		AddFilled("order_side", l.OrderSide).
		AddFilled("start_date", l.StartDate).
		AddFilled("end_date", l.EndDate).
		AddFilled("order_placement_source", l.OrderPlacementSource).
		AddFilled("contract_expiry_type", l.ContractExpiryType).
		AddFilled("retail_portfolio_id", l.UserNativeCurrency).
		AddFilled("user_native_currency", l.RetailPortfolioId).
		AddFilled("limit", l.Limit).
		AddFilled("sort_by", l.SortBy)

	fmt.Println("l.Limit: ", l.Limit)

	return p, nil
}

type EditHistory struct {
	Price                  string `json:"price"`
	Size                   string `json:"size"`
	ReplaceAcceptTimestamp string `json:"replace_accept_timestamp"`
}

type Order struct {
	AverageFilledPrice    string              `json:"average_filled_price"`
	CancelMessage         string              `json:"cancel_message,omitempty"`
	ClientOrderId         string              `json:"client_order_id"`
	CompletionPercentage  string              `json:"completion_percentage"`
	CreatedTime           time.Time           `json:"created_time"`
	EditHistory           []*EditHistory      `json:"edit_history,omitempty"`
	FilledSize            string              `json:"filled_size,omitempty"`
	FilledValue           string              `json:"filled_value"`
	IsLiquidation         string              `json:"is_liquidation,omitempty"`
	LastFillTime          string              `json:"last_fill_time,omitempty"`
	Leverage              string              `json:"leverage,omitempty"`
	MarginType            MarginType          `json:"margin_type,omitempty"`
	NumberOfFills         string              `json:"number_of_fills"`
	OrderConfiguration    *OrderConfiguration `json:"order_configuration"`
	OrderId               string              `json:"order_id"`
	OrderPlacementSource  string              `json:"order_placement_source,omitempty"`
	OrderType             OrderType           `json:"order_type,omitempty"`
	OutstandingHoldAmount string              `json:"outstanding_hold_amount,omitempty"`
	PendingCancel         bool                `json:"pending_cancel"`
	ProductId             string              `json:"product_id"`
	ProductType           ProductType         `json:"product_type,omitempty"`
	RejectMessage         string              `json:"reject_message,omitempty"`
	RejectReason          RejectReason        `json:"reject_reason,omitempty"`
	RetailPortfolioId     string              `json:"retail_portfolio_id,omitempty"`
	Settled               string              `json:"settled,omitempty,omitempty"`
	Side                  OrderSide           `json:"side"`
	SizeInclusiveOfFees   bool                `json:"size_inclusive_of_fees"`
	SizeInQuote           bool                `json:"size_in_quote"`
	Status                OrderStatus         `json:"status"`
	TimeInForce           string              `json:"time_in_force,omitempty"`
	TotalFees             string              `json:"total_fees"`
	TotalValueAfterFees   string              `json:"total_value_after_fees"`
	TriggerStatus         TriggerSide         `json:"trigger_status,omitempty"`
	UserId                string              `json:"user_id"`
}
