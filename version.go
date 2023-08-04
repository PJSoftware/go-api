package api

import "fmt"

const pkgVersion = "0.1.1"

// api.Version() returns the current package version
func (a *APIData) Version() string {
	return pkgVersion
}

// api.Ident() returns an ident string based on api name, url, and package version
func (a *APIData) Ident() string {
	rv := fmt.Sprintf("go-api v%s: %s", pkgVersion, a.rootURL)
	if a.name != defaultAPIName {
		rv = a.name + " via " + rv
	}
	return rv
}
