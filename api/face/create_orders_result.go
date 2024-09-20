package face

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_postorder

// Result of the order creation

type SuccessResponse struct {
	OrderId         string `json:"order_id"`                  // The ID of the order.
	ProductId       string `json:"product_id,omitempty"`      // The trading pair (e.g. 'BTC-USD').
	Side            string `json:"side,omitempty"`            // Possible values: [BUY, SELL]. The side of the market that the order is on (e.g. 'BUY', 'SELL').
	ClientOrderId   string `json:"client_order_id,omitempty"` // The unique ID provided for the order (used for identification purposes).
	AttachedOrderId string `json:"attached_order_id,omitempty"`
}

type ErrorResponse struct {
	Error                 string `json:"error"`                    // **(Deprecated)** The reason the order failed to be created
	Message               string `json:"message"`                  // Generic error message explaining why the order was not created
	ErrorDetails          string `json:"error_details"`            // Descriptive error message explaining why the order was not created
	PreviewFailureReason  string `json:"preview_failure_reason"`   // **(Deprecated)** The reason the order failed to be created
	NewOrderFailureReason string `json:"new_order_failure_reason"` // The reason the order failed to be created
}

type CreateOrderResult struct {
	Success            bool                `json:"success"`
	SuccessResponse    *SuccessResponse    `json:"success_response"`
	ErrorResponse      *ErrorResponse      `json:"error_response"`
	OrderConfiguration *OrderConfiguration `json:"order_configuration,omitempty"`
}
