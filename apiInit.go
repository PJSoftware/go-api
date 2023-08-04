package api

const defaultAPIName string = "No Name Specified"

// api.New() returns a new APIData object based on the specified URL. The rootURL
// string should be the base URL for the API
func New(rootURL string) *APIData {
	a := APIData{}
	a.rootURL = rootURL
	a.name = defaultAPIName
	return &a
}

// api.SetName() allows setting the (optional) name of the API for output purposes
func (a *APIData) SetName(name string) {
	a.name = name
}

// api.Name() returns the (optional) name of the API
func (a *APIData) Name() string {
	return a.name
}