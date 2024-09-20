package face

import (
	"time"
)

type AccountType string

const (
	AccountTypeUnspecified AccountType = "ACCOUNT_TYPE_UNSPECIFIED"
	AccountTypeCrypto      AccountType = "ACCOUNT_TYPE_CRYPTO"
	AccountTypeFiat        AccountType = "ACCOUNT_TYPE_FIAT"
	AccountTypeVault       AccountType = "ACCOUNT_TYPE_VAULT"
	AccountTypePerpFutures AccountType = "ACCOUNT_TYPE_PERP_FUTURES"
)

type Account struct {
	// UUID string, Unique identifier for account.
	UUID string `json:"uuid"`
	// Name string, Name for the account.
	Name string `json:"name"`
	// currency string, Currency symbol for the account.
	Currency string `json:"currency"`

	// AvailableBalance object required, Available balance in the account.
	AvailableBalance struct {
		// Value, string required, Amount of currency that this object represents.
		Value string `json:"value"`
		// Currency string required, Denomination of the currency.
		Currency string `json:"currency"`
	} `json:"available_balance"`

	// Default boolean, Whether this account is the user's primary account
	Default bool `json:"default"`
	// Active boolean, Whether this account is active and okay to use.
	Active bool `json:"active"`
	// CreatedAt RFC3339 Timestamp, the Time at which this account was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt RFC3339 Timestamp, the Time at which this account was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt RFC3339 Timestamp, the Time at which this account was deleted.
	DeletedAt time.Time `json:"deleted_at"`
	// Type string, Possible values:
	// [ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_CRYPTO, ACCOUNT_TYPE_FIAT, ACCOUNT_TYPE_VAULT, ACCOUNT_TYPE_PERP_FUTURES]
	// What type the account is.
	Type AccountType `json:"type"`
	// Ready boolean, Whether or not this account is ready to trade.
	Ready bool `json:"ready"`
	// Hold object required, Amount that is being held for pending transfers against the available balance.
	Hold struct {
		// Value string required, Amount of currency that this object represents.
		Value string `json:"value"`
		// Currency string required, Denomination of the currency.
		Currency string `json:"currency"`
	} `json:"hold"`

	// RetailPortfolioId string, The ID of the portfolio this account is associated with.
	RetailPortfolioId string `json:"retail_portfolio_id"`
}

type Accounts struct {
	Accounts []*Account `json:"accounts"`
	// HasNext boolean required, Whether there are additional pages for this query.
	HasNext bool `json:"has_next"`
	// Cursor string. For paginated responses, returns all responses that come after this value.
	Cursor string `json:"cursor"`
	// Size int32, Number of accounts returned
	Size int32 `json:"size"`
}

func (ac *Accounts) GetAccounts() []*Account {
	var out []*Account
	if ac != nil {
		out = ac.Accounts
	}

	return out
}
