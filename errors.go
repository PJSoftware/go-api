package api

import "fmt"

type StatusError struct {
	err    error
	result *Result
}

func (s *StatusError) Error() string {
	return fmt.Sprintf("status %d: error %v", s.result.Status, s.err)
}

func (s *StatusError) ReturnValue() *Result {
	return s.result
}
