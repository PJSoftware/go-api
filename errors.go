package api

import "fmt"

type RequestError struct {
	Err    error
	Return *Result
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: error %v", r.Return.Status, r.Err)
}

func (r *RequestError) ReturnValue() *Result {
	return r.Return
}
