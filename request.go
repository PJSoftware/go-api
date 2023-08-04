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
func (r *Request) AddQuery(key, value string) *Request {
	query := reqQuery{}
	query.key = key
	query.value = value
	r.queries = append(r.queries, query)
	return r
}

// Add Header to a request
func (r *Request) AddHeader(key, value string) *Request {
	header := reqHeader{}
	header.key = key
	header.value = value
	r.headers = append(r.headers, header)
	return r
}

// FormEncoded adds a predefined (Content-Type) header to a request
func (r *Request) FormEncoded() {
	r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
}

// Add a line (in "key=value" format) to the Body of a request
func (r *Request) AddBody(key, value string) *Request {
	body := reqBody{}
	body.key = key
	body.value = value
	r.body = append(r.body, body)
	return r
}

// (*Request).GET() processes a GET call to the API
func (r *Request) GET() (*Result, error) {
	return r.callAPI("GET")
}

// (*Request).POST() processes a POST call to the API
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
			form.Add(body.key, body.value)
		}
		bodyString := strings.NewReader(form.Encode())
		httpReq, err = http.NewRequest(method, epURL, bodyString)
	} else {
		httpReq, err = http.NewRequest(method, epURL, nil)
	}
	
	if err != nil {
		return nil, fmt.Errorf("error creating *http.Request: %v", err)
	}

	httpQuery := httpReq.URL.Query()
	for _, qry := range r.queries {
		httpQuery.Add(qry.key, qry.value)
	}
	httpReq.URL.RawQuery = httpQuery.Encode()

	for _, hdr := range r.headers {
		httpReq.Header.Set(hdr.key, hdr.value)
	}

	response, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error communicating with api: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		err = fmt.Errorf("[%d] %s", response.StatusCode, string(body))
		// TODO: Is this really an error? Should store StatusCode and return it so client software can decide how to handle it!
		return nil, fmt.Errorf("status not okay: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body of response: %v", err)
	}

	rv := &Result{}
	rv.Body = body
	return rv, nil
}