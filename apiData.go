package api

import (
	"fmt"

	"github.com/pjsoftware/go-api/auth"
)

type APIData struct {
	RootURL  string
	AuthData auth.Data
}

func New(rootURL string) *APIData {
	version()
	fmt.Printf("API root: %s\n", rootURL)
	rv := APIData{}
	rv.RootURL = rootURL
	rv.AuthData.AuthType = auth.None
	return &rv
}

func (a *APIData) Get(endPoint string) (*QRYResult, error) {
	qry := &APIQuery{}
	qry.EndPoint = a.RootURL + endPoint
	fmt.Printf("EndPoint for query: %s\n", qry.EndPoint)
	rv, err := qry.get()
	if err != nil {
		return nil, err
	}

	return rv, nil
}
