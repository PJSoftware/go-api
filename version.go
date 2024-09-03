package api

import "fmt"

const pkgName = "go-api"
const pkgVersion = "v0.4.21"

// api.Version() returns the current package version
func (a *APIData) Version() string {
	return pkgVersion
}

// api.Ident() returns an ident string based on api name, url, and package version
func (a *APIData) Ident() string {
	return fmt.Sprintf("go-api %s: %s", pkgVersion, a.rootURL)
}
