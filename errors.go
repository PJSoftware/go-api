package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnsupportedRange = errors.New("unsupported status code")
	ErrInformation = errors.New("for information")
	ErrSuccess = errors.New("success but requires attention")
	ErrRedirection = errors.New("redirection response")
	ErrClient = errors.New("client error")
	ErrServer = errors.New("server error")
)

type QueryError struct {
	res *Response
	err error
}

func newQueryError(res *Response) *QueryError {
	var err error
	code := res.Status
	if code == http.StatusOK { return nil }

	switch {
	case code <= 99: err = ErrUnsupportedRange
	case code <= 199: err = ErrInformation
	case code <= 299: return nil // err = ErrSuccess
	case code <= 399: err = ErrRedirection
	case code <= 499: err = ErrClient
	case code <= 599: err = ErrServer
	default: err = ErrUnsupportedRange
	}
	
	if code > 100 && code <= 199 {
		err = ErrInformation
	}
	return &QueryError{
		res: res,
		err: err,
	}
}

func (s *QueryError) Unwrap() error {
	return s.err
}

func (s *QueryError) Error() string {
	txt := http.StatusText(s.res.Status)
	if txt == "" {
		return fmt.Sprintf("status %d: %v", s.res.Status, s.err)
	}
	return fmt.Sprintf("status %d (%s): %v", s.res.Status, http.StatusText(s.res.Status), s.err)
}

func (s *QueryError) Response() *Response {
	return s.res
}

func (s *QueryError) Status() int {
	return s.res.Status
}
