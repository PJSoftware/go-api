package api

type APIQuery struct {
	EndPoint string
}

func (q *APIQuery) Get() *QRYResult {
	rv := &QRYResult{}
	rv.Body = "---"
	return rv
}
