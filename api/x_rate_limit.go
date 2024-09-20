package api

import (
	"net/http"
	"sync"
)

type XRatelimitHeaders struct {
	XRatelimitLimit     string `json:"x-ratelimit-limit"`
	XRatelimitRemaining string `json:"x-ratelimit-remaining"`
	XRatelimitReset     string `json:"x-ratelimit-reset"`
}

type XRatelimitAPI struct {
	xRatelimitMutes sync.RWMutex
	xRatelimit      map[string]*XRatelimitHeaders
}

// NewXRatelimitAPI creates a new instance of XRatelimitAPI
func NewXRatelimitAPI() *XRatelimitAPI {
	return &XRatelimitAPI{
		xRatelimitMutes: sync.RWMutex{},
		xRatelimit:      map[string]*XRatelimitHeaders{},
	}
}

// Get retrieves the rate limit headers for a given key
func (x *XRatelimitAPI) Get(key string) (*XRatelimitHeaders, bool) {
	x.xRatelimitMutes.RLock()
	res, find := x.xRatelimit[key]
	x.xRatelimitMutes.RUnlock()

	return res, find
}

// Add reads the limit's header to congregate the rates
func (x *XRatelimitAPI) Add(key string, header http.Header) {
	rate := &XRatelimitHeaders{
		XRatelimitLimit:     header.Get("x-ratelimit-limit"),
		XRatelimitRemaining: header.Get("x-ratelimit-remaining"),
		XRatelimitReset:     header.Get("x-ratelimit-reset"),
	}

	x.xRatelimitMutes.Lock()
	x.xRatelimit[key] = rate
	x.xRatelimitMutes.Unlock()
}
