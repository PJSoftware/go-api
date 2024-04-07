package api

import "fmt"

var ErrTimeout = fmt.Errorf("timeout duration exceeded")

// PackageError is essentially just a "decorator"-style error
type PackageError struct {
	error
}

func (pe *PackageError) Unwrap() error {
	return pe.error
}
