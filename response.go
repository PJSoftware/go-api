package api

import (
	"fmt"
	"net/http"
)

type Response struct {
	HTTP *http.Response
	Body []byte
}

func newResponse(httpRes *http.Response, body []byte) *Response {
	return &Response{
		HTTP: httpRes,
		Body: body,
	}
}

func (r *Response) String() string {
	return fmt.Sprintf("status: %d; body: '%s'", r.HTTP.StatusCode, r.Body)
}
