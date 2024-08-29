package api

import (
	"fmt"
	"net/http"
)

type Response struct {
	HTTPResponse *http.Response
	Body []byte
}

func newResponse(httpRes *http.Response, body []byte) *Response {
	return &Response{
		HTTPResponse: httpRes,
		Body: body,
	}
}

func (r *Response) String() string {
	return fmt.Sprintf("status: %d; body: '%s'", r.HTTPResponse.StatusCode, r.Body)
}
