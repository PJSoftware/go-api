package api

import "fmt"

const pkgVersion = "0.1.1"

func (a *APIData) Version() string {
	return pkgVersion
}

func (a *APIData) Ident() string {
	rv := fmt.Sprintf("go-api v%s: %s (%s)", pkgVersion, a.rootURL, a.auth.Type())
	if a.name != "" {
		rv = a.name + " via " + rv
	}
	return rv
}
