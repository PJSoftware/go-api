package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Initialise new empty API request on specified endpoint
func (e *Endpoint) NewRequest() *Request {
	req := &Request{}
	req.endPoint = e
	return req
}

// Add Query (passed in GET URL) to a request
func (r *Request) AddQuery(qryName, qryValue string) *Request {
	query := reqQuery{}
	query.name = qryName
	query.value = qryValue
	r.queries = append(r.queries, query)
	return r
}

// Add Header to a request
func (r *Request) AddHeader(hdrName, hdrValue string) *Request {
	header := reqHeader{}
	header.name = hdrName
	header.value = hdrValue
	r.headers = append(r.headers, header)
	return r
}

// FormEncoded adds a predefined (Content-Type) header to a request
func (r *Request) FormEncoded() {
	r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
}

// Add a line to the Body of a request
func (r *Request) AddBody(bodyName, bodyValue string) *Request {
	body := reqBody{}
	body.name = bodyName
	body.value = bodyValue
	r.body = append(r.body, body)
	return r
}

// GET() processes a GET call to the API
func (r *Request) GET() (*Result, error) {
	return r.callAPI("GET")
}

// POST() processes a POST call to the API
func (r *Request) POST() (*Result, error) {
	return r.callAPI("POST")
}

// callAPI() handles the call using the specified method
func (r *Request) callAPI(method string) (*Result, error) {
	var httpReq *http.Request
	var err error
	
	epURL := r.endPoint.URL()
	httpClient := http.Client{}

	if len(r.body) > 0 {
		form := url.Values{}
		for _, body := range r.body {
			form.Add(body.name, body.value)
		}
		bodyString := strings.NewReader(form.Encode())
		httpReq, err = http.NewRequest(method, epURL, bodyString)
	} else {
		httpReq, err = http.NewRequest(method, epURL, nil)
	}
	
	if err != nil {
		return nil, fmt.Errorf("err 01: %v", err)
	}

	httpQuery := httpReq.URL.Query()
	for _, qry := range r.queries {
		httpQuery.Add(qry.name, qry.value)
	}
	httpReq.URL.RawQuery = httpQuery.Encode()

	for _, hdr := range r.headers {
		httpReq.Header.Set(hdr.name, hdr.value)
	}

	response, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("err 02: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		err = fmt.Errorf("[%d] %s", response.StatusCode, string(body))
		return nil, fmt.Errorf("err 03: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("err 04: %v", err)
	}

	rv := &Result{}
	rv.Body = body
	return rv, nil
}