package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/iostrovok/coinbaseapi/api/apierrors"
	"github.com/iostrovok/coinbaseapi/api/auth"
	"github.com/iostrovok/coinbaseapi/api/error_response"
	"github.com/iostrovok/coinbaseapi/api/params"
)

type API struct {
	au        *auth.Auth
	host      *url.URL
	viewDebug bool

	xRatelimit *XRatelimitAPI
}

func New(a *auth.Auth, host string) (*API, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return &API{
		au:   a,
		host: u,

		xRatelimit: NewXRatelimitAPI(),
	}, nil
}

func (api *API) SignRequest(req *http.Request, requestMethod, requestPath string) error {
	sign, err := api.au.JWT(requestMethod, api.host.Host, requestPath)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+sign)

	return nil
}

func (api *API) SetPrintDebugOn() {
	api.viewDebug = true
}

func (api *API) SetPrintDebugOff() {
	api.viewDebug = false
}

func (api *API) GetRequest(url, singPath, callKey string, params *params.Params, inOut any) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "http.NewRequest")
	}

	if params != nil {
		if qs := params.QueryString(); qs != "" {
			req.URL.RawQuery = qs
		}
	}

	if api.viewDebug {
		fmt.Printf("RawQuery: %s\n", req.URL.RawQuery)
		fmt.Printf("URL: %s\n", req.URL.String())
		fmt.Printf("singPath: %s\n", singPath)
	}

	req.Header.Add("Content-Type", "application/json")
	if err := api.SignRequest(req, http.MethodGet, singPath); err != nil {
		return errors.Wrap(err, "api.SignRequest")
	}

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "client.Do")
	}
	defer res.Body.Close()
	if api.viewDebug {
		fmt.Printf("res.StatusCode: %d\n", res.StatusCode)
	}

	api.xRatelimit.Add(callKey, res.Header)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "io.ReadAll")
	}

	if api.viewDebug {
		fmt.Printf("body:\n%v\n", string(body))
	}

	if errorResponse := error_response.Parse(body); errorResponse != nil {
		return errorResponse
	}

	if err := json.Unmarshal(body, &inOut); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	return nil
}

func (api *API) PostRequest(url, singPath, callKey string, params any, inOut any) error {
	var reqBody io.Reader
	if params != nil {
		b, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "PostRequest")
		}

		if api.viewDebug {
			fmt.Printf("\nreqBody:\n%v\n", string(b))
		}

		reqBody = bytes.NewReader(b)
	} else {
		if api.viewDebug {
			fmt.Printf("\nreqBody:\nnill\n")
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return errors.Wrap(err, "PostRequest")
	}

	req.Header.Add("Content-Type", "application/json")
	if err := api.SignRequest(req, http.MethodPost, singPath); err != nil {
		return errors.Wrap(err, "api.SignRequest")
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "client.Do")
	}
	defer res.Body.Close()

	api.xRatelimit.Add(callKey, res.Header)

	if api.viewDebug {
		fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "io.ReadAll")
	}
	if api.viewDebug {
		fmt.Printf("body:\n%v\n", string(body))
	}

	if 400 < res.StatusCode && res.StatusCode < 500 {
		return errors.Wrap(apierrors.UnauthorizedError, string(body))
	}

	if 500 <= res.StatusCode {
		out := &apierrors.ApiError{}
		if err := json.Unmarshal(body, &out); err != nil {
			return errors.Wrap(err, "json.Unmarshal")
		}
		// MethodNotAllowedError = errors.Errorf("Method Not Allowed")
		// {"error":"unknown","error_details":"Method Not Allowed","message":"Method Not Allowed"}
		return errors.Errorf("error: %s, error_details: %s, message: %s", out.Error, out.ErrorDetails, out.Message)
	}

	if errorResponse := error_response.Parse(body); errorResponse != nil {
		return errorResponse
	}

	if err := json.Unmarshal(body, &inOut); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	return nil
}
