package api

import "fmt"

type Response struct {
	Status int
	Body   []byte
}

func newResponse(status int, body []byte) *Response {
	return &Response{
		Status: status,
		Body:   body,
	}
}

func (r *Response) String() string {
	return fmt.Sprintf("status: %d; body: '%s'", r.Status, r.Body)
}