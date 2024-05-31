package api

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