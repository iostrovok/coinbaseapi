package face

type PositionSide string

const (
	PositionSideUnknown  PositionSide = "UNKNOWN"
	PositionSideLongSide PositionSide = "LONG"
	PositionSideShort    PositionSide = "SHORT"
)

type OrderSide string

const (
	OrderSideSELL OrderSide = "SELL"
	OrderSideBUY  OrderSide = "BUY"
)

type OrderStatus string

const (
	OrderStatusPending      OrderStatus = "PENDING"
	OrderStatusOpen         OrderStatus = "OPEN"
	OrderStatusFilled       OrderStatus = "FILLED"
	OrderStatusFailed       OrderStatus = "FAILED"
	OrderStatusUnknown      OrderStatus = "UNKNOWN_ORDER_STATUS"
	OrderStatusQueued       OrderStatus = "QUEUED"
	OrderStatusCancelQueued OrderStatus = "CANCEL_QUEUED"
	OrderStatusCancelled    OrderStatus = "CANCELLED"
	OrderStatusExpired      OrderStatus = "EXPIRED"
)

type ListOrderStatus []OrderStatus

func (l ListOrderStatus) String() []string {
	var out []string
	for _, v := range l {
		out = append(out, string(v))
	}

	return out
}

type TriggerSide string

const (
	TriggerSideUnknown       TriggerSide = "UNKNOWN_TRIGGER_STATUS"
	TriggerSideInvalid       TriggerSide = "INVALID_ORDER_TYPE"
	TriggerSideStopPending   TriggerSide = "STOP_PENDING"
	TriggerSideStopTriggered TriggerSide = "STOP_TRIGGERED"
)

type OrderType string

const (
	OrderTypeUnknown OrderType = "UNKNOWN_ORDER_TYPE"
	OrderTypeMarket  OrderType = "MARKET"
	OrderTypeLimit   OrderType = "LIMIT"
	OrderTypeStop    OrderType = "STOP"
	OrderTypeStopLmt OrderType = "STOP_LIMIT"
	OrderTypeBracket OrderType = "BRACKET"
)

type RejectReason string

const (
	OrderTypeUnspecified       RejectReason = "REJECT_REASON_UNSPECIFIED"
	OrderTypeHoldFailure       RejectReason = "HOLD_FAILURE"
	OrderTypeRateLimitExceeded RejectReason = "RATE_LIMIT_EXCEEDED"
	OrderTypeInsufficientFunds RejectReason = "REJECT_REASON_INSUFFICIENT_FUNDS"
	OrderTypeTooManyOpenOrders RejectReason = "TOO_MANY_OPEN_ORDERS"
)

type MarginType string

const (
	MarginTypeCross    MarginType = "CROSS"
	MarginTypeIsolated MarginType = "ISOLATED"
)
