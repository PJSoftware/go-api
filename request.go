package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// An individual Request is used to communicate with the external API. A Request
// is generated via (*Endpoint).NewRequest()
type Request struct {
	endPoint *Endpoint
	queries  []reqQuery
	headers  []reqHeader
	bodyKV   []reqBody
	bodyTXT  string
	hasBody  bool
	options  ReqOptions
}

type reqQuery keyValuePair
type reqHeader keyValuePair
type reqBody keyValuePair

type keyValuePair struct {
	key   string
	value string
}

type Response struct {
	Status int
	Body   string
}

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
func (r *Request) AddBodyKV(key, value string) *Request {
	body := reqBody{}
	body.key = key
	body.value = value
	r.bodyKV = append(r.bodyKV, body)
	r.hasBody = true
	return r
}

// Set the body of the request to a block of JSON-formatted text
//
// TODO: implement proper error handling here
func (r *Request) SetBodyJSON(v any) *Request {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	r.bodyTXT = string(b)
	r.hasBody = true
	return r
}

// (*Request).RawQueryURL() generates the GET URL that would be generated by the
// Request and its query key/value pairs, and returns it as a string. This can be
// useful for Callback situations.
func (r *Request) RawQueryURL() (string, error) {
	epURL := r.endPoint.URL()
	httpReq, err := http.NewRequest("GET", epURL, nil)
	if err != nil {
		return "", &PackageError{err}
	}

	httpQuery := httpReq.URL.Query()
	for _, qry := range r.queries {
		httpQuery.Add(qry.key, qry.value)
	}
	httpReq.URL.RawQuery = httpQuery.Encode()
	return httpReq.URL.String(), nil
}

// (*Request).GET() processes a GET call to the API
func (r *Request) GET() (*Response, error) {
	return r.callAPI("GET")
}

// (*Request).POST() processes a POST call to the API
func (r *Request) POST() (*Response, error) {
	return r.callAPI("POST")
}

// callAPI() handles the call using the specified method
func (r *Request) callAPI(method string) (*Response, error) {
	var httpReq *http.Request
	var err error

	epURL := r.endPoint.URL()
	httpClient := http.Client{}

	if r.hasBody {
		var bodyString *strings.Reader
		if len(r.bodyTXT) > 0 {
			bodyString = strings.NewReader(r.bodyTXT)

		} else if len(r.bodyKV) > 0 {
			form := url.Values{}
			for _, body := range r.bodyKV {
				form.Add(body.key, body.value)
			}
			bodyString = strings.NewReader(form.Encode())
		}

		httpReq, err = http.NewRequest(method, epURL, bodyString)

	} else {
		httpReq, err = http.NewRequest(method, epURL, nil)

	}

	if err != nil {
		return nil, &PackageError{fmt.Errorf("error in %s(): creating *http.Request: %w", method, err)}
	}

	httpQuery := httpReq.URL.Query()
	for _, qry := range r.queries {
		httpQuery.Add(qry.key, qry.value)
	}
	httpReq.URL.RawQuery = httpQuery.Encode()

	for _, hdr := range r.headers {
		httpReq.Header.Set(hdr.key, hdr.value)
	}

	res, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, &PackageError{fmt.Errorf("error in %s(): communicating with api: %w", method, err)}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &PackageError{fmt.Errorf("error in %s(): reading body of response: %w", method, err)}
	}

	rv := &Response{}
	rv.Status = res.StatusCode
	rv.Body = string(body)

	if rv.Status != http.StatusOK {
		return rv, newQueryError(rv)
	}

	return rv, nil
}
