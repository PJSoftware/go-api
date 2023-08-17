package api

import "fmt"

type StatusError struct {
	err error
	res *Response
}

func (s *StatusError) Error() string {
	return fmt.Sprintf("status %d: error %v", s.res.Status, s.err)
}

func (s *StatusError) ReturnValue() *Response {
	return s.res
}
