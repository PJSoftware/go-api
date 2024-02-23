package api

import (
	"fmt"
)

type PackageError struct {
	err error
}

func newPackageError(err error) *PackageError {
	return &PackageError{err: err}
}

func (pe *PackageError) Unwrap() error {
	return pe.err
}

func (pe *PackageError) Error() string {
	return fmt.Sprintf("%v", pe.err)
}
