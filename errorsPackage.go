package api

// PackageError is essentially just a "decorator"-style error
type PackageError struct {
	error
}

func (pe *PackageError) Unwrap() error {
	return pe.error
}
