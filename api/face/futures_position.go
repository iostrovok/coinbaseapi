package face

import (
	"time"
)

type FuturePosition struct {
	ProductId         string       `json:"product_id"`          // The ticker symbol (e.g. 'BIT-28JUL23-CDE')
	ExpirationTime    time.Time    `json:"expiration_time"`     // The expiry of your position
	Side              PositionSide `json:"side"`                // The side of your position
	NumberOfContracts string       `json:"number_of_contracts"` // The size of your position in contracts
	CurrentPrice      string       `json:"current_price"`       // The current price of the product
	AvgEntryPrice     string       `json:"avg_entry_price"`     // The average entry price at which you entered your current position
	UnrealizedPnl     string       `json:"unrealized_pnl"`      // Your current unrealized PnL for your position
	DailyRealizedPnl  string       `json:"daily_realized_pnl"`  // Your realized PnL from your trades in this product on current trade date
}
