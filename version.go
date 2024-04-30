package api

import "fmt"

const Version = "v0.4.9"

// api.Version() returns the current package version
func (a *APIData) Version() string {
	return Version
}

// api.Ident() returns an ident string based on api name, url, and package version
func (a *APIData) Ident() string {
	return fmt.Sprintf("go-api %s: %s", Version, a.rootURL)
}
