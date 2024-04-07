package api

type Response struct {
	Status int
	Body   string
}

func newResponse(status int, body string) *Response {
	return &Response{
		Status: status,
		Body:   body,
	}
}