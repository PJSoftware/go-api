package api

import "sync"

// Under the skin we need to use a *sync.Map for this
type endpointCache struct {
	cache *sync.Map
}

// Because we probably want to apply rate limiting to our API at the endpoint
// level, we need to cache our endpoints. Code which calls NewEndpoint()
// multiple times needs to return the same Endpoint object every time, so that
// its limiting (etc) can apply correctly across the lifetime of the program.
//
// Additionally, this needs to be thread-safe for running via a goroutine!
var epCache *endpointCache

// initialise our global epCache variable
func init() {
	epCache = newEPCache()
}

func newEPCache() *endpointCache {
	ep := &endpointCache{}
	ep.cache = new(sync.Map)
	return ep
}

func (epc *endpointCache) ValueOf(key string) *Endpoint {
	ep, ok := epc.cache.Load(key)
	if !ok { return nil }
	return ep.(*Endpoint)
}

func (epc *endpointCache) Store(key string, endpoint *Endpoint) {
	epc.cache.Store(key, endpoint)
}
