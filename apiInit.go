package api

func New(rootURL string) *APIData {
	rv := APIData{}
	rv.rootURL = rootURL
	rv.auth.NoAuth()
	return &rv
}

func (a *APIData) SetName(name string) {
	a.name = name
}
