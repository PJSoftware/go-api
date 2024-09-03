package api

import (
	"fmt"
	"log/slog"
	"strings"
)

// APIData is the main export from go-api; it is generated via api.New()
type APIData struct {
	rootURL string
	Options *Options
}

// api.New() returns a new APIData object based on the specified URL. The rootURL
// string should be the base URL for the API
func New(rootURL string) *APIData {
	ad := &APIData{
		rootURL: strings.TrimSuffix(rootURL, "/"),
		Options: &Options{},
	}
	slog.Debug(fmt.Sprintf("new %s", ad.Ident()))
	return ad
}
