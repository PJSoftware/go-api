package api

import "fmt"

const pkgVersion = "0.1.0"

func (a *APIData) Version() string {
	return fmt.Sprintf("go-api v%s: %s (%s)", pkgVersion, a.rootURL, a.auth.Type())
}
