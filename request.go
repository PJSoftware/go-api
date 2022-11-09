package api

func (e *Endpoint) NewRequest() *Request {
	req := &Request{}
	req.endPoint = e
	return req
}

func (r *Request) AddQuery(qryName, qryValue string) *Request {
	query := reqQuery{}
	query.name = qryName
	query.value = qryValue
	r.queries = append(r.queries, query)
	return r
}

func (r *Request) AddHeader(hdrName, hdrValue string) *Request {
	header := reqHeader{}
	header.name = hdrName
	header.value = hdrValue
	r.headers = append(r.headers, header)
	return r
}
