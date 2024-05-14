package api

import "strings"

const libName = "go-api"

// APIData is the main export from go-api; it is generated via api.New()
type APIData struct {
	rootURL string
	Options *Options
}

// api.New() returns a new APIData object based on the specified URL. The rootURL
// string should be the base URL for the API
func New(rootURL string) *APIData {
	return &APIData{
		rootURL: strings.TrimSuffix(rootURL, "/"),
		Options: &Options{},
	}
}
