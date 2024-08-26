package api

import (
	"strings"
	"time"
)

// Because we probably want to apply rate limiting to our API at the endpoint
// level, we need to cache our endpoints. Code which calls NewEndpoint()
// multiple times needs to return the same Endpoint object every time, so that
// its limiting (etc) can apply correctly across the lifetime of the program.
var epCache = make(map[string]*Endpoint)

// Each Endpoint should be individually managed by the client code. An Endpoint
// is generated via api.NewEndpoint()
type Endpoint struct {
	endpointURL string
	parent      *APIData
	rateLimiter *rateLimiter
}

// epCacheKey() generates a key for our endpoint cache.
//
// the returned key currently matches ep.URL() -- but that is not a functional
// requirement, and is likely to change if appropriate.
func (a *APIData) epCacheKey(epURL string) string {
	return a.rootURL + "/" + epURL
}

// (*APIData).NewEndpoint returns a new *Endpoint object
func (a *APIData) NewEndpoint(epURL string) *Endpoint {
	epURL = strings.TrimPrefix(epURL, a.rootURL)
	epURL = strings.TrimPrefix(epURL, "/")

	epCacheKey := a.epCacheKey(epURL)
	if ep, ok := epCache[epCacheKey]; ok {
		return ep
	}

	ep := Endpoint{}
	ep.endpointURL = epURL
	ep.parent = a
	ep.rateLimiter = nil

	epCache[epCacheKey] = &ep
	return &ep
}

// (*Endpoint).URL() returns the full URL for that endpoint
func (ep *Endpoint) URL() string {
	return ep.parent.rootURL + "/" + ep.endpointURL
}

func (ep *Endpoint) SetRateLimit(numberOfCalls int, inDuration time.Duration) {
	// For now, ignore any possible changes to rateLimiter
	if ep.rateLimiter == nil {
		ep.rateLimiter = newRateLimiter(numberOfCalls,inDuration)
	}
}

func (ep *Endpoint) waitForAvailability() {
	if ep.rateLimiter == nil {
		return
	}

	ep.rateLimiter.requestToken()
}