package api

import (
	"fmt"
	"io"
	"net/http"
)

type APIQuery struct {
	EndPoint string
}

func (q *APIQuery) get() (*QRYResult, error) {
	query := http.Client{}
	fmt.Printf("Get endpoint: %s\n", q.EndPoint)
	req, err := http.NewRequest("GET", q.EndPoint, nil)
	if err != nil {
		return nil, fmt.Errorf("err 01: %v", err)
	}

	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := query.Do(req)
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

	rv := &QRYResult{}
	rv.Body = body
	return rv, nil
}
