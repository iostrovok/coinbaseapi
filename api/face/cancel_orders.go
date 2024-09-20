package face

type FailureReason string

const (
	FailureReasonUnknown                        FailureReason = "UNKNOWN_CANCEL_FAILURE_REASON"
	FailureReasonInvalidCancelRequest           FailureReason = "INVALID_CANCEL_REQUEST"
	FailureReasonUnknownCancelOrder             FailureReason = "UNKNOWN_CANCEL_ORDER"
	FailureReasonCommanderRejectedCancelOrder   FailureReason = "COMMANDER_REJECTED_CANCEL_ORDER"
	FailureReasonDuplicateCancelRequest         FailureReason = "DUPLICATE_CANCEL_REQUEST"
	FailureReasonInvalidCancelProductId         FailureReason = "INVALID_CANCEL_PRODUCT_ID"
	FailureReasonInvalidCancelFcmTradingSession FailureReason = "INVALID_CANCEL_FCM_TRADING_SESSION"
	FailureReasonNotAllowedToCancel             FailureReason = "NOT_ALLOWED_TO_CANCEL"
	FailureReasonOrderIsFullyFilled             FailureReason = "ORDER_IS_FULLY_FILLED"
	FailureReasonOrderIsBeingReplaced           FailureReason = "ORDER_IS_BEING_REPLACED"
)

type CancelOrder struct {
	OrderID       string        `json:"order_id"`
	Success       bool          `json:"success"`
	FailureReason FailureReason `json:"failure_reason"`
}
