package api

func (a *APIData) NewEndpoint(epURL string) *Endpoint {
	ep := &Endpoint{}
	ep.endpointURL = epURL
	ep.parent = a

	return ep
}

func (e *Endpoint) URL() string {
	return e.parent.rootURL + e.endpointURL
}
