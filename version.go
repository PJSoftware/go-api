package api

import "fmt"

const pkgVersion = "0.4.6"

// api.Version() returns the current package version
func (a *APIData) Version() string {
	return pkgVersion
}

// api.Ident() returns an ident string based on api name, url, and package version
func (a *APIData) Ident() string {
	return fmt.Sprintf("go-api v%s: %s", pkgVersion, a.rootURL)
}
