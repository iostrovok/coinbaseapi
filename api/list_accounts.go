package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getaccounts

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
	"github.com/iostrovok/coinbaseapi/api/params"
)

const (
	ListAccountsPath    = "/api/v3/brokerage/accounts"
	callListAccountsKey = "ListAccounts"
)

// ListAccountsXRatelimit returns the x-ratelimit headers for the GetOrder API call.
func (api *API) ListAccountsXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callListAccountsKey)
}

// ListAccounts Get a list of authenticated accounts for the current user.
// Parameters:
// - limit int3, The number of accounts to display per page. By default, displays 49 (max 250).
// If has_next is true, additional pages of accounts are available to be fetched. Use the cursor
// parameter to start on a specified page.
// - cursor string, For paginated responses, returns all responses that come after this value.
// - retailPortfolioId string, (Deprecated) Only returns the accounts matching the portfolio ID. Only applicable for legacy keys. CDP keys will default to the key's permissioned portfolio.
func (api *API) ListAccounts(limit int32, cursor, retailPortfolioId string) (*face.Accounts, error) {
	// "https://api.coinbase.com/api/v3/brokerage/accounts?limit=10&cursor=20&retail_portfolio_id=30"
	params := params.NewParams()

	//values := req.URL.Query()
	if limit > 250 {
		return nil, errors.New("limit must be less than 250")
	} else if limit > 0 {
		params.Add("limit", limit)
	}

	if cursor != "" {
		params.Add("cursor", cursor)
	}

	if retailPortfolioId != "" {
		params.Add("retail_portfolio_id", retailPortfolioId)
	}

	u, err := url.JoinPath(api.host.String(), ListAccountsPath)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	accounts := &face.Accounts{}
	err = api.GetRequest(u, ListAccountsPath, callListAccountsKey, params, accounts)

	return accounts, nil
}

// ListAllAccounts is a wrapper over ListAccounts that fetches all accounts
func (api *API) ListAllAccounts() ([]*face.Account, error) {
	var accounts []*face.Account
	cursor := ""
	hasNext := true

	for hasNext {
		nextAccounts, err := api.ListAccounts(250, cursor, "")
		if err != nil {
			return nil, err
		}

		if nextAccounts == nil {
			break
		}

		accounts = append(accounts, nextAccounts.GetAccounts()...)
		cursor = nextAccounts.Cursor
		hasNext = nextAccounts.HasNext
	}

	return accounts, nil
}
