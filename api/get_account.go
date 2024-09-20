package api

// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getaccount

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/face"
)

const (
	GetAccountPath    = "/api/v3/brokerage/accounts"
	callGetAccountKey = "GetAccountPath"
)

// GetAccountXRatelimit returns the x-ratelimit headers for the GetAccount API call.
func (api *API) GetAccountXRatelimit() (*XRatelimitHeaders, bool) {
	return api.xRatelimit.Get(callGetAccountKey)
}

type GetAccountResult struct {
	Account *face.Account `json:"account"`
}

// GetAccount Get a list of information about an account, given an account UUID.
// - accountId string required, The account's UUID.
func (api *API) GetAccount(accountId string) (*face.Account, error) {
	if accountId == "" {
		return nil, errors.New("accountId is empty")
	}

	u, err := url.JoinPath(api.host.String(), GetAccountPath, accountId)
	if err != nil {
		return nil, errors.Wrap(err, "url.JoinPath")
	}

	account := &GetAccountResult{}
	err = api.GetRequest(u, GetAccountPath+"/"+accountId, callGetAccountKey, nil, account)

	return account.Account, err
}
