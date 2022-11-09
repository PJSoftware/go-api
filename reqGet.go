package api

import (
	"fmt"
	"io"
	"net/http"
)

func (r *Request) GET() (*Result, error) {
	epURL := r.endPoint.URL()

	httpClient := http.Client{}
	httpReq, err := http.NewRequest("GET", epURL, nil)
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
