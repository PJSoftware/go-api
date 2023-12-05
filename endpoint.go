package api

import "strings"

// (*APIData).NewEndpoint returns a new *Endpoint object
func (a *APIData) NewEndpoint(epURL string) *Endpoint {
	ep := &Endpoint{}
	epURL = strings.TrimPrefix(epURL, a.rootURL)
	epURL = strings.TrimPrefix(epURL, "/")
	ep.endpointURL = epURL
	ep.parent = a

	return ep
}

// (*Endpoint).URL() returns the full URL for that endpoint
func (e *Endpoint) URL() string {
	return e.parent.rootURL + "/" + e.endpointURL
}
