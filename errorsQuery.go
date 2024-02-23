package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnsupported = errors.New("unsupported status code")
	ErrInformation = errors.New("for information")
	ErrSuccess     = errors.New("success but requires attention")
	ErrRedirection = errors.New("redirection response")
	ErrClient      = errors.New("client error")
	ErrServer      = errors.New("server error")
)

type QueryError struct {
	res *Response
	err error
}

func newQueryError(res *Response) *QueryError {
	var err error
	code := res.Status
	if code == http.StatusOK {
		return nil
	}

	switch {
	case code <= 99:
		err = ErrUnsupported
	case code <= 199:
		err = ErrInformation
	case code <= 299:
		err = ErrSuccess
	case code <= 399:
		err = ErrRedirection
	case code <= 499:
		err = ErrClient
	case code <= 599:
		err = ErrServer
	default:
		err = ErrUnsupported
	}

	return &QueryError{
		res: res,
		err: err,
	}
}

func (qe *QueryError) Unwrap() error {
	return qe.err
}

func (qe *QueryError) Error() string {
	txt := http.StatusText(qe.res.Status)
	if txt == "" {
		return fmt.Sprintf("status %d: %v", qe.res.Status, qe.err)
	}
	return fmt.Sprintf("status %d (%s): %v", qe.res.Status, http.StatusText(qe.res.Status), qe.err)
}

func (qe *QueryError) Response() *Response {
	return qe.res
}

func (qe *QueryError) Status() int {
	return qe.res.Status
}
