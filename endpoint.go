package api

import "strings"

// (*APIData).NewEndpoint returns a new *Endpoint object
//
// Note that if the Root URL is specified with a trailing '/', the endpoint URL
// should not have a leading '/' -- and vice versa.
//
// TODO: We need some internal code to ensure consistency in how we store URL values
func (a *APIData) NewEndpoint(epURL string) *Endpoint {
	ep := &Endpoint{}
	ep.endpointURL = strings.TrimPrefix(epURL, "/")
	ep.parent = a

	return ep
}

// (*Endpoint).URL() returns the full URL for that endpoint
func (e *Endpoint) URL() string {
	return e.parent.rootURL + "/" + e.endpointURL
}
