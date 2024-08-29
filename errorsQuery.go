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
	req *Request
	res *Response
	err error
}

func newQueryError(res *Response, req *Request) *QueryError {
	var err error
	code := res.HTTP.StatusCode
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
		req: req,
		res: res,
		err: err,
	}
}

func (qe *QueryError) Unwrap() error {
	return qe.err
}

func (qe *QueryError) Error() string {
	txt := http.StatusText(qe.res.HTTP.StatusCode)
	if txt == "" {
		return fmt.Sprintf("status %d: %v", qe.res.HTTP.StatusCode, qe.err)
	}
	return fmt.Sprintf("status %d (%s): %v", qe.res.HTTP.StatusCode, txt, qe.err)
}

func (qe *QueryError) Response() *Response {
	return qe.res
}

func (qe *QueryError) Request() *Request {
	return qe.req
}

func (qe *QueryError) Status() int {
	return qe.res.HTTP.StatusCode
}
